package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string
	Name              string
	ContactInfo       string
	LocationLatitude  float64
	LocationLongitude float64
	OwnedMeetups      []Meetup  `gorm:"foreignKey:Owner"`
	Attending         []*Meetup `gorm:"many2many:meetup_user_attend;"`
	Tags              []Tag     `gorm:"many2many:tag_user;"`
}
