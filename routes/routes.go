package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bsodmike/go_starter_api/app"
	"github.com/gorilla/mux"
)

func InitializeRoutes(c *app.Config) {
	router := mux.NewRouter()
	apiRoutes := mux.NewRouter()

	c.Router = router
	c.ApiRoutes = apiRoutes

	// Health check
	router.HandleFunc("/health-check", healthCheckHandler).Methods("GET")

	// API
	apiRouter := apiRoutes.PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	apiRouter.HandleFunc("/", apiRootHandler).Methods("GET")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func apiRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)

	json.NewEncoder(w).Encode(make(map[string]string))
}
