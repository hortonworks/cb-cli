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

// DetailedEnvironmentV1Response detailed environment v1 response
// swagger:model DetailedEnvironmentV1Response
type DetailedEnvironmentV1Response struct {

	// Name of the admin group to be used for all the services.
	AdminGroupName string `json:"adminGroupName,omitempty"`

	// SSH key for accessing cluster node instances.
	Authentication *EnvironmentAuthenticationV1Response `json:"authentication,omitempty"`

	// AWS Specific parameters.
	Aws *AwsEnvironmentV1Parameters `json:"aws,omitempty"`

	// AZURE Specific parameters.
	Azure *AzureEnvironmentV1Parameters `json:"azure,omitempty"`

	// Backup related specifics of the environment.
	Backup *BackupResponse `json:"backup,omitempty"`

	// Cloud platform of the environment.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Cloud storage validation enabled or not.
	// Enum: [ENABLED DISABLED]
	CloudStorageValidation string `json:"cloudStorageValidation,omitempty"`

	// Create freeipa in environment
	CreateFreeIpa bool `json:"createFreeIpa,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// crn of the creator
	Creator string `json:"creator,omitempty"`

	// Credential of the environment.
	Credential *CredentialV1Response `json:"credential,omitempty"`

	// id of the resource
	Crn string `json:"crn,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// The version of the Cloudbreak build used to create the environment.
	EnvironmentServiceVersion string `json:"environmentServiceVersion,omitempty"`

	// Status of the environment.
	// Enum: [CREATION_INITIATED DELETE_INITIATED UPDATE_INITIATED ENVIRONMENT_INITIALIZATION_IN_PROGRESS ENVIRONMENT_VALIDATION_IN_PROGRESS PREREQUISITES_CREATE_IN_PROGRESS NETWORK_CREATION_IN_PROGRESS NETWORK_DELETE_IN_PROGRESS RDBMS_DELETE_IN_PROGRESS FREEIPA_CREATION_IN_PROGRESS FREEIPA_DELETE_IN_PROGRESS EXPERIENCE_DELETE_IN_PROGRESS CLUSTER_DEFINITION_CLEANUP_PROGRESS UMS_RESOURCE_DELETE_IN_PROGRESS IDBROKER_MAPPINGS_DELETE_IN_PROGRESS S3GUARD_TABLE_DELETE_IN_PROGRESS DATAHUB_CLUSTERS_DELETE_IN_PROGRESS DATALAKE_CLUSTERS_DELETE_IN_PROGRESS PUBLICKEY_CREATE_IN_PROGRESS PUBLICKEY_DELETE_IN_PROGRESS ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS ENVIRONMENT_ENCRYPTION_RESOURCES_INITIALIZED AVAILABLE ARCHIVED CREATE_FAILED DELETE_FAILED UPDATE_FAILED STOP_DATAHUB_STARTED STOP_DATAHUB_FAILED STOP_DATALAKE_STARTED STOP_DATALAKE_FAILED STOP_FREEIPA_STARTED STOP_FREEIPA_FAILED ENV_STOPPED START_DATAHUB_STARTED START_DATAHUB_FAILED START_DATALAKE_STARTED START_DATALAKE_FAILED START_FREEIPA_STARTED START_FREEIPA_FAILED START_SYNCHRONIZE_USERS_STARTED START_SYNCHRONIZE_USERS_FAILED FREEIPA_DELETED_ON_PROVIDER_SIDE LOAD_BALANCER_ENV_UPDATE_STARTED LOAD_BALANCER_ENV_UPDATE_FAILED LOAD_BALANCER_STACK_UPDATE_STARTED LOAD_BALANCER_STACK_UPDATE_FAILED]
	EnvironmentStatus string `json:"environmentStatus,omitempty"`

	// The FreeIPA paramaters
	FreeIpa *FreeIpaResponse `json:"freeIpa,omitempty"`

	// GCP Specific parameters.
	Gcp GcpEnvironmentV1Parameters `json:"gcp,omitempty"`

	// IDBroker mapping source.
	// Enum: [NONE MOCK IDBMMS]
	IDBrokerMappingSource string `json:"idBrokerMappingSource,omitempty"`

	// Location of the environment.
	Location *LocationV1Response `json:"location,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV1Response `json:"network,omitempty"`

	// Parent environment cloud platform
	ParentEnvironmentCloudPlatform string `json:"parentEnvironmentCloudPlatform,omitempty"`

	// Parent environment global identifier
	ParentEnvironmentCrn string `json:"parentEnvironmentCrn,omitempty"`

	// Parent environment name
	ParentEnvironmentName string `json:"parentEnvironmentName,omitempty"`

	// ProxyConfig attached to the environment.
	ProxyConfig *ProxyResponse `json:"proxyConfig,omitempty"`

	// Regions of the environment.
	Regions *CompactRegionV1Response `json:"regions,omitempty"`

	// Security control for FreeIPA and Datalake deployment.
	SecurityAccess *SecurityAccessV1Response `json:"securityAccess,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// Tags for environments.
	Tags *TagResponse `json:"tags,omitempty"`

	// Telemetry related specifics of the environment.
	Telemetry *TelemetryResponse `json:"telemetry,omitempty"`

	// Configuration that the connection going directly or with cluster proxy or with ccm and cluster proxy.
	// Enum: [DIRECT CCM CLUSTER_PROXY CCMV2]
	Tunnel string `json:"tunnel,omitempty"`

	// YARN Specific parameters.
	Yarn YarnEnvironmentV1Parameters `json:"yarn,omitempty"`
}

