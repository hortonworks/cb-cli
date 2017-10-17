// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateStackV2 update stack v2
// swagger:model UpdateStackV2

type UpdateStackV2 struct {

	// instance group adjustment
	InstanceGroupAdjustment *InstanceGroupAdjustmentV2 `json:"instanceGroupAdjustment,omitempty"`

	// status of the scale request
	Status string `json:"status,omitempty"`

	// on stack update, update cluster too
	WithClusterEvent *bool `json:"withClusterEvent,omitempty"`
}

/* polymorph UpdateStackV2 instanceGroupAdjustment false */

/* polymorph UpdateStackV2 status false */

/* polymorph UpdateStackV2 withClusterEvent false */

// Validate validates this update stack v2
func (m *UpdateStackV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroupAdjustment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateStackV2) validateInstanceGroupAdjustment(formats strfmt.Registry) error {

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

var updateStackV2TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYNC","FULL_SYNC","REPAIR_FAILED_NODES","STOPPED","STARTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateStackV2TypeStatusPropEnum = append(updateStackV2TypeStatusPropEnum, v)
	}
}

const (
	// UpdateStackV2StatusSYNC captures enum value "SYNC"
	UpdateStackV2StatusSYNC string = "SYNC"
	// UpdateStackV2StatusFULLSYNC captures enum value "FULL_SYNC"
	UpdateStackV2StatusFULLSYNC string = "FULL_SYNC"
	// UpdateStackV2StatusREPAIRFAILEDNODES captures enum value "REPAIR_FAILED_NODES"
	UpdateStackV2StatusREPAIRFAILEDNODES string = "REPAIR_FAILED_NODES"
	// UpdateStackV2StatusSTOPPED captures enum value "STOPPED"
	UpdateStackV2StatusSTOPPED string = "STOPPED"
	// UpdateStackV2StatusSTARTED captures enum value "STARTED"
	UpdateStackV2StatusSTARTED string = "STARTED"
)

// prop value enum
func (m *UpdateStackV2) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateStackV2TypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStackV2) validateStatus(formats strfmt.Registry) error {

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
func (m *UpdateStackV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateStackV2) UnmarshalBinary(b []byte) error {
	var res UpdateStackV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
