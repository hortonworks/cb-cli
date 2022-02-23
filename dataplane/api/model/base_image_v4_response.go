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

// BaseImageV4Response base image v4 response
// swagger:model BaseImageV4Response
type BaseImageV4Response struct {

	// advertised
	Advertised bool `json:"advertised,omitempty"`

	// base parcel Url
	BaseParcelURL string `json:"baseParcelUrl,omitempty"`

	// cdh stacks
	CdhStacks []*ClouderaManagerStackDetailsV4Response `json:"cdhStacks"`

	// cloudera manager repo
	ClouderaManagerRepo *ClouderaManagerRepositoryV4Response `json:"clouderaManagerRepo,omitempty"`

	// cm build number
	CmBuildNumber string `json:"cmBuildNumber,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// date
	Date string `json:"date,omitempty"`

	// default image
	DefaultImage bool `json:"defaultImage,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// images
	Images map[string]map[string]string `json:"images,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// os type
	OsType string `json:"osType,omitempty"`

	// package versions
	PackageVersions map[string]string `json:"packageVersions,omitempty"`

	// pre warm csd
	PreWarmCsd []string `json:"preWarmCsd"`

	// pre warm parcels
	PreWarmParcels [][]string `json:"preWarmParcels"`

	// published
	Published int64 `json:"published,omitempty"`

	// repository
	Repository map[string]string `json:"repository,omitempty"`

	// source image Id
	SourceImageID string `json:"sourceImageId,omitempty"`

	// stack details
	StackDetails *BaseStackDetailsV4Response `json:"stackDetails,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this base image v4 response
func (m *BaseImageV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCdhStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClouderaManagerRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseImageV4Response) validateCdhStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.CdhStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.CdhStacks); i++ {
		if swag.IsZero(m.CdhStacks[i]) { // not required
			continue
		}

		if m.CdhStacks[i] != nil {
			if err := m.CdhStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cdhStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageV4Response) validateClouderaManagerRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.ClouderaManagerRepo) { // not required
		return nil
	}

	if m.ClouderaManagerRepo != nil {
		if err := m.ClouderaManagerRepo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManagerRepo")
			}
			return err
		}
	}

	return nil
}

func (m *BaseImageV4Response) validateStackDetails(formats strfmt.Registry) error {

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
func (m *BaseImageV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseImageV4Response) UnmarshalBinary(b []byte) error {
	var res BaseImageV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
