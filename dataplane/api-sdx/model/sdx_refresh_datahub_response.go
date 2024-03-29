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

// SdxRefreshDatahubResponse sdx refresh datahub response
// swagger:model SdxRefreshDatahubResponse
type SdxRefreshDatahubResponse struct {

	// List of Data Hubs that are refreshed with CRN
	RefreshedDatahubs []*StackViewV4Response `json:"refreshedDatahubs"`
}

// Validate validates this sdx refresh datahub response
func (m *SdxRefreshDatahubResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefreshedDatahubs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SdxRefreshDatahubResponse) validateRefreshedDatahubs(formats strfmt.Registry) error {

	if swag.IsZero(m.RefreshedDatahubs) { // not required
		return nil
	}

	for i := 0; i < len(m.RefreshedDatahubs); i++ {
		if swag.IsZero(m.RefreshedDatahubs[i]) { // not required
			continue
		}

		if m.RefreshedDatahubs[i] != nil {
			if err := m.RefreshedDatahubs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("refreshedDatahubs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxRefreshDatahubResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxRefreshDatahubResponse) UnmarshalBinary(b []byte) error {
	var res SdxRefreshDatahubResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
