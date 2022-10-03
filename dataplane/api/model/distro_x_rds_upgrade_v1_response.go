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

// DistroXRdsUpgradeV1Response distro x rds upgrade v1 response
// swagger:model DistroXRdsUpgradeV1Response
type DistroXRdsUpgradeV1Response struct {

	// flow identifier
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// target version
	// Enum: [VERSION_11]
	TargetVersion string `json:"targetVersion,omitempty"`
}

// Validate validates this distro x rds upgrade v1 response
func (m *DistroXRdsUpgradeV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowIdentifier(formats); err != nil {
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

func (m *DistroXRdsUpgradeV1Response) validateFlowIdentifier(formats strfmt.Registry) error {

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

var distroXRdsUpgradeV1ResponseTypeTargetVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_11"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		distroXRdsUpgradeV1ResponseTypeTargetVersionPropEnum = append(distroXRdsUpgradeV1ResponseTypeTargetVersionPropEnum, v)
	}
}

const (

	// DistroXRdsUpgradeV1ResponseTargetVersionVERSION11 captures enum value "VERSION_11"
	DistroXRdsUpgradeV1ResponseTargetVersionVERSION11 string = "VERSION_11"
)

// prop value enum
func (m *DistroXRdsUpgradeV1Response) validateTargetVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, distroXRdsUpgradeV1ResponseTypeTargetVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DistroXRdsUpgradeV1Response) validateTargetVersion(formats strfmt.Registry) error {

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
func (m *DistroXRdsUpgradeV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistroXRdsUpgradeV1Response) UnmarshalBinary(b []byte) error {
	var res DistroXRdsUpgradeV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
