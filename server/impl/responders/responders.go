// Package responders contains implementations of middleware.Responder that can
// be useful for any endpoint.
package responders

import (
	"html/template"
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// InternalServerError is a middleware.Responder that returns a generic 500
// Internal Server Error response.
type InternalServerError struct{}

// WriteResponse implements middleware.Responder.
func (i InternalServerError) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(http.StatusInternalServerError)

	if err := p.Produce(w, &models.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong.",
	}); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SoftRedirect is a middleware.Responder that writes an HTML page to the
// response that would redirect the user to the provided URL. This can be used
// instead of a normal HTTP 3XX redirect, as a way to make the browser think the
// navigation is "same-site" and send SameSite=strict cookies to the server. In
// particular, SoftRedirect implements this Stack Overflow answer:
// https://stackoverflow.com/a/42220786/1937836.
type SoftRedirect struct{ URL string }

// WriteResponse implements middleware.Responder.
func (s SoftRedirect) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	if s.URL == "" {
		panic("SoftRedirect: no URL provided")
	}
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if err := redirectTemplate.Execute(w, s.URL); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

var redirectTemplate = template.Must(template.New("").Parse(`<!doctype html>
<html lang="en">
<meta charset="utf-8">
<a href="{{.}}">Click here to redirect</a>
<script>location.assign({{.}});</script>
`))
