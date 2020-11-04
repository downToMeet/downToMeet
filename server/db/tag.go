package db

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID     string
	Name   string
	Users  []User   `gorm:"many2many:tag_user;"`
	Meetup []Meetup `gorm:"many2many:meetup_tag;"`
}
