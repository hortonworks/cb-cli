// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SdxRepairRequest sdx repair request
// swagger:model SdxRepairRequest
type SdxRepairRequest struct {

	// host group name
	HostGroupName string `json:"hostGroupName,omitempty"`

	// host group names
	HostGroupNames []string `json:"hostGroupNames"`

	// nodes ids
	NodesIds []string `json:"nodesIds"`
}

// Validate validates this sdx repair request
func (m *SdxRepairRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SdxRepairRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxRepairRequest) UnmarshalBinary(b []byte) error {
	var res SdxRepairRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
