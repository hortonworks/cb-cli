// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InstanceGroupNetworkV4Request instance group network v4 request
// swagger:model InstanceGroupNetworkV4Request
type InstanceGroupNetworkV4Request struct {

	// provider specific parameters of the specified network
	Aws *InstanceGroupAwsNetworkV4Parameters `json:"aws,omitempty"`

	// provider specific parameters of the specified network
	Azure *InstanceGroupAzureNetworkV4Parameters `json:"azure,omitempty"`

	// provider specific parameters of the specified network
	Gcp *InstanceGroupGcpNetworkV4Parameters `json:"gcp,omitempty"`
}

// Validate validates this instance group network v4 request
func (m *InstanceGroupNetworkV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupNetworkV4Request) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupNetworkV4Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupNetworkV4Request) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupNetworkV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupNetworkV4Request) UnmarshalBinary(b []byte) error {
	var res InstanceGroupNetworkV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
