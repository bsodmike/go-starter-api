package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index" json:"username"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"not null" json:"-"`
	APIToken string `gorm:"not null;unique_index" json:"apiToken"`
	UUID     string `gorm:"not null;unique_index" json:"uuid"`
}

// UserManager struct
type UserManager struct {
	db *DB
}

// NewUserManager - Create a new *UserManager that can be used for managing users.
func NewUserManager(db *DB) (*UserManager, error) {

	db.gormDB.AutoMigrate(&User{})

	usermgr := UserManager{}

	usermgr.db = db

	return &usermgr, nil
}

func remove() {
	bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	uuid.NewV4()
}
