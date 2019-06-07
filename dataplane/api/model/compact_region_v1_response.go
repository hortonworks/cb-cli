// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CompactRegionV1Response compact region v1 response
// swagger:model CompactRegionV1Response
type CompactRegionV1Response struct {

	// regions with displayNames
	DisplayNames map[string]string `json:"displayNames,omitempty"`

	// Regions of the environment.
	// Unique: true
	Names []string `json:"names"`
}

// Validate validates this compact region v1 response
func (m *CompactRegionV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompactRegionV1Response) validateNames(formats strfmt.Registry) error {

	if swag.IsZero(m.Names) { // not required
		return nil
	}

	if err := validate.UniqueItems("names", "body", m.Names); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompactRegionV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompactRegionV1Response) UnmarshalBinary(b []byte) error {
	var res CompactRegionV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
