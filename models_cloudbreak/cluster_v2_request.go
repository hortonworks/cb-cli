// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterV2Request cluster v2 request
// swagger:model ClusterV2Request

type ClusterV2Request struct {

	// ambari specific requests
	Ambari *AmbariV2Request `json:"ambari,omitempty"`

	// send email about the result of the cluster installation
	EmailNeeded *bool `json:"emailNeeded,omitempty"`

	// send email to the requested address
	EmailTo string `json:"emailTo,omitempty"`

	// executor type of cluster
	ExecutorType string `json:"executorType,omitempty"`

	// external file system configuration
	FileSystem *FileSystem `json:"fileSystem,omitempty"`

	// LDAP config name for the cluster
	LdapConfigName string `json:"ldapConfigName,omitempty"`

	// details of the external database for Hadoop components
	RdsConfigs *RdsConfigs `json:"rdsConfigs,omitempty"`
}

/* polymorph ClusterV2Request ambari false */

/* polymorph ClusterV2Request emailNeeded false */

/* polymorph ClusterV2Request emailTo false */

/* polymorph ClusterV2Request executorType false */

/* polymorph ClusterV2Request fileSystem false */

/* polymorph ClusterV2Request ldapConfigName false */

/* polymorph ClusterV2Request rdsConfigs false */

// Validate validates this cluster v2 request
func (m *ClusterV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbari(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExecutorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFileSystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterV2Request) validateAmbari(formats strfmt.Registry) error {

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

var clusterV2RequestTypeExecutorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTAINER","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterV2RequestTypeExecutorTypePropEnum = append(clusterV2RequestTypeExecutorTypePropEnum, v)
	}
}

const (
	// ClusterV2RequestExecutorTypeCONTAINER captures enum value "CONTAINER"
	ClusterV2RequestExecutorTypeCONTAINER string = "CONTAINER"
	// ClusterV2RequestExecutorTypeDEFAULT captures enum value "DEFAULT"
	ClusterV2RequestExecutorTypeDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *ClusterV2Request) validateExecutorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterV2RequestTypeExecutorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterV2Request) validateExecutorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutorTypeEnum("executorType", "body", m.ExecutorType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV2Request) validateFileSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.FileSystem) { // not required
		return nil
	}

	if m.FileSystem != nil {

		if err := m.FileSystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileSystem")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV2Request) validateRdsConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigs) { // not required
		return nil
	}

	if m.RdsConfigs != nil {

		if err := m.RdsConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rdsConfigs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterV2Request) UnmarshalBinary(b []byte) error {
	var res ClusterV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
