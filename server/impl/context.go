package impl

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

// ctx.Value(contextSession{}) is a *sessions.Session.
type contextSession struct{}

func (contextSession) String() string { return "contextSession" }

func WithSession(ctx context.Context, session *sessions.Session) context.Context {
	return context.WithValue(ctx, contextSession{}, session)
}

func SessionFromContext(ctx context.Context) *sessions.Session {
	if v := ctx.Value(contextSession{}); v != nil {
		return v.(*sessions.Session)
	}
	return nil
}

// ctx.Value(contextRequest{}) is a *http.Request.
type contextRequest struct{}

func WithRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, contextRequest{}, r)
}

func RequestFromSession(ctx context.Context) *http.Request {
	if v := ctx.Value(contextRequest{}); v != nil {
		return v.(*http.Request)
	}
	return nil
}

func RequestMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(WithRequest(r.Context(), r))
		h.ServeHTTP(w, r)
	})
}
