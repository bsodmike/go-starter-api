package models

import (
	"fmt"

	gorm "github.com/jinzhu/gorm"

	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB abstraction
type DB struct {
	gormDB  *gorm.DB
	Source  string
	LogMode bool
}

func (d *DB) connect() *DB {
	gormDB, err := gorm.Open("postgres", d.Source)

	if err != nil {
		panic(err)
	}

	if err = gormDB.DB().Ping(); err != nil {
		panic(err)
	}

	if d.LogMode == true {
		gormDB.LogMode(true)
	}

	fmt.Printf("Connected to %s DB %s\n",
		gormDB.Dialect().GetName(), gormDB.Dialect().CurrentDatabase())

	d.gormDB = gormDB
	return d
}

// NewPostgresDB - postgres database
func NewPostgresDB(db *DB) *DB {

	db.connect()
	defer db.gormDB.Close()

	if !(db.gormDB.HasTable(&User{})) {
		db.gormDB.CreateTable(&User{})
	}

	db.gormDB.AutoMigrate(&User{})

	return db
}
