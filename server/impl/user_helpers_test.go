package impl_test

import (
	"fmt"
	"sync/atomic"

	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
)

var emailNumber int64 = 0

func newEmail() string {
	num := atomic.AddInt64(&emailNumber, 1)
	return fmt.Sprintf("%v@email.com", num)
}

func createUser() *db.User {
	newUser := db.User{
		Email:       newEmail(),
		Name:        newEmail(),
	}
	if err := testImpl.DB().Create(&newUser).Error; err != nil {
		panic(err)
	}
	return &newUser
}
