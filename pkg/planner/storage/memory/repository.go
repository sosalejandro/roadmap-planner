package memory

import (
	"errors"
	"roadmap-planner/pkg/planner/domain/entities"
	"roadmap-planner/pkg/planner/storage"
)

// IdeaPlannerMemoryStorage keeps data in memory
// using the entities.IdeaPlanner - entities.ListManager implementation - interface
type IdeaPlannerMemoryStorage struct {
	lM map[storage.Name]entities.ListManager
}

// AddIdea saves that given idea to the listName within the repository
func (s *IdeaPlannerMemoryStorage) AddIdea(listName storage.Name, idea entities.Idea) error {
	if err := s.lM[listName].AddToList(&idea); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// CreateList instantiates an entities.ListManager
func (s *IdeaPlannerMemoryStorage) CreateList(name storage.Name) entities.ListManager {
	return entities.NewIdeaPlanner(entities.NewList(string(name)))
}

// PromoteToSublist promotes idea at i space slot to Sublist
// returns sublist name and error
func (s *IdeaPlannerMemoryStorage) PromoteToSublist(listName storage.Name, i int) (name string, err error) {
	var sublist *entities.Sublist
	if sublist, err = s.lM[listName].PromoteToSublist(i); err != nil {
		err = errors.New(err.Error())
		return
	}
	name = sublist.String()
	return
}

// InsertAtList inserts a new idea at i space slot in given listName
func (s *IdeaPlannerMemoryStorage) InsertAtList(listName storage.Name, i int, idea entities.Idea) error {
	if err := s.lM[listName].InsertAtList(i, &idea); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
