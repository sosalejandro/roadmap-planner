package entities

import (
	"roadmap-planner/pkg/planner/domain/value-objects"
)

// TODO: Add missing documentation to methods.
// TODO: Move ListManager to own file.
// TODO: Add tests to ListManager file

// ListManager is an interface which declares how to manipulate Lists and ideas
type ListManager interface {
	AddToList(idea value_objects.BaseNode) error
	InsertAtList(i int, idea value_objects.BaseNode) error
	promoteToSubList(idea value_objects.BaseNode) (*Sublist, error)
}

type IdeaPlanner struct {
	list *BaseList
}

func (ip *IdeaPlanner) AddToList(idea value_objects.BaseNode) error {
	//TODO implement me
	panic("implement me")
}

func (ip *IdeaPlanner) InsertAtList(i int, idea value_objects.BaseNode) error {
	//TODO implement me
	panic("implement me")
}

func (ip *IdeaPlanner) promoteToSubList(idea value_objects.BaseNode) (*Sublist, error) {
	//TODO implement me
	panic("implement me")
}
