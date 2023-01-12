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

// DependentHostGroupsV4Response dependent host groups v4 response
// swagger:model DependentHostGroupsV4Response
type DependentHostGroupsV4Response struct {

	// dependent host groups
	DependentHostGroups map[string][]string `json:"dependentHostGroups,omitempty"`
}

// Validate validates this dependent host groups v4 response
func (m *DependentHostGroupsV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependentHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DependentHostGroupsV4Response) validateDependentHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.DependentHostGroups) { // not required
		return nil
	}

	for k := range m.DependentHostGroups {

		if err := validate.UniqueItems("dependentHostGroups"+"."+k, "body", m.DependentHostGroups[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DependentHostGroupsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DependentHostGroupsV4Response) UnmarshalBinary(b []byte) error {
	var res DependentHostGroupsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
