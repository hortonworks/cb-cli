// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RdsConfig rds config
// swagger:model RdsConfig

type RdsConfig struct {

	// Password to use for the jdbc connection
	// Required: true
	ConnectionPassword *string `json:"connectionPassword"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	// Pattern: ^jdbc:(postgresql|mysql|oracle)://[-\w\.]*:\d{1,5}/?\w*
	ConnectionURL *string `json:"connectionURL"`

	// Username to use for the jdbc connection
	// Required: true
	ConnectionUserName *string `json:"connectionUserName"`

	// URL that points to the jar of the connection driver(connector)
	// Max Length: 150
	// Min Length: 0
	ConnectorJarURL *string `json:"connectorJarUrl,omitempty"`

	// Name of the RDS configuration resource
	// Required: true
	// Max Length: 50
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Oracle specific properties
	OracleParameters *Oracle `json:"oracleParameters,omitempty"`

	// Type of RDS, aka the service name that will use the RDS like HIVE, DRUID, SUPERSET, RANGER, etc.
	// Required: true
	// Max Length: 12
	// Min Length: 3
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)
	Type *string `json:"type"`
}

/* polymorph RdsConfig connectionPassword false */

/* polymorph RdsConfig connectionURL false */

/* polymorph RdsConfig connectionUserName false */

/* polymorph RdsConfig connectorJarUrl false */

/* polymorph RdsConfig name false */

/* polymorph RdsConfig oracleParameters false */

/* polymorph RdsConfig type false */

// Validate validates this rds config
func (m *RdsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectorJarURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOracleParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsConfig) validateConnectionPassword(formats strfmt.Registry) error {

	if err := validate.Required("connectionPassword", "body", m.ConnectionPassword); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	if err := validate.Pattern("connectionURL", "body", string(*m.ConnectionURL), `^jdbc:(postgresql|mysql|oracle)://[-\w\.]*:\d{1,5}/?\w*`); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectionUserName(formats strfmt.Registry) error {

	if err := validate.Required("connectionUserName", "body", m.ConnectionUserName); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectorJarURL(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectorJarURL) { // not required
		return nil
	}

	if err := validate.MinLength("connectorJarUrl", "body", string(*m.ConnectorJarURL), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("connectorJarUrl", "body", string(*m.ConnectorJarURL), 150); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateOracleParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.OracleParameters) { // not required
		return nil
	}

	if m.OracleParameters != nil {

		if err := m.OracleParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleParameters")
			}
			return err
		}
	}

	return nil
}

func (m *RdsConfig) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 12); err != nil {
		return err
	}

	if err := validate.Pattern("type", "body", string(*m.Type), `(^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsConfig) UnmarshalBinary(b []byte) error {
	var res RdsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
