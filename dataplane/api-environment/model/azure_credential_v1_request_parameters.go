// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AzureCredentialV1RequestParameters azure credential v1 request parameters
// swagger:model AzureCredentialV1RequestParameters
type AzureCredentialV1RequestParameters struct {

	// app based
	AppBased *AppBasedV1Request `json:"appBased,omitempty"`

	// role based
	RoleBased *RoleBasedV1Request `json:"roleBased,omitempty"`

	// subscription Id
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this azure credential v1 request parameters
func (m *AzureCredentialV1RequestParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppBased(formats); err != nil {
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

func (m *AzureCredentialV1RequestParameters) validateAppBased(formats strfmt.Registry) error {

	if swag.IsZero(m.AppBased) { // not required
		return nil
	}

	if m.AppBased != nil {
		if err := m.AppBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

func (m *AzureCredentialV1RequestParameters) validateRoleBased(formats strfmt.Registry) error {

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
func (m *AzureCredentialV1RequestParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCredentialV1RequestParameters) UnmarshalBinary(b []byte) error {
	var res AzureCredentialV1RequestParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
