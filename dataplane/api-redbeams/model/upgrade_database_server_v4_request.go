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

// UpgradeDatabaseServerV4Request Request for upgrading a database server in a provider to a higher major version
// swagger:model UpgradeDatabaseServerV4Request
type UpgradeDatabaseServerV4Request struct {

	// The major version to which the database server should be upgraded
	// Required: true
	// Enum: [VERSION_11]
	UpgradeTargetMajorVersion *string `json:"upgradeTargetMajorVersion"`
}

// Validate validates this upgrade database server v4 request
func (m *UpgradeDatabaseServerV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpgradeTargetMajorVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var upgradeDatabaseServerV4RequestTypeUpgradeTargetMajorVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_11"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		upgradeDatabaseServerV4RequestTypeUpgradeTargetMajorVersionPropEnum = append(upgradeDatabaseServerV4RequestTypeUpgradeTargetMajorVersionPropEnum, v)
	}
}

const (

	// UpgradeDatabaseServerV4RequestUpgradeTargetMajorVersionVERSION11 captures enum value "VERSION_11"
	UpgradeDatabaseServerV4RequestUpgradeTargetMajorVersionVERSION11 string = "VERSION_11"
)

// prop value enum
func (m *UpgradeDatabaseServerV4Request) validateUpgradeTargetMajorVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, upgradeDatabaseServerV4RequestTypeUpgradeTargetMajorVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpgradeDatabaseServerV4Request) validateUpgradeTargetMajorVersion(formats strfmt.Registry) error {

	if err := validate.Required("upgradeTargetMajorVersion", "body", m.UpgradeTargetMajorVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateUpgradeTargetMajorVersionEnum("upgradeTargetMajorVersion", "body", *m.UpgradeTargetMajorVersion); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeDatabaseServerV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeDatabaseServerV4Request) UnmarshalBinary(b []byte) error {
	var res UpgradeDatabaseServerV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
