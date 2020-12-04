package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/impl/responders"
	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

/*
Before running the query in GetMeetup, do

CREATE EXTENSION earthdistance CASCADE;

as a superuser (probably 'postgres')
*/

// GetMeetup implements the GET /meetup endpoint.
func (i *Implementation) GetMeetup(params operations.GetMeetupParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)

	var tagIds []int
	var dbTags []db.Tag
	if err := tx.Where("name IN ?", params.Tags).Find(&dbTags).Error; err != nil {
		logger.WithError(err).Error("Unable to fetch tags")
		return responders.InternalServerError{}
	}
	for _, tag := range dbTags {
		tagIds = append(tagIds, int(tag.ID))
	}

	var meetups []*db.Meetup
	var err error
	if len(params.Tags) == 0 {
		err = tx.Raw(`
			SELECT * FROM (
				SELECT *,
					earth_distance(ll_to_earth(?, ?),
					ll_to_earth(location_lat, location_lon)) AS distance_from_me
				FROM meetups
			) AS m
			WHERE m.distance_from_me < ? AND m.time >= CURRENT_TIMESTAMP
			ORDER BY m.distance_from_me
			LIMIT 100
		`, params.Lat, params.Lon, params.Radius*1000).Scan(&meetups).Error
	} else {
		var tagIds []uint
		var dbTags []db.Tag
		if err := tx.Where("name IN ?", params.Tags).Find(&dbTags).Error; err != nil {
			logger.WithError(err).Error("Unable to fetch tags")
			return responders.InternalServerError{}
		}
		for _, tag := range dbTags {
			tagIds = append(tagIds, tag.ID)
		}

		err = tx.Raw(`
			SELECT * FROM (
				SELECT DISTINCT ON (m.id) *
				FROM (
					SELECT *,
						earth_distance(ll_to_earth(?, ?),
						ll_to_earth(location_lat, location_lon)) AS distance_from_me
					FROM meetups
				) AS m
				JOIN meetup_tag AS mt ON m.id = mt.meetup_id
				WHERE m.distance_from_me < ? AND mt.tag_id IN ? AND m.time >= CURRENT_TIMESTAMP
				ORDER BY m.id
				LIMIT 100
			) AS m
			ORDER BY m.distance_from_me
		`, params.Lat, params.Lon, params.Radius*1000, tagIds).Scan(&meetups).Error
	}

	if err != nil {
		logger.WithError(err).Error("Unable to find meetups that fit the given parameters")
		return operations.NewGetMeetupOK().WithPayload([]*models.Meetup{})
	}

	var idStr string
	if id := SessionFromContext(ctx).Values[UserID]; id != nil {
		if _, err := db.UserIDFromString(id.(string)); err != nil {
			logger.Error("Session has invalid user ID")
			return responders.InternalServerError{}
		}
		idStr = id.(string)
	}

	// Convert each returned dbMeetup to a models.Meetup
	var modelMeetups []*models.Meetup
	for _, meetupID := range meetups {
		var meetup db.Meetup
		err := tx.First(&meetup, meetupID).Error
		if err != nil {
			logger.WithError(err).Error("Unable to get meetup from ID")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("Attendees").Find(&meetup.Attendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("PendingAttendees").Find(&meetup.PendingAttendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup pending attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("RejectedAttendees").Find(&meetup.RejectedAttendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup rejected attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("Tags").Find(&meetup.Tags); err != nil {
			logger.WithError(err).Error("Unable to find user tags")
			return responders.InternalServerError{}
		}
		modelMeetups = append(modelMeetups, dbMeetupToModelMeetup(&meetup, idStr))
	}
	return operations.NewGetMeetupOK().WithPayload(modelMeetups)
}

func (i *Implementation) GetMeetupRemote(params operations.GetMeetupRemoteParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)

	var meetups []*db.Meetup
	var err error
	if len(params.Tags) == 0 {
		err = tx.Raw(`
			SELECT *
			FROM meetups AS m
			WHERE m.location_lat IS NULL AND m.time >= CURRENT_TIMESTAMP
			ORDER BY m.time
			LIMIT 100
		`).Scan(&meetups).Error
	} else {
		var tagIds []uint
		var dbTags []db.Tag
		if err := tx.Where("name IN ?", params.Tags).Find(&dbTags).Error; err != nil {
			logger.WithError(err).Error("Unable to fetch tags")
			return responders.InternalServerError{}
		}
		for _, tag := range dbTags {
			tagIds = append(tagIds, tag.ID)
		}

		err = tx.Raw(`
			SELECT * FROM (
				SELECT DISTINCT ON (m.id) *
				FROM meetups AS m
				JOIN meetup_tag AS mt ON m.id = mt.meetup_id
				WHERE m.location_lat IS NULL AND mt.tag_id IN ? AND m.time >= CURRENT_TIMESTAMP
				ORDER BY m.id
				LIMIT 100
			) AS m
			ORDER BY m.time
		`, tagIds).Scan(&meetups).Error
	}

	if err != nil {
		logger.WithError(err).Error("Unable to find meetups that fit the given parameters")
		return operations.NewGetMeetupRemoteOK().WithPayload([]*models.Meetup{})
	}

	var idStr string
	if id := SessionFromContext(ctx).Values[UserID]; id != nil {
		if _, err := db.UserIDFromString(id.(string)); err != nil {
			logger.Error("Session has invalid user ID")
			return responders.InternalServerError{}
		}
		idStr = id.(string)
	}

	// Convert each returned dbMeetup to a models.Meetup
	var modelMeetups []*models.Meetup
	for _, meetupID := range meetups {
		var meetup db.Meetup
		err := tx.First(&meetup, meetupID).Error
		if err != nil {
			logger.WithError(err).Error("Unable to get meetup from ID")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("Attendees").Find(&meetup.Attendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("PendingAttendees").Find(&meetup.PendingAttendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup pending attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("RejectedAttendees").Find(&meetup.RejectedAttendees); err != nil {
			logger.WithError(err).Error("Unable to find meetup rejected attendee information")
			return responders.InternalServerError{}
		}

		if err = tx.Model(&meetup).Association("Tags").Find(&meetup.Tags); err != nil {
			logger.WithError(err).Error("Unable to find user tags")
			return responders.InternalServerError{}
		}
		modelMeetups = append(modelMeetups, dbMeetupToModelMeetup(&meetup, idStr))
	}
	return operations.NewGetMeetupRemoteOK().WithPayload(modelMeetups)
}

// GetMeetupID implements the GET /meetup/:id endpoint.
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
		return responders.InternalServerError{}
	}

	if err = tx.Model(&dbMeetup).Association("Tags").Find(&dbMeetup.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return responders.InternalServerError{}
	}

	if err = tx.Model(&dbMeetup).Association("Attendees").Find(&dbMeetup.Attendees); err != nil {
		logger.WithError(err).Error("Unable to find meetup attendee information")
		return responders.InternalServerError{}
	}

	if err = tx.Model(&dbMeetup).Association("PendingAttendees").Find(&dbMeetup.PendingAttendees); err != nil {
		logger.WithError(err).Error("Unable to find meetup pending attendee information")
		return responders.InternalServerError{}
	}

	if err = tx.Model(&dbMeetup).Association("RejectedAttendees").Find(&dbMeetup.RejectedAttendees); err != nil {
		logger.WithError(err).Error("Unable to find meetup rejected attendee information")
		return responders.InternalServerError{}
	}

	var idStr string
	if id := SessionFromContext(ctx).Values[UserID]; id == nil {
		idStr = ""
	} else {
		if _, err := db.UserIDFromString(id.(string)); err != nil {
			logger.Error("Session has invalid user ID")
			return responders.InternalServerError{}
		}
		idStr = id.(string)
	}

	return operations.NewGetMeetupIDOK().WithPayload(dbMeetupToModelMeetup(&dbMeetup, idStr))
}

// PostMeetup implements the POST /meetup endpoint.
func (i *Implementation) PostMeetup(params operations.PostMeetupParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	var dbMeetup db.Meetup
	id := SessionFromContext(ctx).Values[UserID]
	if _, err := db.UserIDFromString(id.(string)); err != nil {
		logger.Error("Session has invalid user ID")
		return responders.InternalServerError{}
	}

	modelMeetup := modelMeetupRequestBodyToModelMeetup(params.Meetup, id.(string))
	if err := i.modelMeetupToDBMeetup(&dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to create db meetup object")
		return responders.InternalServerError{}
	}

	tx := i.DB().WithContext(ctx)
	if err := tx.Create(&dbMeetup).Error; err != nil {
		logger.WithError(err).Error("Failed to create meetup")
		return responders.InternalServerError{}
	}

	if err := i.insertMeetupTagsIntoDB(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to insert meetup tags")
		return responders.InternalServerError{}
	}

	return operations.NewPostMeetupCreated().WithPayload(dbMeetupToModelMeetup(&dbMeetup, id.(string)))
}

// PatchMeetupID implements the PATCH /meetup/:id endpoint.
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
		return responders.InternalServerError{}
	}

	session := SessionFromContext(ctx)
	userID := session.Values[UserID].(string)
	if _, err := db.UserIDFromString(userID); err != nil {
		logger.Error("Session has invalid user ID")
		return responders.InternalServerError{}
	}
	if userID != fmt.Sprint(dbMeetup.Owner) {
		logger.Warn("User tried to PATCH a meetup they do not own")
		return operations.NewPatchMeetupIDForbidden().WithPayload(&models.Error{
			Code:    http.StatusForbidden,
			Message: "Forbidden",
		})
	}

	if dbMeetup.Cancelled == true {
		return operations.NewPatchMeetupIDBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "This meetup has been cancelled",
		})
	}

	modelMeetup := modelMeetupRequestBodyToModelMeetup(params.Meetup, userID)
	if err := i.modelMeetupToDBMeetup(&dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to create db meetup object")
		return responders.InternalServerError{}
	}

	if err := i.insertMeetupTagsIntoDB(ctx, &dbMeetup, &modelMeetup); err != nil {
		logger.WithError(err).Error("Failed to insert meetup tags")
		return responders.InternalServerError{}
	}

	if err := i.updateDBMeetup(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to update meetup")
		return responders.InternalServerError{}
	}
	return operations.NewPatchMeetupIDOK().WithPayload(dbMeetupToModelMeetup(&dbMeetup, userID))
}

