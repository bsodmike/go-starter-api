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

func (d *DB) Close() {
	defer d.Close()
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

	if !(db.gormDB.HasTable(&User{})) {
		db.gormDB.CreateTable(&User{})
	}

	if !(db.gormDB.HasTable(&Project{})) {
		db.gormDB.CreateTable(&Project{})
	}

	if !(db.gormDB.HasTable(&Note{})) {
		db.gormDB.CreateTable(&Note{})
	}

	db.gormDB.AutoMigrate(&User{})
	db.gormDB.AutoMigrate(&Project{})
	db.gormDB.AutoMigrate(&Note{})

	return db
}
