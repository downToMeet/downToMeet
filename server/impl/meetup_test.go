package impl_test

import (
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/restapi/operations"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func populateDatabase(i *impl.Implementation) {
	database := i.DB()
	testUser := &db.User{
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
	testTag := &db.Tag{
		Name:    "Mental Health",
	}
	testUser.Tags = append(testUser.Tags, testTag)
	database.Create(testUser)
	testMeetup := &db.Meetup{
		Title:             "Group Painting",
		Time:              time.Unix(0,0),
		Description:       "",
		Tags:              nil,
		MaxCapacity:       10,
		MinCapacity:       1,
		Owner:             testUser.ID,
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
	testMeetup.Tags = append(testMeetup.Tags, testTag)
	testMeetup.Attendees = append(testMeetup.Attendees, testUser)
	database.Create(testMeetup)
}

func TestGetMeetup(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupParams()
	params.Tags = append(params.Tags, "Mental Health")
	params.Lat = 0.000
	params.Lon = 0.000
	params.Radius = 5
	params.HTTPRequest = req

	raw := testImpl.GetMeetup(params)

	require.IsType(t, (*operations.GetMeetupOK)(nil), raw)
	res := raw.(*operations.GetMeetupOK)
	assert.Greaterf(t, len(res.Payload), 0, "I can't test anything if I there are no meetups in range")
	for _, meetup := range res.Payload {
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lat - params.Lat), params.Radius)
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lon - params.Lon), params.Radius)
	}
}
