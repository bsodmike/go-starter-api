package app

import (
	"github.com/bsodmike/go_starter_api/api"
	"github.com/gorilla/mux"
)

type Config struct {
	Router    *mux.Router
	APIRoutes *mux.Router
	API       *api.API
}
