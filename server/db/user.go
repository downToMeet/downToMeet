package db

import (
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string
	Name            string
	ContactInfo     string
	Location        Coordinates `gorm:"embedded;embeddedPrefix:location_"`
	OwnedMeetups    []*Meetup   `gorm:"foreignKey:Owner"`
	Attending       []*Meetup   `gorm:"many2many:meetup_user_attend;"`
	Tags            []*Tag      `gorm:"many2many:tag_user;"`
	PendingApproval []*Meetup   `gorm:"many2many:meetup_user_pending;"`
}

func (u *User) IDString() string {
	if u == nil || u.ID == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(u.ID), 10)
}

func UserIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}
