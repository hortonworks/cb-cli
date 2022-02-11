// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VerticalScaleResponse vertical scale response
// swagger:model VerticalScaleResponse
type VerticalScaleResponse struct {

	// flow identifier
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// request
	Request *VerticalScaleRequest `json:"request,omitempty"`
}

// Validate validates this vertical scale response
func (m *VerticalScaleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerticalScaleResponse) validateFlowIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowIdentifier) { // not required
		return nil
	}

	if m.FlowIdentifier != nil {
		if err := m.FlowIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *VerticalScaleResponse) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VerticalScaleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerticalScaleResponse) UnmarshalBinary(b []byte) error {
	var res VerticalScaleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
