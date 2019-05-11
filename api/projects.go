package api

import (
	"encoding/json"
	"net/http"
)

// ProjectJSON - JSON data expected for login/signup
type ProjectJSON struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func (api *API) ProjectsHandler(w http.ResponseWriter, req *http.Request) {
	projects := api.projects.FindProjects()
	json, _ := json.Marshal(projects)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
