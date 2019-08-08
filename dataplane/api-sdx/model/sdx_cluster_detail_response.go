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

// SdxClusterDetailResponse sdx cluster detail response
// swagger:model SdxClusterDetailResponse
type SdxClusterDetailResponse struct {

	// cluster shape
	// Enum: [CUSTOM LIGHT_DUTY MEDIUM_DUTY_HA]
	ClusterShape string `json:"clusterShape,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// crn
	Crn string `json:"crn,omitempty"`

	// environment crn
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// stack crn
	StackCrn string `json:"stackCrn,omitempty"`

	// stack v4 response
	StackV4Response *StackV4Response `json:"stackV4Response,omitempty"`

	// status
	// Enum: [REQUESTED STACK_CREATION_IN_PROGRESS STACK_DELETED STACK_DELETION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_IN_PROGRESS RUNNING PROVISIONING_FAILED REPAIR_IN_PROGRESS REPAIR_FAILED DELETE_REQUESTED DELETED DELETE_FAILED]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this sdx cluster detail response
func (m *SdxClusterDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackV4Response(formats); err != nil {
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

var sdxClusterDetailResponseTypeClusterShapePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","LIGHT_DUTY","MEDIUM_DUTY_HA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeClusterShapePropEnum = append(sdxClusterDetailResponseTypeClusterShapePropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseClusterShapeCUSTOM captures enum value "CUSTOM"
	SdxClusterDetailResponseClusterShapeCUSTOM string = "CUSTOM"

	// SdxClusterDetailResponseClusterShapeLIGHTDUTY captures enum value "LIGHT_DUTY"
	SdxClusterDetailResponseClusterShapeLIGHTDUTY string = "LIGHT_DUTY"

	// SdxClusterDetailResponseClusterShapeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	SdxClusterDetailResponseClusterShapeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateClusterShapeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeClusterShapePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateClusterShape(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterShape) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterShapeEnum("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	return nil
}

func (m *SdxClusterDetailResponse) validateStackV4Response(formats strfmt.Registry) error {

	if swag.IsZero(m.StackV4Response) { // not required
		return nil
	}

	if m.StackV4Response != nil {
		if err := m.StackV4Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackV4Response")
			}
			return err
		}
	}

	return nil
}

var sdxClusterDetailResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","STACK_CREATION_IN_PROGRESS","STACK_DELETED","STACK_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","RUNNING","PROVISIONING_FAILED","REPAIR_IN_PROGRESS","REPAIR_FAILED","DELETE_REQUESTED","DELETED","DELETE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeStatusPropEnum = append(sdxClusterDetailResponseTypeStatusPropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseStatusREQUESTED captures enum value "REQUESTED"
	SdxClusterDetailResponseStatusREQUESTED string = "REQUESTED"

	// SdxClusterDetailResponseStatusSTACKCREATIONINPROGRESS captures enum value "STACK_CREATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTACKCREATIONINPROGRESS string = "STACK_CREATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusSTACKDELETED captures enum value "STACK_DELETED"
	SdxClusterDetailResponseStatusSTACKDELETED string = "STACK_DELETED"

	// SdxClusterDetailResponseStatusSTACKDELETIONINPROGRESS captures enum value "STACK_DELETION_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTACKDELETIONINPROGRESS string = "STACK_DELETION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusRUNNING captures enum value "RUNNING"
	SdxClusterDetailResponseStatusRUNNING string = "RUNNING"

	// SdxClusterDetailResponseStatusPROVISIONINGFAILED captures enum value "PROVISIONING_FAILED"
	SdxClusterDetailResponseStatusPROVISIONINGFAILED string = "PROVISIONING_FAILED"

	// SdxClusterDetailResponseStatusREPAIRINPROGRESS captures enum value "REPAIR_IN_PROGRESS"
	SdxClusterDetailResponseStatusREPAIRINPROGRESS string = "REPAIR_IN_PROGRESS"

	// SdxClusterDetailResponseStatusREPAIRFAILED captures enum value "REPAIR_FAILED"
	SdxClusterDetailResponseStatusREPAIRFAILED string = "REPAIR_FAILED"

	// SdxClusterDetailResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	SdxClusterDetailResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// SdxClusterDetailResponseStatusDELETED captures enum value "DELETED"
	SdxClusterDetailResponseStatusDELETED string = "DELETED"

	// SdxClusterDetailResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SdxClusterDetailResponseStatusDELETEFAILED string = "DELETE_FAILED"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *SdxClusterDetailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxClusterDetailResponse) UnmarshalBinary(b []byte) error {
	var res SdxClusterDetailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
