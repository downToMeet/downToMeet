package impl

import "net/http"

// Exported for tests.

const (
	StateNonceLen = stateNonceLen
	StateExpiry   = stateExpiry
)

func NewFunctorResponseWriter(fn func(http.ResponseWriter), w http.ResponseWriter) http.ResponseWriter {
	return &functorResponseWriter{fn: fn, w: w}
}
