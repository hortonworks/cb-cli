// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CertificateResponse certificate response
// swagger:model CertificateResponse

type CertificateResponse struct {

	// client certificate used by the gateway
	ClientCertPath string `json:"clientCertPath,omitempty"`

	// client key used by the gateway
	ClientKeyPath string `json:"clientKeyPath,omitempty"`

	// server certificate used by the gateway
	ServerCert string `json:"serverCert,omitempty"`
}

/* polymorph CertificateResponse clientCertPath false */

/* polymorph CertificateResponse clientKeyPath false */

/* polymorph CertificateResponse serverCert false */

// Validate validates this certificate response
func (m *CertificateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateResponse) UnmarshalBinary(b []byte) error {
	var res CertificateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}