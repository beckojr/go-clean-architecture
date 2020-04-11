package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase return a new db instance
func NewDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=kaherecode dbname=kahere password=kaherecode sslmode=disable")
	if err != nil {
		log.Panic("Error connecting to the database", err)
		return nil
	}
	return db
}
