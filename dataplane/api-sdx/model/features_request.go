// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeaturesRequest features request
// swagger:model FeaturesRequest
type FeaturesRequest struct {

	// enable cluster logs collection
	ClusterLogsCollection *FeatureSetting `json:"clusterLogsCollection,omitempty"`

	// Workload analytics (telemetry) settings.
	WorkloadAnalytics *FeatureSetting `json:"workloadAnalytics,omitempty"`
}

// Validate validates this features request
func (m *FeaturesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterLogsCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadAnalytics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturesRequest) validateClusterLogsCollection(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterLogsCollection) { // not required
		return nil
	}

	if m.ClusterLogsCollection != nil {
		if err := m.ClusterLogsCollection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterLogsCollection")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturesRequest) validateWorkloadAnalytics(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadAnalytics) { // not required
		return nil
	}

	if m.WorkloadAnalytics != nil {
		if err := m.WorkloadAnalytics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadAnalytics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturesRequest) UnmarshalBinary(b []byte) error {
	var res FeaturesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
