// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SpecialParametersJSON special parameters Json
// swagger:model SpecialParametersJson
type SpecialParametersJSON struct {

	// platform specific custom parameters
	PlatformSpecificSpecialParameters map[string]map[string]bool `json:"platformSpecificSpecialParameters,omitempty"`

	// custom parameters
	SpecialParameters map[string]bool `json:"specialParameters,omitempty"`
}

// Validate validates this special parameters Json
func (m *SpecialParametersJSON) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpecialParametersJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecialParametersJSON) UnmarshalBinary(b []byte) error {
	var res SpecialParametersJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
