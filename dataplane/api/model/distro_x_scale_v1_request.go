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

// DistroXScaleV1Request distro x scale v1 request
// swagger:model DistroXScaleV1Request
type DistroXScaleV1Request struct {

	// scaling adjustment type
	// Enum: [EXACT PERCENTAGE BEST_EFFORT]
	AdjustmentType string `json:"adjustmentType,omitempty"`

	// scaling adjustment of the instance groups
	// Required: true
	DesiredCount *int32 `json:"desiredCount"`

	// name of the instance group
	// Required: true
	Group *string `json:"group"`

	// request to support scaling operations in multi-availability zone environments
	NetworkScaleRequest *NetworkScaleV1Request `json:"networkScaleRequest,omitempty"`

	// scaling threshold
	Threshold int64 `json:"threshold,omitempty"`
}

// Validate validates this distro x scale v1 request
func (m *DistroXScaleV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdjustmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkScaleRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var distroXScaleV1RequestTypeAdjustmentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","PERCENTAGE","BEST_EFFORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		distroXScaleV1RequestTypeAdjustmentTypePropEnum = append(distroXScaleV1RequestTypeAdjustmentTypePropEnum, v)
	}
}

const (

	// DistroXScaleV1RequestAdjustmentTypeEXACT captures enum value "EXACT"
	DistroXScaleV1RequestAdjustmentTypeEXACT string = "EXACT"

	// DistroXScaleV1RequestAdjustmentTypePERCENTAGE captures enum value "PERCENTAGE"
	DistroXScaleV1RequestAdjustmentTypePERCENTAGE string = "PERCENTAGE"

	// DistroXScaleV1RequestAdjustmentTypeBESTEFFORT captures enum value "BEST_EFFORT"
	DistroXScaleV1RequestAdjustmentTypeBESTEFFORT string = "BEST_EFFORT"
)

// prop value enum
func (m *DistroXScaleV1Request) validateAdjustmentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, distroXScaleV1RequestTypeAdjustmentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DistroXScaleV1Request) validateAdjustmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.AdjustmentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdjustmentTypeEnum("adjustmentType", "body", m.AdjustmentType); err != nil {
		return err
	}

	return nil
}

func (m *DistroXScaleV1Request) validateDesiredCount(formats strfmt.Registry) error {

	if err := validate.Required("desiredCount", "body", m.DesiredCount); err != nil {
		return err
	}

	return nil
}

func (m *DistroXScaleV1Request) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *DistroXScaleV1Request) validateNetworkScaleRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkScaleRequest) { // not required
		return nil
	}

	if m.NetworkScaleRequest != nil {
		if err := m.NetworkScaleRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkScaleRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistroXScaleV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistroXScaleV1Request) UnmarshalBinary(b []byte) error {
	var res DistroXScaleV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
