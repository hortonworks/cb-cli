// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterV2Request cluster v2 request
// swagger:model ClusterV2Request

type ClusterV2Request struct {

	// ambari specific requests for cluster
	AmbariRequest *AmbariV2Request `json:"ambariRequest,omitempty"`

	// byos specific requests for cluster
	ByosRequest *ByosV2Request `json:"byosRequest,omitempty"`

	// send email about the result of the cluster installation
	EmailNeeded *bool `json:"emailNeeded,omitempty"`

	// send email to the requested address
	EmailTo string `json:"emailTo,omitempty"`

	// executor type of cluster
	ExecutorType string `json:"executorType,omitempty"`

	// external file system configuration
	FileSystem *FileSystem `json:"fileSystem,omitempty"`

	// LDAP config for the cluster
	LdapConfig *LdapConfigRequest `json:"ldapConfig,omitempty"`

	// LDAP config id for the cluster
	LdapConfigID int64 `json:"ldapConfigId,omitempty"`

	// RDS configuration ids for the cluster
	// Unique: true
	RdsConfigIds []int64 `json:"rdsConfigIds"`

	// details of the external database for Hadoop components
	// Unique: true
	RdsConfigJsons []*RDSConfig `json:"rdsConfigJsons"`
}

/* polymorph ClusterV2Request ambariRequest false */

/* polymorph ClusterV2Request byosRequest false */

/* polymorph ClusterV2Request emailNeeded false */

/* polymorph ClusterV2Request emailTo false */

/* polymorph ClusterV2Request executorType false */

/* polymorph ClusterV2Request fileSystem false */

/* polymorph ClusterV2Request ldapConfig false */

/* polymorph ClusterV2Request ldapConfigId false */

/* polymorph ClusterV2Request rdsConfigIds false */

/* polymorph ClusterV2Request rdsConfigJsons false */

// Validate validates this cluster v2 request
func (m *ClusterV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariRequest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateByosRequest(formats); err != nil {
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

	if err := m.validateLdapConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigJsons(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterV2Request) validateAmbariRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariRequest) { // not required
		return nil
	}

	if m.AmbariRequest != nil {

		if err := m.AmbariRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariRequest")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV2Request) validateByosRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.ByosRequest) { // not required
		return nil
	}

	if m.ByosRequest != nil {

		if err := m.ByosRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("byosRequest")
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

func (m *ClusterV2Request) validateLdapConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapConfig) { // not required
		return nil
	}

	if m.LdapConfig != nil {

		if err := m.LdapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV2Request) validateRdsConfigIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigIds", "body", m.RdsConfigIds); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV2Request) validateRdsConfigJsons(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigJsons) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigJsons", "body", m.RdsConfigJsons); err != nil {
		return err
	}

	for i := 0; i < len(m.RdsConfigJsons); i++ {

		if swag.IsZero(m.RdsConfigJsons[i]) { // not required
			continue
		}

		if m.RdsConfigJsons[i] != nil {

			if err := m.RdsConfigJsons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rdsConfigJsons" + "." + strconv.Itoa(i))
				}
				return err
			}
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
