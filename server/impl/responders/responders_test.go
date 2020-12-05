package responders_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"

	. "go.timothygu.me/downtomeet/server/impl/responders"
	"go.timothygu.me/downtomeet/server/models"
)

func TestInternalServerError(t *testing.T) {
	called := false
	p := runtime.ProducerFunc(func(w io.Writer, i interface{}) error {
		called = true
		typedCode := models.Error{Code: http.StatusInternalServerError}.Code
		assert.Equal(t, i.(*models.Error).Code, typedCode)
		return runtime.JSONProducer().Produce(w, i)
	})

	w := httptest.NewRecorder()
	InternalServerError{}.WriteResponse(w, p)
	assert.True(t, called, "producer is invoked")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSoftRedirect(t *testing.T) {
	p := runtime.TextProducer()
	w := httptest.NewRecorder()
	SoftRedirect{URL: "https://example.com/"}.WriteResponse(w, p)
	assert.Equal(t, http.StatusOK, w.Code)

	contentType, _, err := runtime.ContentType(w.Header())
	assert.NoError(t, err, "parsing response Content-Type")
	assert.Equal(t, runtime.HTMLMime, contentType)
}

func TestSoftRedirect_NoURL(t *testing.T) {
	p := runtime.TextProducer()
	w := httptest.NewRecorder()
	assert.Panics(t, func() {
		SoftRedirect{URL: ""}.WriteResponse(w, p)
	})
}
