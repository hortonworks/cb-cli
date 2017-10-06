// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityGroupRequest security group request
// swagger:model SecurityGroupRequest

type SecurityGroupRequest struct {

	// type of cloud provider
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// name of the resource
	// Required: true
	// Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	Name *string `json:"name"`

	// Exisiting security group id
	SecurityGroupID string `json:"securityGroupId,omitempty"`

	// list of security rules that relates to the security group
	SecurityRules []*SecurityRuleRequest `json:"securityRules"`
}

/* polymorph SecurityGroupRequest cloudPlatform false */

/* polymorph SecurityGroupRequest description false */

/* polymorph SecurityGroupRequest name false */

/* polymorph SecurityGroupRequest securityGroupId false */

/* polymorph SecurityGroupRequest securityRules false */

// Validate validates this security group request
func (m *SecurityGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupRequest) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupRequest) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupRequest) validateSecurityRules(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityRules) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityRules); i++ {

		if swag.IsZero(m.SecurityRules[i]) { // not required
			continue
		}

		if m.SecurityRules[i] != nil {

			if err := m.SecurityRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupRequest) UnmarshalBinary(b []byte) error {
	var res SecurityGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
