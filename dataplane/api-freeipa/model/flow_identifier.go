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

// FlowIdentifier flow identifier
// swagger:model FlowIdentifier
type FlowIdentifier struct {

	// pollable Id
	PollableID string `json:"pollableId,omitempty"`

	// type
	// Enum: [FLOW FLOW_CHAIN NOT_TRIGGERED]
	Type string `json:"type,omitempty"`
}

// Validate validates this flow identifier
func (m *FlowIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flowIdentifierTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FLOW","FLOW_CHAIN","NOT_TRIGGERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowIdentifierTypeTypePropEnum = append(flowIdentifierTypeTypePropEnum, v)
	}
}

const (

	// FlowIdentifierTypeFLOW captures enum value "FLOW"
	FlowIdentifierTypeFLOW string = "FLOW"

	// FlowIdentifierTypeFLOWCHAIN captures enum value "FLOW_CHAIN"
	FlowIdentifierTypeFLOWCHAIN string = "FLOW_CHAIN"

	// FlowIdentifierTypeNOTTRIGGERED captures enum value "NOT_TRIGGERED"
	FlowIdentifierTypeNOTTRIGGERED string = "NOT_TRIGGERED"
)

// prop value enum
func (m *FlowIdentifier) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, flowIdentifierTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlowIdentifier) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowIdentifier) UnmarshalBinary(b []byte) error {
	var res FlowIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
