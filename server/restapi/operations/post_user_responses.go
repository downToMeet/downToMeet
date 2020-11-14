// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// PostUserOKCode is the HTTP code returned for type PostUserOK
const PostUserOKCode int = 200

/*PostUserOK OK

swagger:response postUserOK
*/
type PostUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUserOK creates PostUserOK with default headers values
func NewPostUserOK() *PostUserOK {

	return &PostUserOK{}
}

// WithPayload adds the payload to the post user o k response
func (o *PostUserOK) WithPayload(payload *models.User) *PostUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user o k response
func (o *PostUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
