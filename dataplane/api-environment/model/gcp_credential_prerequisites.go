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

// GcpCredentialPrerequisites gcp credential prerequisites
// swagger:model GcpCredentialPrerequisites
type GcpCredentialPrerequisites struct {

	// GCP specific 'gcloud' CLI based commands to create prerequisites for Cloudbreak credential creation. The field is base64 encoded.
	// Required: true
	CreationCommand *string `json:"creationCommand"`

	// Policies for experiences.
	Policies map[string]string `json:"policies,omitempty"`
}

// Validate validates this gcp credential prerequisites
func (m *GcpCredentialPrerequisites) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpCredentialPrerequisites) validateCreationCommand(formats strfmt.Registry) error {

	if err := validate.Required("creationCommand", "body", m.CreationCommand); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GcpCredentialPrerequisites) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpCredentialPrerequisites) UnmarshalBinary(b []byte) error {
	var res GcpCredentialPrerequisites
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
