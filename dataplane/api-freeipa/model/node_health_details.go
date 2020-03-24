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

// NodeHealthDetails node health details
// swagger:model NodeHealthDetails
type NodeHealthDetails struct {

	// instance Id
	// Required: true
	InstanceID *string `json:"instanceId"`

	// issues
	// Required: true
	Issues []string `json:"issues"`

	// name
	// Required: true
	Name *string `json:"name"`

	// status
	// Required: true
	// Enum: [REQUESTED CREATED UNREGISTERED REGISTERED DECOMMISSIONED TERMINATED DELETED_ON_PROVIDER_SIDE FAILED STOPPED REBOOTING UNREACHABLE]
	Status *string `json:"status"`
}

// Validate validates this node health details
func (m *NodeHealthDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *NodeHealthDetails) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instanceId", "body", m.InstanceID); err != nil {
		return err
	}

	return nil
}

func (m *NodeHealthDetails) validateIssues(formats strfmt.Registry) error {

	if err := validate.Required("issues", "body", m.Issues); err != nil {
		return err
	}

	return nil
}

func (m *NodeHealthDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var nodeHealthDetailsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATED","UNREGISTERED","REGISTERED","DECOMMISSIONED","TERMINATED","DELETED_ON_PROVIDER_SIDE","FAILED","STOPPED","REBOOTING","UNREACHABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeHealthDetailsTypeStatusPropEnum = append(nodeHealthDetailsTypeStatusPropEnum, v)
	}
}

const (

	// NodeHealthDetailsStatusREQUESTED captures enum value "REQUESTED"
	NodeHealthDetailsStatusREQUESTED string = "REQUESTED"

	// NodeHealthDetailsStatusCREATED captures enum value "CREATED"
	NodeHealthDetailsStatusCREATED string = "CREATED"

	// NodeHealthDetailsStatusUNREGISTERED captures enum value "UNREGISTERED"
	NodeHealthDetailsStatusUNREGISTERED string = "UNREGISTERED"

	// NodeHealthDetailsStatusREGISTERED captures enum value "REGISTERED"
	NodeHealthDetailsStatusREGISTERED string = "REGISTERED"

	// NodeHealthDetailsStatusDECOMMISSIONED captures enum value "DECOMMISSIONED"
	NodeHealthDetailsStatusDECOMMISSIONED string = "DECOMMISSIONED"

	// NodeHealthDetailsStatusTERMINATED captures enum value "TERMINATED"
	NodeHealthDetailsStatusTERMINATED string = "TERMINATED"

	// NodeHealthDetailsStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	NodeHealthDetailsStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// NodeHealthDetailsStatusFAILED captures enum value "FAILED"
	NodeHealthDetailsStatusFAILED string = "FAILED"

	// NodeHealthDetailsStatusSTOPPED captures enum value "STOPPED"
	NodeHealthDetailsStatusSTOPPED string = "STOPPED"

	// NodeHealthDetailsStatusREBOOTING captures enum value "REBOOTING"
	NodeHealthDetailsStatusREBOOTING string = "REBOOTING"

	// NodeHealthDetailsStatusUNREACHABLE captures enum value "UNREACHABLE"
	NodeHealthDetailsStatusUNREACHABLE string = "UNREACHABLE"
)

// prop value enum
func (m *NodeHealthDetails) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodeHealthDetailsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NodeHealthDetails) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeHealthDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeHealthDetails) UnmarshalBinary(b []byte) error {
	var res NodeHealthDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
