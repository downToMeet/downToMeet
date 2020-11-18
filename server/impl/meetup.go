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

	tx := i.DB().WithContext(ctx)

	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewGetMeetupIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Could not access meetup DB")
		return InternalServerError{}
	}

	id := SessionFromContext(ctx).Values[UserID]
	if id != nil {
		// If logged in, include info about whether or not the user was rejected
		if err = tx.Model(&dbMeetup).Association("Attendees").Find(&dbMeetup.Attendees); err != nil {
			logger.WithError(err).Error("Unable to determine whether user was rejected from event")
			return InternalServerError{}
		}
	}

	if err = tx.Model(&dbMeetup).Association("Tags").Find(&dbMeetup.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return InternalServerError{}
	}

	return operations.NewGetMeetupIDOK().WithPayload(dbMeetupToModelMeetup(&dbMeetup, id.(string)))
}

func (i *Implementation) PostMeetup(params operations.PostMeetupParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	var dbMeetup db.Meetup
	id := SessionFromContext(ctx).Values[UserID]
	modelMeetup := modelMeetupRequestBodyToModelMeetup(params.Meetup, id.(string))
	if err := i.modelMeetupToDBMeetup(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to create db meetup object")
		return InternalServerError{}
	}

	if err := i.createDBMeetup(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to create meetup")
		return InternalServerError{}
	}

	if err := i.insertTagsIntoDB(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to insert meetup tags")
		return InternalServerError{}
	}

	return operations.NewPostMeetupCreated().WithPayload(dbMeetupToModelMeetup(&dbMeetup, id.(string)))
}

func (i *Implementation) PatchMeetupID(params operations.PatchMeetupIDParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)
	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewPatchMeetupIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find meetup in DB")
		return InternalServerError{}
	}

	session := SessionFromContext(ctx)
	id := session.Values[UserID]
	if id.(string) != fmt.Sprint(dbMeetup.Owner) {
		logger.Warn("User tried to PATCH an event they do not own")
		return operations.NewPatchMeetupIDForbidden().WithPayload(&models.Error{
			Code:    http.StatusForbidden,
			Message: "Forbidden",
		})
	}

	modelMeetup := modelMeetupRequestBodyToModelMeetup(params.Meetup, id.(string))
	if err := i.modelMeetupToDBMeetup(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to create db meetup object")
		return InternalServerError{}
	}

	if err := i.insertTagsIntoDB(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to insert meetup tags")
		return InternalServerError{}
	}

	if err := i.updateDBMeetup(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to update meetup")
		return InternalServerError{}
	}

	return operations.NewPatchMeetupIDOK().WithPayload(dbMeetupToModelMeetup(&dbMeetup, id.(string)))

}

func (i *Implementation) updateDBMeetup(ctx context.Context, dbMeetup *db.Meetup) error {
	tx := i.DB().WithContext(ctx)
	return tx.Model(dbMeetup).
		Select("title, time, description, max_capacity, min_capacity, location_lat, location_lon, location_url, location_name").
		Updates(dbMeetup).Error
}

func (i *Implementation) createDBMeetup(ctx context.Context, dbMeetup *db.Meetup) error {
	tx := i.DB().WithContext(ctx)
	return tx.Create(dbMeetup).Error
}

func (i *Implementation) insertTagsIntoDB(ctx context.Context, dbMeetup *db.Meetup, modelMeetup *models.Meetup) error {
	// Insert tags into the DB
	tx := i.DB().WithContext(ctx)

	var placeholders []string
	var variables []interface{}
	for _, tag := range modelMeetup.Tags {
		placeholders = append(placeholders, "(CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)")
		variables = append(variables, tag)
	}

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
		ignored1 AS (  -- delete stale tag-meetup associations
			DELETE FROM meetup_tag tm
			WHERE tm.meetup_id = @meetup_id AND tm.tag_id NOT IN (SELECT id FROM all_tags)
		),
		ignored2 AS (  -- insert new tag-meetup associations, if not already exist
			INSERT INTO meetup_tag (tag_id, meetup_id)
			SELECT id, @meetup_id FROM all_tags
			ON CONFLICT (tag_id, meetup_id) DO NOTHING
		)
	SELECT * FROM all_tags
	`, strings.Join(placeholders, ", ")), variables...).Scan(&dbMeetup.Tags).Error
	} else {
		err = tx.Model(dbMeetup).Association("Tags").Clear()
		dbMeetup.Tags = nil
	}
	return err
}

func (i *Implementation) modelMeetupToDBMeetup(ctx context.Context, dbMeetup *db.Meetup, modelMeetup *models.Meetup) error {
	// Update tags first. Do it through SQL since GORM Association mode Replace
	// doesn't work reliably when the "tags" table has a unique name constraint.
	// https://gorm.io/docs/associations.html#Replace-Associations
	dbMeetup.Title = modelMeetup.Title
	dbMeetup.Description = modelMeetup.Description
	dbMeetup.MaxCapacity = *modelMeetup.MaxCapacity
	dbMeetup.MinCapacity = *modelMeetup.MinCapacity
	dbMeetup.Description = modelMeetup.Description

	id, _ := strconv.ParseUint(string(modelMeetup.Owner), 10, 64)
	dbMeetup.Owner = uint(id)

	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, modelMeetup.Time)
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

func modelMeetupRequestBodyToModelMeetup(modelMeetupRequestBody *models.MeetupRequestBody, id string) models.Meetup {
	var modelMeetup models.Meetup
	modelMeetup.Tags = modelMeetupRequestBody.Tags
	modelMeetup.Owner = models.UserID(id)
	modelMeetup.MinCapacity = modelMeetupRequestBody.MinCapacity
	modelMeetup.MaxCapacity = modelMeetupRequestBody.MaxCapacity
	modelMeetup.Time = modelMeetupRequestBody.Time
	modelMeetup.Title = modelMeetupRequestBody.Title
	modelMeetup.Description = modelMeetupRequestBody.Description
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

func dbMeetupToModelMeetup(dbMeetup *db.Meetup, id string) *models.Meetup {
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

	// Determine if user was rejected or not
	var rejected bool
	if dbMeetup.Attendees != nil && id != "" {
		for _, i := range UsersToIDs(dbMeetup.Attendees) {
			if string(i) == id {
				rejected = true
				break
			}
		}
		rejected = false
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
		Attendees: UsersToIDs(dbMeetup.Attendees),
		PendingAttendees: UsersToIDs(dbMeetup.PendingAttendees),
		Rejected: rejected,
	}
}

func UsersToIDs(dbUsers []*db.User) (ids []models.UserID) {
	for _, dbUser := range dbUsers {
		ids = append(ids, models.UserID(dbUser.IDString()))
	}
	return
}
