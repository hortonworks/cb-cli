// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterTemplateResponse cluster template response
// swagger:model ClusterTemplateResponse
type ClusterTemplateResponse struct {

	// cloudplatform which this template is compatible with
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// status
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED]
	Status string `json:"status,omitempty"`

	// stringified template JSON
	// Required: true
	Template *StackV2Request `json:"template"`
}

// Validate validates this cluster template response
func (m *ClusterTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTemplateResponse) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterTemplateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

var clusterTemplateResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateResponseTypeStatusPropEnum = append(clusterTemplateResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterTemplateResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterTemplateResponseStatusDEFAULT string = "DEFAULT"

	// ClusterTemplateResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterTemplateResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterTemplateResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterTemplateResponseStatusUSERMANAGED string = "USER_MANAGED"
)

// prop value enum
func (m *ClusterTemplateResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateResponse) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplateResponse) UnmarshalBinary(b []byte) error {
	var res ClusterTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
