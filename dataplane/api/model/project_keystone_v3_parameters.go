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

// ProjectKeystoneV3Parameters project keystone v3 parameters
// swagger:model ProjectKeystoneV3Parameters
type ProjectKeystoneV3Parameters struct {

	// project domain name
	// Required: true
	ProjectDomainName *string `json:"projectDomainName"`

	// project name
	// Required: true
	ProjectName *string `json:"projectName"`

	// user domain
	// Required: true
	UserDomain *string `json:"userDomain"`
}

// Validate validates this project keystone v3 parameters
func (m *ProjectKeystoneV3Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectKeystoneV3Parameters) validateProjectDomainName(formats strfmt.Registry) error {

	if err := validate.Required("projectDomainName", "body", m.ProjectDomainName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectKeystoneV3Parameters) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("projectName", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectKeystoneV3Parameters) validateUserDomain(formats strfmt.Registry) error {

	if err := validate.Required("userDomain", "body", m.UserDomain); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectKeystoneV3Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectKeystoneV3Parameters) UnmarshalBinary(b []byte) error {
	var res ProjectKeystoneV3Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
