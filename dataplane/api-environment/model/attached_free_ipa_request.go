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

// AttachedFreeIpaRequest attached free ipa request
// swagger:model AttachedFreeIpaRequest
type AttachedFreeIpaRequest struct {

	// Aws specific FreeIpa parameters
	Aws *AttachedFreeIpaRequestAwsParameters `json:"aws,omitempty"`

	// Azure specific FreeIpa parameters
	Azure AttachedFreeIpaRequestAzureParameters `json:"azure,omitempty"`

	// Create freeipa in environment
	// Required: true
	Create *bool `json:"create"`

	// The FreeIPA multi-AZ enabled or not
	EnableMultiAz bool `json:"enableMultiAz,omitempty"`

	// Gcp specific FreeIpa parameters
	Gcp AttachedFreeIpaRequestGcpParameters `json:"gcp,omitempty"`

	// Image parameters for FreeIpa instance creation.
	Image *FreeIpaImageRequest `json:"image,omitempty"`

	// The number of FreeIPA instances to create per group when creating FreeIPA in environment
	InstanceCountByGroup int32 `json:"instanceCountByGroup,omitempty"`

	// Override default FreeIPA instance type
	InstanceType string `json:"instanceType,omitempty"`

	// FreeIpa recipe list
	// Unique: true
	Recipes []string `json:"recipes"`
}

// Validate validates this attached free ipa request
func (m *AttachedFreeIpaRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
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

func (m *AttachedFreeIpaRequest) validateAws(formats strfmt.Registry) error {

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

func (m *AttachedFreeIpaRequest) validateCreate(formats strfmt.Registry) error {

	if err := validate.Required("create", "body", m.Create); err != nil {
		return err
	}

	return nil
}

func (m *AttachedFreeIpaRequest) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *AttachedFreeIpaRequest) validateRecipes(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipes) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipes", "body", m.Recipes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttachedFreeIpaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedFreeIpaRequest) UnmarshalBinary(b []byte) error {
	var res AttachedFreeIpaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
