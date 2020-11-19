// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// PatchMeetupIDAttendeeOKCode is the HTTP code returned for type PatchMeetupIDAttendeeOK
const PatchMeetupIDAttendeeOKCode int = 200

/*PatchMeetupIDAttendeeOK OK

swagger:response patchMeetupIdAttendeeOK
*/
type PatchMeetupIDAttendeeOK struct {

	/*
	  In: Body
	*/
	Payload models.AttendeeStatus `json:"body,omitempty"`
}

// NewPatchMeetupIDAttendeeOK creates PatchMeetupIDAttendeeOK with default headers values
func NewPatchMeetupIDAttendeeOK() *PatchMeetupIDAttendeeOK {

	return &PatchMeetupIDAttendeeOK{}
}

// WithPayload adds the payload to the patch meetup Id attendee o k response
func (o *PatchMeetupIDAttendeeOK) WithPayload(payload models.AttendeeStatus) *PatchMeetupIDAttendeeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id attendee o k response
func (o *PatchMeetupIDAttendeeOK) SetPayload(payload models.AttendeeStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDAttendeeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchMeetupIDAttendeeBadRequestCode is the HTTP code returned for type PatchMeetupIDAttendeeBadRequest
const PatchMeetupIDAttendeeBadRequestCode int = 400

/*PatchMeetupIDAttendeeBadRequest Bad Request

swagger:response patchMeetupIdAttendeeBadRequest
*/
type PatchMeetupIDAttendeeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDAttendeeBadRequest creates PatchMeetupIDAttendeeBadRequest with default headers values
func NewPatchMeetupIDAttendeeBadRequest() *PatchMeetupIDAttendeeBadRequest {

	return &PatchMeetupIDAttendeeBadRequest{}
}

// WithPayload adds the payload to the patch meetup Id attendee bad request response
func (o *PatchMeetupIDAttendeeBadRequest) WithPayload(payload *models.Error) *PatchMeetupIDAttendeeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id attendee bad request response
func (o *PatchMeetupIDAttendeeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDAttendeeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMeetupIDAttendeeForbiddenCode is the HTTP code returned for type PatchMeetupIDAttendeeForbidden
const PatchMeetupIDAttendeeForbiddenCode int = 403

/*PatchMeetupIDAttendeeForbidden Forbidden

swagger:response patchMeetupIdAttendeeForbidden
*/
type PatchMeetupIDAttendeeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDAttendeeForbidden creates PatchMeetupIDAttendeeForbidden with default headers values
func NewPatchMeetupIDAttendeeForbidden() *PatchMeetupIDAttendeeForbidden {

	return &PatchMeetupIDAttendeeForbidden{}
}

// WithPayload adds the payload to the patch meetup Id attendee forbidden response
func (o *PatchMeetupIDAttendeeForbidden) WithPayload(payload *models.Error) *PatchMeetupIDAttendeeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id attendee forbidden response
func (o *PatchMeetupIDAttendeeForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDAttendeeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMeetupIDAttendeeNotFoundCode is the HTTP code returned for type PatchMeetupIDAttendeeNotFound
const PatchMeetupIDAttendeeNotFoundCode int = 404

/*PatchMeetupIDAttendeeNotFound Not found

swagger:response patchMeetupIdAttendeeNotFound
*/
type PatchMeetupIDAttendeeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchMeetupIDAttendeeNotFound creates PatchMeetupIDAttendeeNotFound with default headers values
func NewPatchMeetupIDAttendeeNotFound() *PatchMeetupIDAttendeeNotFound {

	return &PatchMeetupIDAttendeeNotFound{}
}

// WithPayload adds the payload to the patch meetup Id attendee not found response
func (o *PatchMeetupIDAttendeeNotFound) WithPayload(payload *models.Error) *PatchMeetupIDAttendeeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch meetup Id attendee not found response
func (o *PatchMeetupIDAttendeeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMeetupIDAttendeeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
