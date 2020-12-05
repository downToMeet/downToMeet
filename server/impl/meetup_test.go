package impl_test

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.timothygu.me/downtomeet/server/impl/responders"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

const (
	nonexistentMeetupID = "100000"
	nonexistentUserID   = "100000"
)

func TestGetMeetup(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
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
	require.Greaterf(t, len(res.Payload), 0, "I can't test anything if I there are no meetups in range")
	for _, meetup := range res.Payload {
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lat-params.Lat), params.Radius)
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lon-params.Lon), params.Radius)
	}
}

func TestGetMeetupNoTags(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupParams()
	params.Lat = 0.000
	params.Lon = 0.000
	params.Radius = 5
	params.HTTPRequest = req

	raw := testImpl.GetMeetup(params)

	require.IsType(t, (*operations.GetMeetupOK)(nil), raw)
	res := raw.(*operations.GetMeetupOK)
	require.Greaterf(t, len(res.Payload), 0, "I can't test anything if I there are no meetups in range")
	for _, meetup := range res.Payload {
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lat-params.Lat), params.Radius)
		assert.LessOrEqual(t, math.Abs(*meetup.Location.Coordinates.Lon-params.Lon), params.Radius)
	}
}

func TestGetMeetupInvalidUserID(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "INVALID_ID"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupParams()
	params.Tags = append(params.Tags, "Mental Health")
	params.Lat = 0.000
	params.Lon = 0.000
	params.Radius = 5
	params.HTTPRequest = req

	raw := testImpl.GetMeetup(params)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestGetMeetupNone(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupParams()
	params.Tags = append(params.Tags, "Mental Health")
	params.Lat = 2.500
	params.Lon = 2.000
	params.Radius = 0
	params.HTTPRequest = req

	raw := testImpl.GetMeetup(params)

	require.IsType(t, (*operations.GetMeetupOK)(nil), raw)
	res := raw.(*operations.GetMeetupOK)
	require.Equal(t, 0, len(res.Payload))
}

func TestGetMeetupRemote(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupRemoteParams()
	params.Tags = append(params.Tags, "Mental Health")
	params.HTTPRequest = req

	raw := testImpl.GetMeetupRemote(params)

	require.IsType(t, (*operations.GetMeetupRemoteOK)(nil), raw)
	res := raw.(*operations.GetMeetupRemoteOK)
	require.Greaterf(t, len(res.Payload), 0, "I can't test anything if I there are no meetups in range")
}

func TestGetMeetupRemoteNoTags(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupRemoteParams()
	params.HTTPRequest = req

	raw := testImpl.GetMeetupRemote(params)

	require.IsType(t, (*operations.GetMeetupRemoteOK)(nil), raw)
	res := raw.(*operations.GetMeetupRemoteOK)
	require.Greaterf(t, len(res.Payload), 0, "I can't test anything if I there are no meetups in range")
}

func TestGetMeetupRemoteInvalidUserID(t *testing.T) {
	const sessionName = "session"

	req := httptest.NewRequest(http.MethodGet, new(operations.GetMeetupURL).String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "INVALID_ID"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.NewGetMeetupRemoteParams()
	params.Tags = append(params.Tags, "Mental Health")
	params.HTTPRequest = req

	raw := testImpl.GetMeetupRemote(params)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestGetMeetupID(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDOK)(nil), raw)
	res := raw.(*operations.GetMeetupIDOK)
	assert.Equal(t, fmt.Sprint(res.Payload.ID), TestMeetup.IDString())
}

func TestGetMeetupIDNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = nonexistentUserID
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          nonexistentUserID,
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDNotFound)(nil), raw)
	res := raw.(*operations.GetMeetupIDNotFound)
	assert.Equal(t, res.Payload.Code, int32(404))
}

func TestGetMeetupIDCanceled(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = TestMeetupCanceled.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetupCanceled.IDString(),
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDBadRequest)(nil), raw)
	res := raw.(*operations.GetMeetupIDBadRequest)
	assert.Equal(t, res.Payload.Code, int32(400))
}

func TestGetMeetupIDInvalidUserID(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "TheseAreLetters"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestGetMeetupIDNoUserContext(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDOK)(nil), raw)
	res := raw.(*operations.GetMeetupIDOK)
	assert.Equal(t, fmt.Sprint(res.Payload.ID), TestMeetup.IDString())
}

