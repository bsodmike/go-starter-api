package api

import (
	"github.com/bsodmike/go_starter_api/models"
)

// API struct
type API struct {
	users    *models.UserManager
	projects *models.ProjectManager
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
