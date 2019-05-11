package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bsodmike/go_starter_api/auth"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

// UserSignup -
func (api *API) UserSignup(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(jsondata.Username) {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	if api.users.HasUserWithEmail(jsondata.Email) {
		http.Error(w, "User with email already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(jsondata.Username, jsondata.Password, jsondata.Email)

	jsontoken := auth.GetJSONToken(user)

	var token Token
	rawToken := token.extractToken(jsontoken)

	user.APIToken = rawToken
	api.users.UpdateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

func (t *Token) extractToken(jsontoken string) string {
	tokenErr := json.Unmarshal([]byte(jsontoken), &t)
	if tokenErr != nil {
		log.Fatal(tokenErr)
	}

	return t.AccessToken
}
