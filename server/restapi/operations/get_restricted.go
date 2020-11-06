// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRestrictedHandlerFunc turns a function with the right signature into a get restricted handler
type GetRestrictedHandlerFunc func(GetRestrictedParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRestrictedHandlerFunc) Handle(params GetRestrictedParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRestrictedHandler interface for that can handle valid get restricted params
type GetRestrictedHandler interface {
	Handle(GetRestrictedParams, interface{}) middleware.Responder
}

// NewGetRestricted creates a new http.Handler for the get restricted operation
func NewGetRestricted(ctx *middleware.Context, handler GetRestrictedHandler) *GetRestricted {
	return &GetRestricted{Context: ctx, Handler: handler}
}

/*GetRestricted swagger:route GET /restricted getRestricted

Restricted endpoint

This is a sample endpoint that is restricted only to users who are
"logged in".

This is a dummy endpoint for testing purposes. It should be removed soon.


*/
type GetRestricted struct {
	Context *middleware.Context
	Handler GetRestrictedHandler
}

func (o *GetRestricted) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRestrictedParams()

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