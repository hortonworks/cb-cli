// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AwsDatabaseServerV4Parameters aws database server v4 parameters
// swagger:model AwsDatabaseServerV4Parameters
type AwsDatabaseServerV4Parameters struct {

	// Time to retain backups, in days
	BackupRetentionPeriod int32 `json:"backupRetentionPeriod,omitempty"`

	// Version of the database engine (vendor)
	EngineVersion string `json:"engineVersion,omitempty"`

	// Whether to use a multi-AZ deployment
	MultiAZ string `json:"multiAZ,omitempty"`

	// Storage type
	StorageType string `json:"storageType,omitempty"`
}

// Validate validates this aws database server v4 parameters
func (m *AwsDatabaseServerV4Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsDatabaseServerV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsDatabaseServerV4Parameters) UnmarshalBinary(b []byte) error {
	var res AwsDatabaseServerV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
