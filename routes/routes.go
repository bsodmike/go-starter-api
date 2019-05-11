package routes

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/bsodmike/go_starter_api/api"
	"github.com/bsodmike/go_starter_api/app"
	"github.com/bsodmike/go_starter_api/auth"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := auth.JwtMiddleware.CheckJWT(w, r)

	if len(r.Header["Authorization"]) != 0 {
		accessToken := strings.Split(r.Header["Authorization"][0], " ")[1]
		fmt.Printf("Access Token used | %s\n", accessToken)
	}

	// If there was an error, do not call next.
	if err == nil && next != nil {
		next(w, r)
	} else {
		fmt.Printf("Auth error=%s | %s %s | Remote %s\n", err.Error(), r.Method, html.EscapeString(r.URL.Path), r.RemoteAddr)
	}
}

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
		negroni.HandlerFunc(authMiddleware),
		negroni.Wrap(apiRoutes),
	))

	return router
}