// DeleteMeetupID implements the DELETE /meetup/:id endpoint.
func (i *Implementation) DeleteMeetupID(params operations.DeleteMeetupIDParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)

	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewDeleteMeetupIDNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find meetup in DB")
		return responders.InternalServerError{}
	}

	session := SessionFromContext(ctx)
	id := session.Values[UserID]
	if _, err := db.UserIDFromString(id.(string)); err != nil {
		logger.Error("Session has invalid user ID")
		return responders.InternalServerError{}
	}
	if id.(string) != fmt.Sprint(dbMeetup.Owner) {
		logger.Warn("User tried to DELETE a meetup they do not own")
		return operations.NewDeleteMeetupIDForbidden().WithPayload(&models.Error{
			Code:    http.StatusForbidden,
			Message: "Forbidden",
		})
	}

	if dbMeetup.Cancelled == true {
		return operations.NewDeleteMeetupIDBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "This meetup has already been cancelled",
		})
	}

	dbMeetup.Cancelled = true

	if err := i.updateDBMeetup(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to update meetup")
		return responders.InternalServerError{}
	}
	return operations.NewDeleteMeetupIDNoContent()
}

// GetMeetupIdAttendee implements the GET /meetup/:id/attendee endpoint.
func (i *Implementation) GetMeetupIdAttendee(params operations.GetMeetupIDAttendeeParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)
	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewGetMeetupIDAttendeeNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find meetup in DB")
		return responders.InternalServerError{}
	}

	if err = i.fetchAllAttendeeInformationLists(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to fetch attendee information lists")
		return responders.InternalServerError{}
	}

	var attendeeList models.AttendeeList
	if dbMeetup.Attendees != nil {
		var attending []models.UserID
		for _, attendee := range dbMeetup.Attendees {
			attending = append(attending, models.UserID(fmt.Sprint(attendee.ID)))
		}
		attendeeList.Attending = attending
	}
	if dbMeetup.PendingAttendees != nil {
		var pending []models.UserID
		for _, attendee := range dbMeetup.PendingAttendees {
			pending = append(pending, models.UserID(fmt.Sprint(attendee.ID)))
		}
		attendeeList.Pending = pending
	}
	return operations.NewGetMeetupIDAttendeeOK().WithPayload(&attendeeList)
}

