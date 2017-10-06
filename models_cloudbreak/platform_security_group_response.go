// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PlatformSecurityGroupResponse platform security group response
// swagger:model PlatformSecurityGroupResponse

type PlatformSecurityGroupResponse struct {

	// group Id
	GroupID string `json:"groupId,omitempty"`

	// group name
	GroupName string `json:"groupName,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

/* polymorph PlatformSecurityGroupResponse groupId false */

/* polymorph PlatformSecurityGroupResponse groupName false */

/* polymorph PlatformSecurityGroupResponse properties false */

// Validate validates this platform security group response
func (m *PlatformSecurityGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PlatformSecurityGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformSecurityGroupResponse) UnmarshalBinary(b []byte) error {
	var res PlatformSecurityGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
