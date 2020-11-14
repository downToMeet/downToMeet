package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func Get(logger FieldLogger, dsn string) (*gorm.DB, error) {
	var gormLogger gormlogger.Interface
	if logger != nil {
		gormLogger = Logger{logger}
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}
	if err := database.AutoMigrate(Meetup{}, User{}, Tag{}); err != nil {
		return nil, err
	}
	return database, nil
}
