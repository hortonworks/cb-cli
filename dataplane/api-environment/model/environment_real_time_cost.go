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

// EnvironmentRealTimeCost environment real time cost
// swagger:model EnvironmentRealTimeCost
type EnvironmentRealTimeCost struct {

	// datahubs
	Datahubs map[string]RealTimeCost `json:"datahubs,omitempty"`

	// datalake
	Datalake *RealTimeCost `json:"datalake,omitempty"`

	// freeipa
	Freeipa *RealTimeCost `json:"freeipa,omitempty"`

	// hourly cloudera usd
	HourlyClouderaUsd float64 `json:"hourlyClouderaUsd,omitempty"`

	// hourly provider usd
	HourlyProviderUsd float64 `json:"hourlyProviderUsd,omitempty"`
}

// Validate validates this environment real time cost
func (m *EnvironmentRealTimeCost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatahubs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeipa(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentRealTimeCost) validateDatahubs(formats strfmt.Registry) error {

	if swag.IsZero(m.Datahubs) { // not required
		return nil
	}

	for k := range m.Datahubs {

		if err := validate.Required("datahubs"+"."+k, "body", m.Datahubs[k]); err != nil {
			return err
		}
		if val, ok := m.Datahubs[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *EnvironmentRealTimeCost) validateDatalake(formats strfmt.Registry) error {

	if swag.IsZero(m.Datalake) { // not required
		return nil
	}

	if m.Datalake != nil {
		if err := m.Datalake.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalake")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentRealTimeCost) validateFreeipa(formats strfmt.Registry) error {

	if swag.IsZero(m.Freeipa) { // not required
		return nil
	}

	if m.Freeipa != nil {
		if err := m.Freeipa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeipa")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentRealTimeCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentRealTimeCost) UnmarshalBinary(b []byte) error {
	var res EnvironmentRealTimeCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
