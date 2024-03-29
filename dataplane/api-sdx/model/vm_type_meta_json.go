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

// VMTypeMetaJSON Vm type meta Json
// swagger:model VmTypeMetaJson
type VMTypeMetaJSON struct {

	// Virtual machine type configurations
	Configs []*VolumeParameterConfigResponse `json:"configs"`

	// cpu
	CPU int32 `json:"cpu,omitempty"`

	// memory in gb
	MemoryInGb float32 `json:"memoryInGb,omitempty"`

	// Virtual machine type properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Validate validates this Vm type meta Json
func (m *VMTypeMetaJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTypeMetaJSON) validateConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.Configs) { // not required
		return nil
	}

	for i := 0; i < len(m.Configs); i++ {
		if swag.IsZero(m.Configs[i]) { // not required
			continue
		}

		if m.Configs[i] != nil {
			if err := m.Configs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMTypeMetaJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTypeMetaJSON) UnmarshalBinary(b []byte) error {
	var res VMTypeMetaJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
