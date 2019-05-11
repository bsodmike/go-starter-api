package main

import (
	"github.com/bsodmike/go_starter_api/api"
	"github.com/bsodmike/go_starter_api/app"
	"github.com/bsodmike/go_starter_api/models"
	"github.com/bsodmike/go_starter_api/routes"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	config := app.Config{}

	d := models.DB{Source: "host=localhost port=9001 user=dbuser dbname=goapi password=password sslmode=disable", LogMode: true}
	db := models.NewPostgresDB(&d)
	appAPI := api.NewAPI(db)

	router := routes.NewRoutes(&config, appAPI)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(router)

	n := negroni.Classic()
	n.UseHandler(handler)
	n.Run(":3000")

	defer db.Close()
}
