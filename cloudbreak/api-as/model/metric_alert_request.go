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

// MetricAlertRequest metric alert request
// swagger:model MetricAlertRequest
type MetricAlertRequest struct {

	// Definition of the alert
	AlertDefinition string `json:"alertDefinition,omitempty"`

	// Definition of the alert
	AlertDefinitionLabel string `json:"alertDefinitionLabel,omitempty"`

	// Name of the alert
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*$)
	AlertName string `json:"alertName,omitempty"`

	// State of the alert
	// Enum: [OK WARN CRITICAL]
	AlertState string `json:"alertState,omitempty"`

	// Description of the alert
	Description string `json:"description,omitempty"`

	// Period of the alert
	Period int32 `json:"period,omitempty"`

	// Id of the scaling ploicy
	ScalingPolicy *ScalingPolicyRequest `json:"scalingPolicy,omitempty"`
}

// Validate validates this metric alert request
func (m *MetricAlertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricAlertRequest) validateAlertName(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertName) { // not required
		return nil
	}

	if err := validate.Pattern("alertName", "body", string(m.AlertName), `(^[a-zA-Z][-a-zA-Z0-9]*$)`); err != nil {
		return err
	}

	return nil
}

var metricAlertRequestTypeAlertStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","WARN","CRITICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricAlertRequestTypeAlertStatePropEnum = append(metricAlertRequestTypeAlertStatePropEnum, v)
	}
}

const (

	// MetricAlertRequestAlertStateOK captures enum value "OK"
	MetricAlertRequestAlertStateOK string = "OK"

	// MetricAlertRequestAlertStateWARN captures enum value "WARN"
	MetricAlertRequestAlertStateWARN string = "WARN"

	// MetricAlertRequestAlertStateCRITICAL captures enum value "CRITICAL"
	MetricAlertRequestAlertStateCRITICAL string = "CRITICAL"
)

// prop value enum
func (m *MetricAlertRequest) validateAlertStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricAlertRequestTypeAlertStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricAlertRequest) validateAlertState(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertState) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlertStateEnum("alertState", "body", m.AlertState); err != nil {
		return err
	}

	return nil
}

func (m *MetricAlertRequest) validateScalingPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScalingPolicy) { // not required
		return nil
	}

	if m.ScalingPolicy != nil {
		if err := m.ScalingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scalingPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricAlertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricAlertRequest) UnmarshalBinary(b []byte) error {
	var res MetricAlertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
