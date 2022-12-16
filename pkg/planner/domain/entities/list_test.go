package entities

import (
	"reflect"
	"testing"
)

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
