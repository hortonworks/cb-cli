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

// DistroXClusterV1Request distro x cluster v1 request
// swagger:model DistroXClusterV1Request
type DistroXClusterV1Request struct {

	// blueprint name for the cluster
	BlueprintName string `json:"blueprintName,omitempty"`

	// external cloud storage configuration
	CloudStorage *CloudStorageV1Request `json:"cloudStorage,omitempty"`

	// cloudera manager specific requests
	Cm *ClouderaManagerV1Request `json:"cm,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	Databases []string `json:"databases"`

	// gateway
	Gateway *GatewayV1Request `json:"gateway,omitempty"`

	// ambari password
	// Required: true
	// Max Length: 100
	// Min Length: 8
	Password *string `json:"password"`

	// proxy configuration name for the cluster
	Proxy string `json:"proxy,omitempty"`

	// ambari username
	// Required: true
	// Max Length: 15
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	UserName *string `json:"userName"`
}

// Validate validates this distro x cluster v1 request
func (m *DistroXClusterV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCm(formats); err != nil {
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

func (m *DistroXClusterV1Request) validateCloudStorage(formats strfmt.Registry) error {

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

func (m *DistroXClusterV1Request) validateCm(formats strfmt.Registry) error {

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

func (m *DistroXClusterV1Request) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if err := validate.UniqueItems("databases", "body", m.Databases); err != nil {
		return err
	}

	return nil
}

func (m *DistroXClusterV1Request) validateGateway(formats strfmt.Registry) error {

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

func (m *DistroXClusterV1Request) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 100); err != nil {
		return err
	}

	return nil
}

func (m *DistroXClusterV1Request) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	if err := validate.MinLength("userName", "body", string(*m.UserName), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("userName", "body", string(*m.UserName), 15); err != nil {
		return err
	}

	if err := validate.Pattern("userName", "body", string(*m.UserName), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistroXClusterV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistroXClusterV1Request) UnmarshalBinary(b []byte) error {
	var res DistroXClusterV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
