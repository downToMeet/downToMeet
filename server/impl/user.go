package impl

import (
	"github.com/go-openapi/runtime/middleware"

	"go.timothygu.me/downtomeet/server/models"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

func (i *Implementation) GetUserID(params operations.GetUserIDParams) middleware.Responder {
	user := models.User{
		ID: models.UserID(params.ID),
	}
	return operations.NewGetUserIDOK().WithPayload(&user)
}
