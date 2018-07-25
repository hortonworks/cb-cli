// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TemplateV2Request template v2 request
// swagger:model TemplateV2Request

type TemplateV2Request struct {

	// aws specific parameters for template
	AwsParameters *AwsParameters `json:"awsParameters,omitempty"`

	// azure specific parameters for template
	AzureParameters *AzureParameters `json:"azureParameters,omitempty"`

	// custom instancetype definition
	CustomInstanceType *CustomInstanceType `json:"customInstanceType,omitempty"`

	// gcp specific parameters for template
	GcpTemlateParameters GcpParameters `json:"gcpTemlateParameters,omitempty"`

	// type of the instance
	InstanceType string `json:"instanceType,omitempty"`

	// openstack specific parameters for template
	OpenStackParameters OpenStackParameters `json:"openStackParameters,omitempty"`

	// cloud specific parameters for template
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// size of the root volume
	RootVolumeSize int32 `json:"rootVolumeSize,omitempty"`

	// number of volumes
	VolumeCount int32 `json:"volumeCount,omitempty"`

	// size of volumes
	VolumeSize int32 `json:"volumeSize,omitempty"`

	// type of the volumes
	VolumeType string `json:"volumeType,omitempty"`

	// yarn specific parameters for template
	YarnParameters YarnParameters `json:"yarnParameters,omitempty"`
}

/* polymorph TemplateV2Request awsParameters false */

/* polymorph TemplateV2Request azureParameters false */

/* polymorph TemplateV2Request customInstanceType false */

/* polymorph TemplateV2Request gcpTemlateParameters false */

/* polymorph TemplateV2Request instanceType false */

/* polymorph TemplateV2Request openStackParameters false */

/* polymorph TemplateV2Request parameters false */

/* polymorph TemplateV2Request rootVolumeSize false */

/* polymorph TemplateV2Request volumeCount false */

/* polymorph TemplateV2Request volumeSize false */

/* polymorph TemplateV2Request volumeType false */

/* polymorph TemplateV2Request yarnParameters false */

// Validate validates this template v2 request
func (m *TemplateV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAzureParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomInstanceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateV2Request) validateAwsParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsParameters) { // not required
		return nil
	}

	if m.AwsParameters != nil {

		if err := m.AwsParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParameters")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateV2Request) validateAzureParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureParameters) { // not required
		return nil
	}

	if m.AzureParameters != nil {

		if err := m.AzureParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParameters")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateV2Request) validateCustomInstanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomInstanceType) { // not required
		return nil
	}

	if m.CustomInstanceType != nil {

		if err := m.CustomInstanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customInstanceType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateV2Request) UnmarshalBinary(b []byte) error {
	var res TemplateV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
