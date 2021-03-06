// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMeetupIDHandlerFunc turns a function with the right signature into a delete meetup ID handler
type DeleteMeetupIDHandlerFunc func(DeleteMeetupIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMeetupIDHandlerFunc) Handle(params DeleteMeetupIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteMeetupIDHandler interface for that can handle valid delete meetup ID params
type DeleteMeetupIDHandler interface {
	Handle(DeleteMeetupIDParams, interface{}) middleware.Responder
}

// NewDeleteMeetupID creates a new http.Handler for the delete meetup ID operation
func NewDeleteMeetupID(ctx *middleware.Context, handler DeleteMeetupIDHandler) *DeleteMeetupID {
	return &DeleteMeetupID{Context: ctx, Handler: handler}
}

/*DeleteMeetupID swagger:route DELETE /meetup/{id} deleteMeetupId

Delete the specified meetup

If the specified meetup does not exist, an error is returned

*/
type DeleteMeetupID struct {
	Context *middleware.Context
	Handler DeleteMeetupIDHandler
}

func (o *DeleteMeetupID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMeetupIDParams()

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
