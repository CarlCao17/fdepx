package fdepx

import (
	"context"
	"errors"
	"runtime/debug"
	"time"
)

type TaskFunc func(ctx context.Context) (interface{}, error)

type Task struct {
	TaskOption
	Name string // the identifier of the task, should be guaranteed be unique.
	F    TaskFunc
	R    TaskResult
	Stat
	Parent   *Task
	Children []*Task
}

type TaskOption struct {
	ImmuneParent bool               // ImmuneParent
	Retryable    bool               // Retryable false: run only once
	RetryFunc    func(t *Task) bool // control whether to retry
	Timeout      time.Duration      // Timeout 0: no timeout
}

type Stat struct {
	execCnt int           // current execute count, from 0 to MaxRetry-1
	cost    time.Duration // the last success execute cost time
}

type TaskResult struct {
	Resp  interface{}
	Err   error
	Panic PanicInfo
}

type PanicInfo struct {
	Stack  []byte
	Reason interface{}
}

func NewTask(f TaskFunc, options ...TaskOptionFunc) *Task {
	task := &Task{
		F: f,
	}
	for _, o := range options {
		o(task)
	}
	return task
}

type TaskOptionFunc func(t *Task)

func WithName(name string) TaskOptionFunc {
	return func(t *Task) {
		t.Name = name
	}
}

func WithParent(p *Task) TaskOptionFunc {
	return func(t *Task) {
		t.Parent = p
	}
}

func WithChildren(ch ...*Task) TaskOptionFunc {
	return func(t *Task) {
		t.Children = ch
	}
}

func WithMaxRetryTimes(n int) TaskOptionFunc {
	return func(t *Task) {
		t.Retryable = true
		t.RetryFunc = func(t *Task) bool {
			return t.execCnt >= n-1
		}
	}
}

func WithRetryFunc(f func(t *Task) bool) TaskOptionFunc {
	return func(t *Task) {
		t.Retryable = true
		t.RetryFunc = f
	}
}

func WithImmuneParent() TaskOptionFunc {
	return func(t *Task) {
		t.ImmuneParent = true
	}
}

func WithTimeout(timeout time.Duration) TaskOptionFunc {
	return func(t *Task) {
		t.Timeout = timeout
	}
}

func (t *Task) AddChildren(ch ...*Task) {
	for _, c := range ch {
		t.addChild(c)
	}
}

func (t *Task) addChild(c *Task) {
	t.Children = append(t.Children, c)
	c.Parent = t
}

func (t *Task) AddParent(p *Task) {
	t.Parent = p
	p.AddChildren(t)
}

func (t *Task) Traverse(f func(*Task)) {
	visited := make(map[string]bool)
	t.traverse(visited, f)
}

func (t *Task) traverse(visited map[string]bool, f func(*Task)) {
	if visited[t.Name] {
		return
	}
	f(t)
	visited[t.Name] = true
	for _, c := range t.Children {
		c.traverse(visited, f)
	}
}

func (t *Task) Start() {

}

var (
	ErrParentFailed = errors.New("parent task failed and the task is not immune to the parent failure")
	ErrTimeout      = errors.New("task timeout")
)

func (t *Task) Run(ctx context.Context, pErr error) {
	if !t.ImmuneParent && pErr != nil {
		t.R.Err = ErrParentFailed
		return
	}
	taskF := t.doF
	if t.Retryable {
		taskF = Retry(t.doF, t.RetryFunc, t)
	}
	if t.Timeout > 0 {
		newCtx, cancel := context.WithTimeout(ctx, t.Timeout)
		ctx = newCtx
		defer cancel()
	}

	done := make(chan struct{})
	go func() {
		// 默认启用收集 Panic 信息
		defer func() {
			if r := recover(); r != nil {
				t.R.Panic = PanicInfo{
					Stack:  debug.Stack(),
					Reason: r,
				}
			}
		}()
		t.R.Resp, t.R.Err = taskF(ctx)
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		t.R.Err = ErrTimeout
		return
	case <-done:
		return
	}
}

func (t *Task) doF(ctx context.Context) (resp interface{}, err error) {
	start := time.Now()
	resp, err = t.F(ctx)
	if err == nil {
		t.cost = time.Since(start)
	}
	t.execCnt++
	return
}

func Retry(taskF TaskFunc, retryF func(task *Task) bool, t *Task) TaskFunc {
	return func(ctx context.Context) (resp interface{}, err error) {
		var i int
		for i = 0; retryF(t); i++ {
			resp, err = taskF(ctx)
			if err == nil {
				break
			}
		}
		return
	}
}
