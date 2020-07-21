// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetRangerCloudIdentityMappingRequest set ranger cloud identity mapping request
// swagger:model SetRangerCloudIdentityMappingRequest
type SetRangerCloudIdentityMappingRequest struct {

	// azure group mapping
	// Required: true
	AzureGroupMapping map[string]string `json:"azureGroupMapping"`

	// azure user mapping
	// Required: true
	AzureUserMapping map[string]string `json:"azureUserMapping"`
}

// Validate validates this set ranger cloud identity mapping request
func (m *SetRangerCloudIdentityMappingRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureGroupMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureUserMapping(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetRangerCloudIdentityMappingRequest) validateAzureGroupMapping(formats strfmt.Registry) error {

	return nil
}

func (m *SetRangerCloudIdentityMappingRequest) validateAzureUserMapping(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *SetRangerCloudIdentityMappingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetRangerCloudIdentityMappingRequest) UnmarshalBinary(b []byte) error {
	var res SetRangerCloudIdentityMappingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