// Validate validates this detailed environment v1 response
func (m *DetailedEnvironmentV1Response) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCloudStorageValidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDBrokerMappingSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *DetailedEnvironmentV1Response) validateAuthentication(formats strfmt.Registry) error {

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

func (m *DetailedEnvironmentV1Response) validateAws(formats strfmt.Registry) error {

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

func (m *DetailedEnvironmentV1Response) validateAzure(formats strfmt.Registry) error {

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

func (m *DetailedEnvironmentV1Response) validateBackup(formats strfmt.Registry) error {

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

var detailedEnvironmentV1ResponseTypeCloudStorageValidationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedEnvironmentV1ResponseTypeCloudStorageValidationPropEnum = append(detailedEnvironmentV1ResponseTypeCloudStorageValidationPropEnum, v)
	}
}

const (

	// DetailedEnvironmentV1ResponseCloudStorageValidationENABLED captures enum value "ENABLED"
	DetailedEnvironmentV1ResponseCloudStorageValidationENABLED string = "ENABLED"

	// DetailedEnvironmentV1ResponseCloudStorageValidationDISABLED captures enum value "DISABLED"
	DetailedEnvironmentV1ResponseCloudStorageValidationDISABLED string = "DISABLED"
)

// prop value enum
func (m *DetailedEnvironmentV1Response) validateCloudStorageValidationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedEnvironmentV1ResponseTypeCloudStorageValidationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedEnvironmentV1Response) validateCloudStorageValidation(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorageValidation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudStorageValidationEnum("cloudStorageValidation", "body", m.CloudStorageValidation); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

var detailedEnvironmentV1ResponseTypeEnvironmentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATION_INITIATED","DELETE_INITIATED","UPDATE_INITIATED","ENVIRONMENT_INITIALIZATION_IN_PROGRESS","ENVIRONMENT_VALIDATION_IN_PROGRESS","PREREQUISITES_CREATE_IN_PROGRESS","NETWORK_CREATION_IN_PROGRESS","NETWORK_DELETE_IN_PROGRESS","RDBMS_DELETE_IN_PROGRESS","FREEIPA_CREATION_IN_PROGRESS","FREEIPA_DELETE_IN_PROGRESS","EXPERIENCE_DELETE_IN_PROGRESS","CLUSTER_DEFINITION_CLEANUP_PROGRESS","UMS_RESOURCE_DELETE_IN_PROGRESS","IDBROKER_MAPPINGS_DELETE_IN_PROGRESS","S3GUARD_TABLE_DELETE_IN_PROGRESS","DATAHUB_CLUSTERS_DELETE_IN_PROGRESS","DATALAKE_CLUSTERS_DELETE_IN_PROGRESS","PUBLICKEY_CREATE_IN_PROGRESS","PUBLICKEY_DELETE_IN_PROGRESS","ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS","ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS","ENVIRONMENT_ENCRYPTION_RESOURCES_INITIALIZED","AVAILABLE","ARCHIVED","CREATE_FAILED","DELETE_FAILED","UPDATE_FAILED","STOP_DATAHUB_STARTED","STOP_DATAHUB_FAILED","STOP_DATALAKE_STARTED","STOP_DATALAKE_FAILED","STOP_FREEIPA_STARTED","STOP_FREEIPA_FAILED","ENV_STOPPED","START_DATAHUB_STARTED","START_DATAHUB_FAILED","START_DATALAKE_STARTED","START_DATALAKE_FAILED","START_FREEIPA_STARTED","START_FREEIPA_FAILED","START_SYNCHRONIZE_USERS_STARTED","START_SYNCHRONIZE_USERS_FAILED","FREEIPA_DELETED_ON_PROVIDER_SIDE","LOAD_BALANCER_ENV_UPDATE_STARTED","LOAD_BALANCER_ENV_UPDATE_FAILED","LOAD_BALANCER_STACK_UPDATE_STARTED","LOAD_BALANCER_STACK_UPDATE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedEnvironmentV1ResponseTypeEnvironmentStatusPropEnum = append(detailedEnvironmentV1ResponseTypeEnvironmentStatusPropEnum, v)
	}
}

