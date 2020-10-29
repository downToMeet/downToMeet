package impl

import (
	"fmt"
	"sync/atomic"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"go.timothygu.me/downtomeet/models"
	"go.timothygu.me/downtomeet/restapi/operations"
)

var times int32

func GetHello(params operations.GetHelloParams) middleware.Responder {
	if swag.StringValue(params.ID) == "error" {
		return operations.NewGetHelloDefault(400).
			WithPayload(&models.Error{
				Message: "ID is \"error\"",
			})
	}

	return operations.NewGetHelloOK().
		WithPayload(&operations.GetHelloOKBody{
			Hello: fmt.Sprintf("world %d", atomic.AddInt32(&times, 1)),
		})
}
