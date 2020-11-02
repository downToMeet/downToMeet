// This file is safe to edit.

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"go.timothygu.me/downtomeet/server/impl"
	"go.timothygu.me/downtomeet/server/restapi/operations"
)

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate server --target ../../server --name DownToMeet --spec ../swagger.yml --principal interface{}

var serverOpts struct {
	Production bool `long:"production" description:"Run in production mode"`
}

func configureFlags(api *operations.DownToMeetAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "server options",
			Options:          &serverOpts,
		},
	}
}

func configureAPI(api *operations.DownToMeetAPI) http.Handler {
	if serverOpts.Production {
		api.Middleware = func(builder middleware.Builder) http.Handler {
			return api.Context().RoutesHandler(builder)
		}
	} else {
		// Display documentation at /docs
		api.UseRedoc()
	}

	api.GetHelloHandler = operations.GetHelloHandlerFunc(impl.GetHello)

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
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
