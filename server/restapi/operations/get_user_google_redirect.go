// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserGoogleRedirectHandlerFunc turns a function with the right signature into a get user google redirect handler
type GetUserGoogleRedirectHandlerFunc func(GetUserGoogleRedirectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserGoogleRedirectHandlerFunc) Handle(params GetUserGoogleRedirectParams) middleware.Responder {
	return fn(params)
}

// GetUserGoogleRedirectHandler interface for that can handle valid get user google redirect params
type GetUserGoogleRedirectHandler interface {
	Handle(GetUserGoogleRedirectParams) middleware.Responder
}

// NewGetUserGoogleRedirect creates a new http.Handler for the get user google redirect operation
func NewGetUserGoogleRedirect(ctx *middleware.Context, handler GetUserGoogleRedirectHandler) *GetUserGoogleRedirect {
	return &GetUserGoogleRedirect{Context: ctx, Handler: handler}
}

/*GetUserGoogleRedirect swagger:route GET /user/google/redirect getUserGoogleRedirect

Google OAuth redirect

If authentication fails, the user is not logged in.

*/
type GetUserGoogleRedirect struct {
	Context *middleware.Context
	Handler GetUserGoogleRedirectHandler
}

func (o *GetUserGoogleRedirect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserGoogleRedirectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}