func TestGetMeetupIDRejectionPopulated(t *testing.T) {
	targetMeetup := createMeetup("Eat", TestUserFriend.ID, []*db.Tag{TestTag}, false)
	// Start state: rejected
	targetMeetup.RejectedAttendees = append(targetMeetup.RejectedAttendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.GetMeetupIDURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
	}

	raw := testImpl.GetMeetupID(params)

	require.IsType(t, (*operations.GetMeetupIDOK)(nil), raw)
	res := raw.(*operations.GetMeetupIDOK)
	assert.Equal(t, fmt.Sprint(res.Payload.ID), targetMeetup.IDString())
	assert.Equal(t, true, res.Payload.Rejected)
}

func TestPostMeetup(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
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
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
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
	assert.Equalf(t, 2, len(res.Payload.Tags), "This meetup was created with two tags.")
}

func TestPostMeetupBadUser(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = nonexistentUserID
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

	assert.IsType(t, responders.InternalServerError{}, raw)
}

func TestPostMeetupInvalidUserId(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupURL)
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "theseareletters"
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

	assert.IsType(t, responders.InternalServerError{}, raw)
}

func TestPatchMeetupID(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
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

func TestPatchMeetupIDCanceled(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = TestMeetupCanceled.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUserFriend.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetupCanceled.IDString(),
		Meetup: &models.MeetupRequestBody{
			Description: newDescription,
			Location: &models.Location{
				Coordinates: &models.Coordinates{
					Lat: TestMeetupCanceled.Location.Coordinates.Lat,
					Lon: TestMeetupCanceled.Location.Coordinates.Lon,
				},
				Name: TestMeetupCanceled.Location.Name,
				URL:  TestMeetupCanceled.Location.URL,
			},
			MaxCapacity: &TestMeetupCanceled.MaxCapacity,
			MinCapacity: &TestMeetupCanceled.MinCapacity,
			Tags:        []string{TestTag.Name},
			Time:        strfmt.DateTime(TestMeetupCanceled.Time),
			Title:       TestMeetupCanceled.Title,
		},
	}
	raw := testImpl.PatchMeetupID(params, nil)
	require.IsType(t, (*operations.PatchMeetupIDBadRequest)(nil), raw)
}

func TestPatchMeetupIDNotFound(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID:          nonexistentMeetupID,
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

func TestPatchMeetupIDForbidden(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = FakeUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
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

func TestPatchMeetupIDInvalidUserID(t *testing.T) {
	const sessionName = "session"
	const newDescription = "UWU"
	url := new(operations.PatchMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "INVALID_ID"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
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
	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestDeleteMeetup(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:       "To Be Deleted",
		Time:        time.Time{},
		Description: "",
		MaxCapacity: 2,
		MinCapacity: 1,
		Owner:       TestUser.ID,
		Location:    db.MeetupLocation{},
		Cancelled:   false,
	}
	testImpl.DB().Create(&doomedMeetup)

	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID:          doomedMeetup.IDString(),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDNoContent)(nil), raw)
}

func TestDeleteMeetupNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID:          nonexistentMeetupID,
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDNotFound)(nil), raw)
}

func TestDeleteMeetupForbidden(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:       "To Be Deleted",
		Time:        time.Time{},
		Description: "",
		MaxCapacity: 2,
		MinCapacity: 1,
		Owner:       TestUser.ID,
		Location:    db.MeetupLocation{},
		Cancelled:   false,
	}
	testImpl.DB().Create(&doomedMeetup)
	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = FakeUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID:          doomedMeetup.IDString(),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDForbidden)(nil), raw)
}

func TestDeleteMeetupDeleted(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:       "To Be Deleted",
		Time:        time.Time{},
		Description: "",
		MaxCapacity: 2,
		MinCapacity: 1,
		Owner:       TestUser.ID,
		Location:    db.MeetupLocation{},
		Cancelled:   false,
	}
	testImpl.DB().Create(&doomedMeetup)

	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID:          doomedMeetup.IDString(),
	}

	raw := testImpl.DeleteMeetupID(params, nil)
	raw = testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, (*operations.DeleteMeetupIDBadRequest)(nil), raw)
}