// PostMeetupIdAttendee implements the POST /meetup/:id/attendee endpoint.
func (i *Implementation) PostMeetupIdAttendee(params operations.PostMeetupIDAttendeeParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)
	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewPostMeetupIDAttendeeNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find meetup in DB")
		return responders.InternalServerError{}
	}

	session := SessionFromContext(ctx)
	id := session.Values[UserID]
	if _, err := db.UserIDFromString(id.(string)); err != nil {
		logger.Error("Session has invalid user ID")
		return responders.InternalServerError{}
	}

	if id.(string) == fmt.Sprint(dbMeetup.Owner) {
		return operations.NewPostMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "User is the owner of this meetup",
		})
	}

	if dbMeetup.Cancelled == true {
		return operations.NewPostMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "This meetup has been cancelled",
		})
	}

	if err = i.fetchAllAttendeeInformationLists(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to fetch attendee information lists")
		return responders.InternalServerError{}
	}

	// Make sure that the user is not already in one of the lists
	idStr := id.(string)
	if idStr != "" {
		if dbMeetup.Attendees != nil {
			for _, attendee := range dbMeetup.Attendees {
				if attendee.IDString() == idStr {
					return operations.NewPostMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
						Code:    http.StatusBadRequest,
						Message: "User is already attending this meetup",
					})
				}
			}
		}
		if dbMeetup.PendingAttendees != nil {
			for _, attendee := range dbMeetup.PendingAttendees {
				if attendee.IDString() == idStr {
					return operations.NewPostMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
						Code:    http.StatusBadRequest,
						Message: "User is already pending approval to attend this meetup",
					})
				}
			}
		}
		if dbMeetup.RejectedAttendees != nil {
			for _, attendee := range dbMeetup.RejectedAttendees {
				if attendee.IDString() == idStr {
					return operations.NewPostMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
						Code:    http.StatusBadRequest,
						Message: "User was rejected from this meetup",
					})
				}
			}
		}
	}

	// Fetch the user
	var dbUser db.User
	userID, _ := db.UserIDFromString(idStr)
	if err = tx.First(&dbUser, userID).Error; err != nil {
		logger.WithError(err).Error("Unable to find current user in DB")
		return responders.InternalServerError{}
	}
	if err := tx.Model(&dbUser).Association("Tags").Find(&dbUser.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return responders.InternalServerError{}
	}
	dbMeetup.PendingAttendees = append(dbMeetup.PendingAttendees, &dbUser)

	// Update the db meetup
	if err = tx.Model(&dbMeetup).Updates(&dbMeetup).Error; err != nil {
		logger.WithError(err).Error("Unable to update DB meetup")
		return responders.InternalServerError{}
	}

	var attendeeStatus models.AttendeeStatus = "pending"
	return operations.NewPostMeetupIDAttendeeOK().WithPayload(attendeeStatus)
}

