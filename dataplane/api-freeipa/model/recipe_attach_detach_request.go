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

// RecipeAttachDetachRequest recipe attach detach request
// swagger:model RecipeAttachDetachRequest
type RecipeAttachDetachRequest struct {

	// Environment CRN
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Recipe names
	// Required: true
	Recipes []string `json:"recipes"`
}

// Validate validates this recipe attach detach request
func (m *RecipeAttachDetachRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeAttachDetachRequest) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *RecipeAttachDetachRequest) validateRecipes(formats strfmt.Registry) error {

	if err := validate.Required("recipes", "body", m.Recipes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecipeAttachDetachRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeAttachDetachRequest) UnmarshalBinary(b []byte) error {
	var res RecipeAttachDetachRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
