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

// EnvironmentV1Request environment v1 request
// swagger:model EnvironmentV1Request
type EnvironmentV1Request struct {

	// Name of the admin group to be used for all the services.
	AdminGroupName string `json:"adminGroupName,omitempty"`

	// SSH key for accessing cluster node instances.
	Authentication *EnvironmentAuthenticationV1Request `json:"authentication,omitempty"`

	// AWS Specific parameters.
	Aws *AwsEnvironmentV1Parameters `json:"aws,omitempty"`

	// AZURE Specific parameters.
	Azure *AzureEnvironmentV1Parameters `json:"azure,omitempty"`

	// Backup related specifics of the environment.
	Backup *BackupRequest `json:"backup,omitempty"`

	// ccm v2 Tls type
	// Enum: [ONE_WAY_TLS TWO_WAY_TLS]
	CcmV2TLSType string `json:"ccmV2TlsType,omitempty"`

	// Cloud storage validation enabled or not.
	// Enum: [ENABLED DISABLED]
	CloudStorageValidation string `json:"cloudStorageValidation,omitempty"`

	// Name of the credential of the environment. If the name is given, the detailed credential is ignored in the request.
	CredentialName string `json:"credentialName,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// The version of the Cloudbreak build used to create the environment.
	EnvironmentServiceVersion string `json:"environmentServiceVersion,omitempty"`

	// Properties for FreeIpa which can be attached to the given environment
	FreeIpa *AttachedFreeIpaRequest `json:"freeIpa,omitempty"`

	// GCP Specific parameters.
	Gcp *GcpEnvironmentV1Parameters `json:"gcp,omitempty"`

	// IDBroker mapping source.
	// Enum: [NONE MOCK IDBMMS]
	IDBrokerMappingSource string `json:"idBrokerMappingSource,omitempty"`

	// Location of the environment.
	// Required: true
	Location *LocationV1Request `json:"location"`

	// name of the resource
	// Required: true
	// Max Length: 28
	// Min Length: 5
	Name *string `json:"name"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV1Request `json:"network,omitempty"`

	// Flag that marks that the request was intented to set the tunnel version by hand and it will not be overwritten by Cloudbreak
	OverrideTunnel bool `json:"overrideTunnel,omitempty"`

	// Parent environment name
	ParentEnvironmentName string `json:"parentEnvironmentName,omitempty"`

	// Name of the proxyconfig of the environment.
	ProxyConfigName string `json:"proxyConfigName,omitempty"`

	// Regions of the environment.
	// Unique: true
	Regions []string `json:"regions"`

	// Security control for FreeIPA and Datalake deployment.
	SecurityAccess *SecurityAccessV1Request `json:"securityAccess,omitempty"`

	// Tags for environments.
	Tags map[string]string `json:"tags,omitempty"`

	// Telemetry related specifics of the environment.
	Telemetry *TelemetryRequest `json:"telemetry,omitempty"`

	// Configuration that the connection going directly or with cluster proxy or with CCM and cluster proxy.
	// Enum: [DIRECT CCM CLUSTER_PROXY CCMV2 CCMV2_JUMPGATE]
	Tunnel string `json:"tunnel,omitempty"`
}

// Validate validates this environment v1 request
func (m *EnvironmentV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCcmV2TLSType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorageValidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDBrokerMappingSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelemetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentV1Request) validateAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateAws(formats strfmt.Registry) error {

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

func (m *EnvironmentV1Request) validateAzure(formats strfmt.Registry) error {

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

func (m *EnvironmentV1Request) validateBackup(formats strfmt.Registry) error {

	if swag.IsZero(m.Backup) { // not required
		return nil
	}

	if m.Backup != nil {
		if err := m.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

var environmentV1RequestTypeCcmV2TLSTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONE_WAY_TLS","TWO_WAY_TLS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentV1RequestTypeCcmV2TLSTypePropEnum = append(environmentV1RequestTypeCcmV2TLSTypePropEnum, v)
	}
}

const (

	// EnvironmentV1RequestCcmV2TLSTypeONEWAYTLS captures enum value "ONE_WAY_TLS"
	EnvironmentV1RequestCcmV2TLSTypeONEWAYTLS string = "ONE_WAY_TLS"

	// EnvironmentV1RequestCcmV2TLSTypeTWOWAYTLS captures enum value "TWO_WAY_TLS"
	EnvironmentV1RequestCcmV2TLSTypeTWOWAYTLS string = "TWO_WAY_TLS"
)

// prop value enum
func (m *EnvironmentV1Request) validateCcmV2TLSTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentV1RequestTypeCcmV2TLSTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentV1Request) validateCcmV2TLSType(formats strfmt.Registry) error {

	if swag.IsZero(m.CcmV2TLSType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCcmV2TLSTypeEnum("ccmV2TlsType", "body", m.CcmV2TLSType); err != nil {
		return err
	}

	return nil
}

var environmentV1RequestTypeCloudStorageValidationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentV1RequestTypeCloudStorageValidationPropEnum = append(environmentV1RequestTypeCloudStorageValidationPropEnum, v)
	}
}

