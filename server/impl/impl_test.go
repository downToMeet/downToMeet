package impl_test

import (
	"bytes"

	"github.com/gorilla/sessions"

	. "go.timothygu.me/downtomeet/server/impl"
)

var testImpl = NewMockImplementation(mockCookieStore())

var (
	mockAuthenticationKey = bytes.Repeat([]byte{'n'}, 32)
	mockEncryptionKey     = bytes.Repeat([]byte{'e'}, 16)
)

func mockCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore(mockAuthenticationKey, mockEncryptionKey)
}
