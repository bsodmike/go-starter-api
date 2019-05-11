package models

import (
	// postgress db driver
	gorm "github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Project struct
type Project struct {
	gorm.Model
	Notes  []Note
	Name   string         `gorm:"not null;unique_index" json:"name"`
	Active bool           `gorm:"default:true" json:"active"`
	Tags   pq.StringArray `gorm:"type:varchar(64)[]"`
}

// ProjectManager struct
type ProjectManager struct {
	db *DB
}

func NewProjectManager(db *DB) (*ProjectManager, error) {
	db.gormDB.AutoMigrate(&Project{})

	projectmgr := ProjectManager{}
	projectmgr.db = db

	return &projectmgr, nil
}

func (state *ProjectManager) FindProjects() []Project {
	projects := []Project{}
	state.db.gormDB.Find(&projects)

	return projects
}

func (state *ProjectManager) AddProject(name string) *Project {
	project := &Project{Name: name}
	state.db.gormDB.Create(&project)

	return project
}
