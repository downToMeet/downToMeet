package impl_test

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"go.timothygu.me/downtomeet/server/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/sessions"

	. "go.timothygu.me/downtomeet/server/impl"
)

var testImpl = NewMockImplementation(mockCookieStore(), rand.NewSource(0))

var (
	mockAuthenticationKey = bytes.Repeat([]byte{'n'}, 32)
	mockEncryptionKey     = bytes.Repeat([]byte{'e'}, 16)
)

func mockCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore(mockAuthenticationKey, mockEncryptionKey)
}

func TestMain(m *testing.M) {
	testImpl.Options.Database = "postgresql://localhost:5432/downtomeet_test"

	u, err := url.Parse(testImpl.Options.Database)
	if err != nil {
		log.WithError(err).Panic("Unable to parse DSN")
	}
	u.Path = "postgres"

	mgrDB, err := gorm.Open(postgres.Open(u.String()), &gorm.Config{
		Logger: db.Logger{log.WithField("source", "gormmgr")},
	})
	mgrDB.Exec(`DROP DATABASE downtomeet_test`)
	defer mgrDB.Exec(`DROP DATABASE downtomeet_test`)
	if err := mgrDB.Exec(`CREATE DATABASE downtomeet_test`).Error; err != nil {
		log.WithError(err).Panic("Unable to create test database")
	}

	actualDB := testImpl.DB()
	if err := actualDB.Exec(`CREATE EXTENSION earthdistance CASCADE`).Error; err != nil {
		log.WithError(err).Panic("Unable to install earthdistance extension")
	}
	populateDatabase(testImpl)

	errCode := m.Run()
	mgrDB.Exec(`DROP DATABASE downtomeet_test`)
	os.Exit(errCode)
}
