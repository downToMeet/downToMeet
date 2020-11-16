package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

func (i *Implementation) GetMeetupID(params operations.GetMeetupIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	id, err := db.MeetupIDFromString(params.ID)
	if err != nil {
		return operations.NewGetMeetupIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	}

	tx := i.DB().WithContext(ctx)

	var dbMeetup db.Meetup
	err = tx.First(&dbMeetup, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewGetMeetupIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Could not access meetup DB")
		return InternalServerError{}
	}

	return operations.NewGetMeetupIDOK().WithPayload(dbMeetupToModelMeetup(&dbMeetup))
}

func (i *Implementation) PostMeetup(params operations.PostMeetupParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	session := SessionFromContext(ctx)
	idStr := session.Values[UserID]
	if idStr == nil {
		logger.Error("Session has no user ID")
		return InternalServerError{}
	}

	_, err := db.UserIDFromString(idStr.(string))
	if err != nil {
		logger.Error("Session has invalid user ID")
		return InternalServerError{}
	}

	var dbMeetup db.Meetup
	modelMeetup := modelMeetupRequestBodyToModelMeetup(params.Meetup)
	if err := i.modelMeetupToDBMeetup(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to create db meetup object")
		return InternalServerError{}
	}

	if err := i.createDBMeetup(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to create meetup")
		return InternalServerError{}
	}

	return operations.NewPostMeetupCreated().WithPayload(dbMeetupToModelMeetup(&dbMeetup))
}

func (i *Implementation) createDBMeetup(ctx context.Context, dbMeetup *db.Meetup) error {
	tx := i.DB().WithContext(ctx)
	return tx.Create(dbMeetup).Error
}

func (i *Implementation) modelMeetupToDBMeetup(ctx context.Context, dbMeetup *db.Meetup, modelMeetup *models.Meetup) error {
	// Update tags first. Do it through SQL since GORM Association mode Replace
	// doesn't work reliably when the "tags" table has a unique name constraint.
	// https://gorm.io/docs/associations.html#Replace-Associations

	tx := i.DB().WithContext(ctx)

	var placeholders []string
	var variables []interface{}
	for _, tag := range modelMeetup.Tags {
		placeholders = append(placeholders, "(CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)")
		variables = append(variables, tag)
	}

	// TODO: is this "meetup_id" right?
	variables = append(variables, sql.Named("meetup_id", dbMeetup.ID))

	var err error
	if len(modelMeetup.Tags) > 0 {
		// https://stackoverflow.com/a/42217872/1937836
		err = tx.Raw(fmt.Sprintf(`
WITH
	input_rows(created_at, updated_at, name) AS (VALUES %s),
	ins AS (       -- insert any tags not yet in the DB
		INSERT INTO tags (created_at, updated_at, name)
		SELECT * FROM input_rows
		ON CONFLICT (name) DO NOTHING
		RETURNING created_at, updated_at, deleted_at, id, name
	),
	all_tags AS (  -- get IDs of tags already in the DB
		SELECT * FROM ins
		UNION ALL
		SELECT t.created_at, t.updated_at, t.deleted_at, t.id, t.name FROM input_rows
		JOIN tags t using (name)
	),
	ignored1 AS (  -- delete stale tag-user associations
		DELETE FROM tag_user tu
		WHERE tu.user_id = @user_id AND tu.tag_id NOT IN (SELECT id FROM all_tags)
	),
	ignored2 AS (  -- insert new tag-user associations, if not already exist
		INSERT INTO tag_user (tag_id, user_id)
		SELECT id, @user_id FROM all_tags
		ON CONFLICT (tag_id, user_id) DO NOTHING
	)
SELECT * FROM all_tags
`, strings.Join(placeholders, ", ")), variables...).Scan(&dbMeetup.Tags).Error
	} else {
		err = tx.Model(dbMeetup).Association("Tags").Clear()
		dbMeetup.Tags = nil
	}
	if err != nil {
		return err
	}

	// Update other fields later.
	dbMeetup.Title = modelMeetup.Title
	dbMeetup.Description = modelMeetup.Description
	dbMeetup.MaxCapacity = *modelMeetup.MaxCapacity
	dbMeetup.MinCapacity = *modelMeetup.MinCapacity

	id, _ := strconv.ParseUint(string(modelMeetup.Owner), 10, 64)
	dbMeetup.Owner = uint(id)

	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, modelMeetup.Time)
	dbMeetup.Time = t

	if modelMeetup.Location != nil {
		var coordinates db.Coordinates
		if modelMeetup.Location.Coordinates != nil {
			coordinates = db.Coordinates{
				Lat: modelMeetup.Location.Coordinates.Lat,
				Lon: modelMeetup.Location.Coordinates.Lon,
			}
		}
		dbMeetup.Location = db.MeetupLocation{
			Coordinates: coordinates,
			Name:        modelMeetup.Location.Name,
			URL:         modelMeetup.Location.URL,
		}
	}

	return nil
}

func modelMeetupRequestBodyToModelMeetup(modelMeetupRequestBody *models.MeetupRequestBody) models.Meetup {
	var modelMeetup models.Meetup
	copy(modelMeetup.Tags, modelMeetupRequestBody.Tags)
	modelMeetup.MinCapacity = modelMeetupRequestBody.MinCapacity
	modelMeetup.MaxCapacity = modelMeetupRequestBody.MaxCapacity
	modelMeetup.Time = modelMeetupRequestBody.Time
	modelMeetup.Title = modelMeetupRequestBody.Title
	if modelMeetupRequestBody.Location != nil && modelMeetupRequestBody.Location.Coordinates != nil {
		coordinates := &models.Coordinates{
			Lat: modelMeetupRequestBody.Location.Coordinates.Lat,
			Lon: modelMeetupRequestBody.Location.Coordinates.Lon,
		}
		modelMeetup.Location = &models.Location{
			Coordinates: coordinates,
			Name:        modelMeetupRequestBody.Location.Name,
			URL:         modelMeetupRequestBody.Location.URL,
		}
	}
	return modelMeetup
}

func dbMeetupToModelMeetup(dbMeetup *db.Meetup) *models.Meetup {
	var location *models.Location
	if dbMeetup.Location.Coordinates.Lat != nil && dbMeetup.Location.Coordinates.Lon != nil {
		coordinates := &models.Coordinates{
			Lat: dbMeetup.Location.Coordinates.Lat,
			Lon: dbMeetup.Location.Coordinates.Lon,
		}
		location = &models.Location{
			Coordinates: coordinates,
		}
	}
	if dbMeetup.Location.URL != "" {
		location.URL = dbMeetup.Location.URL
	}
	if dbMeetup.Location.Name != "" {
		location.Name = dbMeetup.Location.Name
	}

	return &models.Meetup{
		ID:               models.MeetupID(fmt.Sprint(dbMeetup.ID)),
		Title:            dbMeetup.Title,
		Location:         location,
		Time:             dbMeetup.Time.String(),
		Description:      dbMeetup.Description,
		Tags:             tagsToNames(dbMeetup.Tags),
		MinCapacity:      &dbMeetup.MinCapacity,
		MaxCapacity:      &dbMeetup.MaxCapacity,
		Owner:            models.UserID(fmt.Sprint(dbMeetup.Owner)),
		PendingAttendees: UsersToIDs(dbMeetup.Attendees),
	}
}

func UsersToIDs(dbUsers []*db.User) (ids []models.UserID) {
	for _, dbUser := range dbUsers {
		ids = append(ids, models.UserID(dbUser.IDString()))
	}
	return
}
