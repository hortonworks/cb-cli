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

// TelemetryResponse telemetry response
// swagger:model TelemetryResponse
type TelemetryResponse struct {

	// Telemetry features settings
	Features *FeaturesResponse `json:"features,omitempty"`

	// Telemetry fluent settings (overrides).
	FluentAttributes map[string]interface{} `json:"fluentAttributes,omitempty"`

	// Cloud Logging (telemetry) settings.
	Logging *LoggingResponse `json:"logging,omitempty"`

	// Monitoring related (telemetry) settings.
	Monitoring *MonitoringResponse `json:"monitoring,omitempty"`

	// Telemetry anonymization rules (persistent on cluster level) that are applied on shipped logs.
	Rules []*AnonymizationRule `json:"rules"`

	// Workload analytics (telemetry) settings.
	WorkloadAnalytics *WorkloadAnalyticsResponse `json:"workloadAnalytics,omitempty"`
}

// Validate validates this telemetry response
func (m *TelemetryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
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

func (m *TelemetryResponse) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *TelemetryResponse) validateLogging(formats strfmt.Registry) error {

	if swag.IsZero(m.Logging) { // not required
		return nil
	}

	if m.Logging != nil {
		if err := m.Logging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logging")
			}
			return err
		}
	}

	return nil
}

func (m *TelemetryResponse) validateMonitoring(formats strfmt.Registry) error {

	if swag.IsZero(m.Monitoring) { // not required
		return nil
	}

	if m.Monitoring != nil {
		if err := m.Monitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *TelemetryResponse) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TelemetryResponse) validateWorkloadAnalytics(formats strfmt.Registry) error {

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
func (m *TelemetryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TelemetryResponse) UnmarshalBinary(b []byte) error {
	var res TelemetryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
