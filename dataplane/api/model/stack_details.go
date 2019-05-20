// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackDetails stack details
// swagger:model StackDetails
type StackDetails struct {

	// ambari version
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// cloud platform
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// cloudbreak version
	CloudbreakVersion string `json:"cloudbreakVersion,omitempty"`

	// cluster type
	ClusterType string `json:"clusterType,omitempty"`

	// cluster version
	ClusterVersion string `json:"clusterVersion,omitempty"`

	// datalake Id
	DatalakeID int64 `json:"datalakeId,omitempty"`

	// datalake resource Id
	DatalakeResourceID int64 `json:"datalakeResourceId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// detailed status
	DetailedStatus string `json:"detailedStatus,omitempty"`

	// existing network
	ExistingNetwork bool `json:"existingNetwork,omitempty"`

	// existing subnet
	ExistingSubnet bool `json:"existingSubnet,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// image identifier
	ImageIdentifier string `json:"imageIdentifier,omitempty"`

	// instance groups
	InstanceGroups []*InstanceGroupDetails `json:"instanceGroups"`

	// name
	Name string `json:"name,omitempty"`

	// platform variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// prewarmed image
	PrewarmedImage bool `json:"prewarmedImage,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this stack details
func (m *StackDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackDetails) validateInstanceGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackDetails) UnmarshalBinary(b []byte) error {
	var res StackDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
