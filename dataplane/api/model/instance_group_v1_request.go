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

// InstanceGroupV1Request instance group v1 request
// swagger:model InstanceGroupV1Request
type InstanceGroupV1Request struct {

	// aws specific parameters for instance group
	Aws AwsInstanceGroupV1Parameters `json:"aws,omitempty"`

	// azure specific parameters for instance group
	Azure *AzureInstanceGroupV1Parameters `json:"azure,omitempty"`

	// cloud platform
	// Enum: [AWS GCP AZURE YARN MOCK OPENSTACK]
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// gcp
	Gcp GcpInstanceGroupV1Parameters `json:"gcp,omitempty"`

	// minimum nodecount in an instancegroup
	MinimumNodeCount int32 `json:"minimumNodeCount,omitempty"`

	// name of the instance group
	// Required: true
	Name *string `json:"name"`

	// referenced network
	Network *InstanceGroupNetworkV1Request `json:"network,omitempty"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// referenced recipe names
	// Unique: true
	RecipeNames []string `json:"recipeNames"`

	// recovery mode of the hostgroup's nodes
	// Enum: [MANUAL AUTO]
	RecoveryMode string `json:"recoveryMode,omitempty"`

	// type of the instance group scalability, default value is ALLOWED
	// Enum: [ALLOWED FORBIDDEN ONLY_UPSCALE ONLY_DOWNSCALE]
	ScalabilityOption string `json:"scalabilityOption,omitempty"`

	// instancegroup related template
	// Required: true
	Template *InstanceTemplateV1Request `json:"template"`

	// type of the instance group, default value is CORE
	// Enum: [CORE GATEWAY]
	Type string `json:"type,omitempty"`

	// yarn
	Yarn YarnInstanceGroupV1Parameters `json:"yarn,omitempty"`
}

// Validate validates this instance group v1 request
func (m *InstanceGroupV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
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

	if err := m.validateRecipeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalabilityOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
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

func (m *InstanceGroupV1Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

var instanceGroupV1RequestTypeCloudPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP","AZURE","YARN","MOCK","OPENSTACK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV1RequestTypeCloudPlatformPropEnum = append(instanceGroupV1RequestTypeCloudPlatformPropEnum, v)
	}
}

const (

	// InstanceGroupV1RequestCloudPlatformAWS captures enum value "AWS"
	InstanceGroupV1RequestCloudPlatformAWS string = "AWS"

	// InstanceGroupV1RequestCloudPlatformGCP captures enum value "GCP"
	InstanceGroupV1RequestCloudPlatformGCP string = "GCP"

	// InstanceGroupV1RequestCloudPlatformAZURE captures enum value "AZURE"
	InstanceGroupV1RequestCloudPlatformAZURE string = "AZURE"

	// InstanceGroupV1RequestCloudPlatformYARN captures enum value "YARN"
	InstanceGroupV1RequestCloudPlatformYARN string = "YARN"

	// InstanceGroupV1RequestCloudPlatformMOCK captures enum value "MOCK"
	InstanceGroupV1RequestCloudPlatformMOCK string = "MOCK"

	// InstanceGroupV1RequestCloudPlatformOPENSTACK captures enum value "OPENSTACK"
	InstanceGroupV1RequestCloudPlatformOPENSTACK string = "OPENSTACK"
)

// prop value enum
func (m *InstanceGroupV1Request) validateCloudPlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV1RequestTypeCloudPlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV1Request) validateCloudPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudPlatformEnum("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV1Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
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

func (m *InstanceGroupV1Request) validateNodeCount(formats strfmt.Registry) error {

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

func (m *InstanceGroupV1Request) validateRecipeNames(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipeNames", "body", m.RecipeNames); err != nil {
		return err
	}

	return nil
}

var instanceGroupV1RequestTypeRecoveryModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV1RequestTypeRecoveryModePropEnum = append(instanceGroupV1RequestTypeRecoveryModePropEnum, v)
	}
}

const (

	// InstanceGroupV1RequestRecoveryModeMANUAL captures enum value "MANUAL"
	InstanceGroupV1RequestRecoveryModeMANUAL string = "MANUAL"

	// InstanceGroupV1RequestRecoveryModeAUTO captures enum value "AUTO"
	InstanceGroupV1RequestRecoveryModeAUTO string = "AUTO"
)

// prop value enum
func (m *InstanceGroupV1Request) validateRecoveryModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV1RequestTypeRecoveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV1Request) validateRecoveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryModeEnum("recoveryMode", "body", m.RecoveryMode); err != nil {
		return err
	}

	return nil
}

var instanceGroupV1RequestTypeScalabilityOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOWED","FORBIDDEN","ONLY_UPSCALE","ONLY_DOWNSCALE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV1RequestTypeScalabilityOptionPropEnum = append(instanceGroupV1RequestTypeScalabilityOptionPropEnum, v)
	}
}

const (

	// InstanceGroupV1RequestScalabilityOptionALLOWED captures enum value "ALLOWED"
	InstanceGroupV1RequestScalabilityOptionALLOWED string = "ALLOWED"

	// InstanceGroupV1RequestScalabilityOptionFORBIDDEN captures enum value "FORBIDDEN"
	InstanceGroupV1RequestScalabilityOptionFORBIDDEN string = "FORBIDDEN"

	// InstanceGroupV1RequestScalabilityOptionONLYUPSCALE captures enum value "ONLY_UPSCALE"
	InstanceGroupV1RequestScalabilityOptionONLYUPSCALE string = "ONLY_UPSCALE"

	// InstanceGroupV1RequestScalabilityOptionONLYDOWNSCALE captures enum value "ONLY_DOWNSCALE"
	InstanceGroupV1RequestScalabilityOptionONLYDOWNSCALE string = "ONLY_DOWNSCALE"
)

// prop value enum
func (m *InstanceGroupV1Request) validateScalabilityOptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV1RequestTypeScalabilityOptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV1Request) validateScalabilityOption(formats strfmt.Registry) error {

	if swag.IsZero(m.ScalabilityOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateScalabilityOptionEnum("scalabilityOption", "body", m.ScalabilityOption); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV1Request) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

var instanceGroupV1RequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CORE","GATEWAY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV1RequestTypeTypePropEnum = append(instanceGroupV1RequestTypeTypePropEnum, v)
	}
}

const (

	// InstanceGroupV1RequestTypeCORE captures enum value "CORE"
	InstanceGroupV1RequestTypeCORE string = "CORE"

	// InstanceGroupV1RequestTypeGATEWAY captures enum value "GATEWAY"
	InstanceGroupV1RequestTypeGATEWAY string = "GATEWAY"
)

// prop value enum
func (m *InstanceGroupV1Request) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV1RequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV1Request) validateType(formats strfmt.Registry) error {

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
func (m *InstanceGroupV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupV1Request) UnmarshalBinary(b []byte) error {
	var res InstanceGroupV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
