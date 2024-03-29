// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DetachRecipeV4Response detach recipe v4 response
// swagger:model DetachRecipeV4Response
type DetachRecipeV4Response struct {

	// host group name
	HostGroupName string `json:"hostGroupName,omitempty"`

	// recipe name
	RecipeName string `json:"recipeName,omitempty"`
}

// Validate validates this detach recipe v4 response
func (m *DetachRecipeV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetachRecipeV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetachRecipeV4Response) UnmarshalBinary(b []byte) error {
	var res DetachRecipeV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
