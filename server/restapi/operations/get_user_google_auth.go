// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserGoogleAuthHandlerFunc turns a function with the right signature into a get user google auth handler
type GetUserGoogleAuthHandlerFunc func(GetUserGoogleAuthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserGoogleAuthHandlerFunc) Handle(params GetUserGoogleAuthParams) middleware.Responder {
	return fn(params)
}

// GetUserGoogleAuthHandler interface for that can handle valid get user google auth params
type GetUserGoogleAuthHandler interface {
	Handle(GetUserGoogleAuthParams) middleware.Responder
}

// NewGetUserGoogleAuth creates a new http.Handler for the get user google auth operation
func NewGetUserGoogleAuth(ctx *middleware.Context, handler GetUserGoogleAuthHandler) *GetUserGoogleAuth {
	return &GetUserGoogleAuth{Context: ctx, Handler: handler}
}

/*GetUserGoogleAuth swagger:route GET /user/google/auth getUserGoogleAuth

Google OAuth login

Start a Google OAuth login flow here.

*/
type GetUserGoogleAuth struct {
	Context *middleware.Context
	Handler GetUserGoogleAuthHandler
}

func (o *GetUserGoogleAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserGoogleAuthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}