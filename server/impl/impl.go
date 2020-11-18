package impl

import (
	"encoding/gob"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go.timothygu.me/downtomeet/server/db"
	"go.timothygu.me/downtomeet/server/impl/nonce"
)

// An Implementation provides all server endpoints for the app.
type Implementation struct {
	Options struct {
		Production bool   `long:"production" description:"Run in production mode"`
		Database   string `long:"database" description:"URL of Postgres DB" default:"postgresql://localhost:5432/downtomeet"`
	}

	sessionStore     sessions.Store // could be lazily initialized; use SessionStore() instead!
	sessionStoreInit sync.Once
	db               *gorm.DB // could be lazily initialized; use DB() instead!
	dbInit           sync.Once

	nonceGen *nonce.Generator
}

// NewImplementation returns a new Implementation intended for production,
// with a sessions.CookieStore as the internal session store.
func NewImplementation() *Implementation {
	i := new(Implementation)
	randSrc := nonce.NewCryptoRandSource()
	i.nonceGen = nonce.NewGenerator(randSrc)
	return i
}

// NewMockImplementation returns a new Implementation with the provided
// parameters.
func NewMockImplementation(store sessions.Store, randSrc rand.Source) *Implementation {
	i := new(Implementation)
	i.nonceGen = nonce.NewGenerator(randSrc)
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

// DB returns the database associated with the Implementation.
func (i *Implementation) DB() *gorm.DB {
	i.dbInit.Do(func() {
		var err error
		if i.db, err = db.Get(log.StandardLogger(), i.Options.Database); err != nil {
			log.WithError(err).Panic("Unable to initialize the database")
		}
	})
	return i.db
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=SessionKey

// SessionKey is the type used to key session.Values.
type SessionKey int

const (
	UserID        SessionKey = 0 // session.Values[UserID] is a string
	FacebookState SessionKey = 1 // session.Values[FacebookState] is an OAuthState
)

func init() {
	gob.Register(SessionKey(0))
	gob.Register(OAuthState{})
}
