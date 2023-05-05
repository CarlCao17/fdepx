package fdepx

import (
	"testing"

	"github.com/CarlCao17/fdepx/internal/assert"
)

func TestTaskGraph_BuildGraph(t *testing.T) {
	task0 := &Task{
		Name: "task-0",
	}
	subtask00 := &Task{
		Name: "task-00",
	}
	subtask01 := &Task{
		Name: "task-01",
	}
	task0.Children = []*Task{subtask00, subtask01}
	subtask00.Children = []*Task{
		{
			Name: "task-000",
		},
		{
			Name: "task-001",
		},
		{
			Name: "task-002",
		},
	}
	subtask01.Children = []*Task{
		{
			Name: "task-010",
		},
	}
	task1 := &Task{
		Name: "task-1",
	}
	subtask10 := &Task{
		Name: "task-10",
	}
	task1.Children = []*Task{subtask01, subtask10}

	type fields struct {
		graph *TaskGraph
		want  *TaskGraph
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "normal",
			fields: fields{
				graph: NewTaskGraph(task0, task1),
				want: &TaskGraph{
					Roots: []*Task{
						task0, task1,
					},
					inDegrees: map[string]int32{
						"task-0":   0,
						"task-1":   0,
						"task-00":  1,
						"task-01":  2,
						"task-10":  1,
						"task-000": 1,
						"task-001": 1,
						"task-002": 1,
						"task-010": 1,
					},
					nodes: map[string]struct{}{
						"task-0":   {},
						"task-1":   {},
						"task-00":  {},
						"task-01":  {},
						"task-10":  {},
						"task-000": {},
						"task-001": {},
						"task-002": {},
						"task-010": {},
					},
					states: map[string]int32{
						"task-0":   TaskStateInitial,
						"task-1":   TaskStateInitial,
						"task-00":  TaskStateInitial,
						"task-01":  TaskStateInitial,
						"task-10":  TaskStateInitial,
						"task-000": TaskStateInitial,
						"task-001": TaskStateInitial,
						"task-002": TaskStateInitial,
						"task-010": TaskStateInitial,
					},
					Results: make(map[string]TaskResult),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.fields.graph
			g.BuildGraph()
			assert.Equal(t, tt.fields.want, g)
		})
	}
}
