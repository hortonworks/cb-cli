// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IPPoolJSON Ip pool Json
// swagger:model IpPoolJson

type IPPoolJSON struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

/* polymorph IpPoolJson id false */

/* polymorph IpPoolJson name false */

/* polymorph IpPoolJson properties false */

// Validate validates this Ip pool Json
func (m *IPPoolJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IPPoolJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPPoolJSON) UnmarshalBinary(b []byte) error {
	var res IPPoolJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
