// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMeetupHandlerFunc turns a function with the right signature into a get meetup handler
type GetMeetupHandlerFunc func(GetMeetupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeetupHandlerFunc) Handle(params GetMeetupParams) middleware.Responder {
	return fn(params)
}

// GetMeetupHandler interface for that can handle valid get meetup params
type GetMeetupHandler interface {
	Handle(GetMeetupParams) middleware.Responder
}

// NewGetMeetup creates a new http.Handler for the get meetup operation
func NewGetMeetup(ctx *middleware.Context, handler GetMeetupHandler) *GetMeetup {
	return &GetMeetup{Context: ctx, Handler: handler}
}

/*GetMeetup swagger:route GET /meetup getMeetup

Get the list of in-person meetups

If the required parameters were not specified correctly, an error is returned

*/
type GetMeetup struct {
	Context *middleware.Context
	Handler GetMeetupHandler
}

func (o *GetMeetup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMeetupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
