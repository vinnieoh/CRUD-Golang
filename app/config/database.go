package config

import (
	"api/app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DataBaseConnect() *gorm.DB {

	
	 
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.User{})

    return db

}