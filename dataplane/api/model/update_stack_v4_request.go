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

// UpdateStackV4Request update stack v4 request
// swagger:model UpdateStackV4Request
type UpdateStackV4Request struct {

	// instance group adjustment
	InstanceGroupAdjustment *InstanceGroupAdjustmentV4Request `json:"instanceGroupAdjustment,omitempty"`

	// status of the scale request
	// Enum: [SYNC FULL_SYNC REPAIR_FAILED_NODES STOPPED STARTED]
	Status string `json:"status,omitempty"`

	// on stack update, update cluster too
	WithClusterEvent bool `json:"withClusterEvent,omitempty"`
}

// Validate validates this update stack v4 request
func (m *UpdateStackV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroupAdjustment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateStackV4Request) validateInstanceGroupAdjustment(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceGroupAdjustment) { // not required
		return nil
	}

	if m.InstanceGroupAdjustment != nil {
		if err := m.InstanceGroupAdjustment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceGroupAdjustment")
			}
			return err
		}
	}

	return nil
}

var updateStackV4RequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYNC","FULL_SYNC","REPAIR_FAILED_NODES","STOPPED","STARTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateStackV4RequestTypeStatusPropEnum = append(updateStackV4RequestTypeStatusPropEnum, v)
	}
}

const (

	// UpdateStackV4RequestStatusSYNC captures enum value "SYNC"
	UpdateStackV4RequestStatusSYNC string = "SYNC"

	// UpdateStackV4RequestStatusFULLSYNC captures enum value "FULL_SYNC"
	UpdateStackV4RequestStatusFULLSYNC string = "FULL_SYNC"

	// UpdateStackV4RequestStatusREPAIRFAILEDNODES captures enum value "REPAIR_FAILED_NODES"
	UpdateStackV4RequestStatusREPAIRFAILEDNODES string = "REPAIR_FAILED_NODES"

	// UpdateStackV4RequestStatusSTOPPED captures enum value "STOPPED"
	UpdateStackV4RequestStatusSTOPPED string = "STOPPED"

	// UpdateStackV4RequestStatusSTARTED captures enum value "STARTED"
	UpdateStackV4RequestStatusSTARTED string = "STARTED"
)

// prop value enum
func (m *UpdateStackV4Request) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateStackV4RequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStackV4Request) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateStackV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateStackV4Request) UnmarshalBinary(b []byte) error {
	var res UpdateStackV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
