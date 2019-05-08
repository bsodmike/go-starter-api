package main

import (
	"github.com/gorilla/mux"
	"github.com/bsodmike/go_starter_api/app"
	"github.com/bsodmike/go_starter_api/api"
	"github.com/bsodmike/go_starter_api/routes"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func Setup(c *app.Config) *mux.Router {
	routes.InitializeRoutes(c)

	router := c.Router
	apiRoutes := c.ApiRoutes

	router.PathPrefix("/api/v1").Handler(negroni.New(
		negroni.HandlerFunc(api.AuthMiddleware),
		negroni.Wrap(apiRoutes),
	))

	return router
}

func main() {
	config := app.Config{}
	router := Setup(&config)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(router)

	n := negroni.Classic()
	n.UseHandler(handler)
	n.Run(":3000")
}
