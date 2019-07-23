// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CommonLoggingAttributes common logging attributes
// swagger:model CommonLoggingAttributes
type CommonLoggingAttributes struct {

	// group
	Group string `json:"group,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this common logging attributes
func (m *CommonLoggingAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonLoggingAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonLoggingAttributes) UnmarshalBinary(b []byte) error {
	var res CommonLoggingAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
