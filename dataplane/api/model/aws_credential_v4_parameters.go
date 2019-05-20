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

// AwsCredentialV4Parameters aws credential v4 parameters
// swagger:model AwsCredentialV4Parameters
type AwsCredentialV4Parameters struct {

	// gov cloud
	// Required: true
	GovCloud *bool `json:"govCloud"`

	// key based
	KeyBased *KeyBasedCredentialParameters `json:"keyBased,omitempty"`

	// role based
	RoleBased *RoleBasedCredentialParameters `json:"roleBased,omitempty"`
}

// Validate validates this aws credential v4 parameters
func (m *AwsCredentialV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGovCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyBased(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleBased(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCredentialV4Parameters) validateGovCloud(formats strfmt.Registry) error {

	if err := validate.Required("govCloud", "body", m.GovCloud); err != nil {
		return err
	}

	return nil
}

func (m *AwsCredentialV4Parameters) validateKeyBased(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyBased) { // not required
		return nil
	}

	if m.KeyBased != nil {
		if err := m.KeyBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyBased")
			}
			return err
		}
	}

	return nil
}

func (m *AwsCredentialV4Parameters) validateRoleBased(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleBased) { // not required
		return nil
	}

	if m.RoleBased != nil {
		if err := m.RoleBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleBased")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsCredentialV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCredentialV4Parameters) UnmarshalBinary(b []byte) error {
	var res AwsCredentialV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
