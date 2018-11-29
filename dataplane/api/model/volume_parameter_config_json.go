// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VolumeParameterConfigJSON volume parameter config Json
// swagger:model VolumeParameterConfigJson
type VolumeParameterConfigJSON struct {

	// maximum number
	MaximumNumber int32 `json:"maximumNumber,omitempty"`

	// maximum size
	MaximumSize int32 `json:"maximumSize,omitempty"`

	// minimum number
	MinimumNumber int32 `json:"minimumNumber,omitempty"`

	// minimum size
	MinimumSize int32 `json:"minimumSize,omitempty"`

	// volume parameter type
	VolumeParameterType string `json:"volumeParameterType,omitempty"`
}

// Validate validates this volume parameter config Json
func (m *VolumeParameterConfigJSON) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeParameterConfigJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeParameterConfigJSON) UnmarshalBinary(b []byte) error {
	var res VolumeParameterConfigJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
