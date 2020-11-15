package impl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.timothygu.me/downtomeet/server/impl"
)

func TestImplementation_SessionMiddleware_NoCookies(t *testing.T) {
	// Create a new request without any cookies.
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	// Assert that a new session is created.
	called := false
	w := httptest.NewRecorder()
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.True(t, session.IsNew)
		assert.Equal(t, nil, session.Values[UserID])

		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Assert that no cookies were created.
	assert.True(t, called)
	assert.Empty(t, w.Result().Cookies())
}

func TestImplementation_SessionMiddleware_NewCookie(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
	)

	// Create a new request without any cookies.
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	called := false

	// Add a cookie in the handler.
	var origSession *sessions.Session
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		origSession = SessionFromContext(r.Context())
		require.NotNil(t, origSession)
		assert.True(t, origSession.IsNew)

		origSession.Values[UserID] = userID

		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Check for the resulting cookie.
	require.True(t, called)
	require.Len(t, w.Result().Cookies(), 1, "cookies")
	sessionCookie := w.Result().Cookies()[0]
	assert.Equal(t, sessionName, sessionCookie.Name)
	assert.True(t, sessionCookie.Expires.After(time.Now()))

	session := decodeCookie(t, testImpl.SessionStore(), sessionCookie)
	assertEqualSessions(t, origSession, session)
	assert.Equal(t, userID, session.Values[UserID])
}

func TestImplementation_SessionMiddleware_DeleteCookie(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
	)

	// Create a new request with a session cookie.
	origSession, origCookie :=
		encodeSession(t, testImpl.SessionStore(), sessionName, map[interface{}]interface{}{
			UserID: userID,
		})
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(origCookie)

	w := httptest.NewRecorder()
	called := false
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		origSession = SessionFromContext(r.Context())
		require.NotNil(t, origSession)
		assert.False(t, origSession.IsNew)
		origSession.Options.MaxAge = -1

		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Check that the resulting cookie is already expired.
	require.True(t, called)
	require.Len(t, w.Result().Cookies(), 1, "cookies")
	sessionCookie := w.Result().Cookies()[0]
	assert.Equal(t, sessionName, sessionCookie.Name)
	assert.True(t, sessionCookie.Expires.Before(time.Now()))

	// Make sure the expired cookie does not have a user ID.
	session := decodeCookie(t, testImpl.SessionStore(), sessionCookie)
	assert.Nil(t, session.Values[UserID])
}

func TestImplementation_SessionMiddleware_InvalidCookie(t *testing.T) {
	const sessionName = "session"

	// Create a new request with an invalid cookie.
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  sessionName,
		Value: "invalid",
	})

	// Assert that a new session is created.
	called := false
	w := httptest.NewRecorder()
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.True(t, session.IsNew)
		assert.Equal(t, nil, session.Values[UserID])

		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Assert that an expired cookie is sent.
	assert.True(t, called)
	require.Len(t, w.Result().Cookies(), 1, "cookies")
	sessionCookie := w.Result().Cookies()[0]
	assert.Equal(t, sessionName, sessionCookie.Name)
	assert.True(t, sessionCookie.Expires.Before(time.Now()))

	// Make sure the expired cookie does not have a user ID.
	session := decodeCookie(t, testImpl.SessionStore(), sessionCookie)
	assert.Nil(t, session.Values[UserID])
}

func TestImplementation_SessionMiddleware_InvalidCookieUpdate(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
	)

	// Create a new request with an invalid cookie.
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  sessionName,
		Value: "invalid",
	})

	// Assert that a new session is created.
	called := false
	w := httptest.NewRecorder()
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.True(t, session.IsNew)
		assert.Equal(t, nil, session.Values[UserID])

		session.Values[UserID] = userID
		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Assert that a valid cookie is sent.
	require.True(t, called)
	require.Len(t, w.Result().Cookies(), 1, "cookies")
	sessionCookie := w.Result().Cookies()[0]
	assert.Equal(t, sessionName, sessionCookie.Name)
	assert.True(t, sessionCookie.Expires.After(time.Now()))

	session := decodeCookie(t, testImpl.SessionStore(), sessionCookie)
	assert.Equal(t, userID, session.Values[UserID])
}

func TestImplementation_SessionMiddleware_UpdateExisting(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
		userID2     = "user ID 2"
	)

	// Create a new request with a session cookie.
	origSession, origCookie :=
		encodeSession(t, testImpl.SessionStore(), sessionName, map[interface{}]interface{}{
			UserID: userID,
		})
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(origCookie)

	// Assert that the restored session is the same as what was saved, and store
	// a new user ID.
	w := httptest.NewRecorder()
	called := false
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.False(t, session.IsNew)
		assertEqualSessions(t, origSession, session)
		assert.Equal(t, userID, session.Values[UserID])

		session.Values[UserID] = userID2

		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(w, r)

	// Check for the resulting cookie.
	require.True(t, called)
	require.Len(t, w.Result().Cookies(), 1, "cookies")
	sessionCookie := w.Result().Cookies()[0]
	assert.Equal(t, sessionName, sessionCookie.Name)
	assert.True(t, sessionCookie.Expires.After(time.Now()))

	session := decodeCookie(t, testImpl.SessionStore(), sessionCookie)
	assert.Equal(t, userID2, session.Values[UserID])
}

// Helpers...

func decodeCookie(t testing.TB, store sessions.Store, cookie *http.Cookie) *sessions.Session {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(cookie)

	session, err := store.New(r, cookie.Name)
	require.NoError(t, err)
	require.False(t, session.IsNew)
	return session
}

func encodeSession(t testing.TB, store sessions.Store, sessionName string, values map[interface{}]interface{}) (*sessions.Session, *http.Cookie) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	session, err := store.New(r, sessionName)
	require.NoError(t, err)
	session.Values = values

	w := httptest.NewRecorder()
	require.NoError(t, testImpl.SessionStore().Save(r, w, session))
	require.NotEmpty(t, len(w.Result().Cookies()))
	cookie := w.Result().Cookies()[0]
	return session, cookie
}

func assertEqualSessions(t assert.TestingT, expected, actual *sessions.Session) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.EqualValues(t, expected.Values, actual.Values)
}
