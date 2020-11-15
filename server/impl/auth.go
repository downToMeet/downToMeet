package impl

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"

	"go.timothygu.me/downtomeet/server/restapi/operations"
)

// SessionMiddleware augments the request's context with a session that can be
// fetched using SessionFromContext(r.Context()). Additionally, if the handler
// wrote values into the session or if there was a preexisting session, the
// updated cookie added to the response.
func (i *Implementation) SessionMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, sessionErr := i.SessionStore().New(r, "session")
		r = r.WithContext(WithSession(r.Context(), session))

		fn := func(w http.ResponseWriter) {
			// Delete any empty sessions.
			if len(session.Values) == 0 {
				session.Options.MaxAge = -1
			}

			// Save the session if the request included a (possibly invalid)
			// session, or if we made any changes to it.
			if sessionErr != nil || !session.IsNew || len(session.Values) > 0 {
				// If we are deleting the cookie anyway, remove its values.
				if session.Options.MaxAge < 0 {
					session.Values = nil
				}

				if err := session.Save(r, w); err != nil {
					log.WithContext(r.Context()).WithError(err).Warn("Failed to save cookies")
				}
			}
		}
		handler.ServeHTTP(&functorResponseWriter{w: w, fn: fn}, r)
	})
}

// Deprecated: This is a dummy endpoint that should be removed.
func (i *Implementation) GetSetCookie(params operations.GetSetCookieParams) middleware.Responder {
	session := SessionFromContext(params.HTTPRequest.Context())
	session.Values[UserID] = swag.StringValue(params.UserID)

	return operations.NewGetSetCookieOK().
		WithPayload(fmt.Sprintf("setting user ID to %s", swag.StringValue(params.UserID)))
}

// Deprecated: This is a dummy endpoint that should be removed.
func (i *Implementation) GetRestricted(params operations.GetRestrictedParams, _ interface{}) middleware.Responder {
	session := SessionFromContext(params.HTTPRequest.Context())
	return operations.NewGetRestrictedOK().
		WithPayload(fmt.Sprintf("user ID is %s", session.Values[UserID]))
}

// https://kevin.burke.dev/kevin/how-to-write-go-middleware/

// functorResponseWriter is a http.ResponseWriter that calls fn before writing
// headers. The returned writer is a http.Flusher as well, but not an
// http.Pusher (since we only support HTTP/1).
type functorResponseWriter struct {
	w           http.ResponseWriter
	fn          func(http.ResponseWriter)
	wroteHeader bool
}

func (w *functorResponseWriter) tryCallFunctor() {
	if !w.wroteHeader {
		w.fn(w.w)
	}
}

func (w *functorResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *functorResponseWriter) Write(b []byte) (int, error) {
	w.tryCallFunctor()
	return w.w.Write(b)
}

func (w *functorResponseWriter) WriteHeader(statusCode int) {
	w.tryCallFunctor()
	w.w.WriteHeader(statusCode)
}

func (w *functorResponseWriter) Flush() {
	if f, ok := w.w.(http.Flusher); ok {
		f.Flush()
	}
}
