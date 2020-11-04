package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "postgresql://localhost:5432/downtomeet"

var instance *gorm.DB

func GetDB() (*gorm.DB, error) {
	if instance == nil {
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		if err := database.AutoMigrate(Meetup{}, User{}, Tag{}); err != nil {
			log.Println(err)
		}
		instance = database
	}
	return instance, nil
}
