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

// LoadBalancerResponse load balancer response
// swagger:model LoadBalancerResponse
type LoadBalancerResponse struct {

	// The AWS resource id for the load balancer.
	AwsResourceID *AwsLoadBalancerResponse `json:"awsResourceId,omitempty"`

	// The Azure load balancer cloud resource information.
	AzureResourceID *AzureLoadBalancerResponse `json:"azureResourceId,omitempty"`

	// The AWS generated DNS name for the load balancer.
	CloudDNS string `json:"cloudDns,omitempty"`

	// The registered FQDN of the load balancer.
	Fqdn string `json:"fqdn,omitempty"`

	// The GCP load balancer information
	GcpResourceID *GcpLoadBalancerResponse `json:"gcpResourceId,omitempty"`

	// The frontend ip address for the load balancer.
	IP string `json:"ip,omitempty"`

	// The list of target instances the load balancer routes traffic to.
	// Required: true
	Targets []*TargetGroupResponse `json:"targets"`

	// Whether the load balancer is internet-facing (public), or only accessible over private endpoints.
	// Required: true
	// Enum: [PUBLIC PRIVATE GATEWAY_PRIVATE OUTBOUND]
	Type *string `json:"type"`
}

// Validate validates this load balancer response
func (m *LoadBalancerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
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

func (m *LoadBalancerResponse) validateAwsResourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsResourceID) { // not required
		return nil
	}

	if m.AwsResourceID != nil {
		if err := m.AwsResourceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsResourceId")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerResponse) validateAzureResourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureResourceID) { // not required
		return nil
	}

	if m.AzureResourceID != nil {
		if err := m.AzureResourceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureResourceId")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerResponse) validateGcpResourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.GcpResourceID) { // not required
		return nil
	}

	if m.GcpResourceID != nil {
		if err := m.GcpResourceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpResourceId")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerResponse) validateTargets(formats strfmt.Registry) error {

	if err := validate.Required("targets", "body", m.Targets); err != nil {
		return err
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var loadBalancerResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE","GATEWAY_PRIVATE","OUTBOUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loadBalancerResponseTypeTypePropEnum = append(loadBalancerResponseTypeTypePropEnum, v)
	}
}

const (

	// LoadBalancerResponseTypePUBLIC captures enum value "PUBLIC"
	LoadBalancerResponseTypePUBLIC string = "PUBLIC"

	// LoadBalancerResponseTypePRIVATE captures enum value "PRIVATE"
	LoadBalancerResponseTypePRIVATE string = "PRIVATE"

	// LoadBalancerResponseTypeGATEWAYPRIVATE captures enum value "GATEWAY_PRIVATE"
	LoadBalancerResponseTypeGATEWAYPRIVATE string = "GATEWAY_PRIVATE"

	// LoadBalancerResponseTypeOUTBOUND captures enum value "OUTBOUND"
	LoadBalancerResponseTypeOUTBOUND string = "OUTBOUND"
)

// prop value enum
func (m *LoadBalancerResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loadBalancerResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LoadBalancerResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerResponse) UnmarshalBinary(b []byte) error {
	var res LoadBalancerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
