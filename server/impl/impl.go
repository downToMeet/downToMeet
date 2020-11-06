package impl

import (
	"encoding/gob"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

// An Implementation provides all server endpoints for the app.
type Implementation struct {
	Options struct {
		Production bool `long:"production" description:"Run in production mode"`
	}

	sessionStore     sessions.Store // could be lazily initialized; use SessionStore() instead!
	sessionStoreInit sync.Once
}

// NewImplementation returns a new Implementation intended for production,
// with a sessions.CookieStore as the internal session store.
func NewImplementation() *Implementation {
	return new(Implementation)
}

// NewImplementation returns a new Implementation with the provided session store.
func NewImplementationWithSessionStore(store sessions.Store) *Implementation {
	i := new(Implementation)
	i.sessionStoreInit.Do(func() {
		i.sessionStore = store
	})
	return i
}

// TODO: change these before deployment.
var (
	authenticationKey = []byte("notsecurenotsecurenotsecurenotse")
	encryptionKey     = []byte("notsecurenotsecu")
)

// SessionStore returns the internal session store.
func (i *Implementation) SessionStore() sessions.Store {
	// Initialize sessionStore lazily since we need to access i.Options.
	i.sessionStoreInit.Do(func() {
		store := sessions.NewCookieStore(authenticationKey, encryptionKey)
		store.Options.HttpOnly = true
		store.Options.Secure = i.Options.Production
		store.Options.SameSite = http.SameSiteStrictMode

		i.sessionStore = store
	})
	return i.sessionStore
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=SessionKey

// SessionKey is the type used to key session.Values.
type SessionKey int

const (
	UserID SessionKey = iota // session.Values[UserID] is a string
)

func init() {
	gob.Register(SessionKey(0))
}
