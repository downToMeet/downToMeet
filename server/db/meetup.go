package db

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

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

func (m *Meetup) IDString() string {
	if m == nil || m.ID == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(m.ID), 10)
}

func MeetupIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}

type Coordinates struct {
	Lat, Lon *float64
}

type MeetupLocation struct {
	Coordinates
	URL string
}
