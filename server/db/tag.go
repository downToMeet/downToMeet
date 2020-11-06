package db

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name   string
	Users  []User   `gorm:"many2many:tag_user;"`
	Meetup []Meetup `gorm:"many2many:meetup_tag;"`
}
