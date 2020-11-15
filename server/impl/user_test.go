package impl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

func TestImplementation_GetUserLogout_Noop(t *testing.T) {
	const sessionName = "session"

	r := httptest.NewRequest(http.MethodGet, new(operations.GetUserLogoutURL).String(), nil)
	session, err := testImpl.SessionStore().New(r, sessionName)
	require.NoError(t, err)
	r = r.WithContext(impl.WithSession(r.Context(), session))

	params := operations.NewGetUserLogoutParams()
	require.NoError(t, params.BindRequest(r, nil))
	w := httptest.NewRecorder()
	testImpl.GetUserLogout(params).WriteResponse(w, runtime.JSONProducer())

	assert.Less(t, session.Options.MaxAge, 0)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestImplementation_GetUserLogout_Logout(t *testing.T) {
	const (
		sessionName = "session"
		userID      = "user ID"
	)

	session, origCookie := encodeSession(t, testImpl.SessionStore(), sessionName, map[interface{}]interface{}{
		impl.UserID: userID,
	})
	require.Equal(t, userID, session.Values[impl.UserID])

	r := httptest.NewRequest(http.MethodGet, new(operations.GetUserLogoutURL).String(), nil)
	r.AddCookie(origCookie)
	r = r.WithContext(impl.WithSession(r.Context(), session))

	params := operations.NewGetUserLogoutParams()
	require.NoError(t, params.BindRequest(r, nil))
	w := httptest.NewRecorder()
	testImpl.GetUserLogout(params).WriteResponse(w, runtime.JSONProducer())

	assert.Less(t, session.Options.MaxAge, 0)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