const (

	// DetailedEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED captures enum value "CREATION_INITIATED"
	DetailedEnvironmentV1ResponseEnvironmentStatusCREATIONINITIATED string = "CREATION_INITIATED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED captures enum value "DELETE_INITIATED"
	DetailedEnvironmentV1ResponseEnvironmentStatusDELETEINITIATED string = "DELETE_INITIATED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED captures enum value "UPDATE_INITIATED"
	DetailedEnvironmentV1ResponseEnvironmentStatusUPDATEINITIATED string = "UPDATE_INITIATED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTINITIALIZATIONINPROGRESS captures enum value "ENVIRONMENT_INITIALIZATION_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTINITIALIZATIONINPROGRESS string = "ENVIRONMENT_INITIALIZATION_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTVALIDATIONINPROGRESS captures enum value "ENVIRONMENT_VALIDATION_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTVALIDATIONINPROGRESS string = "ENVIRONMENT_VALIDATION_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusPREREQUISITESCREATEINPROGRESS captures enum value "PREREQUISITES_CREATE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusPREREQUISITESCREATEINPROGRESS string = "PREREQUISITES_CREATE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS captures enum value "NETWORK_CREATION_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusNETWORKCREATIONINPROGRESS string = "NETWORK_CREATION_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS captures enum value "NETWORK_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusNETWORKDELETEINPROGRESS string = "NETWORK_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS captures enum value "RDBMS_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusRDBMSDELETEINPROGRESS string = "RDBMS_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS captures enum value "FREEIPA_CREATION_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPACREATIONINPROGRESS string = "FREEIPA_CREATION_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS captures enum value "FREEIPA_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEINPROGRESS string = "FREEIPA_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusEXPERIENCEDELETEINPROGRESS captures enum value "EXPERIENCE_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusEXPERIENCEDELETEINPROGRESS string = "EXPERIENCE_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusCLUSTERDEFINITIONCLEANUPPROGRESS captures enum value "CLUSTER_DEFINITION_CLEANUP_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusCLUSTERDEFINITIONCLEANUPPROGRESS string = "CLUSTER_DEFINITION_CLEANUP_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusUMSRESOURCEDELETEINPROGRESS captures enum value "UMS_RESOURCE_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusUMSRESOURCEDELETEINPROGRESS string = "UMS_RESOURCE_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusIDBROKERMAPPINGSDELETEINPROGRESS captures enum value "IDBROKER_MAPPINGS_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusIDBROKERMAPPINGSDELETEINPROGRESS string = "IDBROKER_MAPPINGS_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusS3GUARDTABLEDELETEINPROGRESS captures enum value "S3GUARD_TABLE_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusS3GUARDTABLEDELETEINPROGRESS string = "S3GUARD_TABLE_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusDATAHUBCLUSTERSDELETEINPROGRESS captures enum value "DATAHUB_CLUSTERS_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusDATAHUBCLUSTERSDELETEINPROGRESS string = "DATAHUB_CLUSTERS_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusDATALAKECLUSTERSDELETEINPROGRESS captures enum value "DATALAKE_CLUSTERS_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusDATALAKECLUSTERSDELETEINPROGRESS string = "DATALAKE_CLUSTERS_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusPUBLICKEYCREATEINPROGRESS captures enum value "PUBLICKEY_CREATE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusPUBLICKEYCREATEINPROGRESS string = "PUBLICKEY_CREATE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusPUBLICKEYDELETEINPROGRESS captures enum value "PUBLICKEY_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusPUBLICKEYDELETEINPROGRESS string = "PUBLICKEY_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTRESOURCEENCRYPTIONINITIALIZATIONINPROGRESS captures enum value "ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTRESOURCEENCRYPTIONINITIALIZATIONINPROGRESS string = "ENVIRONMENT_RESOURCE_ENCRYPTION_INITIALIZATION_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTRESOURCEENCRYPTIONDELETEINPROGRESS captures enum value "ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTRESOURCEENCRYPTIONDELETEINPROGRESS string = "ENVIRONMENT_RESOURCE_ENCRYPTION_DELETE_IN_PROGRESS"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTENCRYPTIONRESOURCESINITIALIZED captures enum value "ENVIRONMENT_ENCRYPTION_RESOURCES_INITIALIZED"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVIRONMENTENCRYPTIONRESOURCESINITIALIZED string = "ENVIRONMENT_ENCRYPTION_RESOURCES_INITIALIZED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusAVAILABLE captures enum value "AVAILABLE"
	DetailedEnvironmentV1ResponseEnvironmentStatusAVAILABLE string = "AVAILABLE"

	// DetailedEnvironmentV1ResponseEnvironmentStatusARCHIVED captures enum value "ARCHIVED"
	DetailedEnvironmentV1ResponseEnvironmentStatusARCHIVED string = "ARCHIVED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusCREATEFAILED captures enum value "CREATE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusCREATEFAILED string = "CREATE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusDELETEFAILED captures enum value "DELETE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusDELETEFAILED string = "DELETE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusUPDATEFAILED string = "UPDATE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBSTARTED captures enum value "STOP_DATAHUB_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBSTARTED string = "STOP_DATAHUB_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBFAILED captures enum value "STOP_DATAHUB_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATAHUBFAILED string = "STOP_DATAHUB_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKESTARTED captures enum value "STOP_DATALAKE_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKESTARTED string = "STOP_DATALAKE_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKEFAILED captures enum value "STOP_DATALAKE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPDATALAKEFAILED string = "STOP_DATALAKE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPASTARTED captures enum value "STOP_FREEIPA_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPASTARTED string = "STOP_FREEIPA_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPAFAILED captures enum value "STOP_FREEIPA_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTOPFREEIPAFAILED string = "STOP_FREEIPA_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusENVSTOPPED captures enum value "ENV_STOPPED"
	DetailedEnvironmentV1ResponseEnvironmentStatusENVSTOPPED string = "ENV_STOPPED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBSTARTED captures enum value "START_DATAHUB_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBSTARTED string = "START_DATAHUB_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBFAILED captures enum value "START_DATAHUB_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATAHUBFAILED string = "START_DATAHUB_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKESTARTED captures enum value "START_DATALAKE_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKESTARTED string = "START_DATALAKE_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKEFAILED captures enum value "START_DATALAKE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTDATALAKEFAILED string = "START_DATALAKE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPASTARTED captures enum value "START_FREEIPA_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPASTARTED string = "START_FREEIPA_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPAFAILED captures enum value "START_FREEIPA_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTFREEIPAFAILED string = "START_FREEIPA_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTSYNCHRONIZEUSERSSTARTED captures enum value "START_SYNCHRONIZE_USERS_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTSYNCHRONIZEUSERSSTARTED string = "START_SYNCHRONIZE_USERS_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusSTARTSYNCHRONIZEUSERSFAILED captures enum value "START_SYNCHRONIZE_USERS_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusSTARTSYNCHRONIZEUSERSFAILED string = "START_SYNCHRONIZE_USERS_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEDONPROVIDERSIDE captures enum value "FREEIPA_DELETED_ON_PROVIDER_SIDE"
	DetailedEnvironmentV1ResponseEnvironmentStatusFREEIPADELETEDONPROVIDERSIDE string = "FREEIPA_DELETED_ON_PROVIDER_SIDE"

	// DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERENVUPDATESTARTED captures enum value "LOAD_BALANCER_ENV_UPDATE_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERENVUPDATESTARTED string = "LOAD_BALANCER_ENV_UPDATE_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERENVUPDATEFAILED captures enum value "LOAD_BALANCER_ENV_UPDATE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERENVUPDATEFAILED string = "LOAD_BALANCER_ENV_UPDATE_FAILED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERSTACKUPDATESTARTED captures enum value "LOAD_BALANCER_STACK_UPDATE_STARTED"
	DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERSTACKUPDATESTARTED string = "LOAD_BALANCER_STACK_UPDATE_STARTED"

	// DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERSTACKUPDATEFAILED captures enum value "LOAD_BALANCER_STACK_UPDATE_FAILED"
	DetailedEnvironmentV1ResponseEnvironmentStatusLOADBALANCERSTACKUPDATEFAILED string = "LOAD_BALANCER_STACK_UPDATE_FAILED"
)

