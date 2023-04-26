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

// DatabasePropertiesV4Response Response for the database properties of a database server
// swagger:model DatabasePropertiesV4Response
type DatabasePropertiesV4Response struct {

	// The format of the username for the database connection
	// Required: true
	// Enum: [USERNAME_ONLY USERNAME_WITH_HOSTNAME]
	ConnectionNameFormat *string `json:"connectionNameFormat"`
}

// Validate validates this database properties v4 response
func (m *DatabasePropertiesV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionNameFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var databasePropertiesV4ResponseTypeConnectionNameFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USERNAME_ONLY","USERNAME_WITH_HOSTNAME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		databasePropertiesV4ResponseTypeConnectionNameFormatPropEnum = append(databasePropertiesV4ResponseTypeConnectionNameFormatPropEnum, v)
	}
}

const (

	// DatabasePropertiesV4ResponseConnectionNameFormatUSERNAMEONLY captures enum value "USERNAME_ONLY"
	DatabasePropertiesV4ResponseConnectionNameFormatUSERNAMEONLY string = "USERNAME_ONLY"

	// DatabasePropertiesV4ResponseConnectionNameFormatUSERNAMEWITHHOSTNAME captures enum value "USERNAME_WITH_HOSTNAME"
	DatabasePropertiesV4ResponseConnectionNameFormatUSERNAMEWITHHOSTNAME string = "USERNAME_WITH_HOSTNAME"
)

// prop value enum
func (m *DatabasePropertiesV4Response) validateConnectionNameFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, databasePropertiesV4ResponseTypeConnectionNameFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DatabasePropertiesV4Response) validateConnectionNameFormat(formats strfmt.Registry) error {

	if err := validate.Required("connectionNameFormat", "body", m.ConnectionNameFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateConnectionNameFormatEnum("connectionNameFormat", "body", *m.ConnectionNameFormat); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabasePropertiesV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabasePropertiesV4Response) UnmarshalBinary(b []byte) error {
	var res DatabasePropertiesV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
