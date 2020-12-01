package impl_test

import (
	"bytes"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"go.timothygu.me/downtomeet/server/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/gorilla/sessions"

	. "go.timothygu.me/downtomeet/server/impl"
)

var testImpl = NewMockImplementation(mockCookieStore(), rand.NewSource(0))

var (
	mockAuthenticationKey = bytes.Repeat([]byte{'n'}, 32)
	mockEncryptionKey     = bytes.Repeat([]byte{'e'}, 16)
	TestUser *db.User
	FakeUser *db.User
	TestMeetup *db.Meetup
	TestTag *db.Tag
)

func mockCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore(mockAuthenticationKey, mockEncryptionKey)
}

func populateDatabase(i *Implementation) {
	database := i.DB()
	TestUser = &db.User{
		Email:           "trediehs@g.ucla.edu",
		Name:            "Tim Rediehs",
		ContactInfo:     "trediehs@g.ucla.edu",
		ProfilePic:      nil,
		FacebookID:      nil,
		GoogleID:        nil,
		Location:        db.Coordinates{
			Lat: swag.Float64(0),
			Lon: swag.Float64(0),
		},
		OwnedMeetups:    nil,
		Attending:       nil,
		Tags:            nil,
		PendingApproval: nil,
	}
	TestTag = &db.Tag{
		Name:    "Mental Health",
	}
	TestUser.Tags = append(TestUser.Tags, TestTag)
	database.Create(TestUser)
	FakeUser = &db.User{
		Model:           gorm.Model{},
		Email:           "",
		Name:            "",
		ContactInfo:     "",
		Location:        db.Coordinates{},
	}
	testImpl.DB().Create(&FakeUser)
	TestMeetup = &db.Meetup{
		Title:             "Group Painting",
		Time:              time.Unix(0,0),
		Description:       "",
		Tags:              nil,
		MaxCapacity:       10,
		MinCapacity:       1,
		Owner:             TestUser.ID,
		Location:          db.MeetupLocation{
			Coordinates: db.Coordinates{
				Lat: swag.Float64(0),
				Lon: swag.Float64(0),
			},
			URL:         "",
			Name:        "Null Island",
		},
		Cancelled:         false,
	}
	TestMeetup.Tags = append(TestMeetup.Tags, TestTag)
	TestMeetup.Attendees = append(TestMeetup.Attendees, TestUser)
	database.Create(TestMeetup)
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
