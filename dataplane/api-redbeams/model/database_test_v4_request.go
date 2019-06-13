// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DatabaseTestV4Request database test v4 request
// swagger:model DatabaseTestV4Request
type DatabaseTestV4Request struct {

	// Unsaved database config to be tested for connectivity
	Database *DatabaseV4Request `json:"database,omitempty"`

	// Identifiers of saved database config to be tested for connectivity
	ExistingDatabase *DatabaseV4Identifiers `json:"existingDatabase,omitempty"`
}

// Validate validates this database test v4 request
func (m *DatabaseTestV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingDatabase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTestV4Request) validateDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseTestV4Request) validateExistingDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.ExistingDatabase) { // not required
		return nil
	}

	if m.ExistingDatabase != nil {
		if err := m.ExistingDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDatabase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTestV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTestV4Request) UnmarshalBinary(b []byte) error {
	var res DatabaseTestV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
