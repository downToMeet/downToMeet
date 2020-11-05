// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetMeetupParams creates a new GetMeetupParams object
// no default values defined in spec.
func NewGetMeetupParams() GetMeetupParams {

	return GetMeetupParams{}
}

// GetMeetupParams contains all the bound params for the get meetup operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMeetup
type GetMeetupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The latitude of the center of search
	  Required: true
	  In: query
	*/
	Lat float64
	/*The longitude of the center of search
	  Required: true
	  In: query
	*/
	Lon float64
	/*Desired search radius (kilometers)
	  Required: true
	  Minimum: 0
	  In: query
	*/
	Radius float64
	/*Interests to search for
	  Required: true
	  In: query
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMeetupParams() beforehand.
func (o *GetMeetupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLat, qhkLat, _ := qs.GetOK("lat")
	if err := o.bindLat(qLat, qhkLat, route.Formats); err != nil {
		res = append(res, err)
	}

	qLon, qhkLon, _ := qs.GetOK("lon")
	if err := o.bindLon(qLon, qhkLon, route.Formats); err != nil {
		res = append(res, err)
	}

	qRadius, qhkRadius, _ := qs.GetOK("radius")
	if err := o.bindRadius(qRadius, qhkRadius, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLat binds and validates parameter Lat from query.
func (o *GetMeetupParams) bindLat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("lat", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("lat", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("lat", "query", "float64", raw)
	}
	o.Lat = value

	return nil
}

// bindLon binds and validates parameter Lon from query.
func (o *GetMeetupParams) bindLon(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("lon", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("lon", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("lon", "query", "float64", raw)
	}
	o.Lon = value

	return nil
}

// bindRadius binds and validates parameter Radius from query.
func (o *GetMeetupParams) bindRadius(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("radius", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("radius", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("radius", "query", "float64", raw)
	}
	o.Radius = value

	if err := o.validateRadius(formats); err != nil {
		return err
	}

	return nil
}

// validateRadius carries on validations for parameter Radius
func (o *GetMeetupParams) validateRadius(formats strfmt.Registry) error {

	if err := validate.Minimum("radius", "query", float64(o.Radius), 0, false); err != nil {
		return err
	}

	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetMeetupParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("tags", "query", rawData)
	}

	var qvTags string
	if len(rawData) > 0 {
		qvTags = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	tagsIC := swag.SplitByFormat(qvTags, "")

	if len(tagsIC) == 0 {
		return errors.Required("tags", "query", tagsIC)
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
