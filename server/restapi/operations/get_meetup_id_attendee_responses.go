// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// GetMeetupIDAttendeeOKCode is the HTTP code returned for type GetMeetupIDAttendeeOK
const GetMeetupIDAttendeeOKCode int = 200

/*GetMeetupIDAttendeeOK OK

swagger:response getMeetupIdAttendeeOK
*/
type GetMeetupIDAttendeeOK struct {

	/*
	  In: Body
	*/
	Payload *models.AttendeeList `json:"body,omitempty"`
}

// NewGetMeetupIDAttendeeOK creates GetMeetupIDAttendeeOK with default headers values
func NewGetMeetupIDAttendeeOK() *GetMeetupIDAttendeeOK {

	return &GetMeetupIDAttendeeOK{}
}

// WithPayload adds the payload to the get meetup Id attendee o k response
func (o *GetMeetupIDAttendeeOK) WithPayload(payload *models.AttendeeList) *GetMeetupIDAttendeeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meetup Id attendee o k response
func (o *GetMeetupIDAttendeeOK) SetPayload(payload *models.AttendeeList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeetupIDAttendeeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMeetupIDAttendeeBadRequestCode is the HTTP code returned for type GetMeetupIDAttendeeBadRequest
const GetMeetupIDAttendeeBadRequestCode int = 400

/*GetMeetupIDAttendeeBadRequest Bad Request

swagger:response getMeetupIdAttendeeBadRequest
*/
type GetMeetupIDAttendeeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMeetupIDAttendeeBadRequest creates GetMeetupIDAttendeeBadRequest with default headers values
func NewGetMeetupIDAttendeeBadRequest() *GetMeetupIDAttendeeBadRequest {

	return &GetMeetupIDAttendeeBadRequest{}
}

// WithPayload adds the payload to the get meetup Id attendee bad request response
func (o *GetMeetupIDAttendeeBadRequest) WithPayload(payload *models.Error) *GetMeetupIDAttendeeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meetup Id attendee bad request response
func (o *GetMeetupIDAttendeeBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeetupIDAttendeeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMeetupIDAttendeeForbiddenCode is the HTTP code returned for type GetMeetupIDAttendeeForbidden
const GetMeetupIDAttendeeForbiddenCode int = 403

/*GetMeetupIDAttendeeForbidden Forbidden

swagger:response getMeetupIdAttendeeForbidden
*/
type GetMeetupIDAttendeeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMeetupIDAttendeeForbidden creates GetMeetupIDAttendeeForbidden with default headers values
func NewGetMeetupIDAttendeeForbidden() *GetMeetupIDAttendeeForbidden {

	return &GetMeetupIDAttendeeForbidden{}
}

// WithPayload adds the payload to the get meetup Id attendee forbidden response
func (o *GetMeetupIDAttendeeForbidden) WithPayload(payload *models.Error) *GetMeetupIDAttendeeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meetup Id attendee forbidden response
func (o *GetMeetupIDAttendeeForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeetupIDAttendeeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMeetupIDAttendeeNotFoundCode is the HTTP code returned for type GetMeetupIDAttendeeNotFound
const GetMeetupIDAttendeeNotFoundCode int = 404

/*GetMeetupIDAttendeeNotFound Not found

swagger:response getMeetupIdAttendeeNotFound
*/
type GetMeetupIDAttendeeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMeetupIDAttendeeNotFound creates GetMeetupIDAttendeeNotFound with default headers values
func NewGetMeetupIDAttendeeNotFound() *GetMeetupIDAttendeeNotFound {

	return &GetMeetupIDAttendeeNotFound{}
}

// WithPayload adds the payload to the get meetup Id attendee not found response
func (o *GetMeetupIDAttendeeNotFound) WithPayload(payload *models.Error) *GetMeetupIDAttendeeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meetup Id attendee not found response
func (o *GetMeetupIDAttendeeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeetupIDAttendeeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}