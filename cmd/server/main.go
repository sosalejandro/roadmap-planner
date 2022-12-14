package main

import (
	"fmt"
	"log"
	"net/http"
	"roadmap-planner/pkg/planner/http/rest"
	"roadmap-planner/pkg/planner/storage/memory"

	"github.com/gorilla/mux"
)

// TODO: Add Gremlin provider
// TODO: Add performance tests
// TODO: Add open api specs
// TODO: Add e2e tests
// TODO:
// TODO:

func main() {
	r := mux.NewRouter()
	RegisterListRoutes(r)

	fmt.Println("The list management app is on tap now: http://localhost:8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}

func RegisterListRoutes(r *mux.Router) {
	cV1 := initListWriteControllerV1()
	rest.RegisterListWriteRoutes(r, cV1, "v1")
}

func initListWriteControllerV1() rest.Controller {
	r := new(memory.IdeaPlannerMemoryStorage)
	s := rest.NewService(r)
	c := rest.Controller(rest.NewWriteController(s))
	return c
}
