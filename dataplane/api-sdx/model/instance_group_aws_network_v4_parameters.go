// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InstanceGroupAwsNetworkV4Parameters instance group aws network v4 parameters
// swagger:model InstanceGroupAwsNetworkV4Parameters
type InstanceGroupAwsNetworkV4Parameters struct {

	// endpoint gateway subnet ids
	EndpointGatewaySubnetIds []string `json:"endpointGatewaySubnetIds"`

	// subnet ids
	SubnetIds []string `json:"subnetIds"`
}

// Validate validates this instance group aws network v4 parameters
func (m *InstanceGroupAwsNetworkV4Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkV4Parameters) UnmarshalBinary(b []byte) error {
	var res InstanceGroupAwsNetworkV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
