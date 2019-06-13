// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateDatabaseV4Request create database v4 request
// swagger:model CreateDatabaseV4Request
type CreateDatabaseV4Request struct {

	// Name of the database configuration resource
	DatabaseName string `json:"databaseName,omitempty"`

	// ID of the environment of the resource
	EnvironmentID string `json:"environmentId,omitempty"`

	// Name of the database server
	ExistingDatabaseServerName string `json:"existingDatabaseServerName,omitempty"`

	// Type of database, aka the service name that will use the database like HIVE, DRUID, SUPERSET, RANGER, etc.
	Type string `json:"type,omitempty"`
}

// Validate validates this create database v4 request
func (m *CreateDatabaseV4Request) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDatabaseV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDatabaseV4Request) UnmarshalBinary(b []byte) error {
	var res CreateDatabaseV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
