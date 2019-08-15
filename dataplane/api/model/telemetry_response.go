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

// TelemetryResponse telemetry response
// swagger:model TelemetryResponse
type TelemetryResponse struct {

	// Cloud Logging (telemetry) settings.
	Logging *LoggingResponse `json:"logging,omitempty"`

	// enable cluster deployment log reporting.
	// Enum: [ENABLED DISABLED]
	ReportDeploymentLogs string `json:"reportDeploymentLogs,omitempty"`

	// Workload analytics (telemetry) settings.
	WorkloadAnalytics *WorkloadAnalyticsResponse `json:"workloadAnalytics,omitempty"`
}

// Validate validates this telemetry response
func (m *TelemetryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportDeploymentLogs(formats); err != nil {
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

var telemetryResponseTypeReportDeploymentLogsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		telemetryResponseTypeReportDeploymentLogsPropEnum = append(telemetryResponseTypeReportDeploymentLogsPropEnum, v)
	}
}

const (

	// TelemetryResponseReportDeploymentLogsENABLED captures enum value "ENABLED"
	TelemetryResponseReportDeploymentLogsENABLED string = "ENABLED"

	// TelemetryResponseReportDeploymentLogsDISABLED captures enum value "DISABLED"
	TelemetryResponseReportDeploymentLogsDISABLED string = "DISABLED"
)

// prop value enum
func (m *TelemetryResponse) validateReportDeploymentLogsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, telemetryResponseTypeReportDeploymentLogsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TelemetryResponse) validateReportDeploymentLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportDeploymentLogs) { // not required
		return nil
	}

	// value enum
	if err := m.validateReportDeploymentLogsEnum("reportDeploymentLogs", "body", m.ReportDeploymentLogs); err != nil {
		return err
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
