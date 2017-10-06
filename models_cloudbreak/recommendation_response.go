// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecommendationResponse recommendation response
// swagger:model RecommendationResponse

type RecommendationResponse struct {

	// recommendations
	Recommendations map[string]VMTypeJSON `json:"recommendations,omitempty"`

	// virtual machines
	// Unique: true
	VirtualMachines []*VMTypeJSON `json:"virtualMachines"`
}

/* polymorph RecommendationResponse recommendations false */

/* polymorph RecommendationResponse virtualMachines false */

// Validate validates this recommendation response
func (m *RecommendationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecommendations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVirtualMachines(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecommendationResponse) validateRecommendations(formats strfmt.Registry) error {

	if swag.IsZero(m.Recommendations) { // not required
		return nil
	}

	if err := validate.Required("recommendations", "body", m.Recommendations); err != nil {
		return err
	}

	return nil
}

func (m *RecommendationResponse) validateVirtualMachines(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualMachines) { // not required
		return nil
	}

	if err := validate.UniqueItems("virtualMachines", "body", m.VirtualMachines); err != nil {
		return err
	}

	for i := 0; i < len(m.VirtualMachines); i++ {

		if swag.IsZero(m.VirtualMachines[i]) { // not required
			continue
		}

		if m.VirtualMachines[i] != nil {

			if err := m.VirtualMachines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("virtualMachines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecommendationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecommendationResponse) UnmarshalBinary(b []byte) error {
	var res RecommendationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
