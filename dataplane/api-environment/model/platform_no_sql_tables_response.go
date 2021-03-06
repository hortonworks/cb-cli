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

// PlatformNoSQLTablesResponse platform no Sql tables response
// swagger:model PlatformNoSqlTablesResponse
type PlatformNoSQLTablesResponse struct {

	// no Sql tables
	NoSQLTables []*PlatformNoSQLTableResponse `json:"noSqlTables"`
}

// Validate validates this platform no Sql tables response
func (m *PlatformNoSQLTablesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoSQLTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformNoSQLTablesResponse) validateNoSQLTables(formats strfmt.Registry) error {

	if swag.IsZero(m.NoSQLTables) { // not required
		return nil
	}

	for i := 0; i < len(m.NoSQLTables); i++ {
		if swag.IsZero(m.NoSQLTables[i]) { // not required
			continue
		}

		if m.NoSQLTables[i] != nil {
			if err := m.NoSQLTables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("noSqlTables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformNoSQLTablesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformNoSQLTablesResponse) UnmarshalBinary(b []byte) error {
	var res PlatformNoSQLTablesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
