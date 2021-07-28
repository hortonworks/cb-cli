// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InstanceGroupAwsNetworkParameters instance group aws network parameters
// swagger:model InstanceGroupAwsNetworkParameters
type InstanceGroupAwsNetworkParameters struct {

	// subnet ids
	SubnetIds []string `json:"subnetIds"`
}

// Validate validates this instance group aws network parameters
func (m *InstanceGroupAwsNetworkParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupAwsNetworkParameters) UnmarshalBinary(b []byte) error {
	var res InstanceGroupAwsNetworkParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
