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

// DatabaseTestV4Response database test v4 response
// swagger:model DatabaseTestV4Response
type DatabaseTestV4Response struct {

	// Result of database connection test
	// Required: true
	Result *string `json:"result"`
}

// Validate validates this database test v4 response
func (m *DatabaseTestV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTestV4Response) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTestV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTestV4Response) UnmarshalBinary(b []byte) error {
	var res DatabaseTestV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
