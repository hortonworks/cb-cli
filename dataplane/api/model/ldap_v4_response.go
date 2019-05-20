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

// LdapV4Response ldap v4 response
// swagger:model LdapV4Response
type LdapV4Response struct {

	// LDAP group for administrators
	AdminGroup string `json:"adminGroup,omitempty"`

	// bind distinguished name for connection test and group search (e.g. cn=admin,dc=example,dc=org)
	BindDn *SecretResponse `json:"bindDn,omitempty"`

	// password for the provided bind DN
	BindPassword *SecretResponse `json:"bindPassword,omitempty"`

	// Self-signed certificate of LDAPS server
	Certificate string `json:"certificate,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// directory type of server LDAP or ACTIVE_DIRECTORY and the default is ACTIVE_DIRECTORY
	// Enum: [LDAP ACTIVE_DIRECTORY]
	DirectoryType string `json:"directoryType,omitempty"`

	// domain in LDAP server (e.g. ad.seq.com).
	Domain string `json:"domain,omitempty"`

	// Environments of the resource
	// Unique: true
	Environments []string `json:"environments"`

	// Group Member Attribute (defaults to member)
	GroupMemberAttribute string `json:"groupMemberAttribute,omitempty"`

	// Group Id Attribute (defaults to cn)
	GroupNameAttribute string `json:"groupNameAttribute,omitempty"`

	// Group Object Class (defaults to groupOfNames)
	GroupObjectClass string `json:"groupObjectClass,omitempty"`

	// template for group search for authorization (e.g. dc=hadoop,dc=apache,dc=org)
	GroupSearchBase string `json:"groupSearchBase,omitempty"`

	// public host or IP address of LDAP server
	// Required: true
	Host *string `json:"host"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// port of LDAP server (typically: 389 or 636 for LDAPS)
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Port *int32 `json:"port"`

	// determines the protocol (LDAP or LDAP over SSL)
	Protocol string `json:"protocol,omitempty"`

	// template for pattern based user search for authentication (e.g. cn={0},dc=hadoop,dc=apache,dc=org)
	// Required: true
	UserDnPattern *string `json:"userDnPattern"`

	// attribute name for simplified search filter (e.g. sAMAccountName in case of AD, UID or cn for LDAP).
	UserNameAttribute string `json:"userNameAttribute,omitempty"`

	// User Object Class (defaults to person)
	UserObjectClass string `json:"userObjectClass,omitempty"`

	// template for user search for authentication (e.g. dc=hadoop,dc=apache,dc=org)
	// Required: true
	UserSearchBase *string `json:"userSearchBase"`
}

// Validate validates this ldap v4 response
func (m *LdapV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDnPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapV4Response) validateBindDn(formats strfmt.Registry) error {

	if swag.IsZero(m.BindDn) { // not required
		return nil
	}

	if m.BindDn != nil {
		if err := m.BindDn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bindDn")
			}
			return err
		}
	}

	return nil
}

func (m *LdapV4Response) validateBindPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.BindPassword) { // not required
		return nil
	}

	if m.BindPassword != nil {
		if err := m.BindPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bindPassword")
			}
			return err
		}
	}

	return nil
}

func (m *LdapV4Response) validateDescription(formats strfmt.Registry) error {

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

var ldapV4ResponseTypeDirectoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LDAP","ACTIVE_DIRECTORY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapV4ResponseTypeDirectoryTypePropEnum = append(ldapV4ResponseTypeDirectoryTypePropEnum, v)
	}
}

const (

	// LdapV4ResponseDirectoryTypeLDAP captures enum value "LDAP"
	LdapV4ResponseDirectoryTypeLDAP string = "LDAP"

	// LdapV4ResponseDirectoryTypeACTIVEDIRECTORY captures enum value "ACTIVE_DIRECTORY"
	LdapV4ResponseDirectoryTypeACTIVEDIRECTORY string = "ACTIVE_DIRECTORY"
)

// prop value enum
func (m *LdapV4Response) validateDirectoryTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ldapV4ResponseTypeDirectoryTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LdapV4Response) validateDirectoryType(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectoryTypeEnum("directoryType", "body", m.DirectoryType); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	if err := validate.UniqueItems("environments", "body", m.Environments); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validateUserDnPattern(formats strfmt.Registry) error {

	if err := validate.Required("userDnPattern", "body", m.UserDnPattern); err != nil {
		return err
	}

	return nil
}

func (m *LdapV4Response) validateUserSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("userSearchBase", "body", m.UserSearchBase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapV4Response) UnmarshalBinary(b []byte) error {
	var res LdapV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
