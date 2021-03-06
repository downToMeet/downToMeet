// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMeetupRemoteHandlerFunc turns a function with the right signature into a get meetup remote handler
type GetMeetupRemoteHandlerFunc func(GetMeetupRemoteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeetupRemoteHandlerFunc) Handle(params GetMeetupRemoteParams) middleware.Responder {
	return fn(params)
}

// GetMeetupRemoteHandler interface for that can handle valid get meetup remote params
type GetMeetupRemoteHandler interface {
	Handle(GetMeetupRemoteParams) middleware.Responder
}

// NewGetMeetupRemote creates a new http.Handler for the get meetup remote operation
func NewGetMeetupRemote(ctx *middleware.Context, handler GetMeetupRemoteHandler) *GetMeetupRemote {
	return &GetMeetupRemote{Context: ctx, Handler: handler}
}

/*GetMeetupRemote swagger:route GET /meetup/remote getMeetupRemote

Get the list of remote meetups

If the required parameters were not specified correctly, an error is returned

*/
type GetMeetupRemote struct {
	Context *middleware.Context
	Handler GetMeetupRemoteHandler
}

func (o *GetMeetupRemote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMeetupRemoteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
