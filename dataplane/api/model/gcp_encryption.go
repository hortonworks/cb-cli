// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GcpEncryption gcp encryption
// swagger:model GcpEncryption
type GcpEncryption struct {

	// encryption key for vm
	Key string `json:"key,omitempty"`

	// encryption method for the key (RAW|RSA)
	KeyEncryptionMethod string `json:"keyEncryptionMethod,omitempty"`

	// encryption type for vm (DEFAULT|CUSTOM)
	Type string `json:"type,omitempty"`
}

// Validate validates this gcp encryption
func (m *GcpEncryption) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpEncryption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpEncryption) UnmarshalBinary(b []byte) error {
	var res GcpEncryption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
