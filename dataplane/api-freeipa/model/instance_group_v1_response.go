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

// InstanceGroupV1Response instance group v1 response
// swagger:model InstanceGroupV1Response
type InstanceGroupV1Response struct {

	// instancegroup related template
	// Required: true
	InstanceTemplate *InstanceTemplateV1Response `json:"instanceTemplate"`

	// meta data
	// Unique: true
	MetaData []*InstanceMetaDataV1Response `json:"metaData"`

	// name of the instance group
	// Required: true
	Name *string `json:"name"`

	// referenced network
	// Required: true
	Network *InstanceGroupNetworkV1Response `json:"network"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// instancegroup related securitygroup
	SecurityGroup *SecurityGroupV1Response `json:"securityGroup,omitempty"`

	// type of the instance group, default value is CORE
	// Enum: [MASTER SLAVE]
	Type string `json:"type,omitempty"`
}

// Validate validates this instance group v1 response
func (m *InstanceGroupV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupV1Response) validateInstanceTemplate(formats strfmt.Registry) error {

	if err := validate.Required("instanceTemplate", "body", m.InstanceTemplate); err != nil {
		return err
	}

	if m.InstanceTemplate != nil {
		if err := m.InstanceTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instanceTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupV1Response) validateMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if err := validate.UniqueItems("metaData", "body", m.MetaData); err != nil {
		return err
	}

	for i := 0; i < len(m.MetaData); i++ {
		if swag.IsZero(m.MetaData[i]) { // not required
			continue
		}

		if m.MetaData[i] != nil {
			if err := m.MetaData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metaData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupV1Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV1Response) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupV1Response) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(*m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(*m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV1Response) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

var instanceGroupV1ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MASTER","SLAVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV1ResponseTypeTypePropEnum = append(instanceGroupV1ResponseTypeTypePropEnum, v)
	}
}

const (

	// InstanceGroupV1ResponseTypeMASTER captures enum value "MASTER"
	InstanceGroupV1ResponseTypeMASTER string = "MASTER"

	// InstanceGroupV1ResponseTypeSLAVE captures enum value "SLAVE"
	InstanceGroupV1ResponseTypeSLAVE string = "SLAVE"
)

// prop value enum
func (m *InstanceGroupV1Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV1ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV1Response) validateType(formats strfmt.Registry) error {

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
func (m *InstanceGroupV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupV1Response) UnmarshalBinary(b []byte) error {
	var res InstanceGroupV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
