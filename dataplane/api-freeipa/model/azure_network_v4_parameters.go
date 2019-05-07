// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureNetworkV4Parameters azure network v4 parameters
// swagger:model AzureNetworkV4Parameters
type AzureNetworkV4Parameters struct {

	// network Id
	NetworkID string `json:"networkId,omitempty"`

	// no firewall rules
	NoFirewallRules bool `json:"noFirewallRules,omitempty"`

	// no public Ip
	NoPublicIP bool `json:"noPublicIp,omitempty"`

	// resource group name
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// subnet Id
	SubnetID string `json:"subnetId,omitempty"`
}

// Validate validates this azure network v4 parameters
func (m *AzureNetworkV4Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureNetworkV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureNetworkV4Parameters) UnmarshalBinary(b []byte) error {
	var res AzureNetworkV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
