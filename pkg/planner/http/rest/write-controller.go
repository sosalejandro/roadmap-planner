package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type WriteController struct {
	list Service
}

func NewWriteController(list Service) *WriteController {
	return &WriteController{list: list}
}

func (c *WriteController) AddToList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idea := params["idea"]

	// validate idea is not empty or whitespace
	if len(strings.TrimSpace(idea)) == 0 {
		err := fmt.Sprintf("idea param is empty")
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New idea has been added.")
}

func (c *WriteController) InsertAtList(w http.ResponseWriter, r *http.Request) {

}

func (c *WriteController) PromoteToSublist(w http.ResponseWriter, r *http.Request) {

}

func (c *WriteController) CreateList(w http.ResponseWriter, r *http.Request) {

}
