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

// CredentialV1Response credential v1 response
// swagger:model CredentialV1Response
type CredentialV1Response struct {
	CredentialBase

	// provider specific identifier of the account/subscription/project that is used by Cloudbreak
	AccountID string `json:"accountId,omitempty"`

	// provider specific attributes of the credential
	Attributes *SecretResponse `json:"attributes,omitempty"`

	// custom parameters for Azure credential
	Azure *AzureCredentialV1ResponseParameters `json:"azure,omitempty"`

	// creation time of the credential in long
	Created int64 `json:"created,omitempty"`

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// global identifiers of the resource
	Crn string `json:"crn,omitempty"`

	// gov cloud
	GovCloud bool `json:"govCloud,omitempty"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// type of credential
	// Required: true
	// Enum: [ENVIRONMENT AUDIT]
	Type *string `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CredentialV1Response) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CredentialBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CredentialBase = aO0

	// AO1
	var dataAO1 struct {
		AccountID string `json:"accountId,omitempty"`

		Attributes *SecretResponse `json:"attributes,omitempty"`

		Azure *AzureCredentialV1ResponseParameters `json:"azure,omitempty"`

		Created int64 `json:"created,omitempty"`

		Creator string `json:"creator,omitempty"`

		Crn string `json:"crn,omitempty"`

		GovCloud bool `json:"govCloud,omitempty"`

		Name *string `json:"name"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AccountID = dataAO1.AccountID

	m.Attributes = dataAO1.Attributes

	m.Azure = dataAO1.Azure

	m.Created = dataAO1.Created

	m.Creator = dataAO1.Creator

	m.Crn = dataAO1.Crn

	m.GovCloud = dataAO1.GovCloud

	m.Name = dataAO1.Name

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CredentialV1Response) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CredentialBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AccountID string `json:"accountId,omitempty"`

		Attributes *SecretResponse `json:"attributes,omitempty"`

		Azure *AzureCredentialV1ResponseParameters `json:"azure,omitempty"`

		Created int64 `json:"created,omitempty"`

		Creator string `json:"creator,omitempty"`

		Crn string `json:"crn,omitempty"`

		GovCloud bool `json:"govCloud,omitempty"`

		Name *string `json:"name"`

		Type *string `json:"type"`
	}

	dataAO1.AccountID = m.AccountID

	dataAO1.Attributes = m.Attributes

	dataAO1.Azure = m.Azure

	dataAO1.Created = m.Created

	dataAO1.Creator = m.Creator

	dataAO1.Crn = m.Crn

	dataAO1.GovCloud = m.GovCloud

	dataAO1.Name = m.Name

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this credential v1 response
func (m *CredentialV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CredentialBase
	if err := m.CredentialBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialV1Response) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialV1Response) validateAzure(formats strfmt.Registry) error {

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

func (m *CredentialV1Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var credentialV1ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENVIRONMENT","AUDIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		credentialV1ResponseTypeTypePropEnum = append(credentialV1ResponseTypeTypePropEnum, v)
	}
}

// property enum
func (m *CredentialV1Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, credentialV1ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CredentialV1Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialV1Response) UnmarshalBinary(b []byte) error {
	var res CredentialV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
