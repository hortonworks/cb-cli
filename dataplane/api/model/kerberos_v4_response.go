// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KerberosV4Response kerberos v4 response
// swagger:model KerberosV4Response
type KerberosV4Response struct {

	// kerberos admin user
	Admin *SecretResponse `json:"admin,omitempty"`

	// kerberos admin server URL
	AdminURL string `json:"adminUrl,omitempty"`

	// container dn
	ContainerDn string `json:"containerDn,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Ambari kerberos descriptor
	Descriptor *SecretResponse `json:"descriptor,omitempty"`

	// cluster instances will set this as the domain part of their hostname
	Domain string `json:"domain,omitempty"`

	// Environments of the resource
	// Unique: true
	Environments []string `json:"environments"`

	// id
	ID int64 `json:"id,omitempty"`

	// Ambari kerberos krb5.conf template
	Krb5Conf *SecretResponse `json:"krb5Conf,omitempty"`

	// ldap Url
	LdapURL string `json:"ldapUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// comma separated list of nameservers' IP address which will be used by cluster instances
	// Pattern: (^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(,((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))*$)
	NameServers string `json:"nameServers,omitempty"`

	// kerberos admin password
	Password *SecretResponse `json:"password,omitempty"`

	// kerberos principal
	Principal *SecretResponse `json:"principal,omitempty"`

	// realm
	Realm string `json:"realm,omitempty"`

	// tcp allowed
	TCPAllowed bool `json:"tcpAllowed,omitempty"`

	// type
	// Required: true
	// Enum: [ACTIVE_DIRECTORY MIT FREEIPA AMBARI_DESCRIPTOR]
	Type *string `json:"type"`

	// kerberos KDC server URL
	URL string `json:"url,omitempty"`

	// Allows to select either a trusting SSL connection or a validating (non-trusting) SSL connection to KDC
	VerifyKdcTrust bool `json:"verifyKdcTrust,omitempty"`
}

// Validate validates this kerberos v4 response
func (m *KerberosV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKrb5Conf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KerberosV4Response) validateAdmin(formats strfmt.Registry) error {

	if swag.IsZero(m.Admin) { // not required
		return nil
	}

	if m.Admin != nil {
		if err := m.Admin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Response) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *KerberosV4Response) validateDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.Descriptor) { // not required
		return nil
	}

	if m.Descriptor != nil {
		if err := m.Descriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Response) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	if err := validate.UniqueItems("environments", "body", m.Environments); err != nil {
		return err
	}

	return nil
}

func (m *KerberosV4Response) validateKrb5Conf(formats strfmt.Registry) error {

	if swag.IsZero(m.Krb5Conf) { // not required
		return nil
	}

	if m.Krb5Conf != nil {
		if err := m.Krb5Conf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("krb5Conf")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Response) validateNameServers(formats strfmt.Registry) error {

	if swag.IsZero(m.NameServers) { // not required
		return nil
	}

	if err := validate.Pattern("nameServers", "body", string(m.NameServers), `(^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(,((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))*$)`); err != nil {
		return err
	}

	return nil
}

func (m *KerberosV4Response) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if m.Password != nil {
		if err := m.Password.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password")
			}
			return err
		}
	}

	return nil
}

func (m *KerberosV4Response) validatePrincipal(formats strfmt.Registry) error {

	if swag.IsZero(m.Principal) { // not required
		return nil
	}

	if m.Principal != nil {
		if err := m.Principal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

var kerberosV4ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE_DIRECTORY","MIT","FREEIPA","AMBARI_DESCRIPTOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kerberosV4ResponseTypeTypePropEnum = append(kerberosV4ResponseTypeTypePropEnum, v)
	}
}

const (

	// KerberosV4ResponseTypeACTIVEDIRECTORY captures enum value "ACTIVE_DIRECTORY"
	KerberosV4ResponseTypeACTIVEDIRECTORY string = "ACTIVE_DIRECTORY"

	// KerberosV4ResponseTypeMIT captures enum value "MIT"
	KerberosV4ResponseTypeMIT string = "MIT"

	// KerberosV4ResponseTypeFREEIPA captures enum value "FREEIPA"
	KerberosV4ResponseTypeFREEIPA string = "FREEIPA"

	// KerberosV4ResponseTypeAMBARIDESCRIPTOR captures enum value "AMBARI_DESCRIPTOR"
	KerberosV4ResponseTypeAMBARIDESCRIPTOR string = "AMBARI_DESCRIPTOR"
)

// prop value enum
func (m *KerberosV4Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kerberosV4ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KerberosV4Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KerberosV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KerberosV4Response) UnmarshalBinary(b []byte) error {
	var res KerberosV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
