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

// FreeIpaUpgradeOptionsV1 free ipa upgrade options v1
// swagger:model FreeIpaUpgradeOptionsV1
type FreeIpaUpgradeOptionsV1 struct {

	// current image
	CurrentImage *ImageInfoResponse `json:"currentImage,omitempty"`

	// images
	Images []*ImageInfoResponse `json:"images"`
}

// Validate validates this free ipa upgrade options v1
func (m *FreeIpaUpgradeOptionsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FreeIpaUpgradeOptionsV1) validateCurrentImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentImage) { // not required
		return nil
	}

	if m.CurrentImage != nil {
		if err := m.CurrentImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentImage")
			}
			return err
		}
	}

	return nil
}

func (m *FreeIpaUpgradeOptionsV1) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FreeIpaUpgradeOptionsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeIpaUpgradeOptionsV1) UnmarshalBinary(b []byte) error {
	var res FreeIpaUpgradeOptionsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
