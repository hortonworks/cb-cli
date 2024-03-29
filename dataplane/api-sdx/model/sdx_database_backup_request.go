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

// SdxDatabaseBackupRequest sdx database backup request
// swagger:model SdxDatabaseBackupRequest
type SdxDatabaseBackupRequest struct {

	// The ID of the database backup.
	// Required: true
	BackupID *string `json:"backupId"`

	// The location where the database backup will be stored.
	// Required: true
	BackupLocation *string `json:"backupLocation"`

	// The conditional parameter for whether connections to the database will be closed during backup or not.
	// Required: true
	CloseConnections *bool `json:"closeConnections"`

	// The maximum duration allowed for a backup or restore
	DatabaseMaxDurationInMin int32 `json:"databaseMaxDurationInMin,omitempty"`

	// The names of the Databases not to be backed up, or blank for all
	SkipDatabaseNames []string `json:"skipDatabaseNames"`
}

// Validate validates this sdx database backup request
func (m *SdxDatabaseBackupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloseConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxDatabaseBackupRequest) validateBackupID(formats strfmt.Registry) error {

	if err := validate.Required("backupId", "body", m.BackupID); err != nil {
		return err
	}

	return nil
}

func (m *SdxDatabaseBackupRequest) validateBackupLocation(formats strfmt.Registry) error {

	if err := validate.Required("backupLocation", "body", m.BackupLocation); err != nil {
		return err
	}

	return nil
}

func (m *SdxDatabaseBackupRequest) validateCloseConnections(formats strfmt.Registry) error {

	if err := validate.Required("closeConnections", "body", m.CloseConnections); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxDatabaseBackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxDatabaseBackupRequest) UnmarshalBinary(b []byte) error {
	var res SdxDatabaseBackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
