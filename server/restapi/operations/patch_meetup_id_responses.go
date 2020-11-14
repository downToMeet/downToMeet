// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// PatchMeetupIDOKCode is the HTTP code returned for type PatchMeetupIDOK
const PatchMeetupIDOKCode int = 200

/*PatchMeetupIDOK OK

swagger:response patchMeetupIdOK
*/
type PatchMeetupIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Meetup `json:"body,omitempty"`
}

// NewPatchMeetupIDOK creates PatchMeetupIDOK with default headers values
func NewPatchMeetupIDOK() *PatchMeetupIDOK {

	return &PatchMeetupIDOK{}
}

// WithPayload adds the payload to the patch meetup Id o k response
func (o *PatchMeetupIDOK) WithPayload(payload *models.Meetup) *PatchMeetupIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id o k response
func (o *PatchMeetupIDOK) SetPayload(payload *models.Meetup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMeetupIDBadRequestCode is the HTTP code returned for type PatchMeetupIDBadRequest
const PatchMeetupIDBadRequestCode int = 400

/*PatchMeetupIDBadRequest Bad Request

swagger:response patchMeetupIdBadRequest
*/
type PatchMeetupIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDBadRequest creates PatchMeetupIDBadRequest with default headers values
func NewPatchMeetupIDBadRequest() *PatchMeetupIDBadRequest {

	return &PatchMeetupIDBadRequest{}
}

// WithPayload adds the payload to the patch meetup Id bad request response
func (o *PatchMeetupIDBadRequest) WithPayload(payload *models.Error) *PatchMeetupIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id bad request response
func (o *PatchMeetupIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMeetupIDForbiddenCode is the HTTP code returned for type PatchMeetupIDForbidden
const PatchMeetupIDForbiddenCode int = 403

/*PatchMeetupIDForbidden Forbidden

swagger:response patchMeetupIdForbidden
*/
type PatchMeetupIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDForbidden creates PatchMeetupIDForbidden with default headers values
func NewPatchMeetupIDForbidden() *PatchMeetupIDForbidden {

	return &PatchMeetupIDForbidden{}
}

// WithPayload adds the payload to the patch meetup Id forbidden response
func (o *PatchMeetupIDForbidden) WithPayload(payload *models.Error) *PatchMeetupIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id forbidden response
func (o *PatchMeetupIDForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMeetupIDNotFoundCode is the HTTP code returned for type PatchMeetupIDNotFound
const PatchMeetupIDNotFoundCode int = 404

/*PatchMeetupIDNotFound Not found

swagger:response patchMeetupIdNotFound
*/
type PatchMeetupIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDNotFound creates PatchMeetupIDNotFound with default headers values
func NewPatchMeetupIDNotFound() *PatchMeetupIDNotFound {

	return &PatchMeetupIDNotFound{}
}

// WithPayload adds the payload to the patch meetup Id not found response
func (o *PatchMeetupIDNotFound) WithPayload(payload *models.Error) *PatchMeetupIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id not found response
func (o *PatchMeetupIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
