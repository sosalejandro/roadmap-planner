package entities

import (
	"roadmap-planner/pkg/planner/domain/value-objects"
	"strings"
)

// TODO: Increase range of tests to evaluate slightly more invalid states.
// TODO: Add CI Workflow template
// TODO: Add missing documentation to methods.
// TODO: Get 100% coverage.
// TODO: Add more business logic
// TODO: Research CLI apps and re usable interfaces for CLI/API
// TODO: Create sublist logic.

type BaseList interface {
	value_objects.BaseNode
	AddIdea(idea value_objects.BaseNode) bool
	GetIdeas() string
	MergeList(ideas []value_objects.BaseNode) bool
	InsertAt(i int, idea value_objects.BaseNode) bool
	DeleteAt(i int) bool
	Swap(s, t int) bool
	Len() int
}

// List is a representation of an entity which stores ideas
type List struct {
	*value_objects.Node
	ideas []value_objects.BaseNode
}

type Sublist struct {
	List
}

func (l *List) Len() int {
	return len(l.ideas)
}

// Idea is a representation of a task which could be mutated into other entities.
type Idea struct {
	*value_objects.Node
}

func (i *Idea) String() string {
	// Funny how referencing the pointer allow to access the base.String() which is found
	// in the Node.
	return i.Node.String()
}

// NewIdea initializes a new Idea taking a string name.
func NewIdea(name string) *Idea {
	n := Idea{value_objects.NewNode(name)}
	return &n
}

// NewList initializes a new List taking a string name.
func NewList(name string) *List {
	n := value_objects.NewNode(name)
	return &List{
		Node:  n,
		ideas: make([]value_objects.BaseNode, 0),
	}
}

// AddIdea pushes an Idea to the end of the List.
func (l *List) AddIdea(idea value_objects.BaseNode) bool {
	var ok bool
	if idea, ok = idea.(value_objects.BaseNode); !ok {
		return false
	}

	l.ideas = append(l.ideas, idea)
	return true
}

// InsertAt pushes an Idea at i position.
func (l *List) InsertAt(i int, idea value_objects.BaseNode) bool {
	if i < 0 || i > l.Len() {
		return false
	}

	if i == 0 {
		l.ideas = append([]value_objects.BaseNode{idea}, l.ideas...)
		return true
	}

	if i == l.Len() {
		l.AddIdea(idea)
	}

	right := []value_objects.BaseNode{idea}
	right = append(right, l.ideas[i:]...)
	l.ideas = append(l.ideas[:i], right...)

	return true
}

// TODO: Add len validation in test

// DeleteAt deletes an Idea at i position.
func (l *List) DeleteAt(i int) bool {

	if i < 0 || i >= l.Len() {
		return false
	}

	if i == 0 {
		l.ideas = l.ideas[i+1:]
		return true
	}

	if i == l.Len()-1 {
		l.ideas = l.ideas[:i]
		return true
	}

	l.ideas = append(l.ideas[:i], l.ideas[i+1:]...)
	return true
}

// MergeList takes a list to be merged with the current.
func (l *List) MergeList(ideas []value_objects.BaseNode) bool {
	if len(ideas) == 0 {
		return false
	}

	l.ideas = append(l.ideas, ideas...)
	return true
}

// Swap moves Idea at s position into t position.
func (l *List) Swap(s, t int) bool {
	validate := func(x int) bool {
		return x >= l.Len() || x < 0
	}

	if v1, v2 := validate(s), validate(t); v1 || v2 {
		return false
	}

	l.ideas[s], l.ideas[t] = l.ideas[t], l.ideas[s]
	return true
}

// GetIdeas retrieves the name of all ideas.
func (l *List) GetIdeas() string {
	var sb strings.Builder
	for k, idea := range l.ideas {
		sb.WriteString(idea.String())
		if k < l.Len()-1 {
			sb.WriteRune(',')
			continue
		}
		sb.WriteRune('.')
	}

	return sb.String()
}
