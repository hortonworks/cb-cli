// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserResponseJSON user response Json
// swagger:model UserResponseJson

type UserResponseJSON struct {

	// id
	ID int64 `json:"id,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

/* polymorph UserResponseJson id false */

/* polymorph UserResponseJson userId false */

/* polymorph UserResponseJson username false */

// Validate validates this user response Json
func (m *UserResponseJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UserResponseJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserResponseJSON) UnmarshalBinary(b []byte) error {
	var res UserResponseJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
