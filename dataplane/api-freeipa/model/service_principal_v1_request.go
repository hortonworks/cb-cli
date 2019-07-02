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

// ServicePrincipalV1Request service principal v1 request
// swagger:model ServicePrincipalV1Request
type ServicePrincipalV1Request struct {

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Hostname where the service is running
	// Required: true
	ServerHostName *string `json:"serverHostName"`

	// Service requesting keytab
	// Required: true
	ServiceName *string `json:"serviceName"`
}

// Validate validates this service principal v1 request
func (m *ServicePrincipalV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePrincipalV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *ServicePrincipalV1Request) validateServerHostName(formats strfmt.Registry) error {

	if err := validate.Required("serverHostName", "body", m.ServerHostName); err != nil {
		return err
	}

	return nil
}

func (m *ServicePrincipalV1Request) validateServiceName(formats strfmt.Registry) error {

	if err := validate.Required("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicePrincipalV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePrincipalV1Request) UnmarshalBinary(b []byte) error {
	var res ServicePrincipalV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
