package storage

import "roadmap-planner/pkg/planner/domain/entities"

// Name represents the ID for lists and base nodes
// applicable to the entities.ListManager interface
type Name string

type Repository interface {
	// AddIdea saves that given idea to the repository
	AddIdea(listName Name, idea entities.Idea) error
	// CreateList instantiates a list
	CreateList(name Name) entities.ListManager
	// PromoteToSublist promotes idea at i space slot to Sublist
	PromoteToSublist(listName Name, i int) (name string, err error)
	// InsertAtList inserts a new idea at i space slot
	InsertAtList(listName Name, i int, idea entities.Idea) error
}
