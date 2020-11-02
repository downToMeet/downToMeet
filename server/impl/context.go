package impl

import (
	"context"

	"github.com/gorilla/sessions"
)

// contextKey is the type used to key values stored in a context.
type contextKey int

const (
	contextSession contextKey = iota // ctx.Value(contextSession) is a *sessions.Session
)

func WithSession(ctx context.Context, session *sessions.Session) context.Context {
	return context.WithValue(ctx, contextSession, session)
}

func SessionFromContext(ctx context.Context) *sessions.Session {
	if v := ctx.Value(contextSession); v != nil {
		return v.(*sessions.Session)
	}
	return nil
}
