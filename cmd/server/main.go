package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"roadmap-planner/pkg/planner/http/rest"
	"roadmap-planner/pkg/planner/storage/memory"
)

func main() {
	router := mux.NewRouter()
	RegisterListRoutes(router)

	fmt.Println("The list management app is on tap now: http://localhost:8080")
	srv := &http.Server{
		Addr:    "8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

func RegisterListRoutes(r *mux.Router) {
	cV1 := initListControllerV1()
	rest.RegisterListRoutes(r, cV1, "v1")
}

func initListControllerV1() rest.Controller {
	r := new(memory.IdeaPlannerMemoryStorage)
	s := rest.NewService(r)
	c := rest.Controller(rest.NewWriteController(s))
	return c
}
