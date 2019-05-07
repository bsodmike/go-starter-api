package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bsodmike/go_starter_api/api"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/health-check", api.HealthCheckHandler).Methods("GET")

	apiRouter := a.Router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/", api.ApiRootHandler).Methods("GET")
}

func main() {
	a := App{}

	a.Initialize()

	a.Run(":3000")
}
