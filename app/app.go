package app

import (
	"github.com/gorilla/mux"
)

type Config struct {
	Router    *mux.Router
	ApiRoutes *mux.Router
}
