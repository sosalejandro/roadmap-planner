package entities

import (
	"reflect"
	value_objects "roadmap-planner/pkg/planner/domain/value-objects"
	"testing"
)

var ideasList = []value_objects.BaseNode{
	NewIdea("Create repository"),
	NewIdea("Brainstorm use cases"),
	NewIdea("Create data models"),
	NewIdea("Define relationships"),
	NewIdea("Define scope of the problem"),
	NewIdea("Create sequence diagrams"),
	NewIdea("Create C4 Diagrams"),
	NewIdea("Create CI Pipeline"),
	NewIdea("Create CD Pipeline"),
	NewIdea("Setup actions to display tests"),
	NewIdea("Add PR template"),
	NewIdea("Create K8s infrastructure"),
	NewIdea("Define context"),
	NewIdea("Benchmark performance"),
	NewIdea("Create integration tests"),
	NewIdea("Create E2E tests"),
	NewIdea("Create Postman Documentation"),
	NewIdea("Create Postman Tests (E2E)"),
	NewIdea("Create Swagger Docs"),
	NewIdea("Add OAuth2"),
	NewIdea("Add Google Calendar integration"),
	NewIdea("Create software documentation"),
}

func TestNewIdea(t *testing.T) {
	for _, tt := range newIdeaTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdea(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewList(t *testing.T) {
	for _, tt := range newListTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AddIdea(t *testing.T) {
	for _, tt := range listAddIdeaTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			if got := l.AddIdea(tt.args.idea); got != tt.want.success ||
				!reflect.DeepEqual(l.ideas, tt.want.state) {
				t.Errorf("l.AddIdea(%s) = %v, want %v", tt.args.idea, got, tt.want.success)
			}
		})
	}
}

func TestList_InsertAt(t *testing.T) {
	for _, tt := range listInsertAtTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			got := l.InsertAt(tt.args.i, tt.args.idea)

			if !tt.want.success && got != tt.want.success {
				t.Errorf("InsertAt() expect %v, got %v", tt.want.success, got)
			}

			if got != tt.want.success &&
				!reflect.DeepEqual(l.ideas[tt.args.i], tt.want.idea) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want.success)
			}
		})
	}
}

func TestList_DeleteAt(t *testing.T) {
	for _, tt := range listDeleteAtTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			if got := l.DeleteAt(tt.args.i); got != tt.want {
				t.Errorf("DeleteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdea_String(t *testing.T) {
	for _, tt := range ideaStringMethodTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_MergeList(t *testing.T) {
	for _, tt := range listMergeListsTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			if got := l.MergeList(tt.args.ideas); got != tt.want {
				t.Errorf("MergeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Swap(t *testing.T) {
	for _, tt := range listSwapTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			if got := l.Swap(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_GetIdeas(t *testing.T) {
	for _, tt := range listGetIdeasTests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Node:  tt.fields.node,
				ideas: tt.fields.ideas,
			}
			if got := l.GetIdeas(); got != tt.want {
				t.Errorf("GetIdeas() = %v, want %v", got, tt.want)
			}
		})
	}
}
