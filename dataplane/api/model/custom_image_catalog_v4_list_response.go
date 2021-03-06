// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustomImageCatalogV4ListResponse custom image catalog v4 list response
// swagger:model CustomImageCatalogV4ListResponse
type CustomImageCatalogV4ListResponse struct {

	// responses
	Responses []*CustomImageCatalogV4ListItemResponse `json:"responses"`
}

// Validate validates this custom image catalog v4 list response
func (m *CustomImageCatalogV4ListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomImageCatalogV4ListResponse) validateResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	for i := 0; i < len(m.Responses); i++ {
		if swag.IsZero(m.Responses[i]) { // not required
			continue
		}

		if m.Responses[i] != nil {
			if err := m.Responses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomImageCatalogV4ListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomImageCatalogV4ListResponse) UnmarshalBinary(b []byte) error {
	var res CustomImageCatalogV4ListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
