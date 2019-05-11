package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/bsodmike/go_starter_api/auth"
	"github.com/bsodmike/go_starter_api/models"
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

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
