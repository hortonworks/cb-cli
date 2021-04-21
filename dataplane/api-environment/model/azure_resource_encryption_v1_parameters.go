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

// AzureResourceEncryptionV1Parameters azure resource encryption v1 parameters
// swagger:model AzureResourceEncryptionV1Parameters
type AzureResourceEncryptionV1Parameters struct {

	// Resource Id of the disk encryption set used to encrypt Azure disks.
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// Name of the Azure resource group of the Customer Managed Key to encrypt Azure resources
	EncryptionKeyResourceGroupName string `json:"encryptionKeyResourceGroupName,omitempty"`

	// URL of the Customer Managed Key to encrypt Azure resources
	// Pattern: ^https://[a-zA-Z-][0-9a-zA-Z-]*\.vault\.(azure\.net|azure\.cn|usgovcloudapi\.net|microsoftazure\.de)/keys/[0-9a-zA-Z-]+/[0-9A-Za-z]+
	EncryptionKeyURL string `json:"encryptionKeyUrl,omitempty"`
}

// Validate validates this azure resource encryption v1 parameters
func (m *AzureResourceEncryptionV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionKeyURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureResourceEncryptionV1Parameters) validateEncryptionKeyURL(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionKeyURL) { // not required
		return nil
	}

	if err := validate.Pattern("encryptionKeyUrl", "body", string(m.EncryptionKeyURL), `^https://[a-zA-Z-][0-9a-zA-Z-]*\.vault\.(azure\.net|azure\.cn|usgovcloudapi\.net|microsoftazure\.de)/keys/[0-9a-zA-Z-]+/[0-9A-Za-z]+`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureResourceEncryptionV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureResourceEncryptionV1Parameters) UnmarshalBinary(b []byte) error {
	var res AzureResourceEncryptionV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
