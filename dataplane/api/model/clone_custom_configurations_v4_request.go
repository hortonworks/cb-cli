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

// CloneCustomConfigurationsV4Request clone custom configurations v4 request
// swagger:model CloneCustomConfigurationsV4Request
type CloneCustomConfigurationsV4Request struct {

	// unique name of the custom configs
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// Runtime version that custom configs point to
	RuntimeVersion string `json:"runtimeVersion,omitempty"`
}

// Validate validates this clone custom configurations v4 request
func (m *CloneCustomConfigurationsV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloneCustomConfigurationsV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[^;\/%]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloneCustomConfigurationsV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloneCustomConfigurationsV4Request) UnmarshalBinary(b []byte) error {
	var res CloneCustomConfigurationsV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
