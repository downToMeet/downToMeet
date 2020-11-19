// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDownToMeetAPI creates a new DownToMeet instance
func NewDownToMeetAPI(spec *loads.Document) *DownToMeetAPI {
	return &DownToMeetAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),
		TxtProducer:  runtime.TextProducer(),

		DeleteMeetupIDHandler: DeleteMeetupIDHandlerFunc(func(params DeleteMeetupIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DeleteMeetupID has not yet been implemented")
		}),
		GetHelloHandler: GetHelloHandlerFunc(func(params GetHelloParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHello has not yet been implemented")
		}),
		GetMeetupHandler: GetMeetupHandlerFunc(func(params GetMeetupParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMeetup has not yet been implemented")
		}),
		GetMeetupIDHandler: GetMeetupIDHandlerFunc(func(params GetMeetupIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMeetupID has not yet been implemented")
		}),
		GetMeetupIDAttendeeHandler: GetMeetupIDAttendeeHandlerFunc(func(params GetMeetupIDAttendeeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetMeetupIDAttendee has not yet been implemented")
		}),
		GetRestrictedHandler: GetRestrictedHandlerFunc(func(params GetRestrictedParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetRestricted has not yet been implemented")
		}),
		GetSetCookieHandler: GetSetCookieHandlerFunc(func(params GetSetCookieParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSetCookie has not yet been implemented")
		}),
		GetUserFacebookAuthHandler: GetUserFacebookAuthHandlerFunc(func(params GetUserFacebookAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation GetUserFacebookAuth has not yet been implemented")
		}),
		GetUserFacebookRedirectHandler: GetUserFacebookRedirectHandlerFunc(func(params GetUserFacebookRedirectParams) middleware.Responder {
			return middleware.NotImplemented("operation GetUserFacebookRedirect has not yet been implemented")
		}),
		GetUserIDHandler: GetUserIDHandlerFunc(func(params GetUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetUserID has not yet been implemented")
		}),
		GetUserLogoutHandler: GetUserLogoutHandlerFunc(func(params GetUserLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation GetUserLogout has not yet been implemented")
		}),
		GetUserMeHandler: GetUserMeHandlerFunc(func(params GetUserMeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetUserMe has not yet been implemented")
		}),
		PatchMeetupIDHandler: PatchMeetupIDHandlerFunc(func(params PatchMeetupIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PatchMeetupID has not yet been implemented")
		}),
		PatchMeetupIDAttendeeHandler: PatchMeetupIDAttendeeHandlerFunc(func(params PatchMeetupIDAttendeeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PatchMeetupIDAttendee has not yet been implemented")
		}),
		PatchUserIDHandler: PatchUserIDHandlerFunc(func(params PatchUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PatchUserID has not yet been implemented")
		}),
		PostMeetupHandler: PostMeetupHandlerFunc(func(params PostMeetupParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostMeetup has not yet been implemented")
		}),
		PostMeetupIDAttendeeHandler: PostMeetupIDAttendeeHandlerFunc(func(params PostMeetupIDAttendeeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostMeetupIDAttendee has not yet been implemented")
		}),
		PostUserHandler: PostUserHandlerFunc(func(params PostUserParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUser has not yet been implemented")
		}),

		// Applies when the "COOKIE" query is set
		CookieSessionAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (cookieSession) COOKIE from query param [COOKIE] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DownToMeetAPI the down to meet API */
type DownToMeetAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - text/plain
	TxtProducer runtime.Producer

	// CookieSessionAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key COOKIE provided in the query
	CookieSessionAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DeleteMeetupIDHandler sets the operation handler for the delete meetup ID operation
	DeleteMeetupIDHandler DeleteMeetupIDHandler
	// GetHelloHandler sets the operation handler for the get hello operation
	GetHelloHandler GetHelloHandler
	// GetMeetupHandler sets the operation handler for the get meetup operation
	GetMeetupHandler GetMeetupHandler
	// GetMeetupIDHandler sets the operation handler for the get meetup ID operation
	GetMeetupIDHandler GetMeetupIDHandler
	// GetMeetupIDAttendeeHandler sets the operation handler for the get meetup ID attendee operation
	GetMeetupIDAttendeeHandler GetMeetupIDAttendeeHandler
	// GetRestrictedHandler sets the operation handler for the get restricted operation
	GetRestrictedHandler GetRestrictedHandler
	// GetSetCookieHandler sets the operation handler for the get set cookie operation
	GetSetCookieHandler GetSetCookieHandler
	// GetUserFacebookAuthHandler sets the operation handler for the get user facebook auth operation
	GetUserFacebookAuthHandler GetUserFacebookAuthHandler
	// GetUserFacebookRedirectHandler sets the operation handler for the get user facebook redirect operation
	GetUserFacebookRedirectHandler GetUserFacebookRedirectHandler
	// GetUserIDHandler sets the operation handler for the get user ID operation
	GetUserIDHandler GetUserIDHandler
	// GetUserLogoutHandler sets the operation handler for the get user logout operation
	GetUserLogoutHandler GetUserLogoutHandler
	// GetUserMeHandler sets the operation handler for the get user me operation
	GetUserMeHandler GetUserMeHandler
	// PatchMeetupIDHandler sets the operation handler for the patch meetup ID operation
	PatchMeetupIDHandler PatchMeetupIDHandler
	// PatchMeetupIDAttendeeHandler sets the operation handler for the patch meetup ID attendee operation
	PatchMeetupIDAttendeeHandler PatchMeetupIDAttendeeHandler
	// PatchUserIDHandler sets the operation handler for the patch user ID operation
	PatchUserIDHandler PatchUserIDHandler
	// PostMeetupHandler sets the operation handler for the post meetup operation
	PostMeetupHandler PostMeetupHandler
	// PostMeetupIDAttendeeHandler sets the operation handler for the post meetup ID attendee operation
	PostMeetupIDAttendeeHandler PostMeetupIDAttendeeHandler
	// PostUserHandler sets the operation handler for the post user operation
	PostUserHandler PostUserHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *DownToMeetAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *DownToMeetAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *DownToMeetAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DownToMeetAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DownToMeetAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DownToMeetAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DownToMeetAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DownToMeetAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DownToMeetAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DownToMeetAPI