// PatchMeetupIdAttendee implements the PATCH /meetup/:id/attendee endpoint.
func (i *Implementation) PatchMeetupIdAttendee(params operations.PatchMeetupIDAttendeeParams, _ interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	logger := log.WithContext(ctx)

	tx := i.DB().WithContext(ctx)
	var dbMeetup db.Meetup
	err := tx.First(&dbMeetup, params.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operations.NewPatchMeetupIDAttendeeNotFound().WithPayload(&models.Error{
			Code:    http.StatusNotFound,
			Message: "Specified meetup not found.",
		})
	} else if err != nil {
		logger.WithError(err).Error("Failed to find meetup in DB")
		return responders.InternalServerError{}
	}

	session := SessionFromContext(ctx)
	id := session.Values[UserID]
	attendeeId, err := db.UserIDFromString(params.PatchMeetupAttendeeBody.Attendee)
	status := params.PatchMeetupAttendeeBody.AttendeeStatus
	if attendeeId != 0 && err != nil {
		logger.Error("Trying to patch an invalid user ID")
		return responders.InternalServerError{}
	}

	if _, err = db.UserIDFromString(id.(string)); err != nil {
		logger.Error("Session has invalid user ID")
		return responders.InternalServerError{}
	}

	if params.PatchMeetupAttendeeBody.Attendee == id.(string) {
		return operations.NewPatchMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "To patch current user, omit attendee field",
		})
	}

	if id.(string) == fmt.Sprint(dbMeetup.Owner) && (attendeeId == 0) {
		return operations.NewPatchMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "User is the owner of this meetup",
		})
	}

	if (status == "attending" || status == "rejected") && id.(string) != fmt.Sprint(dbMeetup.Owner) {
		return operations.NewPatchMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "Only the owner of this meetup can approve or reject an attendee",
		})
	}

	if (status == "none" || status == "pending") && attendeeId != 0 {
		return operations.NewPatchMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "A user can only change their own attendee status to either none or pending",
		})
	}

	if dbMeetup.Cancelled == true {
		return operations.NewPatchMeetupIDAttendeeBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "This meetup has been cancelled",
		})
	}

	if err = i.fetchAllAttendeeInformationLists(ctx, &dbMeetup); err != nil {
		logger.WithError(err).Error("Failed to fetch attendee information lists")
		return responders.InternalServerError{}
	}

	// Remove the user from whichever list they are currently in
	var idStr string
	if attendeeId != 0 {
		idStr = params.PatchMeetupAttendeeBody.Attendee
	} else {
		idStr = id.(string)
	}
	found := false

	// Fetch the user
	var dbUser db.User
	userID, _ := db.UserIDFromString(idStr)
	if err = tx.First(&dbUser, userID).Error; err != nil {
		logger.WithError(err).Error("Unable to find current user in DB")
		return responders.InternalServerError{}
	}
	if err := tx.Model(&dbUser).Association("Tags").Find(&dbUser.Tags); err != nil {
		logger.WithError(err).Error("Unable to find user tags")
		return responders.InternalServerError{}
	}

	if dbMeetup.Attendees != nil {
		for _, attendee := range dbMeetup.Attendees {
			if attendee.IDString() == idStr {
				// Remove the user from this array
				found = true
				err := tx.Model(&dbMeetup).Association("Attendees").Delete(&dbUser)
				if err != nil {
					logger.Error("Unable to remove attendee from attendee list")
					return responders.InternalServerError{}
				}
				break
			}
		}
	}
	if dbMeetup.PendingAttendees != nil && !found {
		for _, attendee := range dbMeetup.PendingAttendees {
			if attendee.IDString() == idStr {
				// Remove the user from this array
				found = true
				err := tx.Model(&dbMeetup).Association("PendingAttendees").Delete(&dbUser)
				if err != nil {
					logger.Error("Unable to remove attendee from pending attendee list")
					return responders.InternalServerError{}
				}
				break
			}
		}
	}
	if dbMeetup.RejectedAttendees != nil && !found {
		for _, attendee := range dbMeetup.RejectedAttendees {
			if attendee.IDString() == idStr {
				// Remove the user from this array
				found = true
				err := tx.Model(&dbMeetup).Association("RejectedAttendees").Delete(&dbUser)
				if err != nil {
					logger.Error("Unable to remove attendee from rejected attendee list")
					return responders.InternalServerError{}
				}
				break
			}
		}
	}

	// Add user to the appropriate array
	// If attendeeStatus is "none", we don't add them anywhere
	if status == "pending" {
		dbMeetup.PendingAttendees = append(dbMeetup.PendingAttendees, &dbUser)
	} else if status == "attending" {
		dbMeetup.Attendees = append(dbMeetup.Attendees, &dbUser)
	} else if status == "rejected" {
		dbMeetup.RejectedAttendees = append(dbMeetup.RejectedAttendees, &dbUser)
	}

	// Update the db meetup
	if err = tx.Model(&dbMeetup).Select("Attendees", "PendingAttendees", "RejectedAttendees").Updates(&dbMeetup).Error; err != nil {
		logger.WithError(err).Error("Unable to update DB meetup")
		return responders.InternalServerError{}
	}

	return operations.NewPatchMeetupIDAttendeeOK().WithPayload(status)
}

