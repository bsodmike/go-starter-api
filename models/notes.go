package models

import (
	// postgress db driver
	gorm "github.com/jinzhu/gorm"
)

// Note struct
type Note struct {
	gorm.Model
	ProjectID uint
	Name      string `gorm:"not null;unique_index" json:"name"`
	Active    bool   `gorm:"default:true" json:"active"`
}

// NoteManager struct
type NoteManager struct {
	db *DB
}
