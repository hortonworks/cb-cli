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

// AzureMarketplaceTermsRequest azure marketplace terms request
// swagger:model AzureMarketplaceTermsRequest
type AzureMarketplaceTermsRequest struct {

	// accepted
	// Required: true
	Accepted *bool `json:"accepted"`
}

// Validate validates this azure marketplace terms request
func (m *AzureMarketplaceTermsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccepted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureMarketplaceTermsRequest) validateAccepted(formats strfmt.Registry) error {

	if err := validate.Required("accepted", "body", m.Accepted); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureMarketplaceTermsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureMarketplaceTermsRequest) UnmarshalBinary(b []byte) error {
	var res AzureMarketplaceTermsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
