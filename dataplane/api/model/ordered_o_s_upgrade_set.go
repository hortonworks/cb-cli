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

// OrderedOSUpgradeSet ordered o s upgrade set
// swagger:model OrderedOSUpgradeSet
type OrderedOSUpgradeSet struct {

	// instance ids
	// Unique: true
	InstanceIds []string `json:"instanceIds"`

	// order
	Order int32 `json:"order,omitempty"`
}

// Validate validates this ordered o s upgrade set
func (m *OrderedOSUpgradeSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderedOSUpgradeSet) validateInstanceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("instanceIds", "body", m.InstanceIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderedOSUpgradeSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderedOSUpgradeSet) UnmarshalBinary(b []byte) error {
	var res OrderedOSUpgradeSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
