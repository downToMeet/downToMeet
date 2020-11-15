// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserFacebookAuthHandlerFunc turns a function with the right signature into a get user facebook auth handler
type GetUserFacebookAuthHandlerFunc func(GetUserFacebookAuthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserFacebookAuthHandlerFunc) Handle(params GetUserFacebookAuthParams) middleware.Responder {
	return fn(params)
}

// GetUserFacebookAuthHandler interface for that can handle valid get user facebook auth params
type GetUserFacebookAuthHandler interface {
	Handle(GetUserFacebookAuthParams) middleware.Responder
}

// NewGetUserFacebookAuth creates a new http.Handler for the get user facebook auth operation
func NewGetUserFacebookAuth(ctx *middleware.Context, handler GetUserFacebookAuthHandler) *GetUserFacebookAuth {
	return &GetUserFacebookAuth{Context: ctx, Handler: handler}
}

/*GetUserFacebookAuth swagger:route GET /user/facebook/auth getUserFacebookAuth

Facebook OAuth login

Start a Facebook OAuth login flow here.

*/
type GetUserFacebookAuth struct {
	Context *middleware.Context
	Handler GetUserFacebookAuthHandler
}

func (o *GetUserFacebookAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserFacebookAuthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
