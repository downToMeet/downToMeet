// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserFacebookRedirectHandlerFunc turns a function with the right signature into a get user facebook redirect handler
type GetUserFacebookRedirectHandlerFunc func(GetUserFacebookRedirectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserFacebookRedirectHandlerFunc) Handle(params GetUserFacebookRedirectParams) middleware.Responder {
	return fn(params)
}

// GetUserFacebookRedirectHandler interface for that can handle valid get user facebook redirect params
type GetUserFacebookRedirectHandler interface {
	Handle(GetUserFacebookRedirectParams) middleware.Responder
}

// NewGetUserFacebookRedirect creates a new http.Handler for the get user facebook redirect operation
func NewGetUserFacebookRedirect(ctx *middleware.Context, handler GetUserFacebookRedirectHandler) *GetUserFacebookRedirect {
	return &GetUserFacebookRedirect{Context: ctx, Handler: handler}
}

/*GetUserFacebookRedirect swagger:route GET /user/facebook/redirect getUserFacebookRedirect

Facebook OAuth redirect

If authentication fails, the user is not logged in.

*/
type GetUserFacebookRedirect struct {
	Context *middleware.Context
	Handler GetUserFacebookRedirectHandler
}

func (o *GetUserFacebookRedirect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserFacebookRedirectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
