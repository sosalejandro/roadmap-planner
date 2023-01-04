package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	CreateListPath       = "list/create/{name}"
	AddListPath          = "list/{id}/create/{idea}"
	InsertAtListPath     = "list/{id}/{subId}/create/{idea}"
	PromoteToSublistPath = "list/{id}/promote/{ideaId}"
)

type Controller interface {
	AddToList(w http.ResponseWriter, r *http.Request)
	InsertAtList(w http.ResponseWriter, r *http.Request)
	PromoteToSublist(w http.ResponseWriter, r *http.Request)
	CreateList(w http.ResponseWriter, r *http.Request)
}

func RegisterListWriteRoutes(r *mux.Router, c Controller, ver string) {
	protected := r.NewRoute().Subrouter()
	protected.HandleFunc(createRoute(ver, CreateListPath), c.CreateList).Methods("POST")
	protected.HandleFunc(createRoute(ver, AddListPath), c.AddToList).Methods("POST")
	protected.HandleFunc(createRoute(ver, InsertAtListPath), c.InsertAtList).Methods("POST")
	protected.HandleFunc(createRoute(ver, PromoteToSublistPath), c.PromoteToSublist).Methods("PUT")
}

func createRoute(ver string, path string) string {
	return fmt.Sprintf("/api/%s/%s", ver, path)
}
