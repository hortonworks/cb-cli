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

// SdxInternalClusterRequest sdx internal cluster request
// swagger:model SdxInternalClusterRequest
type SdxInternalClusterRequest struct {

	// AWS options.
	Aws *SdxAwsRequest `json:"aws,omitempty"`

	// Azure options.
	Azure *SdxAzureRequest `json:"azure,omitempty"`

	// Details about the cloud storage type and location.
	CloudStorage *SdxCloudStorageRequest `json:"cloudStorage,omitempty"`

	// The shape of the cluster such as Micro Duty, Light Duty, Medium Duty...
	// Required: true
	// Enum: [CUSTOM LIGHT_DUTY MEDIUM_DUTY_HA ENTERPRISE MICRO_DUTY]
	ClusterShape *string `json:"clusterShape"`

	// Custom instance group options.
	CustomInstanceGroups []*SdxInstanceGroupRequest `json:"customInstanceGroups"`

	// Option to enable multi AZ.
	EnableMultiAz bool `json:"enableMultiAz,omitempty"`

	// Option to enable ranger raz.
	EnableRangerRaz bool `json:"enableRangerRaz,omitempty"`

	// The name of the environment.
	// Required: true
	Environment *string `json:"environment"`

	// External database options.
	ExternalDatabase *SdxDatabaseRequest `json:"externalDatabase,omitempty"`

	// Java version to be forced on virtual machines
	JavaVersion int32 `json:"javaVersion,omitempty"`

	// Recipes.
	// Unique: true
	Recipes []*SdxRecipe `json:"recipes"`

	// Runtime version.
	Runtime string `json:"runtime,omitempty"`

	// Stack request.
	StackV4Request *StackV4Request `json:"stackV4Request,omitempty"`

	// Tags.
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sdx internal cluster request
func (m *SdxInternalClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackV4Request(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxInternalClusterRequest) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateAzure(formats strfmt.Registry) error {

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

func (m *SdxInternalClusterRequest) validateCloudStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorage) { // not required
		return nil
	}

	if m.CloudStorage != nil {
		if err := m.CloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudStorage")
			}
			return err
		}
	}

	return nil
}

var sdxInternalClusterRequestTypeClusterShapePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","LIGHT_DUTY","MEDIUM_DUTY_HA","ENTERPRISE","MICRO_DUTY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxInternalClusterRequestTypeClusterShapePropEnum = append(sdxInternalClusterRequestTypeClusterShapePropEnum, v)
	}
}

const (

	// SdxInternalClusterRequestClusterShapeCUSTOM captures enum value "CUSTOM"
	SdxInternalClusterRequestClusterShapeCUSTOM string = "CUSTOM"

	// SdxInternalClusterRequestClusterShapeLIGHTDUTY captures enum value "LIGHT_DUTY"
	SdxInternalClusterRequestClusterShapeLIGHTDUTY string = "LIGHT_DUTY"

	// SdxInternalClusterRequestClusterShapeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	SdxInternalClusterRequestClusterShapeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"

	// SdxInternalClusterRequestClusterShapeENTERPRISE captures enum value "ENTERPRISE"
	SdxInternalClusterRequestClusterShapeENTERPRISE string = "ENTERPRISE"

	// SdxInternalClusterRequestClusterShapeMICRODUTY captures enum value "MICRO_DUTY"
	SdxInternalClusterRequestClusterShapeMICRODUTY string = "MICRO_DUTY"
)

// prop value enum
func (m *SdxInternalClusterRequest) validateClusterShapeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxInternalClusterRequestTypeClusterShapePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxInternalClusterRequest) validateClusterShape(formats strfmt.Registry) error {

	if err := validate.Required("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	// value enum
	if err := m.validateClusterShapeEnum("clusterShape", "body", *m.ClusterShape); err != nil {
		return err
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateCustomInstanceGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomInstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomInstanceGroups); i++ {
		if swag.IsZero(m.CustomInstanceGroups[i]) { // not required
			continue
		}

		if m.CustomInstanceGroups[i] != nil {
			if err := m.CustomInstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SdxInternalClusterRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateExternalDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDatabase) { // not required
		return nil
	}

	if m.ExternalDatabase != nil {
		if err := m.ExternalDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDatabase")
			}
			return err
		}
	}

	return nil
}

func (m *SdxInternalClusterRequest) validateRecipes(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipes) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipes", "body", m.Recipes); err != nil {
		return err
	}

	for i := 0; i < len(m.Recipes); i++ {
		if swag.IsZero(m.Recipes[i]) { // not required
			continue
		}

		if m.Recipes[i] != nil {
			if err := m.Recipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SdxInternalClusterRequest) validateStackV4Request(formats strfmt.Registry) error {

	if swag.IsZero(m.StackV4Request) { // not required
		return nil
	}

	if m.StackV4Request != nil {
		if err := m.StackV4Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackV4Request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxInternalClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxInternalClusterRequest) UnmarshalBinary(b []byte) error {
	var res SdxInternalClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
