package impl

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// InternalServerError is a middleware.Responder that returns a generic 500
// Internal Server Error response.
type InternalServerError struct{}

func (i InternalServerError) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(http.StatusInternalServerError)

	if err := p.Produce(w, &models.Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong.",
	}); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
