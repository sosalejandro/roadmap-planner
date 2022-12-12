package entities

import (
	"reflect"
	"roadmap-planner/pkg/planner/domain/value-objects"
	"testing"
)

func TestNewIdea(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Idea
	}{
		{
			name: "Creates an Idea named 'Create graph model'",
			args: args{name: "Create graph model"},
			want: func() *Idea {
				n := Idea(*value_objects.NewNode("Create graph model"))
				return &n
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdea(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewList(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		{
			name: "Creates a List named 'Create an application to manage organization'",
			args: args{name: "Create an application to manage organization"},
			want: &List{
				Node: value_objects.NewNode(
					"Create an application to manage organization"),
				ideas: make([]value_objects.BaseNode, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AddIdea(t *testing.T) {
	type fields struct {
		Node  *value_objects.Node
		Ideas []value_objects.BaseNode
	}
	type args struct {
		idea *Idea
	}
	type want struct {
		success bool
		state   []value_objects.BaseNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Pushes 'Make benchmarks after logic'",
			fields: fields{
				Node:  value_objects.NewNode("Makes benchmark after logic"),
				Ideas: make([]value_objects.BaseNode, 0),
			},
			args: args{idea: NewIdea("Make benchmarks after logic")},
			want: want{
				success: false,
				state:   []value_objects.BaseNode{NewIdea("Make benchmarks after logic")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.Node,
				ideas: tt.fields.Ideas,
			}
			if got := l.AddIdea(tt.args.idea); got != tt.want.success &&
				!reflect.DeepEqual(l.ideas, tt.want.state) {
				t.Errorf("l.AddIdea(%s) = %v, want %v", tt.args.idea, got, tt.want.success)
			}
		})
	}
}

func TestList_InsertAt(t *testing.T) {
	type fields struct {
		Node  *value_objects.Node
		ideas []value_objects.BaseNode
	}
	type args struct {
		i    int
		idea *Idea
	}
	type want struct {
		success bool
		idea    *Idea
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "Inserts at index 0 successfully",
			fields: fields{
				Node: value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{
					NewIdea("Create Diagrams"),
					NewIdea("Create Business Logic"),
					NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i:    0,
				idea: NewIdea("Create GitHub Repo"),
			},
			want: want{
				success: true,
				idea:    NewIdea("Create GitHub Repo"),
			},
		},
		{
			name: "Inserts at index 2 successfully",
			fields: fields{
				Node: value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{
					NewIdea("Create Diagrams"),
					NewIdea("Create Business Logic"),
					NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i:    2,
				idea: NewIdea("Create GitHub Repo"),
			},
			want: want{
				success: true,
				idea:    NewIdea("Create GitHub Repo"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.Node,
				ideas: tt.fields.ideas,
			}
			if got := l.InsertAt(tt.args.i, tt.args.idea); got != tt.want.success ||
				!reflect.DeepEqual(l.ideas[tt.args.i], tt.want.idea) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want.success)
			}
		})
	}
}

func TestList_DeleteAt(t *testing.T) {
	type fields struct {
		Node  *value_objects.Node
		ideas []value_objects.BaseNode
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Deletes at index 2 successfully",
			fields: fields{
				Node: value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{
					NewIdea("Create Diagrams"),
					NewIdea("Create Business Logic"),
					NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i: 2,
			},
			want: true,
		},
		{
			name: "Negative index returns false",
			fields: fields{
				Node:  value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{},
			},
			args: args{
				i: -1,
			},
			want: false,
		},
		{
			name: "Positive index on empty List returns false",
			fields: fields{
				Node:  value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{},
			},
			args: args{
				i: 2,
			},
			want: false,
		},
		{
			name: "Deletes at index 0 successfully",
			fields: fields{
				Node: value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{
					NewIdea("Create Diagrams"),
					NewIdea("Create Business Logic"),
					NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i: 0,
			},
			want: true,
		},
		{
			name: "Deletes at index 1 successfully",
			fields: fields{
				Node: value_objects.NewNode("Create app"),
				ideas: []value_objects.BaseNode{
					NewIdea("Create Diagrams"),
					NewIdea("Create Business Logic"),
					NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.Node,
				ideas: tt.fields.ideas,
			}
			if got := l.DeleteAt(tt.args.i); got != tt.want {
				t.Errorf("DeleteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
