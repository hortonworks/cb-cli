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

// LoggingV1Request logging v1 request
// swagger:model LoggingV1Request
type LoggingV1Request struct {

	// Logging component specific attributes of the environment
	Attributes *LoggingAttributesHolder `json:"attributes,omitempty"`

	// Enable telemetry component feature of the environment
	Enabled bool `json:"enabled,omitempty"`

	// Logging output type of the environment
	// Enum: [s3 wasb abfs gcs cloudwatch]
	Output string `json:"output,omitempty"`
}

// Validate validates this logging v1 request
func (m *LoggingV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggingV1Request) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

var loggingV1RequestTypeOutputPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["s3","wasb","abfs","gcs","cloudwatch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loggingV1RequestTypeOutputPropEnum = append(loggingV1RequestTypeOutputPropEnum, v)
	}
}

const (

	// LoggingV1RequestOutputS3 captures enum value "s3"
	LoggingV1RequestOutputS3 string = "s3"

	// LoggingV1RequestOutputWasb captures enum value "wasb"
	LoggingV1RequestOutputWasb string = "wasb"

	// LoggingV1RequestOutputAbfs captures enum value "abfs"
	LoggingV1RequestOutputAbfs string = "abfs"

	// LoggingV1RequestOutputGcs captures enum value "gcs"
	LoggingV1RequestOutputGcs string = "gcs"

	// LoggingV1RequestOutputCloudwatch captures enum value "cloudwatch"
	LoggingV1RequestOutputCloudwatch string = "cloudwatch"
)

// prop value enum
func (m *LoggingV1Request) validateOutputEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loggingV1RequestTypeOutputPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LoggingV1Request) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutputEnum("output", "body", m.Output); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoggingV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoggingV1Request) UnmarshalBinary(b []byte) error {
	var res LoggingV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
