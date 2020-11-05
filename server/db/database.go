package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "postgresql://localhost:5432/downtomeet"

var instance *gorm.DB
var once sync.Once

func GetDB() (*gorm.DB, error) {
	once.Do(func() {
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			// If we can't connect to the DB, there's no point doing anything else.
			log.Fatalf("Can't connect to the database: %s\n", err.Error())
		}
		if mErr := database.AutoMigrate(Meetup{}, User{}, Tag{}, MeetupLocation{}); mErr != nil {
			log.Println(err)
		}
		instance = database
	})
	return instance, nil
}
