// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InstanceGroupAwsNetworkV1Parameters instance group aws network v1 parameters
// swagger:model InstanceGroupAwsNetworkV1Parameters
type InstanceGroupAwsNetworkV1Parameters struct {

	// endpoint gateway subnet ids
	EndpointGatewaySubnetIds []string `json:"endpointGatewaySubnetIds"`

	// subnet ids
	SubnetIds []string `json:"subnetIds"`
}

// Validate validates this instance group aws network v1 parameters
func (m *InstanceGroupAwsNetworkV1Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkV1Parameters) UnmarshalBinary(b []byte) error {
	var res InstanceGroupAwsNetworkV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
