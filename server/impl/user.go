package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

func (i *Implementation) GetUserMe(params operations.GetUserMeParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	session := SessionFromContext(ctx)
	idStr := session.Values[UserID]
	if idStr == nil {
		logger.Error("Session has no user ID")
		return InternalServerError{}
	}

	id, err := db.UserIDFromString(idStr.(string))
	if err != nil {
		logger.Error("Session has invalid user ID")
		return InternalServerError{}
	}

	tx := i.DB().WithContext(ctx)

	var dbUser db.User
	if err := tx.First(&dbUser, id).Error; err != nil {
		logger.WithError(err).Error("Unable to find session's ID in DB")
		return InternalServerError{}
	}

	if err := tx.Model(&dbUser).Association("Tags").Find(&dbUser.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return InternalServerError{}
	}

	return operations.NewGetUserMeOK().WithPayload(dbUserToModelUser(&dbUser))
}

func (i *Implementation) GetUserID(params operations.GetUserIDParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	id, err := db.UserIDFromString(params.ID)
	if err != nil {
		return operations.NewGetUserIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified user not found.",
		})
	}

	tx := i.DB().WithContext(ctx)

	var dbUser db.User
	err = tx.First(&dbUser, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewGetUserIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified user not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Could not access user DB")
		return InternalServerError{}
	}

	session := SessionFromContext(params.HTTPRequest.Context())
	if session.Values[UserID] != dbUser.IDString() {
		modelUser := &models.User{
			ID:   models.UserID(dbUser.IDString()),
			Name: dbUser.Name,
		}
		return operations.NewGetUserIDOK().WithPayload(modelUser)
	}

	if err := tx.Model(&dbUser).Association("Tags").Find(&dbUser.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return InternalServerError{}
	}

	return operations.NewGetUserIDOK().WithPayload(dbUserToModelUser(&dbUser))
}

func (i *Implementation) PatchUserID(params operations.PatchUserIDParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	session := SessionFromContext(ctx)
	if id := session.Values[UserID]; id != params.ID {
		logger.Warn("User tried to PATCH someone else")
		return operations.NewPatchUserIDForbidden().WithPayload(&models.Error{
			Code:    http.StatusForbidden,
			Message: "Forbidden.",
		})
	}

	tx := i.DB().WithContext(ctx)
	var dbUser db.User
	err := tx.First(&dbUser, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewPatchUserIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified user not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find user in DB")
		return InternalServerError{}
	}

	if err := i.updateDBUser(ctx, &dbUser, params.UpdatedUser); err != nil {
		logger.WithError(err).Error("Failed to update user")
		return InternalServerError{}
	}

	return operations.NewPatchUserIDOK().WithPayload(dbUserToModelUser(&dbUser))
}

func (i *Implementation) PostUser(params operations.PostUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)
	dbUser := &db.User{
		Name: params.Name,
	}
	if err := tx.Create(dbUser).Error; err != nil {
		logger.WithError(err).Error("Can't create item in DB")
		return InternalServerError{}
	}

	return operations.NewPostUserOK().WithPayload(dbUserToModelUser(dbUser))
}

func (i *Implementation) GetUserLogout(params operations.GetUserLogoutParams) middleware.Responder {
	session := SessionFromContext(params.HTTPRequest.Context())
	session.Options.MaxAge = -1
	return operations.NewGetUserLogoutNoContent()
}

func (i *Implementation) updateDBUser(ctx context.Context, dbUser *db.User, modelUser *models.User) error {
	// Update tags first. Do it through SQL since GORM Association mode Replace
	// doesn't work reliably when the "tags" table has a unique name constraint.
	// https://gorm.io/docs/associations.html#Replace-Associations

	tx := i.DB().WithContext(ctx)

	var placeholders []string
	var variables []interface{}
	for _, tag := range modelUser.Interests {
		placeholders = append(placeholders, "(CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)")
		variables = append(variables, tag)
	}
	variables = append(variables, sql.Named("user_id", dbUser.ID))

	var err error
	if len(modelUser.Interests) > 0 {
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
		JOIN tags t USING (name)
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
`, strings.Join(placeholders, ", ")), variables...).Scan(&dbUser.Tags).Error
	} else {
		err = tx.Model(dbUser).Association("Tags").Clear()
		dbUser.Tags = nil
	}
	if err != nil {
		return err
	}

	// Update other fields later.

	dbUser.Name = modelUser.Name
	dbUser.ContactInfo = modelUser.ContactInfo
	if modelUser.Location != nil {
		dbUser.Location = db.Coordinates{
			Lat: modelUser.Location.Lat,
			Lon: modelUser.Location.Lon,
		}
	} else {
		dbUser.Location = db.Coordinates{}
	}

	return tx.
		Model(dbUser).
		Select("Name", "ContactInfo", "location_lat", "location_lon").
		Updates(dbUser).Error
}

func dbUserToModelUser(dbUser *db.User) *models.User {
	interests := tagsToNames(dbUser.Tags)

	var location *models.Coordinates
	if dbUser.Location.Lat != nil && dbUser.Location.Lon != nil {
		location = &models.Coordinates{
			Lat: dbUser.Location.Lat,
			Lon: dbUser.Location.Lon,
		}
	}

	var connections []string
	if dbUser.FacebookID != nil {
		connections = append(connections, "Facebook")
	}
	if dbUser.GoogleID != nil {
		connections = append(connections, "Google")
	}

	return &models.User{
		ID:              models.UserID(fmt.Sprint(dbUser.ID)),
		Name:            dbUser.Name,
		Email:           dbUser.Email,
		Connections:     connections,
		ContactInfo:     dbUser.ContactInfo,
		Location:        location,
		Interests:       interests,
		Attending:       meetupsToIDs(dbUser.Attending),
		OwnedMeetups:    meetupsToIDs(dbUser.OwnedMeetups),
		PendingApproval: meetupsToIDs(dbUser.PendingApproval),
	}
}

func meetupsToIDs(dbMeetups []*db.Meetup) (ids []models.MeetupID) {
	for _, dbMeetup := range dbMeetups {
		ids = append(ids, models.MeetupID(dbMeetup.IDString()))
	}
	return
}

func tagsToNames(dbTags []*db.Tag) (tags []string) {
	for _, dbTag := range dbTags {
		tags = append(tags, dbTag.Name)
	}
	return
}
