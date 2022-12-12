package planner

// ListManager is an interface which declares how to manipulate Lists and ideas
type ListManager interface {
}

// List is a representation of an entity which stores ideas
type List struct {
	*Node
	ideas []Idea
}

func (l *List) Len() int {
	return len(l.ideas)
}

// Idea is a representation of a task which could be mutated into other entities.
type Idea struct {
	*Node
}

// NewIdea initializes a new Idea taking a string name.
func NewIdea(name string) *Idea {
	n := NewNode(name)
	return &Idea{n}
}

// NewList initializes a new List taking a string name.
func NewList(name string) *List {
	n := NewNode(name)
	return &List{
		Node:  n,
		ideas: make([]Idea, 0),
	}
}

// AddIdea pushes an Idea to the end of the List.
func (l *List) AddIdea(idea Idea) bool {
	if idea == (Idea{}) {
		return false
	}

	l.ideas = append(l.ideas, idea)
	return true
}

// InsertAt pushes an Idea at i position.
func (l *List) InsertAt(i int, idea Idea) bool {
	if i < 0 || i > l.Len() {
		return false
	}

	if i == 0 {
		l.ideas = append([]Idea{idea}, l.ideas...)
		return true
	}

	if i == l.Len() {
		l.AddIdea(idea)
	}

	right := []Idea{idea}
	right = append(right, l.ideas[i:]...)
	l.ideas = append(l.ideas[:i], right...)

	return true
}

// DeleteAt deletes an Idea at i position.
func (l *List) DeleteAt(i int) bool {

	if i < 0 || i > l.Len() {
		return false
	}

	if i == 0 {
		l.ideas = l.ideas[i+1:]
	}

	l.ideas = append(l.ideas[:i], l.ideas[i+1:]...)
	return true
}

// TODO: MergeList takes a list to be merged with the current.
// TODO: Swap moves Idea at s position into t position.
// TODO: GetIdeas retrieves the name of all ideas.
