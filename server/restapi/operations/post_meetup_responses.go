// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go.timothygu.me/downtomeet/server/models"
)

// PostMeetupCreatedCode is the HTTP code returned for type PostMeetupCreated
const PostMeetupCreatedCode int = 201

/*PostMeetupCreated Created

swagger:response postMeetupCreated
*/
type PostMeetupCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Meetup `json:"body,omitempty"`
}

// NewPostMeetupCreated creates PostMeetupCreated with default headers values
func NewPostMeetupCreated() *PostMeetupCreated {

	return &PostMeetupCreated{}
}

// WithPayload adds the payload to the post meetup created response
func (o *PostMeetupCreated) WithPayload(payload *models.Meetup) *PostMeetupCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post meetup created response
func (o *PostMeetupCreated) SetPayload(payload *models.Meetup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMeetupCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMeetupBadRequestCode is the HTTP code returned for type PostMeetupBadRequest
const PostMeetupBadRequestCode int = 400

/*PostMeetupBadRequest Bad Request

swagger:response postMeetupBadRequest
*/
type PostMeetupBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMeetupBadRequest creates PostMeetupBadRequest with default headers values
func NewPostMeetupBadRequest() *PostMeetupBadRequest {

	return &PostMeetupBadRequest{}
}

// WithPayload adds the payload to the post meetup bad request response
func (o *PostMeetupBadRequest) WithPayload(payload *models.Error) *PostMeetupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post meetup bad request response
func (o *PostMeetupBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMeetupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
