// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostGroupResponse host group response
// swagger:model HostGroupResponse

type HostGroupResponse struct {

	// instance group or resource constraint for a hostgroup
	// Required: true
	Constraint *Constraint `json:"constraint"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// metadata of hosts
	// Unique: true
	Metadata []*HostMetadata `json:"metadata"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// referenced recipe ids
	// Unique: true
	RecipeIds []int64 `json:"recipeIds"`

	// referenced recipes
	// Unique: true
	Recipes []*RecipeResponse `json:"recipes"`

	// recovery mode of the hostgroup's nodes
	RecoveryMode string `json:"recoveryMode,omitempty"`
}

/* polymorph HostGroupResponse constraint false */

/* polymorph HostGroupResponse id false */

/* polymorph HostGroupResponse metadata false */

/* polymorph HostGroupResponse name false */

/* polymorph HostGroupResponse recipeIds false */

/* polymorph HostGroupResponse recipes false */

/* polymorph HostGroupResponse recoveryMode false */

// Validate validates this host group response
func (m *HostGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipeIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecoveryMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupResponse) validateConstraint(formats strfmt.Registry) error {

	if err := validate.Required("constraint", "body", m.Constraint); err != nil {
		return err
	}

	if m.Constraint != nil {

		if err := m.Constraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraint")
			}
			return err
		}
	}

	return nil
}

func (m *HostGroupResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.UniqueItems("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {

		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {

			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostGroupResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupResponse) validateRecipeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipeIds", "body", m.RecipeIds); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupResponse) validateRecipes(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipes) { // not required
		return nil
	}

	if err := validate.UniqueItems("recipes", "body", m.Recipes); err != nil {
		return err
	}

	for i := 0; i < len(m.Recipes); i++ {

		if swag.IsZero(m.Recipes[i]) { // not required
			continue
		}

		if m.Recipes[i] != nil {

			if err := m.Recipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var hostGroupResponseTypeRecoveryModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostGroupResponseTypeRecoveryModePropEnum = append(hostGroupResponseTypeRecoveryModePropEnum, v)
	}
}

const (
	// HostGroupResponseRecoveryModeMANUAL captures enum value "MANUAL"
	HostGroupResponseRecoveryModeMANUAL string = "MANUAL"
	// HostGroupResponseRecoveryModeAUTO captures enum value "AUTO"
	HostGroupResponseRecoveryModeAUTO string = "AUTO"
)

// prop value enum
func (m *HostGroupResponse) validateRecoveryModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hostGroupResponseTypeRecoveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HostGroupResponse) validateRecoveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryModeEnum("recoveryMode", "body", m.RecoveryMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostGroupResponse) UnmarshalBinary(b []byte) error {
	var res HostGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
