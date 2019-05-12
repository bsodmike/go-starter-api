package api

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/bsodmike/go_starter_api/auth"
	"github.com/bsodmike/go_starter_api/models"
)

// API struct
type API struct {
	users    *models.UserManager
	projects *models.ProjectManager
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// NewAPI - Returns and instance of *API
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	projectmgr, _ := models.NewProjectManager(db)

	return &API{
		users:    usermgr,
		projects: projectmgr,
	}
}

func (api *API) AuthJWTMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := auth.JwtMiddleware.CheckJWT(w, r)

	if len(r.Header["Authorization"]) != 0 {
		accessToken := strings.Split(r.Header["Authorization"][0], " ")[1]
		fmt.Printf("Access Token used | %s\n", accessToken)

		query, _ := api.users.FindUserByAccessToken(accessToken)
		if query.RecordNotFound() {
			fmt.Printf("User does not exist for bearer token! | token=%s | %s %s | Remote %s\n", accessToken, r.Method, html.EscapeString(r.URL.Path), r.RemoteAddr)

			RespondWithError(w, http.StatusUnauthorized, "User does not exist for bearer token!")
			return
		}
	}

	// If there was an error, do not call next.
	if err == nil && next != nil {
		next(w, r)
	} else {
		fmt.Printf("Auth error=%s | %s %s | Remote %s\n", err.Error(), r.Method, html.EscapeString(r.URL.Path), r.RemoteAddr)
	}
}
