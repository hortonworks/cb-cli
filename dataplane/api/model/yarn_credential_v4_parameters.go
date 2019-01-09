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

// YarnCredentialV4Parameters yarn credential v4 parameters
// swagger:model YarnCredentialV4Parameters
type YarnCredentialV4Parameters struct {

	// ambari user
	// Required: true
	AmbariUser *string `json:"ambariUser"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`
}

// Validate validates this yarn credential v4 parameters
func (m *YarnCredentialV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *YarnCredentialV4Parameters) validateAmbariUser(formats strfmt.Registry) error {

	if err := validate.Required("ambariUser", "body", m.AmbariUser); err != nil {
		return err
	}

	return nil
}

func (m *YarnCredentialV4Parameters) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *YarnCredentialV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *YarnCredentialV4Parameters) UnmarshalBinary(b []byte) error {
	var res YarnCredentialV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
