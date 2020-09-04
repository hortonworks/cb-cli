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

// MetricAlertResponse metric alert response
// swagger:model MetricAlertResponse
type MetricAlertResponse struct {

	// Definition of the alert
	AlertDefinition string `json:"alertDefinition,omitempty"`

	// Definition label of the alert
	AlertDefinitionLabel string `json:"alertDefinitionLabel,omitempty"`

	// Name of the alert
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*$)
	AlertName string `json:"alertName,omitempty"`

	// State of the alert
	// Enum: [OK WARN CRITICAL]
	AlertState string `json:"alertState,omitempty"`

	// Description of the alert
	Description string `json:"description,omitempty"`

	// Id of the alert
	ID int64 `json:"id,omitempty"`

	// Period of the alert
	Period int32 `json:"period,omitempty"`

	// Id of the scaling ploicy
	ScalingPolicy *ScalingPolicyRequest `json:"scalingPolicy,omitempty"`

	// Id of the scaling ploicy
	ScalingPolicyID int64 `json:"scalingPolicyId,omitempty"`
}

// Validate validates this metric alert response
func (m *MetricAlertResponse) Validate(formats strfmt.Registry) error {
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

func (m *MetricAlertResponse) validateAlertName(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertName) { // not required
		return nil
	}

	if err := validate.Pattern("alertName", "body", string(m.AlertName), `(^[a-zA-Z][-a-zA-Z0-9]*$)`); err != nil {
		return err
	}

	return nil
}

var metricAlertResponseTypeAlertStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","WARN","CRITICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricAlertResponseTypeAlertStatePropEnum = append(metricAlertResponseTypeAlertStatePropEnum, v)
	}
}

const (

	// MetricAlertResponseAlertStateOK captures enum value "OK"
	MetricAlertResponseAlertStateOK string = "OK"

	// MetricAlertResponseAlertStateWARN captures enum value "WARN"
	MetricAlertResponseAlertStateWARN string = "WARN"

	// MetricAlertResponseAlertStateCRITICAL captures enum value "CRITICAL"
	MetricAlertResponseAlertStateCRITICAL string = "CRITICAL"
)

// prop value enum
func (m *MetricAlertResponse) validateAlertStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, metricAlertResponseTypeAlertStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MetricAlertResponse) validateAlertState(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertState) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlertStateEnum("alertState", "body", m.AlertState); err != nil {
		return err
	}

	return nil
}

func (m *MetricAlertResponse) validateScalingPolicy(formats strfmt.Registry) error {

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
func (m *MetricAlertResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricAlertResponse) UnmarshalBinary(b []byte) error {
	var res MetricAlertResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
