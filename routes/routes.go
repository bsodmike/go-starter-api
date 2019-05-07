package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bsodmike/go_starter_api/app"
	"github.com/gorilla/mux"
)

func InitializeRoutes(c *app.Config) *mux.Router {
	mux := mux.NewRouter()
	c.Router = mux

	mux.HandleFunc("/health-check", HealthCheckHandler).Methods("GET")

	apiRouter := mux.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/", ApiRootHandler).Methods("GET")

	return mux
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func ApiRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)

	json.NewEncoder(w).Encode(make(map[string]string))
}
