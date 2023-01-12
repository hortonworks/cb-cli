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

// UpgradeDatabaseServerV4Response Response for upgrading a database server
// swagger:model UpgradeDatabaseServerV4Response
type UpgradeDatabaseServerV4Response struct {

	// The current version of the database server
	// Required: true
	// Enum: [VERSION_FAMILY_9 VERSION_9_6 VERSION_10 VERSION_11 VERSION_12 VERSION_13 VERSION_14]
	CurrentVersion *string `json:"currentVersion"`

	// The id of the flow or flow chain that was triggered as part of the process.
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// The status reason if upgrade flow is not started due to some validation
	Reason string `json:"reason,omitempty"`
}

// Validate validates this upgrade database server v4 response
func (m *UpgradeDatabaseServerV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var upgradeDatabaseServerV4ResponseTypeCurrentVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_FAMILY_9","VERSION_9_6","VERSION_10","VERSION_11","VERSION_12","VERSION_13","VERSION_14"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		upgradeDatabaseServerV4ResponseTypeCurrentVersionPropEnum = append(upgradeDatabaseServerV4ResponseTypeCurrentVersionPropEnum, v)
	}
}

const (

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSIONFAMILY9 captures enum value "VERSION_FAMILY_9"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSIONFAMILY9 string = "VERSION_FAMILY_9"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION96 captures enum value "VERSION_9_6"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION96 string = "VERSION_9_6"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION10 captures enum value "VERSION_10"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION10 string = "VERSION_10"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION11 captures enum value "VERSION_11"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION11 string = "VERSION_11"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION12 captures enum value "VERSION_12"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION12 string = "VERSION_12"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION13 captures enum value "VERSION_13"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION13 string = "VERSION_13"

	// UpgradeDatabaseServerV4ResponseCurrentVersionVERSION14 captures enum value "VERSION_14"
	UpgradeDatabaseServerV4ResponseCurrentVersionVERSION14 string = "VERSION_14"
)

// prop value enum
func (m *UpgradeDatabaseServerV4Response) validateCurrentVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, upgradeDatabaseServerV4ResponseTypeCurrentVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpgradeDatabaseServerV4Response) validateCurrentVersion(formats strfmt.Registry) error {

	if err := validate.Required("currentVersion", "body", m.CurrentVersion); err != nil {
		return err
	}

	// value enum
	if err := m.validateCurrentVersionEnum("currentVersion", "body", *m.CurrentVersion); err != nil {
		return err
	}

	return nil
}

func (m *UpgradeDatabaseServerV4Response) validateFlowIdentifier(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UpgradeDatabaseServerV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeDatabaseServerV4Response) UnmarshalBinary(b []byte) error {
	var res UpgradeDatabaseServerV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
