package database

import (
	"Example/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	mydb *gorm.DB
}

func InitDB() (*gorm.DB, error) {
	db, _ := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=dbexample user=example password=myexample sslmode=disable")
	return db, nil
}

func CreateTable(db *gorm.DB) {
	db.CreateTable(&models.User{})
}
