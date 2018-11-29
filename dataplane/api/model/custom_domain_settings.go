// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustomDomainSettings custom domain settings
// swagger:model CustomDomainSettings
type CustomDomainSettings struct {

	// using the cluster name to create subdomain
	ClusterNameAsSubdomain *bool `json:"clusterNameAsSubdomain,omitempty"`

	// custom domain name for the nodes in the stack
	CustomDomain string `json:"customDomain,omitempty"`

	// custom hostname for nodes in the stack
	CustomHostname string `json:"customHostname,omitempty"`

	// using the hostgroup names to create hostnames
	HostgroupNameAsHostname *bool `json:"hostgroupNameAsHostname,omitempty"`
}

// Validate validates this custom domain settings
func (m *CustomDomainSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomDomainSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomDomainSettings) UnmarshalBinary(b []byte) error {
	var res CustomDomainSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
