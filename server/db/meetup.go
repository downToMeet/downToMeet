package db

import (
	"time"

	"gorm.io/gorm"
)

type Meetup struct {
	gorm.Model
	Title       string
	ContactInfo string
	Time        time.Time
	Description string
	Tags        []Tag `gorm:"many2many:meetup_tag;"`
	MaxCapacity int64
	MinCapacity int64
	Owner       string
	Attendees   []User         `gorm:"many2many:meetup_user_attend;"`
	Location    MeetupLocation `gorm:"embedded;embeddedPrefix:location_"`
}

type MeetupLocation struct {
	Lat, Lon float64
	URL      string
	Name     string
}
