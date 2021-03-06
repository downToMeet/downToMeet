// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostMeetupIDAttendeeHandlerFunc turns a function with the right signature into a post meetup ID attendee handler
type PostMeetupIDAttendeeHandlerFunc func(PostMeetupIDAttendeeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMeetupIDAttendeeHandlerFunc) Handle(params PostMeetupIDAttendeeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostMeetupIDAttendeeHandler interface for that can handle valid post meetup ID attendee params
type PostMeetupIDAttendeeHandler interface {
	Handle(PostMeetupIDAttendeeParams, interface{}) middleware.Responder
}

// NewPostMeetupIDAttendee creates a new http.Handler for the post meetup ID attendee operation
func NewPostMeetupIDAttendee(ctx *middleware.Context, handler PostMeetupIDAttendeeHandler) *PostMeetupIDAttendee {
	return &PostMeetupIDAttendee{Context: ctx, Handler: handler}
}

/*PostMeetupIDAttendee swagger:route POST /meetup/{id}/attendee postMeetupIdAttendee

Post the current user's attendee status (to "pending") for the specified meetup

If the specified meetup does not exist, an error is returned

*/
type PostMeetupIDAttendee struct {
	Context *middleware.Context
	Handler PostMeetupIDAttendeeHandler
}

func (o *PostMeetupIDAttendee) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMeetupIDAttendeeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
