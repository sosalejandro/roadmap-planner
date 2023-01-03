package rest

import (
	"errors"
	"fmt"
	"roadmap-planner/pkg/planner/domain/entities"
	value_objects "roadmap-planner/pkg/planner/domain/value-objects"
	"roadmap-planner/pkg/planner/storage"
)

type Service interface {
	// AddIdea saves an idea into the list
	AddIdea(listName string, name string) error
	// InsertAtList places an idea with given name at i space slot.
	InsertAtList(listName string, i int, name string) error
	// PromoteToSublist retrieves i space slot BaseNode and promotes it
	// to sublist or retrieves if it already exists as a sublist
	PromoteToSublist(listName string, i int) (name string, err error)
	// CreateList instantiates a list with given name
	CreateList(name string) (listName string, err error)
}

type service struct {
	lR storage.Repository
}

func NewService(r storage.Repository) Service {
	return &service{r}
}

func (s *service) AddIdea(listName string, name string) error {
	if err := s.lR.AddIdea(storage.Name(listName), *entities.NewIdea(name)); err != nil {
		return fmt.Errorf("something went wrong")
	}

	return nil
}

func (s *service) InsertAtList(listName string, i int, name string) error {
	if err := s.lR.InsertAtList(storage.Name(listName), i, *entities.NewIdea(name)); err != nil {
		return fmt.Errorf("something went wrong")
	}

	return nil
}

func (s *service) PromoteToSublist(listName string, i int) (name string, err error) {
	if name, err = s.lR.PromoteToSublist(storage.Name(listName), i); err != nil {
		err = fmt.Errorf("something went wrong")
	}

	return
}

func (s *service) CreateList(name string) (listName string, err error) {
	list := s.lR.CreateList(storage.Name(name))

	if list == nil {
		err = errors.New("something went wrong")
		return
	}

	listName = list.(value_objects.BaseNode).String()
	return
}
