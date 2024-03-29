// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayV4Response gateway v4 response
// swagger:model GatewayV4Response
type GatewayV4Response struct {

	// gateway signing public key
	GatewaySigningPublicKey string `json:"gatewaySigningPublicKey,omitempty"`

	// Knox gateway type
	// Enum: [CENTRAL INDIVIDUAL]
	GatewayType string `json:"gatewayType,omitempty"`

	// Knox SSO type
	// Enum: [SSO_PROVIDER SSO_PROVIDER_FROM_UMS NONE PAM]
	SsoType string `json:"ssoType,omitempty"`

	// SSO Provider certificate
	TokenCert string `json:"tokenCert,omitempty"`

	// Topology definitions of the gateway.
	Topologies []*GatewayTopologyV4Response `json:"topologies"`
}

// Validate validates this gateway v4 response
func (m *GatewayV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewayType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gatewayV4ResponseTypeGatewayTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CENTRAL","INDIVIDUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayV4ResponseTypeGatewayTypePropEnum = append(gatewayV4ResponseTypeGatewayTypePropEnum, v)
	}
}

const (

	// GatewayV4ResponseGatewayTypeCENTRAL captures enum value "CENTRAL"
	GatewayV4ResponseGatewayTypeCENTRAL string = "CENTRAL"

	// GatewayV4ResponseGatewayTypeINDIVIDUAL captures enum value "INDIVIDUAL"
	GatewayV4ResponseGatewayTypeINDIVIDUAL string = "INDIVIDUAL"
)

// prop value enum
func (m *GatewayV4Response) validateGatewayTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gatewayV4ResponseTypeGatewayTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GatewayV4Response) validateGatewayType(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGatewayTypeEnum("gatewayType", "body", m.GatewayType); err != nil {
		return err
	}

	return nil
}

var gatewayV4ResponseTypeSsoTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSO_PROVIDER","SSO_PROVIDER_FROM_UMS","NONE","PAM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gatewayV4ResponseTypeSsoTypePropEnum = append(gatewayV4ResponseTypeSsoTypePropEnum, v)
	}
}

const (

	// GatewayV4ResponseSsoTypeSSOPROVIDER captures enum value "SSO_PROVIDER"
	GatewayV4ResponseSsoTypeSSOPROVIDER string = "SSO_PROVIDER"

	// GatewayV4ResponseSsoTypeSSOPROVIDERFROMUMS captures enum value "SSO_PROVIDER_FROM_UMS"
	GatewayV4ResponseSsoTypeSSOPROVIDERFROMUMS string = "SSO_PROVIDER_FROM_UMS"

	// GatewayV4ResponseSsoTypeNONE captures enum value "NONE"
	GatewayV4ResponseSsoTypeNONE string = "NONE"

	// GatewayV4ResponseSsoTypePAM captures enum value "PAM"
	GatewayV4ResponseSsoTypePAM string = "PAM"
)

// prop value enum
func (m *GatewayV4Response) validateSsoTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, gatewayV4ResponseTypeSsoTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GatewayV4Response) validateSsoType(formats strfmt.Registry) error {

	if swag.IsZero(m.SsoType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSsoTypeEnum("ssoType", "body", m.SsoType); err != nil {
		return err
	}

	return nil
}

func (m *GatewayV4Response) validateTopologies(formats strfmt.Registry) error {

	if swag.IsZero(m.Topologies) { // not required
		return nil
	}

	for i := 0; i < len(m.Topologies); i++ {
		if swag.IsZero(m.Topologies[i]) { // not required
			continue
		}

		if m.Topologies[i] != nil {
			if err := m.Topologies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayV4Response) UnmarshalBinary(b []byte) error {
	var res GatewayV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
