// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WorkloadAnalyticsV4Response workload analytics v4 response
// swagger:model WorkloadAnalyticsV4Response
type WorkloadAnalyticsV4Response struct {

	// stack related telemetry dynamic component settings
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// telemetry - workload altus service (databus) endpoint url
	DatabusEndpoint string `json:"databusEndpoint,omitempty"`

	// enable telemetry component
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this workload analytics v4 response
func (m *WorkloadAnalyticsV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkloadAnalyticsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkloadAnalyticsV4Response) UnmarshalBinary(b []byte) error {
	var res WorkloadAnalyticsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
