package db

import (
	"strconv"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name    string    `gorm:"uniqueIndex"`
	Users   []*User   `gorm:"many2many:tag_user;"`
	Meetups []*Meetup `gorm:"many2many:meetup_tag;"`
}

func (t *Tag) IDString() string {
	if t == nil || t.ID == 0 {
		return ""
	}
	return strconv.FormatUint(uint64(t.ID), 10)
}

func TagIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}
