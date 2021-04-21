// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClientTestBaseV1Response client test base v1 response
// swagger:model ClientTestBaseV1Response
type ClientTestBaseV1Response struct {

	// Result of the check.
	Result bool `json:"result,omitempty"`
}

// Validate validates this client test base v1 response
func (m *ClientTestBaseV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientTestBaseV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientTestBaseV1Response) UnmarshalBinary(b []byte) error {
	var res ClientTestBaseV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
