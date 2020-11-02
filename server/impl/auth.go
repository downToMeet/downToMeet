package impl

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"

	"go.timothygu.me/downtomeet/server/restapi/operations"
)

// SessionMiddleware augments the request's context with a session that can be fetched
// using SessionFromContext(r.Context()).
func (i *Implementation) SessionMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := i.SessionStore().New(r, "session")
		r = r.WithContext(WithSession(r.Context(), session))

		handler.ServeHTTP(w, r)
	})
}

// SaveSession creates a new responder that saves the session before calling
// responder.WriteResponse.
func SaveSession(r *http.Request, responder middleware.Responder) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
		ctx := r.Context()
		session := SessionFromContext(ctx)
		if err := session.Save(r, w); err != nil {
			log.WithContext(ctx).WithError(err).Println("Failed to save cookies")
		}

		responder.WriteResponse(w, p)
	})
}

// Deprecated: This is a dummy endpoint that should be removed.
func (i *Implementation) GetSetCookie(params operations.GetSetCookieParams) middleware.Responder {
	session := SessionFromContext(params.HTTPRequest.Context())
	session.Values[UserID] = swag.StringValue(params.UserID)

	return SaveSession(
		params.HTTPRequest,
		operations.NewGetSetCookieOK().
			WithPayload(fmt.Sprintf("setting user ID to %s", swag.StringValue(params.UserID))))
}

// Deprecated: This is a dummy endpoint that should be removed.
func (i *Implementation) GetRestricted(params operations.GetRestrictedParams, _ interface{}) middleware.Responder {
	session := SessionFromContext(params.HTTPRequest.Context())
	return operations.NewGetRestrictedOK().
		WithPayload(fmt.Sprintf("user ID is %s", session.Values[UserID]))
}
