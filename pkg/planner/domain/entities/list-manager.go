package entities

import value_objects "roadmap-planner/pkg/planner/domain/value-objects"

// ListManager is an interface which declares how to manipulate Lists and ideas
type ListManager interface {
	// AddToList adds an Idea to the current list returning
	// an error specifying what's going on.
	AddToList(idea value_objects.BaseNode) error
	// InsertAtList takes an i specifying the index place of the sublist and an idea returning
	// an error specifying what's going on.
	InsertAtList(i int, idea value_objects.BaseNode) error
	// PromoteToSublist takes an i specifying the index place of the Idea
	// to be promoted to Sublist, returning also the error specifying
	// what's going on.
	PromoteToSublist(i int) (*Sublist, error)
}
