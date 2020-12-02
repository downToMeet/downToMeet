package impl_test

import (
	"go.timothygu.me/downtomeet/server/db"
	"gorm.io/gorm"
)

func createUser(email string) *db.User {
	newUser := db.User{
		Model:           gorm.Model{},
		Email:           email,
		Name:            email,
		ContactInfo:     "",
		Location:        db.Coordinates{},
	}
	testImpl.DB().Create(&newUser)
	return &newUser
}
