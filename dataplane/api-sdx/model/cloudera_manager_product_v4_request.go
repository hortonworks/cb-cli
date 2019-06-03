// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClouderaManagerProductV4Request cloudera manager product v4 request
// swagger:model ClouderaManagerProductV4Request
type ClouderaManagerProductV4Request struct {

	// name of the Cloudera manager product
	// Required: true
	Name *string `json:"name"`

	// parcel url of the Cloudera manager product
	Parcel string `json:"parcel,omitempty"`

	// version of the Cloudera manager product
	Version string `json:"version,omitempty"`
}

// Validate validates this cloudera manager product v4 request
func (m *ClouderaManagerProductV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClouderaManagerProductV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClouderaManagerProductV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClouderaManagerProductV4Request) UnmarshalBinary(b []byte) error {
	var res ClouderaManagerProductV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
