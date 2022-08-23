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

// RdsUpgradeV4Response rds upgrade v4 response
// swagger:model RdsUpgradeV4Response
type RdsUpgradeV4Response struct {

	// flow identifier
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// response type
	// Enum: [TRIGGERED SKIP ERROR]
	ResponseType string `json:"responseType,omitempty"`

	// target version
	// Enum: [VERSION_11]
	TargetVersion string `json:"targetVersion,omitempty"`
}

// Validate validates this rds upgrade v4 response
func (m *RdsUpgradeV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsUpgradeV4Response) validateFlowIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowIdentifier) { // not required
		return nil
	}

	if m.FlowIdentifier != nil {
		if err := m.FlowIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowIdentifier")
			}
			return err
		}
	}

	return nil
}

var rdsUpgradeV4ResponseTypeResponseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRIGGERED","SKIP","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rdsUpgradeV4ResponseTypeResponseTypePropEnum = append(rdsUpgradeV4ResponseTypeResponseTypePropEnum, v)
	}
}

const (

	// RdsUpgradeV4ResponseResponseTypeTRIGGERED captures enum value "TRIGGERED"
	RdsUpgradeV4ResponseResponseTypeTRIGGERED string = "TRIGGERED"

	// RdsUpgradeV4ResponseResponseTypeSKIP captures enum value "SKIP"
	RdsUpgradeV4ResponseResponseTypeSKIP string = "SKIP"

	// RdsUpgradeV4ResponseResponseTypeERROR captures enum value "ERROR"
	RdsUpgradeV4ResponseResponseTypeERROR string = "ERROR"
)

// prop value enum
func (m *RdsUpgradeV4Response) validateResponseTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rdsUpgradeV4ResponseTypeResponseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RdsUpgradeV4Response) validateResponseType(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResponseTypeEnum("responseType", "body", m.ResponseType); err != nil {
		return err
	}

	return nil
}

var rdsUpgradeV4ResponseTypeTargetVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_11"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rdsUpgradeV4ResponseTypeTargetVersionPropEnum = append(rdsUpgradeV4ResponseTypeTargetVersionPropEnum, v)
	}
}

const (

	// RdsUpgradeV4ResponseTargetVersionVERSION11 captures enum value "VERSION_11"
	RdsUpgradeV4ResponseTargetVersionVERSION11 string = "VERSION_11"
)

// prop value enum
func (m *RdsUpgradeV4Response) validateTargetVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rdsUpgradeV4ResponseTypeTargetVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RdsUpgradeV4Response) validateTargetVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetVersionEnum("targetVersion", "body", m.TargetVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsUpgradeV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsUpgradeV4Response) UnmarshalBinary(b []byte) error {
	var res RdsUpgradeV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
