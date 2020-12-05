package impl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
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

const someURL = "https://example.com/"

func TestWithRequest_Single(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, someURL, nil)
	ctx := WithRequest(context.Background(), r)
	assert.Equal(t, r, RequestFromContext(ctx))
}

func TestWithRequest_Override(t *testing.T) {
	r1 := httptest.NewRequest(http.MethodGet, someURL, nil)
	r2 := httptest.NewRequest(http.MethodGet, someURL, nil)

	ctx1 := WithRequest(context.Background(), r1)
	ctx2 := WithRequest(ctx1, r2)

	assert.Equal(t, r1, RequestFromContext(ctx1))
	assert.Equal(t, r2, RequestFromContext(ctx2)) // the second call to WithRequest should override the first
}

func TestRequestFromContext_Nil(t *testing.T) {
	var nilRequest *http.Request
	assert.Equal(t, nilRequest, RequestFromContext(context.Background()))
}

func TestRequestMiddleware(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, someURL, nil)
	require.Nil(t, RequestFromContext(request.Context()))

	called := false
	RequestMiddleware(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		called = true
		assert.Equal(t, request, RequestFromContext(r.Context()))
	})).ServeHTTP(httptest.NewRecorder(), request)
	require.True(t, called)
}
