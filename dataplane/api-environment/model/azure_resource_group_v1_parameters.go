// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureResourceGroupV1Parameters azure resource group v1 parameters
// swagger:model AzureResourceGroupV1Parameters
type AzureResourceGroupV1Parameters struct {

	// Name of an existing azure resource group.
	Name string `json:"name,omitempty"`

	// Resource group usage: single resource group for all resources where possible or use multiple resource groups.
	// Enum: [SINGLE MULTIPLE]
	ResourceGroupUsage string `json:"resourceGroupUsage,omitempty"`
}

// Validate validates this azure resource group v1 parameters
func (m *AzureResourceGroupV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceGroupUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var azureResourceGroupV1ParametersTypeResourceGroupUsagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE","MULTIPLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureResourceGroupV1ParametersTypeResourceGroupUsagePropEnum = append(azureResourceGroupV1ParametersTypeResourceGroupUsagePropEnum, v)
	}
}

const (

	// AzureResourceGroupV1ParametersResourceGroupUsageSINGLE captures enum value "SINGLE"
	AzureResourceGroupV1ParametersResourceGroupUsageSINGLE string = "SINGLE"

	// AzureResourceGroupV1ParametersResourceGroupUsageMULTIPLE captures enum value "MULTIPLE"
	AzureResourceGroupV1ParametersResourceGroupUsageMULTIPLE string = "MULTIPLE"
)

// prop value enum
func (m *AzureResourceGroupV1Parameters) validateResourceGroupUsageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, azureResourceGroupV1ParametersTypeResourceGroupUsagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AzureResourceGroupV1Parameters) validateResourceGroupUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroupUsage) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceGroupUsageEnum("resourceGroupUsage", "body", m.ResourceGroupUsage); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureResourceGroupV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureResourceGroupV1Parameters) UnmarshalBinary(b []byte) error {
	var res AzureResourceGroupV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
