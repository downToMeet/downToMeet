package db

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MeetupLocation struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
	MeetupID string
	Lon int
	Lat int
	Url string
	Name string
}
