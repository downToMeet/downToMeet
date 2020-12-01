package impl_test

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)



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

func TestGetMeetupID(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = fmt.Sprintf("%v", TestMeetup.ID)
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDOK)(nil), raw)
	res := raw.(*operations.GetMeetupIDOK)
	assert.Equal(t, fmt.Sprintf("%v", res.Payload.ID), fmt.Sprintf("%v", TestMeetup.ID))
}

func TestGetMeetupIDNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	nonexistentUserID := fmt.Sprintf("%v", TestUser.ID + 100000)
	url.ID = nonexistentUserID
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID: nonexistentUserID,
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDNotFound)(nil), raw)
	res := raw.(*operations.GetMeetupIDNotFound)
	assert.Equal(t, res.Payload.Code, int32(404))
}

func TestPostMeetup(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	eventTitle := "Pet Jacob's Dog Ryder"
	params := operations.PostMeetupParams{
		HTTPRequest: req,
		Meetup: &models.MeetupRequestBody{
			Description: "",
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: swag.Float64(40),
					Lon: swag.Float64(40),
				},
				Name: "",
				URL:  "",
			},
			MaxCapacity: swag.Int64(1),
			MinCapacity: swag.Int64(2),
			Tags:        nil,
			Time:        strfmt.DateTime{},
			Title:       eventTitle,
		},
	}

	raw := testImpl.PostMeetup(params, nil)

	require.IsType(t, (*operations.PostMeetupCreated)(nil), raw)
	res := raw.(*operations.PostMeetupCreated)
	assert.Equal(t, res.Payload.Title, eventTitle)
}

func TestPostMeetupTags(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	eventTitle := "Pet Jacob's Dog Ryder"
	params := operations.PostMeetupParams{
		HTTPRequest: req,
		Meetup: &models.MeetupRequestBody{
			Description: "",
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: swag.Float64(40),
					Lon: swag.Float64(40),
				},
				Name: "",
				URL:  "",
			},
			MaxCapacity: swag.Int64(1),
			MinCapacity: swag.Int64(2),
			Tags:        []string{"Mental Health", "Magic: The Gathering"},
			Time:        strfmt.DateTime{},
			Title:       eventTitle,
		},
	}


	raw := testImpl.PostMeetup(params, nil)

	require.IsType(t, (*operations.PostMeetupCreated)(nil), raw)
	res := raw.(*operations.PostMeetupCreated)
	assert.Equal(t, res.Payload.Title, eventTitle)
	meetup := db.Meetup{}
	testImpl.DB().Preload("Tags").First(&meetup, fmt.Sprintf("%v", res.Payload.ID))
	assert.Equalf(t, 2, len(meetup.Tags),"This meetup was created with two tags.")
}

func TestPostMeetupBadUser(t *testing.T) {
	const sessionName = "session"
	const nonexistentUserId = 10000000
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", nonexistentUserId)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	eventTitle := "Pet Jacob's Dog Ryder"
	params := operations.PostMeetupParams{
		HTTPRequest: req,
		Meetup: &models.MeetupRequestBody{
			Description: "",
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: swag.Float64(40),
					Lon: swag.Float64(40),
				},
				Name: "",
				URL:  "",
			},
			MaxCapacity: swag.Int64(1),
			MinCapacity: swag.Int64(2),
			Tags:        nil,
			Time:        strfmt.DateTime{},
			Title:       eventTitle,
		},
	}

	raw := testImpl.PostMeetup(params, nil)

	assert.IsType(t, impl.InternalServerError{}, raw)
}

