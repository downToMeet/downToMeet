// This file is safe to edit.

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/swag"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate server --target ../../server --name DownToMeet --spec ../swagger.yml --principal interface{}

var Impl = impl.NewImplementation()

func configureFlags(api *operations.DownToMeetAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "server options",
			Options:          &Impl.Options,
		},
	}
}

func configureAPI(api *operations.DownToMeetAPI) http.Handler {
	log.AddHook(impl.RequestLogHook{})
	api.Logger = log.Printf

	if Impl.Options.Production {
		api.Middleware = func(builder middleware.Builder) http.Handler {
			return api.Context().RoutesHandler(builder)
		}
	} else {
		// Display documentation at /docs
		api.UseSwaggerUI()
	}

	api.HTMLProducer = runtime.TextProducer()

	_ = Impl.DB() // ensure we are connected to the database

	api.GetUserIDHandler = operations.GetUserIDHandlerFunc(Impl.GetUserID)
	api.PatchUserIDHandler = operations.PatchUserIDHandlerFunc(Impl.PatchUserID)
	api.GetUserMeHandler = operations.GetUserMeHandlerFunc(Impl.GetUserMe)
	api.GetUserLogoutHandler = operations.GetUserLogoutHandlerFunc(Impl.GetUserLogout)
	api.GetUserFacebookAuthHandler = operations.GetUserFacebookAuthHandlerFunc(Impl.GetUserFacebookAuth)
	api.GetUserFacebookRedirectHandler = operations.GetUserFacebookRedirectHandlerFunc(Impl.GetUserFacebookRedirect)
	api.GetUserGoogleAuthHandler = operations.GetUserGoogleAuthHandlerFunc(Impl.GetUserGoogleAuth)
	api.GetUserGoogleRedirectHandler = operations.GetUserGoogleRedirectHandlerFunc(Impl.GetUserGoogleRedirect)

	api.GetMeetupIDHandler = operations.GetMeetupIDHandlerFunc(Impl.GetMeetupID)
	api.GetMeetupRemoteHandler = operations.GetMeetupRemoteHandlerFunc(Impl.GetMeetupRemote)
	api.PostMeetupHandler = operations.PostMeetupHandlerFunc(Impl.PostMeetup)
	api.PatchMeetupIDHandler = operations.PatchMeetupIDHandlerFunc(Impl.PatchMeetupID)
	api.DeleteMeetupIDHandler = operations.DeleteMeetupIDHandlerFunc(Impl.DeleteMeetupID)
	api.GetMeetupHandler = operations.GetMeetupHandlerFunc(Impl.GetMeetup)

	api.GetMeetupIDAttendeeHandler = operations.GetMeetupIDAttendeeHandlerFunc(Impl.GetMeetupIDAttendee)
	api.PostMeetupIDAttendeeHandler = operations.PostMeetupIDAttendeeHandlerFunc(Impl.PostMeetupIDAttendee)
	api.PatchMeetupIDAttendeeHandler = operations.PatchMeetupIDAttendeeHandlerFunc(Impl.PatchMeetupIDAttendee)

	api.APIKeyAuthenticator = func(name, in string, authentication security.TokenAuthentication) runtime.Authenticator {
		if name == "COOKIE" {
			return security.HttpAuthenticator(func(r *http.Request) (authenticated bool, principal interface{}, err error) {
				session := impl.SessionFromContext(r.Context())
				authenticated = session.Values[impl.UserID] != nil
				return authenticated, session, err
			})
		}

		return security.APIKeyAuth(name, in, authentication)
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return Impl.SessionMiddleware(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // TODO: update for deployment
		AllowedMethods:   []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
		Debug:            !Impl.Options.Production,
	})
	c.Log = log.WithField("source", "cors")
	return impl.RequestMiddleware(c.Handler(handler))
}