func (o *DownToMeetAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.CookieSessionAuth == nil {
		unregistered = append(unregistered, "COOKIEAuth")
	}

	if o.DeleteMeetupIDHandler == nil {
		unregistered = append(unregistered, "DeleteMeetupIDHandler")
	}
	if o.GetHelloHandler == nil {
		unregistered = append(unregistered, "GetHelloHandler")
	}
	if o.GetMeetupHandler == nil {
		unregistered = append(unregistered, "GetMeetupHandler")
	}
	if o.GetMeetupIDHandler == nil {
		unregistered = append(unregistered, "GetMeetupIDHandler")
	}
	if o.GetMeetupIDAttendeeHandler == nil {
		unregistered = append(unregistered, "GetMeetupIDAttendeeHandler")
	}
	if o.GetRestrictedHandler == nil {
		unregistered = append(unregistered, "GetRestrictedHandler")
	}
	if o.GetSetCookieHandler == nil {
		unregistered = append(unregistered, "GetSetCookieHandler")
	}
	if o.GetUserFacebookAuthHandler == nil {
		unregistered = append(unregistered, "GetUserFacebookAuthHandler")
	}
	if o.GetUserFacebookRedirectHandler == nil {
		unregistered = append(unregistered, "GetUserFacebookRedirectHandler")
	}
	if o.GetUserIDHandler == nil {
		unregistered = append(unregistered, "GetUserIDHandler")
	}
	if o.GetUserLogoutHandler == nil {
		unregistered = append(unregistered, "GetUserLogoutHandler")
	}
	if o.GetUserMeHandler == nil {
		unregistered = append(unregistered, "GetUserMeHandler")
	}
	if o.PatchMeetupIDHandler == nil {
		unregistered = append(unregistered, "PatchMeetupIDHandler")
	}
	if o.PatchMeetupIDAttendeeHandler == nil {
		unregistered = append(unregistered, "PatchMeetupIDAttendeeHandler")
	}
	if o.PatchUserIDHandler == nil {
		unregistered = append(unregistered, "PatchUserIDHandler")
	}
	if o.PostMeetupHandler == nil {
		unregistered = append(unregistered, "PostMeetupHandler")
	}
	if o.PostMeetupIDAttendeeHandler == nil {
		unregistered = append(unregistered, "PostMeetupIDAttendeeHandler")
	}
	if o.PostUserHandler == nil {
		unregistered = append(unregistered, "PostUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DownToMeetAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DownToMeetAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "cookieSession":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CookieSessionAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DownToMeetAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DownToMeetAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DownToMeetAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/plain":
			result["text/plain"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DownToMeetAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the down to meet API
func (o *DownToMeetAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DownToMeetAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/meetup/{id}"] = NewDeleteMeetupID(o.context, o.DeleteMeetupIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/hello"] = NewGetHello(o.context, o.GetHelloHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meetup"] = NewGetMeetup(o.context, o.GetMeetupHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meetup/{id}"] = NewGetMeetupID(o.context, o.GetMeetupIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meetup/{id}/attendee"] = NewGetMeetupIDAttendee(o.context, o.GetMeetupIDAttendeeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/restricted"] = NewGetRestricted(o.context, o.GetRestrictedHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/set-cookie"] = NewGetSetCookie(o.context, o.GetSetCookieHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/facebook/auth"] = NewGetUserFacebookAuth(o.context, o.GetUserFacebookAuthHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/facebook/redirect"] = NewGetUserFacebookRedirect(o.context, o.GetUserFacebookRedirectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{id}"] = NewGetUserID(o.context, o.GetUserIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = NewGetUserLogout(o.context, o.GetUserLogoutHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/me"] = NewGetUserMe(o.context, o.GetUserMeHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/meetup/{id}"] = NewPatchMeetupID(o.context, o.PatchMeetupIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/meetup/{id}/attendee"] = NewPatchMeetupIDAttendee(o.context, o.PatchMeetupIDAttendeeHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/user/{id}"] = NewPatchUserID(o.context, o.PatchUserIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/meetup"] = NewPostMeetup(o.context, o.PostMeetupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/meetup/{id}/attendee"] = NewPostMeetupIDAttendee(o.context, o.PostMeetupIDAttendeeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = NewPostUser(o.context, o.PostUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DownToMeetAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DownToMeetAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DownToMeetAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DownToMeetAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DownToMeetAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
