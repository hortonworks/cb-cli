// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BaseImageResponse base image response
// swagger:model BaseImageResponse

type BaseImageResponse struct {

	// date
	Date string `json:"date,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// hdf stacks
	HdfStacks []*StackDetailsJSON `json:"hdfStacks"`

	// hdp stacks
	HdpStacks []*StackDetailsJSON `json:"hdpStacks"`

	// images
	Images map[string]map[string]string `json:"images,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// os type
	OsType string `json:"osType,omitempty"`

	// repo
	Repo map[string]string `json:"repo,omitempty"`

	// stack details
	StackDetails *StackDetailsJSON `json:"stackDetails,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

/* polymorph BaseImageResponse date false */

/* polymorph BaseImageResponse description false */

/* polymorph BaseImageResponse hdfStacks false */

/* polymorph BaseImageResponse hdpStacks false */

/* polymorph BaseImageResponse images false */

/* polymorph BaseImageResponse os false */

/* polymorph BaseImageResponse osType false */

/* polymorph BaseImageResponse repo false */

/* polymorph BaseImageResponse stackDetails false */

/* polymorph BaseImageResponse uuid false */

/* polymorph BaseImageResponse version false */

// Validate validates this base image response
func (m *BaseImageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHdfStacks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHdpStacks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseImageResponse) validateHdfStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.HdfStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.HdfStacks); i++ {

		if swag.IsZero(m.HdfStacks[i]) { // not required
			continue
		}

		if m.HdfStacks[i] != nil {

			if err := m.HdfStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hdfStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageResponse) validateHdpStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.HdpStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.HdpStacks); i++ {

		if swag.IsZero(m.HdpStacks[i]) { // not required
			continue
		}

		if m.HdpStacks[i] != nil {

			if err := m.HdpStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hdpStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageResponse) validateStackDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.StackDetails) { // not required
		return nil
	}

	if m.StackDetails != nil {

		if err := m.StackDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseImageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseImageResponse) UnmarshalBinary(b []byte) error {
	var res BaseImageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
