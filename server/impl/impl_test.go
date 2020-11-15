package impl_test

import (
	"bytes"
	"math/rand"

	"github.com/gorilla/sessions"

	. "go.timothygu.me/downtomeet/server/impl"
)

var testImpl = NewMockImplementation(mockCookieStore(), rand.NewSource(0))

var (
	mockAuthenticationKey = bytes.Repeat([]byte{'n'}, 32)
	mockEncryptionKey     = bytes.Repeat([]byte{'e'}, 16)
)

func mockCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore(mockAuthenticationKey, mockEncryptionKey)
}