func TestDeleteMeetupBadUserContext(t *testing.T) {
	// Create a meetup to delete
	doomedMeetup := db.Meetup{
		Title:       "To Be Deleted",
		Time:        time.Time{},
		Description: "",
		MaxCapacity: 2,
		MinCapacity: 1,
		Owner:       TestUser.ID,
		Location:    db.MeetupLocation{},
		Cancelled:   false,
	}
	testImpl.DB().Create(&doomedMeetup)

	const sessionName = "session"
	url := new(operations.DeleteMeetupIDURL)
	url.ID = doomedMeetup.IDString()
	req := httptest.NewRequest(http.MethodDelete, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "TheseAreLetters"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.DeleteMeetupIDParams{
		HTTPRequest: req,
		ID:          doomedMeetup.IDString(),
	}

	raw := testImpl.DeleteMeetupID(params, nil)

	assert.IsType(t, responders.InternalServerError{}, raw)
}

func TestGetMeetupIdAttendee(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDAttendeeURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          TestMeetup.IDString(),
	}
	var numExpectedAttendees int64
	testImpl.DB().Table("meetup_user_attend").Where("meetup_id = ?", TestMeetup.ID).Count(&numExpectedAttendees)

	raw := testImpl.GetMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.GetMeetupIDAttendeeOK)(nil), raw)
	res := raw.(*operations.GetMeetupIDAttendeeOK)
	assert.Equal(t, numExpectedAttendees, int64(len(res.Payload.Attending)))
}

func TestGetMeetupIDAttendeesNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.GetMeetupIDAttendeeURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodGet, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.GetMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          nonexistentMeetupID,
	}
	raw := testImpl.GetMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.GetMeetupIDAttendeeNotFound)(nil), raw)
}

func TestPostMeetupIDAttendee(t *testing.T) {
	newUser := createUser()
	targetMeetup := createMeetup("Eat", newUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PostMeetupIDAttendeeOK)(nil), raw)
	var count int64
	testImpl.DB().Table("meetup_user_pending").Where("meetup_id = ? AND user_id = ?", targetMeetup.ID, TestUser.ID).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestPostMeetupIDAttendeeNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          nonexistentMeetupID,
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PostMeetupIDAttendeeNotFound)(nil), raw)
}

func TestPostMeetupIdAttendeeUserNotFound(t *testing.T) {
	newUser := createUser()
	targetMeetup := createMeetup("Eat", newUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = nonexistentUserID
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, responders.InternalServerError{}, raw)
}

// TestPostMeetupIdAttendeeAlreadyInvolved checks that a BadRequest response is returned if the user is already
// involved in a meetup
func TestPostMeetupIDAttendeeAlreadyInvolved(t *testing.T) {
	rejectedUser := createUser()
	acceptedUser := createUser()
	pendingUser := createUser()
	targetMeetup := createMeetup("Harass Random People", TestUser.ID, []*db.Tag{TestTag}, false)
	targetMeetup.PendingAttendees = append(targetMeetup.PendingAttendees, pendingUser)
	targetMeetup.Attendees = append(targetMeetup.Attendees, acceptedUser)
	targetMeetup.RejectedAttendees = append(targetMeetup.RejectedAttendees, rejectedUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = TestMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)

	for _, user := range []*db.User{pendingUser, rejectedUser, acceptedUser} {
		session.Values[impl.UserID] = fmt.Sprint(user.ID)
		req = req.WithContext(impl.WithSession(req.Context(), session))

		params := operations.PostMeetupIDAttendeeParams{
			HTTPRequest: req,
			ID:          fmt.Sprint(targetMeetup.ID),
		}
		raw := testImpl.PostMeetupIDAttendee(params, nil)

		assert.IsType(t, (*operations.PostMeetupIDAttendeeBadRequest)(nil), raw)
	}
}

func TestPostMeetupIdAttendeeAlreadyOwner(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = ownerUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PostMeetupIDAttendeeBadRequest)(nil), raw)
}

