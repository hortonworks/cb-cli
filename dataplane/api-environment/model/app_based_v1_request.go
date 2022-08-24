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

// AppBasedV1Request app based v1 request
// swagger:model AppBasedV1Request
type AppBasedV1Request struct {

	// access key
	// Required: true
	AccessKey *string `json:"accessKey"`

	// authentication type
	// Enum: [SECRET CERTIFICATE]
	AuthenticationType string `json:"authenticationType,omitempty"`

	// secret key
	SecretKey *string `json:"secretKey,omitempty"`
}

// Validate validates this app based v1 request
func (m *AppBasedV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppBasedV1Request) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("accessKey", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

var appBasedV1RequestTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SECRET","CERTIFICATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appBasedV1RequestTypeAuthenticationTypePropEnum = append(appBasedV1RequestTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// AppBasedV1RequestAuthenticationTypeSECRET captures enum value "SECRET"
	AppBasedV1RequestAuthenticationTypeSECRET string = "SECRET"

	// AppBasedV1RequestAuthenticationTypeCERTIFICATE captures enum value "CERTIFICATE"
	AppBasedV1RequestAuthenticationTypeCERTIFICATE string = "CERTIFICATE"
)

// prop value enum
func (m *AppBasedV1Request) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, appBasedV1RequestTypeAuthenticationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AppBasedV1Request) validateAuthenticationType(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTypeEnum("authenticationType", "body", m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppBasedV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppBasedV1Request) UnmarshalBinary(b []byte) error {
	var res AppBasedV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
