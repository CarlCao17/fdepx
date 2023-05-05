package fdepx

type TaskGraph struct {
	TaskGraphOptions
	Roots     []*Task
	inDegrees map[string]int32
	nodes     map[string]struct{}
	states    map[string]int32
	Results   map[string]TaskResult
}

const (
	TaskStateInitial = 0
	TaskStateRunning = 1
	TaskStateSucceed = 2
	TaskStateFailed  = 3
)

type TaskGraphOptions struct {
	FastFail bool // whether the child task should fast fail when parent task failed
}

func NewTaskGraph(roots ...*Task) *TaskGraph {
	return &TaskGraph{
		Roots:     roots,
		inDegrees: make(map[string]int32),
		nodes:     make(map[string]struct{}),
		states:    make(map[string]int32),
		Results:   make(map[string]TaskResult),
	}
}

func (g *TaskGraph) AddRootTask(r *Task) {
	g.Roots = append(g.Roots, r)
}

func (g *TaskGraph) BuildGraph() {
	for _, r := range g.Roots {
		g.inDegrees[r.Name] = 0
		r.Traverse(func(t *Task) {
			id := t.Name
			g.nodes[id] = struct{}{}
			g.states[id] = TaskStateInitial
			for _, c := range t.Children {
				g.inDegrees[c.Name]++
			}
		})
	}
}

func (g *TaskGraph) Run() {

}
