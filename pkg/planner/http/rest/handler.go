package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"roadmap-planner/pkg/planner/storage/memory"
)

type Controller interface {
	AddToList(w http.ResponseWriter, r *http.Request)
	InsertAtList(w http.ResponseWriter, r *http.Request)
	PromoteToSublist(w http.ResponseWriter, r *http.Request)
	CreateList(w http.ResponseWriter, r *http.Request)
}

func Router() {
	router := mux.NewRouter()
	c := initController()

	router.HandleFunc("/api/v1/list/create/{name}", c.CreateList).Methods("POST")
	router.HandleFunc("/api/v1/list/{id}/create/{idea}", c.AddToList).Methods("POST")
	router.HandleFunc("/api/v1/list/{id}/{subId}/create/{idea}", c.InsertAtList).Methods("POST")
	router.HandleFunc("/api/v1/list/{id}/promote/{ideaId}", c.PromoteToSublist).Methods("PUT")

	//return router
}

func initController() Controller {
	r := new(memory.IdeaPlannerMemoryStorage)
	s := NewService(r)
	c := Controller(NewWriteController(s))
	return c
}
