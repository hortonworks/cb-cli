// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomConfigurationsV4Response custom configurations v4 response
// swagger:model CustomConfigurationsV4Response
type CustomConfigurationsV4Response struct {

	// account
	// Required: true
	Account *string `json:"account"`

	// list of properties
	// Required: true
	// Unique: true
	Configurations []*CustomConfigurationPropertyParameters `json:"configurations"`

	// created
	// Required: true
	Created *int64 `json:"created"`

	// unique crn of the custom configs
	// Required: true
	Crn *string `json:"crn"`

	// unique name of the custom configs
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// Runtime version that custom configs point to
	RuntimeVersion string `json:"runtimeVersion,omitempty"`
}

// Validate validates this custom configurations v4 response
func (m *CustomConfigurationsV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfigurationsV4Response) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	return nil
}

func (m *CustomConfigurationsV4Response) validateConfigurations(formats strfmt.Registry) error {

	if err := validate.Required("configurations", "body", m.Configurations); err != nil {
		return err
	}

	if err := validate.UniqueItems("configurations", "body", m.Configurations); err != nil {
		return err
	}

	for i := 0; i < len(m.Configurations); i++ {
		if swag.IsZero(m.Configurations[i]) { // not required
			continue
		}

		if m.Configurations[i] != nil {
			if err := m.Configurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomConfigurationsV4Response) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *CustomConfigurationsV4Response) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *CustomConfigurationsV4Response) validateName(formats strfmt.Registry) error {

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
func (m *CustomConfigurationsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomConfigurationsV4Response) UnmarshalBinary(b []byte) error {
	var res CustomConfigurationsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
