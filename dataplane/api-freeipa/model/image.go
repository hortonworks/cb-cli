// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Image image
// swagger:model Image
type Image struct {

	// advertised
	// Read Only: true
	Advertised *bool `json:"advertised,omitempty"`

	// created
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// date
	// Read Only: true
	Date string `json:"date,omitempty"`

	// description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// images
	// Read Only: true
	Images map[string]map[string]string `json:"images,omitempty"`

	// os
	// Read Only: true
	Os string `json:"os,omitempty"`

	// os type
	// Read Only: true
	OsType string `json:"os_type,omitempty"`

	// package versions
	// Read Only: true
	PackageVersions map[string]string `json:"package-versions,omitempty"`

	// uuid
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
