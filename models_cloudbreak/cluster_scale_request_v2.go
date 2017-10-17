// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterScaleRequestV2 cluster scale request v2
// swagger:model ClusterScaleRequestV2

type ClusterScaleRequestV2 struct {

	// scaling adjustment of the host groups
	// Required: true
	DesiredCount *int32 `json:"desiredCount"`

	// name of the host group
	// Required: true
	Group *string `json:"group"`

	// validate node count during downscale
	ValidateNodeCount *bool `json:"validateNodeCount,omitempty"`

	// on cluster update, update stack too
	WithStackUpdate *bool `json:"withStackUpdate,omitempty"`
}

/* polymorph ClusterScaleRequestV2 desiredCount false */

/* polymorph ClusterScaleRequestV2 group false */

/* polymorph ClusterScaleRequestV2 validateNodeCount false */

/* polymorph ClusterScaleRequestV2 withStackUpdate false */

// Validate validates this cluster scale request v2
func (m *ClusterScaleRequestV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterScaleRequestV2) validateDesiredCount(formats strfmt.Registry) error {

	if err := validate.Required("desiredCount", "body", m.DesiredCount); err != nil {
		return err
	}

	return nil
}

func (m *ClusterScaleRequestV2) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterScaleRequestV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterScaleRequestV2) UnmarshalBinary(b []byte) error {
	var res ClusterScaleRequestV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
