package db

import (
	"log"

	"github.com/HashimovH/accounts-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/accounts"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
