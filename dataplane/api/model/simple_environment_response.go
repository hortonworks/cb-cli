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

// SimpleEnvironmentResponse simple environment response
// swagger:model SimpleEnvironmentResponse
type SimpleEnvironmentResponse struct {

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Name of the credential of the environment.
	CredentialName string `json:"credentialName,omitempty"`

	// Names of the datalake clusters created in the environment.
	// Unique: true
	DatalakeClusterNames []string `json:"datalakeClusterNames"`

	// Datalake cluster resources registered to the environment.
	// Unique: true
	DatalakeResourcesNames []string `json:"datalakeResourcesNames"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Location of the environment.
	Location *LocationResponse `json:"location,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Regions of the environment.
	Regions *CompactRegionResponse `json:"regions,omitempty"`

	// Names of the workload clusters created in the environment.
	// Unique: true
	WorkloadClusterNames []string `json:"workloadClusterNames"`

	// workspace
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

// Validate validates this simple environment response
func (m *SimpleEnvironmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeResourcesNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleEnvironmentResponse) validateDatalakeClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeClusterNames", "body", m.DatalakeClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateDatalakeResourcesNames(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeResourcesNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("datalakeResourcesNames", "body", m.DatalakeResourcesNames); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if m.Regions != nil {
		if err := m.Regions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regions")
			}
			return err
		}
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateWorkloadClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("workloadClusterNames", "body", m.WorkloadClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *SimpleEnvironmentResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleEnvironmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleEnvironmentResponse) UnmarshalBinary(b []byte) error {
	var res SimpleEnvironmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
