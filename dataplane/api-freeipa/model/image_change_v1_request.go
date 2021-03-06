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

// ImageChangeV1Request image change v1 request
// swagger:model ImageChangeV1Request
type ImageChangeV1Request struct {

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// image settings
	// Required: true
	ImageSettings *ImageSettingsV1Request `json:"imageSettings"`
}

// Validate validates this image change v1 request
func (m *ImageChangeV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageChangeV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *ImageChangeV1Request) validateImageSettings(formats strfmt.Registry) error {

	if err := validate.Required("imageSettings", "body", m.ImageSettings); err != nil {
		return err
	}

	if m.ImageSettings != nil {
		if err := m.ImageSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageChangeV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageChangeV1Request) UnmarshalBinary(b []byte) error {
	var res ImageChangeV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
