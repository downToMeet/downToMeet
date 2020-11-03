// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// GetUserMeOKCode is the HTTP code returned for type GetUserMeOK
const GetUserMeOKCode int = 200

/*GetUserMeOK OK

swagger:response getUserMeOK
*/
type GetUserMeOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserMeOK creates GetUserMeOK with default headers values
func NewGetUserMeOK() *GetUserMeOK {

	return &GetUserMeOK{}
}

// WithPayload adds the payload to the get user me o k response
func (o *GetUserMeOK) WithPayload(payload *models.User) *GetUserMeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user me o k response
func (o *GetUserMeOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserMeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserMeForbiddenCode is the HTTP code returned for type GetUserMeForbidden
const GetUserMeForbiddenCode int = 403

/*GetUserMeForbidden Forbidden

swagger:response getUserMeForbidden
*/
type GetUserMeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserMeForbidden creates GetUserMeForbidden with default headers values
func NewGetUserMeForbidden() *GetUserMeForbidden {

	return &GetUserMeForbidden{}
}

// WithPayload adds the payload to the get user me forbidden response
func (o *GetUserMeForbidden) WithPayload(payload *models.Error) *GetUserMeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user me forbidden response
func (o *GetUserMeForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserMeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserMeInternalServerErrorCode is the HTTP code returned for type GetUserMeInternalServerError
const GetUserMeInternalServerErrorCode int = 500

/*GetUserMeInternalServerError Internal server error

swagger:response getUserMeInternalServerError
*/
type GetUserMeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserMeInternalServerError creates GetUserMeInternalServerError with default headers values
func NewGetUserMeInternalServerError() *GetUserMeInternalServerError {

	return &GetUserMeInternalServerError{}
}

// WithPayload adds the payload to the get user me internal server error response
func (o *GetUserMeInternalServerError) WithPayload(payload *models.Error) *GetUserMeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user me internal server error response
func (o *GetUserMeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserMeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}