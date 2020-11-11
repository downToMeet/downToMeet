package impl_test

import (
	"context"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"

	. "go.timothygu.me/downtomeet/server/impl"
)

func TestWithSession_Single(t *testing.T) {
	store := mockCookieStore()
	session := sessions.NewSession(store, "session")

	ctx := WithSession(context.Background(), session)
	assert.Equal(t, session, SessionFromContext(ctx))
}

func TestWithSession_Override(t *testing.T) {
	store := mockCookieStore()
	session1 := sessions.NewSession(store, "session1")
	session2 := sessions.NewSession(store, "session2")

	ctx1 := WithSession(context.Background(), session1)
	ctx2 := WithSession(ctx1, session2)

	assert.Equal(t, session1, SessionFromContext(ctx1))
	assert.Equal(t, session2, SessionFromContext(ctx2)) // the second call to WithSession should override the first
}

func TestSessionFromContext_Nil(t *testing.T) {
	var nilSession *sessions.Session
	assert.Equal(t, nilSession, SessionFromContext(context.Background()))
}
