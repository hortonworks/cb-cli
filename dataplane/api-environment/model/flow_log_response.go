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

// FlowLogResponse flow log response
// swagger:model FlowLogResponse
type FlowLogResponse struct {

	// created
	Created int64 `json:"created,omitempty"`

	// current state
	CurrentState string `json:"currentState,omitempty"`

	// end time
	EndTime int64 `json:"endTime,omitempty"`

	// finalized
	Finalized bool `json:"finalized,omitempty"`

	// flow chain Id
	FlowChainID string `json:"flowChainId,omitempty"`

	// flow Id
	FlowID string `json:"flowId,omitempty"`

	// flow trigger user crn
	FlowTriggerUserCrn string `json:"flowTriggerUserCrn,omitempty"`

	// next event
	NextEvent string `json:"nextEvent,omitempty"`

	// node Id
	NodeID string `json:"nodeId,omitempty"`

	// resource Id
	ResourceID int64 `json:"resourceId,omitempty"`

	// state status
	// Enum: [PENDING SUCCESSFUL FAILED]
	StateStatus string `json:"stateStatus,omitempty"`
}

// Validate validates this flow log response
func (m *FlowLogResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flowLogResponseTypeStateStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowLogResponseTypeStateStatusPropEnum = append(flowLogResponseTypeStateStatusPropEnum, v)
	}
}

const (

	// FlowLogResponseStateStatusPENDING captures enum value "PENDING"
	FlowLogResponseStateStatusPENDING string = "PENDING"

	// FlowLogResponseStateStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	FlowLogResponseStateStatusSUCCESSFUL string = "SUCCESSFUL"

	// FlowLogResponseStateStatusFAILED captures enum value "FAILED"
	FlowLogResponseStateStatusFAILED string = "FAILED"
)

// prop value enum
func (m *FlowLogResponse) validateStateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, flowLogResponseTypeStateStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlowLogResponse) validateStateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.StateStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateStatusEnum("stateStatus", "body", m.StateStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowLogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowLogResponse) UnmarshalBinary(b []byte) error {
	var res FlowLogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
