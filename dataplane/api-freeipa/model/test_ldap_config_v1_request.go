// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TestLdapConfigV1Request test ldap config v1 request
// swagger:model TestLdapConfigV1Request
type TestLdapConfigV1Request struct {

	// name of the resource
	ExistingLdapName string `json:"existingLdapName,omitempty"`

	// Request that contains the minimal set of fields to test LDAP connectivity
	Ldap *MinimalLdapConfigV1Request `json:"ldap,omitempty"`
}

// Validate validates this test ldap config v1 request
func (m *TestLdapConfigV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestLdapConfigV1Request) validateLdap(formats strfmt.Registry) error {

	if swag.IsZero(m.Ldap) { // not required
		return nil
	}

	if m.Ldap != nil {
		if err := m.Ldap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestLdapConfigV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestLdapConfigV1Request) UnmarshalBinary(b []byte) error {
	var res TestLdapConfigV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
