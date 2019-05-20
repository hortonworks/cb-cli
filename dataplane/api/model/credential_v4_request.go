// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CredentialV4Request credential v4 request
// swagger:model CredentialV4Request
type CredentialV4Request struct {
	CredentialV4Base

	// custom parameters for Azure credential
	Azure *AzureCredentialV4RequestParameters `json:"azure,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CredentialV4Request) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CredentialV4Base
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CredentialV4Base = aO0

	// AO1
	var dataAO1 struct {
		Azure *AzureCredentialV4RequestParameters `json:"azure,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Azure = dataAO1.Azure

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CredentialV4Request) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CredentialV4Base)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Azure *AzureCredentialV4RequestParameters `json:"azure,omitempty"`
	}

	dataAO1.Azure = m.Azure

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this credential v4 request
func (m *CredentialV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CredentialV4Base
	if err := m.CredentialV4Base.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialV4Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialV4Request) UnmarshalBinary(b []byte) error {
	var res CredentialV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
