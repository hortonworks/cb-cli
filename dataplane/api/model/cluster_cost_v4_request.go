// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterCostV4Request cluster cost v4 request
// swagger:model ClusterCostV4Request
type ClusterCostV4Request struct {

	// cluster crns
	ClusterCrns []string `json:"clusterCrns"`

	// environment crns
	EnvironmentCrns []string `json:"environmentCrns"`
}

// Validate validates this cluster cost v4 request
func (m *ClusterCostV4Request) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCostV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCostV4Request) UnmarshalBinary(b []byte) error {
	var res ClusterCostV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
