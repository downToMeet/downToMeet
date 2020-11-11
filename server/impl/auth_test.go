package impl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.timothygu.me/downtomeet/server/impl"
)

func TestSaveSession(t *testing.T) {
	const sessionName = "session"

	// Set up request with a session.
	store := mockCookieStore()
	session := sessions.NewSession(store, sessionName)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r = r.WithContext(WithSession(r.Context(), session))

	// Write the result of SaveSession to a new ResponseRecorder.
	w := httptest.NewRecorder()
	SaveSession(r, emptyHandler{}).WriteResponse(w, runtime.JSONProducer())

	// Check if the response has the cookies we expect.
	require.Len(t, w.Result().Cookies(), 1)
	assert.Equal(t, sessionName, w.Result().Cookies()[0].Name)
}

func TestImplementation_SessionMiddleware_Nil(t *testing.T) {
	// Create a new request without any cookies.
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	// Assert that a new session is created.
	called := false
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.True(t, session.IsNew)
		assert.Equal(t, nil, session.Values[UserID])
	})).ServeHTTP(httptest.NewRecorder(), r)

	assert.True(t, called)
}

func TestImplementation_SessionMiddleware_Stored(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
	)

	// Create a session and a corresponding origCookie.
	var (
		origSession *sessions.Session
		origCookie  *http.Cookie
		err         error
	)
	{
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		origSession, err = testImpl.SessionStore().New(r, sessionName)
		require.NoError(t, err)
		origSession.Values[UserID] = userID

		w := httptest.NewRecorder()
		require.NoError(t, testImpl.SessionStore().Save(r, w, origSession))
		require.NotEmpty(t, len(w.Result().Cookies()))
		origCookie = w.Result().Cookies()[0]
	}

	// Create a new request with a session cookie.
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(origCookie)

	// Assert that the restored session is the same as what was saved.
	called := false
	testImpl.SessionMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		session := SessionFromContext(r.Context())
		require.NotNil(t, session)
		assert.False(t, session.IsNew)
		assertEqualSessions(t, origSession, session)
		assert.Equal(t, userID, session.Values[UserID])
	})).ServeHTTP(httptest.NewRecorder(), r)

	assert.True(t, called)
}

// Helpers...

// A middleware.Responder and http.Handler that always writes an empty 200 response.
type emptyHandler struct{}

func (emptyHandler) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(200)
}

func (emptyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func assertEqualSessions(t assert.TestingT, expected, actual *sessions.Session) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.EqualValues(t, expected.Values, actual.Values)
}
