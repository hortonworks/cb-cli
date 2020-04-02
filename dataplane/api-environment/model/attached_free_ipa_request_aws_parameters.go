// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AttachedFreeIpaRequestAwsParameters attached free ipa request aws parameters
// swagger:model AttachedFreeIpaRequestAwsParameters
type AttachedFreeIpaRequestAwsParameters struct {

	// Aws spot instance related parameters.
	Spot *AttachedFreeIpaRequestAwsSpotParameters `json:"spot,omitempty"`
}

// Validate validates this attached free ipa request aws parameters
func (m *AttachedFreeIpaRequestAwsParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttachedFreeIpaRequestAwsParameters) validateSpot(formats strfmt.Registry) error {

	if swag.IsZero(m.Spot) { // not required
		return nil
	}

	if m.Spot != nil {
		if err := m.Spot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttachedFreeIpaRequestAwsParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedFreeIpaRequestAwsParameters) UnmarshalBinary(b []byte) error {
	var res AttachedFreeIpaRequestAwsParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
