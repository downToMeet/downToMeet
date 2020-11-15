// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// User user
//
// swagger:model user
type User struct {

	// attending
	Attending []MeetupID `json:"attending,omitempty"`

	// connections
	Connections []string `json:"connections,omitempty"`

	// contact info
	ContactInfo string `json:"contactInfo,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	// Required: true
	ID UserID `json:"id"`

	// interests
	Interests []string `json:"interests,omitempty"`

	// location
	Location *Coordinates `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owned meetups
	OwnedMeetups []MeetupID `json:"ownedMeetups,omitempty"`

	// pending approval
	PendingApproval []MeetupID `json:"pendingApproval,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttending(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnedMeetups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingApproval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAttending(formats strfmt.Registry) error {

	if swag.IsZero(m.Attending) { // not required
		return nil
	}

	for i := 0; i < len(m.Attending); i++ {

		if err := m.Attending[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attending" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *User) validateLocation(formats strfmt.Registry) error {

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

func (m *User) validateOwnedMeetups(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnedMeetups) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnedMeetups); i++ {

		if err := m.OwnedMeetups[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownedMeetups" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *User) validatePendingApproval(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingApproval) { // not required
		return nil
	}

	for i := 0; i < len(m.PendingApproval); i++ {

		if err := m.PendingApproval[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pendingApproval" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
