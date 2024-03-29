// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Versions versions
// swagger:model Versions
type Versions struct {

	// freeipa
	// Read Only: true
	Freeipa []*FreeIpaVersions `json:"freeipa"`
}

// Validate validates this versions
func (m *Versions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFreeipa(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Versions) validateFreeipa(formats strfmt.Registry) error {

	if swag.IsZero(m.Freeipa) { // not required
		return nil
	}

	for i := 0; i < len(m.Freeipa); i++ {
		if swag.IsZero(m.Freeipa[i]) { // not required
			continue
		}

		if m.Freeipa[i] != nil {
			if err := m.Freeipa[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("freeipa" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Versions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Versions) UnmarshalBinary(b []byte) error {
	var res Versions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
