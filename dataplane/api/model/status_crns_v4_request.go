// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StatusCrnsV4Request status crns v4 request
// swagger:model StatusCrnsV4Request
type StatusCrnsV4Request struct {

	// the unique crns of the resource
	Crns []string `json:"crns"`
}

// Validate validates this status crns v4 request
func (m *StatusCrnsV4Request) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusCrnsV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusCrnsV4Request) UnmarshalBinary(b []byte) error {
	var res StatusCrnsV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
