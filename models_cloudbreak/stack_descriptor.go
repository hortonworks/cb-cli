// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackDescriptor stack descriptor
// swagger:model StackDescriptor

type StackDescriptor struct {

	// ambari
	Ambari *AmbariInfoJSON `json:"ambari,omitempty"`

	// min ambari
	MinAmbari string `json:"minAmbari,omitempty"`

	// mpacks
	Mpacks []*ManagementPackEntry `json:"mpacks"`

	// repo
	Repo *StackRepoDetailsJSON `json:"repo,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

/* polymorph StackDescriptor ambari false */

/* polymorph StackDescriptor minAmbari false */

/* polymorph StackDescriptor mpacks false */

/* polymorph StackDescriptor repo false */

/* polymorph StackDescriptor version false */

// Validate validates this stack descriptor
func (m *StackDescriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbari(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMpacks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackDescriptor) validateAmbari(formats strfmt.Registry) error {

	if swag.IsZero(m.Ambari) { // not required
		return nil
	}

	if m.Ambari != nil {

		if err := m.Ambari.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambari")
			}
			return err
		}
	}

	return nil
}

func (m *StackDescriptor) validateMpacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Mpacks) { // not required
		return nil
	}

	for i := 0; i < len(m.Mpacks); i++ {

		if swag.IsZero(m.Mpacks[i]) { // not required
			continue
		}

		if m.Mpacks[i] != nil {

			if err := m.Mpacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mpacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackDescriptor) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackDescriptor) UnmarshalBinary(b []byte) error {
	var res StackDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
