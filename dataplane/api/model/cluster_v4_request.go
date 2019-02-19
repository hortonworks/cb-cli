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

// ClusterV4Request cluster v4 request
// swagger:model ClusterV4Request
type ClusterV4Request struct {

	// ambari specific requests
	Ambari *AmbariV4Request `json:"ambari,omitempty"`

	// external cloud storage configuration
	CloudStorage *CloudStorageV4Request `json:"cloudStorage,omitempty"`

	// custom containers
	CustomContainer *CustomContainerV4Request `json:"customContainer,omitempty"`

	// custom queue for yarn orchestrator
	CustomQueue string `json:"customQueue,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	Databases []string `json:"databases"`

	// executor type of cluster
	// Enum: [CONTAINER DEFAULT]
	ExecutorType string `json:"executorType,omitempty"`

	// gateway
	Gateway *GatewayV4Request `json:"gateway,omitempty"`

	// kerberos name
	KerberosName string `json:"kerberosName,omitempty"`

	// LDAP config name for the cluster
	LdapName string `json:"ldapName,omitempty"`

	// proxy configuration name for the cluster
	ProxyName string `json:"proxyName,omitempty"`
}

// Validate validates this cluster v4 request
func (m *ClusterV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbari(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterV4Request) validateAmbari(formats strfmt.Registry) error {

	if swag.IsZero(m.Ambari) { // not required
		return nil
	}

	if m.Ambari != nil {
		if err := m.Ambari.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambari")
			}
			return err
		}
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

var clusterV4RequestTypeExecutorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTAINER","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterV4RequestTypeExecutorTypePropEnum = append(clusterV4RequestTypeExecutorTypePropEnum, v)
	}
}

const (

	// ClusterV4RequestExecutorTypeCONTAINER captures enum value "CONTAINER"
	ClusterV4RequestExecutorTypeCONTAINER string = "CONTAINER"

	// ClusterV4RequestExecutorTypeDEFAULT captures enum value "DEFAULT"
	ClusterV4RequestExecutorTypeDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *ClusterV4Request) validateExecutorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterV4RequestTypeExecutorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterV4Request) validateExecutorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutorTypeEnum("executorType", "body", m.ExecutorType); err != nil {
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
