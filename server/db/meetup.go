package db

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Meetup struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	ContactInfo string
	Time        string
	Description string
	Tags        []Tag `gorm:"many2many:meetup_tag;"`
	MaxCapacity int64
	MinCapacity int64
	Owner       string
	Attendees   []User `gorm:"many2many:meetup_user_attend;"`
}
