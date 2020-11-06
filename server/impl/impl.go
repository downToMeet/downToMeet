package impl

import (
	"encoding/gob"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

type Implementation struct {
	Options struct {
		Production bool `long:"production" description:"Run in production mode"`
	}

	sessionStore     *sessions.CookieStore // lazily initialized; use SessionStore() instead!
	sessionStoreInit sync.Once
}

func NewImplementation() *Implementation {
	return new(Implementation)
}

// TODO: change these before deployment.
var (
	authenticationKey = []byte("notsecurenotsecurenotsecurenotse")
	encryptionKey     = []byte("notsecurenotsecu")
)

func (i *Implementation) SessionStore() sessions.Store {
	// Initialize sessionStore lazily since we need to access i.Options.
	i.sessionStoreInit.Do(func() {
		i.sessionStore = sessions.NewCookieStore(authenticationKey, encryptionKey)
		i.sessionStore.Options.HttpOnly = true
		i.sessionStore.Options.Secure = i.Options.Production
		i.sessionStore.Options.SameSite = http.SameSiteStrictMode
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
