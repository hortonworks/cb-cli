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

// StackImageChangeRequest stack image change request
// swagger:model StackImageChangeRequest
type StackImageChangeRequest struct {

	// custom image catalog URL
	ImageCatalogName string `json:"imageCatalogName,omitempty"`

	// virtual machine image id from ImageCatalog, machines of the cluster will be started from this image
	// Required: true
	// Max Length: 2147483647
	// Min Length: 1
	ImageID *string `json:"imageId"`
}

// Validate validates this stack image change request
func (m *StackImageChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackImageChangeRequest) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageId", "body", m.ImageID); err != nil {
		return err
	}

	if err := validate.MinLength("imageId", "body", string(*m.ImageID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("imageId", "body", string(*m.ImageID), 2147483647); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackImageChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackImageChangeRequest) UnmarshalBinary(b []byte) error {
	var res StackImageChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
