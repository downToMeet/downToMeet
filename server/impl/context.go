package impl

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

// ctx.Value(contextSession{}) is a *sessions.Session.
type contextSession struct{}

func (contextSession) String() string { return "contextSession" }

// WithSession returns a new context with the session attached, that can later
// be retrieved using SessionFromContext. The following is guaranteed to hold:
//  SessionFromContext(WithSession(ctx, s)) == s
func WithSession(ctx context.Context, session *sessions.Session) context.Context {
	return context.WithValue(ctx, contextSession{}, session)
}

// SessionFromContext retrieves the session previously stored in ctx with
// WithSession. If no such session exists, SessionFromContext returns nil.
func SessionFromContext(ctx context.Context) *sessions.Session {
	if v := ctx.Value(contextSession{}); v != nil {
		return v.(*sessions.Session)
	}
	return nil
}

// ctx.Value(contextRequest{}) is a *http.Request.
type contextRequest struct{}

// WithRequest returns a new context with the given request attached, that can
// later be retrieved using RequestFromContext. The following is guaranteed to
// hold:
//  RequestFromContext(WithRequest(ctx, r)) == r
func WithRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, contextRequest{}, r)
}

// RequestFromContext retrieves the request previously stored in ctx with
// WithRequest. If no such request exists, RequestFromContext returns nil.
func RequestFromContext(ctx context.Context) *http.Request {
	if v := ctx.Value(contextRequest{}); v != nil {
		return v.(*http.Request)
	}
	return nil
}

// RequestMiddleware returns a handler that ensures that the request's context
// has the request attached, so that it can later be retrieved using
// RequestFromContext.
func RequestMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(WithRequest(r.Context(), r))
		h.ServeHTTP(w, r)
	})
}
