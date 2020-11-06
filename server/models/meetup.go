// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Meetup meetup
//
// swagger:model meetup
type Meetup struct {

	// attendees
	Attendees []UserID `json:"attendees"`

	// canceled
	Canceled bool `json:"canceled,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	ID MeetupID `json:"id"`

	// location
	Location *Location `json:"location,omitempty"`

	// max capacity
	MaxCapacity int64 `json:"maxCapacity,omitempty"`

	// min capacity
	// Minimum: 0
	MinCapacity *int64 `json:"minCapacity,omitempty"`

	// owner
	Owner UserID `json:"owner,omitempty"`

	// pending attendees
	PendingAttendees []UserID `json:"pendingAttendees"`

	// rejected
	Rejected bool `json:"rejected,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// time
	Time string `json:"time,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this meetup
func (m *Meetup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttendees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingAttendees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Meetup) validateAttendees(formats strfmt.Registry) error {

	if swag.IsZero(m.Attendees) { // not required
		return nil
	}

	for i := 0; i < len(m.Attendees); i++ {

		if err := m.Attendees[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attendees" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Meetup) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Meetup) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Meetup) validateMinCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.MinCapacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("minCapacity", "body", int64(*m.MinCapacity), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Meetup) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if err := m.Owner.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner")
		}
		return err
	}

	return nil
}

func (m *Meetup) validatePendingAttendees(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingAttendees) { // not required
		return nil
	}

	for i := 0; i < len(m.PendingAttendees); i++ {

		if err := m.PendingAttendees[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingAttendees" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Meetup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Meetup) UnmarshalBinary(b []byte) error {
	var res Meetup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}