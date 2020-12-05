package impl_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/go-openapi/swag"
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
	TestUser              *db.User
	TestUserFriend        *db.User
	TestUserFriend2       *db.User
	FakeUser              *db.User
	TestMeetup            *db.Meetup
	TestMeetupCanceled    *db.Meetup
	TestMeetupRemote      *db.Meetup
	TestTag               *db.Tag
)

func mockCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore(mockAuthenticationKey, mockEncryptionKey)
}

func populateDatabase(i *Implementation) {
	database := i.DB()
	TestUser = &db.User{
		Email:       "trediehs@g.ucla.edu",
		Name:        "Tim Rediehs",
		ContactInfo: "trediehs@g.ucla.edu",
		ProfilePic:  nil,
		FacebookID:  nil,
		GoogleID:    nil,
		Location: db.Coordinates{
			Lat: swag.Float64(0),
			Lon: swag.Float64(0),
		},
		OwnedMeetups:    nil,
		Attending:       nil,
		Tags:            nil,
		PendingApproval: nil,
	}
	TestTag = &db.Tag{
		Name: "Mental Health",
	}
	TestUser.Tags = append(TestUser.Tags, TestTag)
	database.Create(TestUser)
	TestUserFriend = &db.User{
		Email:       "timothygu99@gmail.com",
		Name:        "Tim Gu",
		ContactInfo: "timothygu99@gmail.com",
		ProfilePic:  nil,
		FacebookID:  nil,
		GoogleID:    nil,
		Location: db.Coordinates{
			Lat: swag.Float64(0),
			Lon: swag.Float64(0),
		},
		OwnedMeetups:    nil,
		Attending:       nil,
		Tags:            nil,
		PendingApproval: nil,
	}
	TestUserFriend.Tags = append(TestUserFriend.Tags, TestTag)
	database.Create(TestUserFriend)
	TestUserFriend2 = &db.User{
		Email:       "person@human.human",
		Name:        "Person Personson",
		ContactInfo: "person@human.human",
		ProfilePic:  nil,
		FacebookID:  nil,
		GoogleID:    nil,
		Location: db.Coordinates{
			Lat: swag.Float64(0),
			Lon: swag.Float64(0),
		},
		OwnedMeetups:    nil,
		Attending:       nil,
		Tags:            nil,
		PendingApproval: nil,
	}
	TestUserFriend2.Tags = append(TestUserFriend2.Tags, TestTag)
	database.Create(TestUserFriend2)
	FakeUser = &db.User{
		Model:       gorm.Model{},
		Email:       "",
		Name:        "",
		ContactInfo: "",
		Location:    db.Coordinates{},
	}
	testImpl.DB().Create(&FakeUser)
	TestMeetup = &db.Meetup{
		Title:       "Group Painting",
		Time:        time.Now().Add(24 * time.Hour),
		Description: "",
		Tags:        nil,
		MaxCapacity: 10,
		MinCapacity: 1,
		Owner:       TestUser.ID,
		Location: db.MeetupLocation{
			Coordinates: db.Coordinates{
				Lat: swag.Float64(0),
				Lon: swag.Float64(0),
			},
			URL:  "",
			Name: "Null Island",
		},
		Cancelled: false,
	}
	TestMeetupCanceled = &db.Meetup{
		Title:       "Basketball",
		Time:        time.Now().Add(24 * time.Hour),
		Description: "",
		Tags:        nil,
		MaxCapacity: 10,
		MinCapacity: 1,
		Owner:       TestUserFriend.ID,
		Location: db.MeetupLocation{
			Coordinates: db.Coordinates{
				Lat: swag.Float64(0),
				Lon: swag.Float64(0),
			},
			URL:  "",
			Name: "Null Island",
		},
		Cancelled: true,
	}
	TestMeetupRemote = &db.Meetup{
		Title:       "Cry",
		Time:        time.Now().Add(24 * time.Hour),
		Description: "",
		Tags:        nil,
		MaxCapacity: 10,
		MinCapacity: 1,
		Owner:       TestUserFriend.ID,
		Location: db.MeetupLocation{
			URL: "https://hack.uclaacm.com/",
		},
		Cancelled: false,
	}
	TestMeetup.Tags = append(TestMeetup.Tags, TestTag)
	TestMeetupRemote.Tags = append(TestMeetupRemote.Tags, TestTag)
	TestMeetupCanceled.Tags = append(TestMeetupCanceled.Tags, TestTag)
	TestMeetup.Attendees = append(TestMeetup.Attendees, TestUser, TestUserFriend)
	TestMeetup.PendingAttendees = append(TestMeetup.PendingAttendees, TestUserFriend2)
	database.Create(TestMeetup)
	database.Create(TestMeetupCanceled)
	database.Create(TestMeetupRemote)
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
