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

// SdxAwsSpotParameters sdx aws spot parameters
// swagger:model SdxAwsSpotParameters
type SdxAwsSpotParameters struct {

	// percentage
	// Maximum: 100
	// Minimum: 0
	Percentage *int32 `json:"percentage,omitempty"`
}

// Validate validates this sdx aws spot parameters
func (m *SdxAwsSpotParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxAwsSpotParameters) validatePercentage(formats strfmt.Registry) error {

	if swag.IsZero(m.Percentage) { // not required
		return nil
	}

	if err := validate.MinimumInt("percentage", "body", int64(*m.Percentage), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("percentage", "body", int64(*m.Percentage), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxAwsSpotParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxAwsSpotParameters) UnmarshalBinary(b []byte) error {
	var res SdxAwsSpotParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
