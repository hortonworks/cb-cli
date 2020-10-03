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

// AllocateDatabaseServerV4Request Request for allocating a new database server in a provider
// swagger:model AllocateDatabaseServerV4Request
type AllocateDatabaseServerV4Request struct {

	// AWS-specific parameters for the database stack
	Aws AwsDBStackV4Parameters `json:"aws,omitempty"`

	// CRN of the cluster of the database server
	// Required: true
	ClusterCrn *string `json:"clusterCrn"`

	// Database server information for the database stack
	// Required: true
	DatabaseServer *DatabaseServerV4StackRequest `json:"databaseServer"`

	// CRN of the environment of the database server
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Azure-specific parameters for the database stack
	Gcp GcpDBStackV4Parameters `json:"gcp,omitempty"`

	// Name of the database stack
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name string `json:"name,omitempty"`

	// Network information for the database stack
	Network *NetworkV4StackRequest `json:"network,omitempty"`
}

// Validate validates this allocate database server v4 request
func (m *AllocateDatabaseServerV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocateDatabaseServerV4Request) validateClusterCrn(formats strfmt.Registry) error {

	if err := validate.Required("clusterCrn", "body", m.ClusterCrn); err != nil {
		return err
	}

	return nil
}

func (m *AllocateDatabaseServerV4Request) validateDatabaseServer(formats strfmt.Registry) error {

	if err := validate.Required("databaseServer", "body", m.DatabaseServer); err != nil {
		return err
	}

	if m.DatabaseServer != nil {
		if err := m.DatabaseServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseServer")
			}
			return err
		}
	}

	return nil
}

func (m *AllocateDatabaseServerV4Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *AllocateDatabaseServerV4Request) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *AllocateDatabaseServerV4Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocateDatabaseServerV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocateDatabaseServerV4Request) UnmarshalBinary(b []byte) error {
	var res AllocateDatabaseServerV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
