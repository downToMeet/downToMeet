// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// DeleteMeetupIDNoContentCode is the HTTP code returned for type DeleteMeetupIDNoContent
const DeleteMeetupIDNoContentCode int = 204

/*DeleteMeetupIDNoContent No Content

swagger:response deleteMeetupIdNoContent
*/
type DeleteMeetupIDNoContent struct {
}

// NewDeleteMeetupIDNoContent creates DeleteMeetupIDNoContent with default headers values
func NewDeleteMeetupIDNoContent() *DeleteMeetupIDNoContent {

	return &DeleteMeetupIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteMeetupIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteMeetupIDBadRequestCode is the HTTP code returned for type DeleteMeetupIDBadRequest
const DeleteMeetupIDBadRequestCode int = 400

/*DeleteMeetupIDBadRequest Bad Request

swagger:response deleteMeetupIdBadRequest
*/
type DeleteMeetupIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMeetupIDBadRequest creates DeleteMeetupIDBadRequest with default headers values
func NewDeleteMeetupIDBadRequest() *DeleteMeetupIDBadRequest {

	return &DeleteMeetupIDBadRequest{}
}

// WithPayload adds the payload to the delete meetup Id bad request response
func (o *DeleteMeetupIDBadRequest) WithPayload(payload *models.Error) *DeleteMeetupIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete meetup Id bad request response
func (o *DeleteMeetupIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMeetupIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMeetupIDForbiddenCode is the HTTP code returned for type DeleteMeetupIDForbidden
const DeleteMeetupIDForbiddenCode int = 403

/*DeleteMeetupIDForbidden Forbidden

swagger:response deleteMeetupIdForbidden
*/
type DeleteMeetupIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMeetupIDForbidden creates DeleteMeetupIDForbidden with default headers values
func NewDeleteMeetupIDForbidden() *DeleteMeetupIDForbidden {

	return &DeleteMeetupIDForbidden{}
}

// WithPayload adds the payload to the delete meetup Id forbidden response
func (o *DeleteMeetupIDForbidden) WithPayload(payload *models.Error) *DeleteMeetupIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete meetup Id forbidden response
func (o *DeleteMeetupIDForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMeetupIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMeetupIDNotFoundCode is the HTTP code returned for type DeleteMeetupIDNotFound
const DeleteMeetupIDNotFoundCode int = 404

/*DeleteMeetupIDNotFound Not found

swagger:response deleteMeetupIdNotFound
*/
type DeleteMeetupIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteMeetupIDNotFound creates DeleteMeetupIDNotFound with default headers values
func NewDeleteMeetupIDNotFound() *DeleteMeetupIDNotFound {

	return &DeleteMeetupIDNotFound{}
}

// WithPayload adds the payload to the delete meetup Id not found response
func (o *DeleteMeetupIDNotFound) WithPayload(payload *models.Error) *DeleteMeetupIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete meetup Id not found response
func (o *DeleteMeetupIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMeetupIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