func TestPostMeetupIdAttendeeCanceled(t *testing.T) {
	canceledMeetup := createMeetup("Eat", TestUser.ID, []*db.Tag{TestTag}, true)
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = canceledMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUserFriend.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          canceledMeetup.IDString(),
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PostMeetupIDAttendeeBadRequest)(nil), raw)
}

func TestPostMeetupIdAttendeeInvalidUserID(t *testing.T) {
	newUser := createUser()
	targetMeetup := createMeetup("Eat", newUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PostMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPost, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "INVALID_ID"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PostMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
	}
	raw := testImpl.PostMeetupIDAttendee(params, nil)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestPatchMeetupIdAttendeeAddPending(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	// Start state: attending
	targetMeetup.Attendees = append(targetMeetup.Attendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       "",
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeOK)(nil), raw)
	res := raw.(*operations.PatchMeetupIDAttendeeOK)
	assert.Equal(t, models.AttendeeStatus("pending"), res.Payload)
}

func TestPatchMeetupIDAttendeeApproveUser(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	// Start state: rejected
	targetMeetup.RejectedAttendees = append(targetMeetup.RejectedAttendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = ownerUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       TestUser.IDString(),
			AttendeeStatus: "attending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeOK)(nil), raw)
	res := raw.(*operations.PatchMeetupIDAttendeeOK)
	assert.Equal(t, models.AttendeeStatus("attending"), res.Payload)
}

func TestPatchMeetupIDAttendeeRejectUser(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	// Start state: pending
	targetMeetup.PendingAttendees = append(targetMeetup.PendingAttendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = ownerUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       TestUser.IDString(),
			AttendeeStatus: "rejected",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeOK)(nil), raw)
	res := raw.(*operations.PatchMeetupIDAttendeeOK)
	assert.Equal(t, models.AttendeeStatus("rejected"), res.Payload)
}

func TestPatchMeetupIDAttendeeNotFound(t *testing.T) {
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = nonexistentMeetupID
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          nonexistentMeetupID,
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       "",
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeNotFound)(nil), raw)
}

func TestPatchMeetupIDAttendeeNoPatchOwner(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	targetMeetup.PendingAttendees = append(targetMeetup.PendingAttendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = ownerUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			AttendeeStatus: "attending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeBadRequest)(nil), raw)
}

func TestPatchMeetupIDOnlyOwnerApprove(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	targetMeetup.PendingAttendees = append(targetMeetup.PendingAttendees, TestUser)
	if err := testImpl.DB().Model(&targetMeetup).Updates(&targetMeetup).Error; err != nil {
		t.Fatal("I Couldn't update the test db")
	}
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			AttendeeStatus: "attending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeBadRequest)(nil), raw)
}

func TestPatchMeetupIDCancel(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, true)
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       "",
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeBadRequest)(nil), raw)
}

func TestPatchMeetupIdAttendeeAddPendingInvalidUserID(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = "TheseAreLetters"
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       "",
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestPatchMeetupIdAttendeeAddPendingInvalidAttendeeID(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       "INVALID_ID",
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, responders.InternalServerError{}, raw)
}

func TestPatchMeetupIdAttendeeAddPendingNoOwnerInAttendeeField(t *testing.T) {
	ownerUser := createUser()
	targetMeetup := createMeetup("Eat", ownerUser.ID, []*db.Tag{TestTag}, false)
	const sessionName = "session"
	url := new(operations.PatchMeetupIDAttendeeURL)
	url.ID = targetMeetup.IDString()
	req := httptest.NewRequest(http.MethodPatch, url.String(), nil)
	session, err := testImpl.SessionStore().New(req, sessionName)
	require.NoError(t, err)
	session.Values[impl.UserID] = TestUser.IDString()
	req = req.WithContext(impl.WithSession(req.Context(), session))

	params := operations.PatchMeetupIDAttendeeParams{
		HTTPRequest: req,
		ID:          targetMeetup.IDString(),
		PatchMeetupAttendeeBody: &models.PatchMeetupAttendeeBody{
			Attendee:       ownerUser.IDString(),
			AttendeeStatus: "pending",
		},
	}
	raw := testImpl.PatchMeetupIDAttendee(params, nil)

	require.IsType(t, (*operations.PatchMeetupIDAttendeeBadRequest)(nil), raw)
}
