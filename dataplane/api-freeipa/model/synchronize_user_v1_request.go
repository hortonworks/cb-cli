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

// SynchronizeUserV1Request synchronize user v1 request
// swagger:model SynchronizeUserV1Request
type SynchronizeUserV1Request struct {

	// Optional environment crns to sync
	// Unique: true
	Environments []string `json:"environments"`
}

// Validate validates this synchronize user v1 request
func (m *SynchronizeUserV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SynchronizeUserV1Request) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	if err := validate.UniqueItems("environments", "body", m.Environments); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SynchronizeUserV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SynchronizeUserV1Request) UnmarshalBinary(b []byte) error {
	var res SynchronizeUserV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
