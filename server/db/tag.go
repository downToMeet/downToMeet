package db

import (
	"strconv"

	"gorm.io/gorm"
)

// Tag is the database model for tags.
type Tag struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Users   []*User   `gorm:"many2many:tag_user;"`
	Meetups []*Meetup `gorm:"many2many:meetup_tag;"`
}

// IDString returns the tag's ID, represented as a string.
func (t *Tag) IDString() string {
	if t == nil || t.ID == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(t.ID), 10)
}

// TagIDFromString reverses Tag.IDString.
func TagIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}
