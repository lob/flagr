// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Flag flag
// swagger:model flag
type Flag struct {

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// enabled data records will get data logging in the metrics pipeline, for example, kafka.
	// Required: true
	DataRecordsEnabled *bool `json:"dataRecordsEnabled"`

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// it will override the entityType in the evaluation logs if it's not empty
	EntityType string `json:"entityType,omitempty"`

	// id
	// Read Only: true
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// unique key representation of the flag
	// Min Length: 1
	Key string `json:"key,omitempty"`

	// flag usage details in markdown format
	Notes string `json:"notes,omitempty"`

	// segments
	Segments []*Segment `json:"segments"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// variants
	Variants []*Variant `json:"variants"`
}

// Validate validates this flag
func (m *Flag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataRecordsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Flag) validateDataRecordsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("dataRecordsEnabled", "body", m.DataRecordsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateKey(formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := validate.MinLength("key", "body", string(m.Key), 1); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateSegments(formats strfmt.Registry) error {

	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for i := 0; i < len(m.Segments); i++ {
		if swag.IsZero(m.Segments[i]) { // not required
			continue
		}

		if m.Segments[i] != nil {
			if err := m.Segments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Flag) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Flag) validateVariants(formats strfmt.Registry) error {

	if swag.IsZero(m.Variants) { // not required
		return nil
	}

	for i := 0; i < len(m.Variants); i++ {
		if swag.IsZero(m.Variants[i]) { // not required
			continue
		}

		if m.Variants[i] != nil {
			if err := m.Variants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Flag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flag) UnmarshalBinary(b []byte) error {
	var res Flag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