// prop value enum
func (m *DetailedEnvironmentV1Response) validateEnvironmentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedEnvironmentV1ResponseTypeEnvironmentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedEnvironmentV1Response) validateEnvironmentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentStatusEnum("environmentStatus", "body", m.EnvironmentStatus); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateFreeIpa(formats strfmt.Registry) error {

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

var detailedEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","MOCK","IDBMMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum = append(detailedEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum, v)
	}
}

const (

	// DetailedEnvironmentV1ResponseIDBrokerMappingSourceNONE captures enum value "NONE"
	DetailedEnvironmentV1ResponseIDBrokerMappingSourceNONE string = "NONE"

	// DetailedEnvironmentV1ResponseIDBrokerMappingSourceMOCK captures enum value "MOCK"
	DetailedEnvironmentV1ResponseIDBrokerMappingSourceMOCK string = "MOCK"

	// DetailedEnvironmentV1ResponseIDBrokerMappingSourceIDBMMS captures enum value "IDBMMS"
	DetailedEnvironmentV1ResponseIDBrokerMappingSourceIDBMMS string = "IDBMMS"
)

// prop value enum
func (m *DetailedEnvironmentV1Response) validateIDBrokerMappingSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedEnvironmentV1ResponseTypeIDBrokerMappingSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedEnvironmentV1Response) validateIDBrokerMappingSource(formats strfmt.Registry) error {

	if swag.IsZero(m.IDBrokerMappingSource) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDBrokerMappingSourceEnum("idBrokerMappingSource", "body", m.IDBrokerMappingSource); err != nil {
		return err
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
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

func (m *DetailedEnvironmentV1Response) validateNetwork(formats strfmt.Registry) error {

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

func (m *DetailedEnvironmentV1Response) validateProxyConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ProxyConfig) { // not required
		return nil
	}

	if m.ProxyConfig != nil {
		if err := m.ProxyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyConfig")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if m.Regions != nil {
		if err := m.Regions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regions")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateSecurityAccess(formats strfmt.Registry) error {

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

func (m *DetailedEnvironmentV1Response) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedEnvironmentV1Response) validateTelemetry(formats strfmt.Registry) error {

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

var detailedEnvironmentV1ResponseTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM","CLUSTER_PROXY","CCMV2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailedEnvironmentV1ResponseTypeTunnelPropEnum = append(detailedEnvironmentV1ResponseTypeTunnelPropEnum, v)
	}
}

const (

	// DetailedEnvironmentV1ResponseTunnelDIRECT captures enum value "DIRECT"
	DetailedEnvironmentV1ResponseTunnelDIRECT string = "DIRECT"

	// DetailedEnvironmentV1ResponseTunnelCCM captures enum value "CCM"
	DetailedEnvironmentV1ResponseTunnelCCM string = "CCM"

	// DetailedEnvironmentV1ResponseTunnelCLUSTERPROXY captures enum value "CLUSTER_PROXY"
	DetailedEnvironmentV1ResponseTunnelCLUSTERPROXY string = "CLUSTER_PROXY"

	// DetailedEnvironmentV1ResponseTunnelCCMV2 captures enum value "CCMV2"
	DetailedEnvironmentV1ResponseTunnelCCMV2 string = "CCMV2"
)

// prop value enum
func (m *DetailedEnvironmentV1Response) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, detailedEnvironmentV1ResponseTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DetailedEnvironmentV1Response) validateTunnel(formats strfmt.Registry) error {

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
func (m *DetailedEnvironmentV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedEnvironmentV1Response) UnmarshalBinary(b []byte) error {
	var res DetailedEnvironmentV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
