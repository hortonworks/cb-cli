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

// ClusterV4Request cluster v4 request
// swagger:model ClusterV4Request
type ClusterV4Request struct {

	// blueprint name for the cluster
	BlueprintName string `json:"blueprintName,omitempty"`

	// external cloud storage configuration
	CloudStorage *CloudStorageRequest `json:"cloudStorage,omitempty"`

	// cloudera manager specific requests
	Cm *ClouderaManagerV4Request `json:"cm,omitempty"`

	// custom configurations name for the stack
	CustomConfigurationsName string `json:"customConfigurationsName,omitempty"`

	// custom containers
	CustomContainer *CustomContainerV4Request `json:"customContainer,omitempty"`

	// custom queue for yarn orchestrator
	CustomQueue string `json:"customQueue,omitempty"`

	// Contains valid Crn for a redbeams database server
	DatabaseServerCrn string `json:"databaseServerCrn,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	Databases []string `json:"databases"`

	// gateway
	Gateway *GatewayV4Request `json:"gateway,omitempty"`

	// ambari password
	// Max Length: 100
	// Min Length: 8
	Password string `json:"password,omitempty"`

	// proxy CRN for the cluster
	ProxyConfigCrn string `json:"proxyConfigCrn,omitempty"`

	// Enables Ranger Raz for the cluster on S3 and ADLSv2.
	RangerRazEnabled bool `json:"rangerRazEnabled,omitempty"`

	// ambari username
	// Max Length: 15
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	UserName string `json:"userName,omitempty"`

	// blueprint validation
	ValidateBlueprint bool `json:"validateBlueprint,omitempty"`
}

// Validate validates this cluster v4 request
func (m *ClusterV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterV4Request) validateCloudStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorage) { // not required
		return nil
	}

	if m.CloudStorage != nil {
		if err := m.CloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudStorage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Request) validateCm(formats strfmt.Registry) error {

	if swag.IsZero(m.Cm) { // not required
		return nil
	}

	if m.Cm != nil {
		if err := m.Cm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cm")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Request) validateCustomContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomContainer) { // not required
		return nil
	}

	if m.CustomContainer != nil {
		if err := m.CustomContainer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customContainer")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Request) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if err := validate.UniqueItems("databases", "body", m.Databases); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV4Request) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Request) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", string(m.Password), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 100); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV4Request) validateUserName(formats strfmt.Registry) error {

	if swag.IsZero(m.UserName) { // not required
		return nil
	}

	if err := validate.MinLength("userName", "body", string(m.UserName), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("userName", "body", string(m.UserName), 15); err != nil {
		return err
	}

	if err := validate.Pattern("userName", "body", string(m.UserName), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterV4Request) UnmarshalBinary(b []byte) error {
	var res ClusterV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
