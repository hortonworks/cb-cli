// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomImageCatalogV4CreateImageRequest custom image catalog v4 create image request
// swagger:model CustomImageCatalogV4CreateImageRequest
type CustomImageCatalogV4CreateImageRequest struct {

	// base url of of parcels in the image
	// Required: true
	// Max Length: 255
	// Min Length: 1
	BaseParcelURL *string `json:"baseParcelUrl"`

	// type of the image - datalake, datahub or freeipa
	// Required: true
	ImageType *string `json:"imageType"`

	// id of the source image serving as the base of the customized image
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Pattern: (^[a-z0-9][-a-z0-9]*[a-z0-9]$)
	SourceImageID *string `json:"sourceImageId"`

	// image references of the custom image in different regions
	// Required: true
	// Unique: true
	VMImages []*CustomImageCatalogV4VMImageRequest `json:"vmImages"`
}

// Validate validates this custom image catalog v4 create image request
func (m *CustomImageCatalogV4CreateImageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseParcelURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomImageCatalogV4CreateImageRequest) validateBaseParcelURL(formats strfmt.Registry) error {

	if err := validate.Required("baseParcelUrl", "body", m.BaseParcelURL); err != nil {
		return err
	}

	if err := validate.MinLength("baseParcelUrl", "body", string(*m.BaseParcelURL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("baseParcelUrl", "body", string(*m.BaseParcelURL), 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomImageCatalogV4CreateImageRequest) validateImageType(formats strfmt.Registry) error {

	if err := validate.Required("imageType", "body", m.ImageType); err != nil {
		return err
	}

	return nil
}

func (m *CustomImageCatalogV4CreateImageRequest) validateSourceImageID(formats strfmt.Registry) error {

	if err := validate.Required("sourceImageId", "body", m.SourceImageID); err != nil {
		return err
	}

	if err := validate.MinLength("sourceImageId", "body", string(*m.SourceImageID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("sourceImageId", "body", string(*m.SourceImageID), 255); err != nil {
		return err
	}

	if err := validate.Pattern("sourceImageId", "body", string(*m.SourceImageID), `(^[a-z0-9][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *CustomImageCatalogV4CreateImageRequest) validateVMImages(formats strfmt.Registry) error {

	if err := validate.Required("vmImages", "body", m.VMImages); err != nil {
		return err
	}

	if err := validate.UniqueItems("vmImages", "body", m.VMImages); err != nil {
		return err
	}

	for i := 0; i < len(m.VMImages); i++ {
		if swag.IsZero(m.VMImages[i]) { // not required
			continue
		}

		if m.VMImages[i] != nil {
			if err := m.VMImages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vmImages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomImageCatalogV4CreateImageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomImageCatalogV4CreateImageRequest) UnmarshalBinary(b []byte) error {
	var res CustomImageCatalogV4CreateImageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
