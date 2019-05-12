package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/bsodmike/go_starter_api/auth"
	"github.com/bsodmike/go_starter_api/models"
)

// UserJSON - JSON data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token - JSON data for JWT access token
type Token struct {
	AccessToken string `json:"access_token"`
}

func (t *Token) extractToken(jsontoken string) string {
	tokenErr := json.Unmarshal([]byte(jsontoken), &t)
	if tokenErr != nil {
		log.Fatal(tokenErr)
	}

	return t.AccessToken
}

// UserSignup -
func (api *API) UserSignup(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" || jsondata.Email == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(jsondata.Username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	emailErr := checkmail.ValidateFormat(jsondata.Email)
	if emailErr != nil {
		http.Error(w, "Invalid email provided", http.StatusBadRequest)
		return
	}

	if api.users.HasUserWithEmail(jsondata.Email) {
		http.Error(w, "User with email already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(jsondata.Username, jsondata.Password, jsondata.Email)

	jsontoken := auth.GetJSONToken(user)

	token := Token{}
	rawToken := token.extractToken(jsontoken)

	user.APIToken = rawToken
	api.users.UpdateUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfoHandler - Return user details as marshalled to JSON.
func (api *API) UserInfoHandler(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	jsonuser, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonuser)
}
