package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type WriteController struct {
	list Service
}

func NewWriteController(list Service) *WriteController {
	return &WriteController{list: list}
}

func (c *WriteController) AddToList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, idea := params["id"], params["idea"]

	// validate idea is not empty or whitespace
	if len(strings.TrimSpace(idea)) == 0 {
		err := fmt.Errorf("idea param is empty")
		throwBadRequest(w, err)
		return
	}

	if err := c.list.AddIdea(id, idea); err != nil {
		throwBadRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New idea has been added.")
}

func (c *WriteController) InsertAtList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, subId, idea := params["id"], params["subId"], params["idea"]

	subIdInt, err := strconv.Atoi(subId)
	if err != nil || subIdInt < 0 {
		throwBadRequest(w, err)
		return
	}

	// validate idea is not empty or whitespace
	if len(strings.TrimSpace(idea)) == 0 {
		err := fmt.Errorf("idea param is empty")
		throwBadRequest(w, err)
		return
	}

	err = c.list.InsertAtList(id, subIdInt, idea)
	if err != nil {
		throwBadRequest(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	msg := fmt.Sprintf("New idea has been added at %s list.", id)
	json.NewEncoder(w).Encode(msg)
}

func (c *WriteController) PromoteToSublist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ideaId := params["id"], params["idea"]

	ideaIdInt, err := strconv.Atoi(ideaId)
	if err != nil || ideaIdInt < 0 {
		throwBadRequest(w, err)
		return
	}

	name, err := c.list.PromoteToSublist(id, ideaIdInt)
	if err != nil {
		throwBadRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	msg := fmt.Sprintf("Idea %s has been promoted to sublist.", name)
	json.NewEncoder(w).Encode(msg)
}

func (c *WriteController) CreateList(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()
	logger.Println("executing create list v1")

	params := mux.Vars(r)
	name := params["name"]

	// validate name is not empty or whitespace
	if len(strings.TrimSpace(name)) == 0 {
		err := fmt.Errorf("idea param is empty")
		throwBadRequest(w, err)
		return
	}

	name, err := c.list.CreateList(name)
	if err != nil {
		throwBadRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	msg := fmt.Sprintf("List `%s` has been created.", name)
	json.NewEncoder(w).Encode(msg)
}

func throwBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
