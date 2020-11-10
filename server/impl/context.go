package impl

import (
	"context"

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
