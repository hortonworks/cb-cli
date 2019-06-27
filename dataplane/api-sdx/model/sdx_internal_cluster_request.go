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

// SdxInternalClusterRequest sdx internal cluster request
// swagger:model SdxInternalClusterRequest
type SdxInternalClusterRequest struct {

	// access cidr
	// Required: true
	AccessCidr *string `json:"accessCidr"`

	// cloud storage
	// Required: true
	CloudStorage *SdxCloudStorageRequest `json:"cloudStorage"`

	// cluster shape
	// Required: true
	ClusterShape *string `json:"clusterShape"`

	// environment
	// Required: true
	Environment *string `json:"environment"`

	// stack v4 request
	StackV4Request *StackV4Request `json:"stackV4Request,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sdx internal cluster request
func (m *SdxInternalClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackV4Request(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxInternalClusterRequest) validateAccessCidr(formats strfmt.Registry) error {

	if err := validate.Required("accessCidr", "body", m.AccessCidr); err != nil {
		return err
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateCloudStorage(formats strfmt.Registry) error {

	if err := validate.Required("cloudStorage", "body", m.CloudStorage); err != nil {
		return err
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

func (m *SdxInternalClusterRequest) validateClusterShape(formats strfmt.Registry) error {

	if err := validate.Required("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateStackV4Request(formats strfmt.Registry) error {

	if swag.IsZero(m.StackV4Request) { // not required
		return nil
	}

	if m.StackV4Request != nil {
		if err := m.StackV4Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackV4Request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxInternalClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxInternalClusterRequest) UnmarshalBinary(b []byte) error {
	var res SdxInternalClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
