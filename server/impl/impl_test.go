package impl_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	. "go.timothygu.me/downtomeet/server/impl"
)

const postgresDSNEnv = "POSTGRES_TEST_DSN"

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
	if v, ok := os.LookupEnv(postgresDSNEnv); ok {
		log.WithField("dsn", v).Info(fmt.Sprintf("Using environment variable %s for test DB", postgresDSNEnv))
		testImpl.Options.Database = v
	}

	u, err := url.Parse(testImpl.Options.Database)
	if err != nil {
		log.WithError(err).WithField("dsn", testImpl.Options.Database).Panic("Unable to parse DSN")
	}
	if u.Path != "/downtomeet_test" {
		log.WithField("dsn", testImpl.Options.Database).Panic(fmt.Sprintf("%s must end in /downtomeet_test", postgresDSNEnv))
	}
	u.Path = "/postgres"

	mgrDB, err := gorm.Open(postgres.Open(u.String()), &gorm.Config{
		Logger: db.Logger{Logger: log.WithField("source", "gormmgr")},
	})
	if err != nil {
		log.WithError(err).WithField("dsn", u.String()).Panic("Unable to connect to PostgreSQL")
	}
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
