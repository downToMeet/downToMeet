// Package db implements DownToMeet server's connection to the PostgreSQL
// database. It also contains the database model definitions as GORM models.
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Get returns a new gorm.DB from the given Data Source Name (DSN) connection
// string.
func Get(logger FieldLogger, dsn string) (*gorm.DB, error) {
	var gormLogger gormlogger.Interface
	if logger != nil {
		gormLogger = Logger{logger.WithField("source", "gorm")}
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
