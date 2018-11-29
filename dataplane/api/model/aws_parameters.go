// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AwsParameters aws parameters
// swagger:model AwsParameters
type AwsParameters struct {

	// should encrypt the vm
	Encrypted *bool `json:"encrypted,omitempty"`

	// encryption for vm
	Encryption *AwsEncryption `json:"encryption,omitempty"`

	// spot price for aws
	SpotPrice float64 `json:"spotPrice,omitempty"`
}

// Validate validates this aws parameters
func (m *AwsParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsParameters) validateEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsParameters) UnmarshalBinary(b []byte) error {
	var res AwsParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
