// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpgradeOptionsV4Response upgrade options v4 response
// swagger:model UpgradeOptionsV4Response
type UpgradeOptionsV4Response struct {

	// current
	Current *ImageInfoV4Response `json:"current,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// upgrade candidates
	UpgradeCandidates []*ImageInfoV4Response `json:"upgradeCandidates"`
}

// Validate validates this upgrade options v4 response
func (m *UpgradeOptionsV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeCandidates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeOptionsV4Response) validateCurrent(formats strfmt.Registry) error {

	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *UpgradeOptionsV4Response) validateUpgradeCandidates(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeCandidates) { // not required
		return nil
	}

	for i := 0; i < len(m.UpgradeCandidates); i++ {
		if swag.IsZero(m.UpgradeCandidates[i]) { // not required
			continue
		}

		if m.UpgradeCandidates[i] != nil {
			if err := m.UpgradeCandidates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("upgradeCandidates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeOptionsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeOptionsV4Response) UnmarshalBinary(b []byte) error {
	var res UpgradeOptionsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
