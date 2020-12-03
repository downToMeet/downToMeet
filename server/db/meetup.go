package db

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Meetup is the database model for a meetup.
type Meetup struct {
	gorm.Model
	Title             string
	Time              time.Time
	Description       string
	Tags              []*Tag `gorm:"many2many:meetup_tag;"`
	MaxCapacity       int64
	MinCapacity       int64
	Owner             uint
	Attendees         []*User        `gorm:"many2many:meetup_user_attend;"`
	Location          MeetupLocation `gorm:"embedded;embeddedPrefix:location_"`
	PendingAttendees  []*User        `gorm:"many2many:meetup_user_pending;"`
	RejectedAttendees []*User        `gorm:"many2many:meetup_user_rejected;"`
	Cancelled         bool
}

// IDString returns the meetup's ID as a string.
func (m *Meetup) IDString() string {
	if m == nil || m.ID == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(m.ID), 10)
}

// MeetupIDFromString reverses Meetup.IDString.
func MeetupIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}

// Coordinates represent WGS 84 coordinates of the earth. It is not a database
// model but rather stored as a part of MeetupLocation.
type Coordinates struct {
	Lat, Lon *float64
}

// MeetupLocation represents where a meetup could be held. It is not a database
// model but rather stored as a part of Meetup.
type MeetupLocation struct {
	Coordinates
	URL  string
	Name string
}
