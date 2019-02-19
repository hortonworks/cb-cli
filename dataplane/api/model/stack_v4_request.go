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

// StackV4Request stack v4 request
// swagger:model StackV4Request
type StackV4Request struct {

	// stack related authentication
	// Required: true
	Authentication *StackAuthenticationV4Request `json:"authentication"`

	// aws specific parameters for stack
	Aws AwsStackV4Parameters `json:"aws,omitempty"`

	// azure specific parameters for stack
	Azure *AzureStackV4Parameters `json:"azure,omitempty"`

	// cluster request object on stack
	Cluster *ClusterV4Request `json:"cluster,omitempty"`

	// settings related to custom domain names
	CustomDomain *CustomDomainSettingsV4Request `json:"customDomain,omitempty"`

	// general configuration parameters for a cluster (e.g. 'name', 'credentialname')
	// Required: true
	Environment *EnvironmentSettingsV4Request `json:"environment"`

	// id of the related flex subscription
	FlexID int64 `json:"flexId,omitempty"`

	// port of the gateway secured proxy
	// Maximum: 65535
	// Minimum: 1025
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// gcp specific parameters for stack
	Gcp GcpStackV4Parameters `json:"gcp,omitempty"`

	// settings for custom images
	Image *ImageSettingsV4Request `json:"image,omitempty"`

	// dynamic properties
	Inputs map[string]interface{} `json:"inputs,omitempty"`

	// collection of instance groupst
	// Required: true
	InstanceGroups []*InstanceGroupV4Request `json:"instanceGroups"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// stack related network
	Network *NetworkV4Request `json:"network,omitempty"`

	// openstack specific parameters for stack
	Openstack OpenStackStackV4Parameters `json:"openstack,omitempty"`

	// placement configuration parameters for a cluster (e.g. 'region', 'availabilityZone')
	Placement *PlacementSettingsV4Request `json:"placement,omitempty"`

	// Shared service request
	SharedService *SharedServiceV4Request `json:"sharedService,omitempty"`

	// stack related tags
	Tags *TagsV4Request `json:"tags,omitempty"`

	// time to live
	TimeToLive int64 `json:"timeToLive,omitempty"`

	// type
	// Enum: [WORKLOAD DATALAKE TEMPLATE]
	Type string `json:"type,omitempty"`

	// openstack specific parameters for stack
	Yarn *YarnStackV4Parameters `json:"yarn,omitempty"`
}

// Validate validates this stack v4 request
func (m *StackV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackV4Request) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateAzure(formats strfmt.Registry) error {

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

func (m *StackV4Request) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateCustomDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomDomain) { // not required
		return nil
	}

	if m.CustomDomain != nil {
		if err := m.CustomDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDomain")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateGatewayPort(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayPort) { // not required
		return nil
	}

	if err := validate.MinimumInt("gatewayPort", "body", int64(m.GatewayPort), 1025, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("gatewayPort", "body", int64(m.GatewayPort), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Request) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Request) validateNetwork(formats strfmt.Registry) error {

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

func (m *StackV4Request) validatePlacement(formats strfmt.Registry) error {

	if swag.IsZero(m.Placement) { // not required
		return nil
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateSharedService(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedService) { // not required
		return nil
	}

	if m.SharedService != nil {
		if err := m.SharedService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedService")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Request) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

var stackV4RequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WORKLOAD","DATALAKE","TEMPLATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackV4RequestTypeTypePropEnum = append(stackV4RequestTypeTypePropEnum, v)
	}
}

const (

	// StackV4RequestTypeWORKLOAD captures enum value "WORKLOAD"
	StackV4RequestTypeWORKLOAD string = "WORKLOAD"

	// StackV4RequestTypeDATALAKE captures enum value "DATALAKE"
	StackV4RequestTypeDATALAKE string = "DATALAKE"

	// StackV4RequestTypeTEMPLATE captures enum value "TEMPLATE"
	StackV4RequestTypeTEMPLATE string = "TEMPLATE"
)

// prop value enum
func (m *StackV4Request) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackV4RequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackV4Request) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Request) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackV4Request) UnmarshalBinary(b []byte) error {
	var res StackV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
