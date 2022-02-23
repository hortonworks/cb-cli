// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OperationV1Status operation v1 status
// swagger:model OperationV1Status
type OperationV1Status struct {

	// User synchronization operation end time
	EndTime int64 `json:"endTime,omitempty"`

	// error information about operation failure
	Error string `json:"error,omitempty"`

	// details about environments where operation failed
	Failure []*FailureDetailsV1 `json:"failure"`

	// User synchronization operation id
	// Required: true
	OperationID *string `json:"operationId"`

	// Operation type
	// Required: true
	// Enum: [USER_SYNC SET_PASSWORD CLEANUP REBOOT REPAIR DOWNSCALE UPSCALE BIND_USER_CREATE UPGRADE UPGRADE_CCM]
	OperationType *string `json:"operationType"`

	// User synchronization operation start time
	StartTime int64 `json:"startTime,omitempty"`

	// User synchronization operation status
	// Enum: [REQUESTED RUNNING COMPLETED FAILED REJECTED TIMEDOUT]
	Status string `json:"status,omitempty"`

	// details about environments where operation succeeded
	Success []*SuccessDetailsV1 `json:"success"`
}

// Validate validates this operation v1 status
func (m *OperationV1Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationV1Status) validateFailure(formats strfmt.Registry) error {

	if swag.IsZero(m.Failure) { // not required
		return nil
	}

	for i := 0; i < len(m.Failure); i++ {
		if swag.IsZero(m.Failure[i]) { // not required
			continue
		}

		if m.Failure[i] != nil {
			if err := m.Failure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OperationV1Status) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

var operationV1StatusTypeOperationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER_SYNC","SET_PASSWORD","CLEANUP","REBOOT","REPAIR","DOWNSCALE","UPSCALE","BIND_USER_CREATE","UPGRADE","UPGRADE_CCM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationV1StatusTypeOperationTypePropEnum = append(operationV1StatusTypeOperationTypePropEnum, v)
	}
}

const (

	// OperationV1StatusOperationTypeUSERSYNC captures enum value "USER_SYNC"
	OperationV1StatusOperationTypeUSERSYNC string = "USER_SYNC"

	// OperationV1StatusOperationTypeSETPASSWORD captures enum value "SET_PASSWORD"
	OperationV1StatusOperationTypeSETPASSWORD string = "SET_PASSWORD"

	// OperationV1StatusOperationTypeCLEANUP captures enum value "CLEANUP"
	OperationV1StatusOperationTypeCLEANUP string = "CLEANUP"

	// OperationV1StatusOperationTypeREBOOT captures enum value "REBOOT"
	OperationV1StatusOperationTypeREBOOT string = "REBOOT"

	// OperationV1StatusOperationTypeREPAIR captures enum value "REPAIR"
	OperationV1StatusOperationTypeREPAIR string = "REPAIR"

	// OperationV1StatusOperationTypeDOWNSCALE captures enum value "DOWNSCALE"
	OperationV1StatusOperationTypeDOWNSCALE string = "DOWNSCALE"

	// OperationV1StatusOperationTypeUPSCALE captures enum value "UPSCALE"
	OperationV1StatusOperationTypeUPSCALE string = "UPSCALE"

	// OperationV1StatusOperationTypeBINDUSERCREATE captures enum value "BIND_USER_CREATE"
	OperationV1StatusOperationTypeBINDUSERCREATE string = "BIND_USER_CREATE"

	// OperationV1StatusOperationTypeUPGRADE captures enum value "UPGRADE"
	OperationV1StatusOperationTypeUPGRADE string = "UPGRADE"

	// OperationV1StatusOperationTypeUPGRADECCM captures enum value "UPGRADE_CCM"
	OperationV1StatusOperationTypeUPGRADECCM string = "UPGRADE_CCM"
)

// prop value enum
func (m *OperationV1Status) validateOperationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, operationV1StatusTypeOperationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OperationV1Status) validateOperationType(formats strfmt.Registry) error {

	if err := validate.Required("operationType", "body", m.OperationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationTypeEnum("operationType", "body", *m.OperationType); err != nil {
		return err
	}

	return nil
}

var operationV1StatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","RUNNING","COMPLETED","FAILED","REJECTED","TIMEDOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationV1StatusTypeStatusPropEnum = append(operationV1StatusTypeStatusPropEnum, v)
	}
}

const (

	// OperationV1StatusStatusREQUESTED captures enum value "REQUESTED"
	OperationV1StatusStatusREQUESTED string = "REQUESTED"

	// OperationV1StatusStatusRUNNING captures enum value "RUNNING"
	OperationV1StatusStatusRUNNING string = "RUNNING"

	// OperationV1StatusStatusCOMPLETED captures enum value "COMPLETED"
	OperationV1StatusStatusCOMPLETED string = "COMPLETED"

	// OperationV1StatusStatusFAILED captures enum value "FAILED"
	OperationV1StatusStatusFAILED string = "FAILED"

	// OperationV1StatusStatusREJECTED captures enum value "REJECTED"
	OperationV1StatusStatusREJECTED string = "REJECTED"

	// OperationV1StatusStatusTIMEDOUT captures enum value "TIMEDOUT"
	OperationV1StatusStatusTIMEDOUT string = "TIMEDOUT"
)

// prop value enum
func (m *OperationV1Status) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, operationV1StatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OperationV1Status) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *OperationV1Status) validateSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Success) { // not required
		return nil
	}

	for i := 0; i < len(m.Success); i++ {
		if swag.IsZero(m.Success[i]) { // not required
			continue
		}

		if m.Success[i] != nil {
			if err := m.Success[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationV1Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationV1Status) UnmarshalBinary(b []byte) error {
	var res OperationV1Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