const (

	// EnvironmentV1RequestCloudStorageValidationENABLED captures enum value "ENABLED"
	EnvironmentV1RequestCloudStorageValidationENABLED string = "ENABLED"

	// EnvironmentV1RequestCloudStorageValidationDISABLED captures enum value "DISABLED"
	EnvironmentV1RequestCloudStorageValidationDISABLED string = "DISABLED"
)

// prop value enum
func (m *EnvironmentV1Request) validateCloudStorageValidationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentV1RequestTypeCloudStorageValidationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentV1Request) validateCloudStorageValidation(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorageValidation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudStorageValidationEnum("cloudStorageValidation", "body", m.CloudStorageValidation); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateDescription(formats strfmt.Registry) error {

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

func (m *EnvironmentV1Request) validateFreeIpa(formats strfmt.Registry) error {

	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateGcp(formats strfmt.Registry) error {

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

var environmentV1RequestTypeIDBrokerMappingSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","MOCK","IDBMMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentV1RequestTypeIDBrokerMappingSourcePropEnum = append(environmentV1RequestTypeIDBrokerMappingSourcePropEnum, v)
	}
}

const (

	// EnvironmentV1RequestIDBrokerMappingSourceNONE captures enum value "NONE"
	EnvironmentV1RequestIDBrokerMappingSourceNONE string = "NONE"

	// EnvironmentV1RequestIDBrokerMappingSourceMOCK captures enum value "MOCK"
	EnvironmentV1RequestIDBrokerMappingSourceMOCK string = "MOCK"

	// EnvironmentV1RequestIDBrokerMappingSourceIDBMMS captures enum value "IDBMMS"
	EnvironmentV1RequestIDBrokerMappingSourceIDBMMS string = "IDBMMS"
)

// prop value enum
func (m *EnvironmentV1Request) validateIDBrokerMappingSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentV1RequestTypeIDBrokerMappingSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentV1Request) validateIDBrokerMappingSource(formats strfmt.Registry) error {

	if swag.IsZero(m.IDBrokerMappingSource) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDBrokerMappingSourceEnum("idBrokerMappingSource", "body", m.IDBrokerMappingSource); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 28); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateSecurityAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityAccess) { // not required
		return nil
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateTelemetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Telemetry) { // not required
		return nil
	}

	if m.Telemetry != nil {
		if err := m.Telemetry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("telemetry")
			}
			return err
		}
	}

	return nil
}

var environmentV1RequestTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM","CLUSTER_PROXY","CCMV2","CCMV2_JUMPGATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentV1RequestTypeTunnelPropEnum = append(environmentV1RequestTypeTunnelPropEnum, v)
	}
}

const (

	// EnvironmentV1RequestTunnelDIRECT captures enum value "DIRECT"
	EnvironmentV1RequestTunnelDIRECT string = "DIRECT"

	// EnvironmentV1RequestTunnelCCM captures enum value "CCM"
	EnvironmentV1RequestTunnelCCM string = "CCM"

	// EnvironmentV1RequestTunnelCLUSTERPROXY captures enum value "CLUSTER_PROXY"
	EnvironmentV1RequestTunnelCLUSTERPROXY string = "CLUSTER_PROXY"

	// EnvironmentV1RequestTunnelCCMV2 captures enum value "CCMV2"
	EnvironmentV1RequestTunnelCCMV2 string = "CCMV2"

	// EnvironmentV1RequestTunnelCCMV2JUMPGATE captures enum value "CCMV2_JUMPGATE"
	EnvironmentV1RequestTunnelCCMV2JUMPGATE string = "CCMV2_JUMPGATE"
)

// prop value enum
func (m *EnvironmentV1Request) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentV1RequestTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentV1Request) validateTunnel(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnel) { // not required
		return nil
	}

	// value enum
	if err := m.validateTunnelEnum("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentV1Request) UnmarshalBinary(b []byte) error {
	var res EnvironmentV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