func TestPatchMeetup(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = fmt.Sprintf("%v", TestMeetup.ID)
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID: fmt.Sprintf("%v", TestMeetup.ID),
		Meetup: &models.MeetupRequestBody{
			Description: newDescription,
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: TestMeetup.Location.Coordinates.Lat,
					Lon: TestMeetup.Location.Coordinates.Lon,
				},
				Name: TestMeetup.Location.Name,
				URL:  TestMeetup.Location.URL,
			},
			MaxCapacity: &TestMeetup.MaxCapacity,
			MinCapacity: &TestMeetup.MinCapacity,
			Tags:        []string{TestTag.Name},
			Time:        strfmt.DateTime(TestMeetup.Time),
			Title:       TestMeetup.Title,
		},
	}
	raw := testImpl.PatchMeetupID(params, nil)
	require.IsType(t, (*operations.PatchMeetupIDOK)(nil), raw)
	res := raw.(*operations.PatchMeetupIDOK)
	assert.Equal(t, res.Payload.Description, newDescription)
}

func TestPatchMeetupNotFound(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	const nonexistentMeetupID = "101202"
	url := new(operations.PatchMeetupIDURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID: nonexistentMeetupID,
		Meetup: &models.MeetupRequestBody{
			Description: newDescription,
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: TestMeetup.Location.Coordinates.Lat,
					Lon: TestMeetup.Location.Coordinates.Lon,
				},
				Name: TestMeetup.Location.Name,
				URL:  TestMeetup.Location.URL,
			},
			MaxCapacity: &TestMeetup.MaxCapacity,
			MinCapacity: &TestMeetup.MinCapacity,
			Tags:        []string{TestTag.Name},
			Time:        strfmt.DateTime(TestMeetup.Time),
			Title:       TestMeetup.Title,
		},
	}
	raw := testImpl.PatchMeetupID(params, nil)
	require.IsType(t, (*operations.PatchMeetupIDNotFound)(nil), raw)
}

func TestPatchMeetupForbidden(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = fmt.Sprintf("%v", TestMeetup.ID)
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", FakeUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID: fmt.Sprintf("%v", TestMeetup.ID),
		Meetup: &models.MeetupRequestBody{
			Description: newDescription,
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: TestMeetup.Location.Coordinates.Lat,
					Lon: TestMeetup.Location.Coordinates.Lon,
				},
				Name: TestMeetup.Location.Name,
				URL:  TestMeetup.Location.URL,
			},
			MaxCapacity: &TestMeetup.MaxCapacity,
			MinCapacity: &TestMeetup.MinCapacity,
			Tags:        []string{TestTag.Name},
			Time:        strfmt.DateTime(TestMeetup.Time),
			Title:       TestMeetup.Title,
		},
	}
	raw := testImpl.PatchMeetupID(params, nil)
	require.IsType(t, (*operations.PatchMeetupIDForbidden)(nil), raw)
}

func TestDeleteMeetup(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:             "To Be Deleted",
		Time:              time.Time{},
		Description:       "",
		MaxCapacity:       2,
		MinCapacity:       1,
		Owner:             TestUser.ID,
		Location:          db.MeetupLocation{},
		Cancelled:         false,
	}
	testImpl.DB().Create(&doomedMeetup)

	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = fmt.Sprintf("%v", TestMeetup.ID)
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID: fmt.Sprintf("%v", doomedMeetup.ID),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDNoContent)(nil), raw)
}

func TestDeleteMeetupNotFound(t *testing.T) {
	const nonExistentMeetupID = 1000000
	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = fmt.Sprintf("%v", nonExistentMeetupID)
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", TestUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID: fmt.Sprintf("%v", nonExistentMeetupID),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDNotFound)(nil), raw)
}

func TestDeleteMeetupForbidden(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:             "To Be Deleted",
		Time:              time.Time{},
		Description:       "",
		MaxCapacity:       2,
		MinCapacity:       1,
		Owner:             TestUser.ID,
		Location:          db.MeetupLocation{},
		Cancelled:         false,
	}
	testImpl.DB().Create(&doomedMeetup)
	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = fmt.Sprintf("%v", TestMeetup.ID)
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	session.Values[impl.UserID] = fmt.Sprintf("%v", FakeUser.ID)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID: fmt.Sprintf("%v", doomedMeetup.ID),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDForbidden)(nil), raw)
}
