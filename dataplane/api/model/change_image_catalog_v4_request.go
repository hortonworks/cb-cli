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

// ChangeImageCatalogV4Request change image catalog v4 request
// swagger:model ChangeImageCatalogV4Request
type ChangeImageCatalogV4Request struct {

	// name of the image catalog
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	ImageCatalog *string `json:"imageCatalog"`
}

// Validate validates this change image catalog v4 request
func (m *ChangeImageCatalogV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageCatalog(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeImageCatalogV4Request) validateImageCatalog(formats strfmt.Registry) error {

	if err := validate.Required("imageCatalog", "body", m.ImageCatalog); err != nil {
		return err
	}

	if err := validate.MinLength("imageCatalog", "body", string(*m.ImageCatalog), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("imageCatalog", "body", string(*m.ImageCatalog), 100); err != nil {
		return err
	}

	if err := validate.Pattern("imageCatalog", "body", string(*m.ImageCatalog), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeImageCatalogV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeImageCatalogV4Request) UnmarshalBinary(b []byte) error {
	var res ChangeImageCatalogV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
