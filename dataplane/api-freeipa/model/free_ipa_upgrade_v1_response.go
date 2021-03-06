// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FreeIpaUpgradeV1Response free ipa upgrade v1 response
// swagger:model FreeIpaUpgradeV1Response
type FreeIpaUpgradeV1Response struct {

	// flow identifier
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// operation Id
	OperationID string `json:"operationId,omitempty"`

	// original image
	OriginalImage *ImageInfoResponse `json:"originalImage,omitempty"`

	// target image
	TargetImage *ImageInfoResponse `json:"targetImage,omitempty"`
}

// Validate validates this free ipa upgrade v1 response
func (m *FreeIpaUpgradeV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FreeIpaUpgradeV1Response) validateFlowIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowIdentifier) { // not required
		return nil
	}

	if m.FlowIdentifier != nil {
		if err := m.FlowIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *FreeIpaUpgradeV1Response) validateOriginalImage(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginalImage) { // not required
		return nil
	}

	if m.OriginalImage != nil {
		if err := m.OriginalImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalImage")
			}
			return err
		}
	}

	return nil
}

func (m *FreeIpaUpgradeV1Response) validateTargetImage(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetImage) { // not required
		return nil
	}

	if m.TargetImage != nil {
		if err := m.TargetImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FreeIpaUpgradeV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeIpaUpgradeV1Response) UnmarshalBinary(b []byte) error {
	var res FreeIpaUpgradeV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