func (i *Implementation) fetchAllAttendeeInformationLists(ctx context.Context, dbMeetup *db.Meetup) error {
	tx := i.DB().WithContext(ctx)

	if err := tx.Model(&dbMeetup).Association("Attendees").Find(&dbMeetup.Attendees); err != nil {
		return err
	}
	if err := tx.Model(&dbMeetup).Association("PendingAttendees").Find(&dbMeetup.PendingAttendees); err != nil {
		return err
	}
	if err := tx.Model(&dbMeetup).Association("RejectedAttendees").Find(&dbMeetup.RejectedAttendees); err != nil {
		return err
	}

	return nil
}

func (i *Implementation) updateDBMeetup(ctx context.Context, dbMeetup *db.Meetup) error {
	tx := i.DB().WithContext(ctx)
	return tx.Model(dbMeetup).
		Select("title, time, description, max_capacity, min_capacity, location_lat, location_lon, location_url, location_name").
		Updates(dbMeetup).Error
}

func (i *Implementation) insertMeetupTagsIntoDB(ctx context.Context, dbMeetup *db.Meetup, modelMeetup *models.Meetup) error {
	// Update tags. Do it through SQL since GORM Association mode Replace
	// doesn't work reliably when the "tags" table has a unique name constraint.
	// https://gorm.io/docs/associations.html#Replace-Associations
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

func (i *Implementation) modelMeetupToDBMeetup(dbMeetup *db.Meetup, modelMeetup *models.Meetup) error {
	dbMeetup.Title = modelMeetup.Title
	dbMeetup.Description = modelMeetup.Description
	dbMeetup.MaxCapacity = swag.Int64Value(modelMeetup.MaxCapacity)
	dbMeetup.MinCapacity = swag.Int64Value(modelMeetup.MinCapacity)
	dbMeetup.Description = modelMeetup.Description
	dbMeetup.Owner, _ = db.UserIDFromString(string(modelMeetup.Owner))
	dbMeetup.Time = time.Time(modelMeetup.Time)

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
	if modelMeetupRequestBody.Location != nil {
		if modelMeetupRequestBody.Location.Coordinates != nil {
			coordinates := &models.Coordinates{
				Lat: modelMeetupRequestBody.Location.Coordinates.Lat,
				Lon: modelMeetupRequestBody.Location.Coordinates.Lon,
			}
			modelMeetup.Location = &models.Location{
				Coordinates: coordinates,
				Name:        modelMeetupRequestBody.Location.Name,
				URL:         modelMeetupRequestBody.Location.URL,
			}
		} else {
			modelMeetup.Location = &models.Location{
				URL: modelMeetupRequestBody.Location.URL,
			}
		}
	}
	return modelMeetup
}

func dbMeetupToModelMeetup(dbMeetup *db.Meetup, userID string) *models.Meetup {
	location := &models.Location{}
	if dbMeetup.Location.Coordinates.Lat != nil && dbMeetup.Location.Coordinates.Lon != nil {
		location.Coordinates = &models.Coordinates{
			Lat: dbMeetup.Location.Coordinates.Lat,
			Lon: dbMeetup.Location.Coordinates.Lon,
		}
	}
	if dbMeetup.Location.URL != "" {
		location.URL = dbMeetup.Location.URL
	}
	if dbMeetup.Location.Name != "" {
		location.Name = dbMeetup.Location.Name
	}

	// Determine if user was rejected or not
	rejected := false
	if dbMeetup.RejectedAttendees != nil && userID != "" {
		for _, attendee := range dbMeetup.RejectedAttendees {
			if attendee.IDString() == userID {
				rejected = true
				break
			}
		}
	}

	return &models.Meetup{
		ID:               models.MeetupID(dbMeetup.IDString()),
		Title:            dbMeetup.Title,
		Location:         location,
		Time:             strfmt.DateTime(dbMeetup.Time),
		Description:      dbMeetup.Description,
		Tags:             tagsToNames(dbMeetup.Tags),
		MinCapacity:      &dbMeetup.MinCapacity,
		MaxCapacity:      &dbMeetup.MaxCapacity,
		Owner:            models.UserID(fmt.Sprint(dbMeetup.Owner)),
		Attendees:        usersToIDs(dbMeetup.Attendees),
		PendingAttendees: usersToIDs(dbMeetup.PendingAttendees),
		Rejected:         rejected,
		Canceled:         dbMeetup.Cancelled,
	}
}

func usersToIDs(dbUsers []*db.User) (ids []models.UserID) {
	for _, dbUser := range dbUsers {
		ids = append(ids, models.UserID(dbUser.IDString()))
	}
	return
}
