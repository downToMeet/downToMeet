package impl

import (
	"github.com/go-openapi/runtime/middleware"

	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

func GetMeetupID(params operations.GetMeetupIDParams) middleware.Responder {
	meetup := models.Meetup{
		ID: models.MeetupID(params.ID),
	}
	return operations.NewGetMeetupIDOK().WithPayload(&meetup)
}
