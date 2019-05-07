package main

import (
	"github.com/bsodmike/go_starter_api/app"
	"github.com/bsodmike/go_starter_api/routes"
	"github.com/urfave/negroni"
)

func main() {
	config := app.Config{}

	routes := routes.InitializeRoutes(&config)

	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
