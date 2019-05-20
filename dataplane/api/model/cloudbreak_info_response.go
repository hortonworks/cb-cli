// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CloudbreakInfoResponse cloudbreak info response
// swagger:model CloudbreakInfoResponse
type CloudbreakInfoResponse struct {

	// info
	Info map[string]interface{} `json:"info,omitempty"`
}

// Validate validates this cloudbreak info response
func (m *CloudbreakInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudbreakInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudbreakInfoResponse) UnmarshalBinary(b []byte) error {
	var res CloudbreakInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
