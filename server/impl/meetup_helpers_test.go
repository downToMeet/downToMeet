package impl_test

import (
	"go.timothygu.me/downtomeet/server/db"
)

func createMeetup(title string, ownerID uint, tags []*db.Tag, canceled bool) *db.Meetup {
	newMeetup := db.Meetup{
		Cancelled:    canceled,
		Description: "",
		MaxCapacity:      2,
		MinCapacity:      1,
		Owner:            ownerID,
		Tags:			  tags,
		Title:            title,
	}
	if err := testImpl.DB().Create(&newMeetup).Error; err != nil {
		panic(err)
	}
	return &newMeetup
}
