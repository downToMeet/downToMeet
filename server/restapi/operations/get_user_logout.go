// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserLogoutHandlerFunc turns a function with the right signature into a get user logout handler
type GetUserLogoutHandlerFunc func(GetUserLogoutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserLogoutHandlerFunc) Handle(params GetUserLogoutParams) middleware.Responder {
	return fn(params)
}

// GetUserLogoutHandler interface for that can handle valid get user logout params
type GetUserLogoutHandler interface {
	Handle(GetUserLogoutParams) middleware.Responder
}

// NewGetUserLogout creates a new http.Handler for the get user logout operation
func NewGetUserLogout(ctx *middleware.Context, handler GetUserLogoutHandler) *GetUserLogout {
	return &GetUserLogout{Context: ctx, Handler: handler}
}

/*GetUserLogout swagger:route GET /user/logout getUserLogout

Log out the user

*/
type GetUserLogout struct {
	Context *middleware.Context
	Handler GetUserLogoutHandler
}

func (o *GetUserLogout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserLogoutParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
