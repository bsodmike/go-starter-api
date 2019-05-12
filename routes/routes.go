package routes

import (
	"github.com/bsodmike/go_starter_api/api"
	"github.com/bsodmike/go_starter_api/app"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewRoutes - Create routes
func NewRoutes(c *app.Config, appAPI *api.API) *mux.Router {
	router := mux.NewRouter()
	apiRoutes := mux.NewRouter()

	// Health check
	router.HandleFunc("/health-check", healthCheckHandler).Methods("GET")

	// Public API
	publicAPIRouter := router.PathPrefix("/api/v1").Subrouter()
	publicAPIRouter.HandleFunc("/", apiRootHandler).Methods("GET")

	u := publicAPIRouter.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", appAPI.UserSignup).Methods("POST")

	// API secured, requires JWT auth.
	apiRouter := apiRoutes.PathPrefix("/api/v1/secured").Subrouter()
	apiRouter.HandleFunc("/", apiRootHandler).Methods("GET")
	apiRouter.HandleFunc("/userinfo", appAPI.UserInfoHandler).Methods("GET")

	// Projects
	apiRouter.HandleFunc("/projects", appAPI.ProjectsHandler).Methods("GET")

	router.PathPrefix("/api/v1/secured").Handler(negroni.New(
		negroni.HandlerFunc(appAPI.AuthJWTMiddleware),
		negroni.Wrap(apiRoutes),
	))

	return router
}
