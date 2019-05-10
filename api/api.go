package api

import (
	"fmt"
	"net/http"

	"github.com/bsodmike/go_starter_api/models"
)

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	next(rw, r)
	// do some stuff after
	fmt.Println("Finished API Auth Middleware")
}

// API -
type API struct {
	users *models.UserManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)

	return &API{
		users: usermgr,
	}
}
