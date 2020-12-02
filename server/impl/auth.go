package impl

import (
	"net/http"

	log "github.com/sirupsen/logrus"
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
		functorMiddleware{h: handler, fn: fn}.ServeHTTP(w, r)
	})
}

// https://kevin.burke.dev/kevin/how-to-write-go-middleware/

type functorMiddleware struct {
	h  http.Handler
	fn func(w http.ResponseWriter)
}

func (m functorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrapped := &functorResponseWriter{w: w, fn: m.fn}
	m.h.ServeHTTP(wrapped, r)
	wrapped.tryCallFunctor()
}

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
