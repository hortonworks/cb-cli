// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MultipleInstanceDeleteRequest multiple instance delete request
// swagger:model MultipleInstanceDeleteRequest
type MultipleInstanceDeleteRequest struct {

	// multiple instance which should be deleted
	Instances []string `json:"instances"`
}

// Validate validates this multiple instance delete request
func (m *MultipleInstanceDeleteRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MultipleInstanceDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultipleInstanceDeleteRequest) UnmarshalBinary(b []byte) error {
	var res MultipleInstanceDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
