package impl_test

import (
	"context"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.timothygu.me/downtomeet/server/impl"
)

func TestWithSession_Single(t *testing.T) {
	store := mockCookieStore()
	session := sessions.NewSession(store, "session")

	ctx := WithSession(context.Background(), session)
	assert.Equal(t, session, SessionFromContext(ctx))
}

func TestWithSession_Overwrite(t *testing.T) {
	store := mockCookieStore()
	session := sessions.NewSession(store, "session")
	session2 := sessions.NewSession(store, "session")

	ctx1 := WithSession(context.Background(), session)
	ctx2 := WithSession(ctx1, session2)

	require.Equal(t, session, SessionFromContext(ctx1))
	assert.Equal(t, session2, SessionFromContext(ctx2)) // the second call should overwrite the first
}

func TestSessionFromContext_Nil(t *testing.T) {
	var nilSession *sessions.Session
	assert.Equal(t, nilSession, SessionFromContext(context.Background()))
}
