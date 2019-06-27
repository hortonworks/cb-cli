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

// InstanceTemplateV4Response instance template v4 response
// swagger:model InstanceTemplateV4Response
type InstanceTemplateV4Response struct {

	// attached volumes
	// Unique: true
	AttachedVolumes []*VolumeV4Response `json:"attachedVolumes"`

	// aws specific parameters for template
	Aws *AwsInstanceTemplateV4Parameters `json:"aws,omitempty"`

	// azure specific parameters for template
	Azure *AzureInstanceTemplateV4Parameters `json:"azure,omitempty"`

	// ephemeral volume
	EphemeralVolume *VolumeV4Response `json:"ephemeralVolume,omitempty"`

	// gcp specific parameters for template
	Gcp *GcpInstanceTemplateV4Parameters `json:"gcp,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// type of the instance
	InstanceType string `json:"instanceType,omitempty"`

	// yarn specific parameters for template
	Mock MockInstanceTemplateV4Parameters `json:"mock,omitempty"`

	// openstack specific parameters for template
	Openstack OpenStackInstanceTemplateV4Parameters `json:"openstack,omitempty"`

	// root volume
	RootVolume *RootVolumeV4Response `json:"rootVolume,omitempty"`

	// yarn specific parameters for template
	Yarn *YarnInstanceTemplateV4Parameters `json:"yarn,omitempty"`
}

// Validate validates this instance template v4 response
func (m *InstanceTemplateV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeralVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceTemplateV4Response) validateAttachedVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.AttachedVolumes) { // not required
		return nil
	}

	if err := validate.UniqueItems("attachedVolumes", "body", m.AttachedVolumes); err != nil {
		return err
	}

	for i := 0; i < len(m.AttachedVolumes); i++ {
		if swag.IsZero(m.AttachedVolumes[i]) { // not required
			continue
		}

		if m.AttachedVolumes[i] != nil {
			if err := m.AttachedVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceTemplateV4Response) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceTemplateV4Response) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceTemplateV4Response) validateEphemeralVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.EphemeralVolume) { // not required
		return nil
	}

	if m.EphemeralVolume != nil {
		if err := m.EphemeralVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeralVolume")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceTemplateV4Response) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceTemplateV4Response) validateRootVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.RootVolume) { // not required
		return nil
	}

	if m.RootVolume != nil {
		if err := m.RootVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootVolume")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceTemplateV4Response) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceTemplateV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceTemplateV4Response) UnmarshalBinary(b []byte) error {
	var res InstanceTemplateV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
