package impl_test

import (
	"fmt"
	"go.timothygu.me/downtomeet/server/db"
	"gorm.io/gorm"
	"sync/atomic"
)

var emailNumber int64 = 0
func newEmail() string {
	num := atomic.AddInt64(&emailNumber, 1)
	return fmt.Sprintf("%v@email.com", num)
}

func createUser() *db.User {
	newUser := db.User{
		Model:           gorm.Model{},
		Email:           newEmail(),
		Name:            newEmail(),
		ContactInfo:     "",
		Location:        db.Coordinates{},
	}
	if err := testImpl.DB().Create(&newUser).Error; err != nil {
		panic(err)
	}
	return &newUser
}
