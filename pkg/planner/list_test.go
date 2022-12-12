package planner

import (
	"reflect"
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
		// TODO: Add test cases.
		{
			name: "Creates an Idea named 'Create graph model'",
			args: args{name: "Create graph model"},
			want: &Idea{&Node{name: "Create graph model"}},
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
				Node: &Node{
					name: "Create an application to manage organization",
				},
				ideas: make([]Idea, 0),
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
		Node  *Node
		Ideas []Idea
	}
	type args struct {
		idea Idea
	}
	type want struct {
		success bool
		state   []Idea
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		// TODO: Add test cases.
		{
			name: "Pushes 'Make benchmarks after logic'",
			fields: fields{
				Node:  NewNode("Makes benchmark after logic"),
				Ideas: make([]Idea, 0),
			},
			args: args{idea: *NewIdea("Make benchmarks after logic")},
			want: want{
				success: false,
				state:   []Idea{*NewIdea("Make benchmarks after logic")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.Node,
				ideas: tt.fields.Ideas,
			}
			if got := l.AddIdea(tt.args.idea); got != tt.want.success && !reflect.DeepEqual(l.ideas, tt.want.state) {
				t.Errorf("l.AddIdea(%s) = %v, want %v", tt.args.idea, got, tt.want.success)
			}
		})
	}
}

func TestList_InsertAt(t *testing.T) {
	type fields struct {
		Node  *Node
		ideas []Idea
	}
	type args struct {
		i    int
		idea Idea
	}
	type want struct {
		success bool
		idea    Idea
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		// TODO: Add test cases.
		{
			name: "Inserts at index 0 successfully",
			fields: fields{
				Node: NewNode("Create app"),
				ideas: []Idea{
					*NewIdea("Create Diagrams"),
					*NewIdea("Create Business Logic"),
					*NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i:    0,
				idea: *NewIdea("Create GitHub Repo"),
			},
			want: want{
				success: true,
				idea:    *NewIdea("Create GitHub Repo"),
			},
		},
		{
			name: "Inserts at index 2 successfully",
			fields: fields{
				Node: NewNode("Create app"),
				ideas: []Idea{
					*NewIdea("Create Diagrams"),
					*NewIdea("Create Business Logic"),
					*NewIdea("Create Folder Structure"),
				},
			},
			args: args{
				i:    2,
				idea: *NewIdea("Create GitHub Repo"),
			},
			want: want{
				success: true,
				idea:    *NewIdea("Create GitHub Repo"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.Node,
				ideas: tt.fields.ideas,
			}
			if got := l.InsertAt(tt.args.i, tt.args.idea); got != tt.want.success || !reflect.DeepEqual(l.ideas[tt.args.i], tt.want.idea) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want.success)
			}
		})
	}
}
