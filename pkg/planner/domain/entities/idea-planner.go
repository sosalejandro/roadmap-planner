package entities

import (
	"errors"
	"fmt"
	"log"
	"roadmap-planner/pkg/planner/domain/value-objects"
)

// TODO: Add missing documentation to methods.
// TODO: Move ListManager to own file.
// TODO: Add tests to ListManager file

// Planner represents the contract interface which this type must comply with.
type Planner interface {
	BaseListRead
	BaseListWrite
}

// IdeaPlanner represents a ListManager complying class to perform
// administration on data nodes.
type IdeaPlanner struct {
	list *Planner
}

// NewIdeaPlanner factor creates a new IdeaPlanner pointer taking a Planner list as
// an instantiating parameter.
func NewIdeaPlanner(list Planner) *IdeaPlanner {
	return &IdeaPlanner{
		list: &list,
	}
}

// AddToList adds an Idea to the current list returning
// an error specifying what's going on.
func (ip *IdeaPlanner) AddToList(idea value_objects.BaseNode) error {
	var list = *ip.list
	if err := list.AddIdea(idea); !err {
		return errors.New(fmt.Sprintf("error inserting %s", idea))
	}
	return nil
}

// InsertAtList takes an i specifying the index place of the sublist and an idea returning
// an error specifying what's going on.
func (ip *IdeaPlanner) InsertAtList(i int, idea value_objects.BaseNode) error {
	list := *ip.list
	var old value_objects.BaseNode
	var found bool
	if old, found = list.GetElementAt(i); !found {
		return errors.New("element not found")
	}
	var sublist = old.(Planner)
	idp := NewIdeaPlanner(sublist)
	return idp.AddToList(idea)
}

// PromoteToSublist takes an i specifying the index place of the Idea
// to be promoted to Sublist, returning also the error specifying
// what's going on.
func (ip *IdeaPlanner) PromoteToSublist(i int) (*Sublist, error) {
	list := *ip.list
	// old Idea or Sublist node under inspection
	var old value_objects.BaseNode
	// found represents the existence of the element at i space slot
	// within the underlying BaseList
	var found bool

	if old, found = list.GetElementAt(i); !found {
		return nil, errors.New("element not found")
	}
	newSublist := NewSublist(old.String())

	if ip.validateSublist(old) {
		// getSublist retrieves the existing Sublist node from the underlying BaseList
		// and possibly returning an error for system states malfunctioning due to
		// data corruption.
		goto getSublist
	} else {
		// createSublist performs the deletion of the existing node at i space slot
		// and inserting the newly created Sublist into the i space slot.
		// Masking a replace operation.
		goto createSublist
	}

createSublist:
	if err := list.DeleteAt(i); !err {
		return nil, errors.New(fmt.Sprintf("error deleting %d index", i))
	}
	if err := list.InsertAt(i, newSublist); !err {
		return nil, errors.New(fmt.Sprintf("error inserting %d index", i))
	}

getSublist:
	if old, found = (*ip.list).GetElementAt(i); !found {
		return nil, errors.New(fmt.Sprintf("element not found at %d index", i))
	}
	return old.(*Sublist), nil
}

// validateSublist defers the possible panic caused by the
// type assertion performed in the old BaseNode
// verifying whether this object is a Sublist node or an Idea node.
func (ip *IdeaPlanner) validateSublist(old value_objects.BaseNode) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	_, ok := old.(*Sublist)
	return ok
}
