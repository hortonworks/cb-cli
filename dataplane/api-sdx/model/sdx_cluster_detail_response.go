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

// SdxClusterDetailResponse sdx cluster detail response
// swagger:model SdxClusterDetailResponse
type SdxClusterDetailResponse struct {

	// Indicates the certificate status on the cluster
	// Enum: [VALID HOST_CERT_EXPIRING]
	CertExpirationState string `json:"certExpirationState,omitempty"`

	// Cloud storage base location.
	CloudStorageBaseLocation string `json:"cloudStorageBaseLocation,omitempty"`

	// Cloud storage file system type.
	// Enum: [WASB_INTEGRATED GCS WASB ADLS ADLS_GEN_2 S3 EFS]
	CloudStorageFileSystemType string `json:"cloudStorageFileSystemType,omitempty"`

	// The shape of the cluster such as Micro Duty, Light Duty, Medium Duty...
	// Enum: [CUSTOM LIGHT_DUTY MEDIUM_DUTY_HA MICRO_DUTY]
	ClusterShape string `json:"clusterShape,omitempty"`

	// Timestamp when the resource was created.
	Created int64 `json:"created,omitempty"`

	// Data Lake crn.
	Crn string `json:"crn,omitempty"`

	// Database engine version.
	DatabaseEngineVersion string `json:"databaseEngineVersion,omitempty"`

	// Database server crn.
	DatabaseServerCrn string `json:"databaseServerCrn,omitempty"`

	// Shows whether the cluster is in detached state.
	Detached bool `json:"detached,omitempty"`

	// Option to enable multi AZ.
	EnableMultiAz bool `json:"enableMultiAz,omitempty"`

	// Environment crn.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The name of the environment.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The id of the flow or flow chain that was triggered as part of the process.
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// Data Lake name.
	Name string `json:"name,omitempty"`

	// Option to enable ranger raz.
	RangerRazEnabled bool `json:"rangerRazEnabled,omitempty"`

	// Runtime version.
	Runtime string `json:"runtime,omitempty"`

	// Data Lake cluster service version.
	SdxClusterServiceVersion string `json:"sdxClusterServiceVersion,omitempty"`

	// Stack crn.
	StackCrn string `json:"stackCrn,omitempty"`

	// Stack response.
	StackV4Response *StackV4Response `json:"stackV4Response,omitempty"`

	// Data Lake status.
	// Enum: [REQUESTED WAIT_FOR_ENVIRONMENT ENVIRONMENT_CREATED STACK_CREATION_IN_PROGRESS STACK_CREATION_FINISHED STACK_DELETED STACK_DELETION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_STARTED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOPPED EXTERNAL_DATABASE_CREATED RUNNING PROVISIONING_FAILED REPAIR_IN_PROGRESS REPAIR_FAILED CHANGE_IMAGE_IN_PROGRESS DATALAKE_UPGRADE_IN_PROGRESS DATALAKE_UPGRADE_FAILED RECOVERY_IN_PROGRESS RECOVERY_FAILED DELETE_REQUESTED DELETED DELETE_FAILED DELETED_ON_PROVIDER_SIDE START_IN_PROGRESS START_FAILED STOP_IN_PROGRESS STOP_FAILED STOPPED CLUSTER_AMBIGUOUS CLUSTER_UNREACHABLE NODE_FAILURE SYNC_FAILED CERT_ROTATION_IN_PROGRESS CERT_ROTATION_FAILED CERT_ROTATION_FINISHED CERT_RENEWAL_IN_PROGRESS CERT_RENEWAL_FAILED CERT_RENEWAL_FINISHED DATALAKE_BACKUP_INPROGRESS DATALAKE_RESTORE_INPROGRESS DATALAKE_RESTORE_FAILED DATALAKE_DETACHED DATALAKE_UPGRADE_CCM_IN_PROGRESS DATALAKE_UPGRADE_CCM_FAILED DATAHUB_REFRESH_IN_PROGRESS DATAHUB_REFRESH_FAILED SALT_PASSWORD_ROTATION_IN_PROGRESS SALT_PASSWORD_ROTATION_FAILED SALT_PASSWORD_ROTATION_FINISHED DATALAKE_UPGRADE_PREPARATION_IN_PROGRESS DATALAKE_UPGRADE_PREPARATION_FAILED]
	Status string `json:"status,omitempty"`

	// Data Lake status reason.
	StatusReason string `json:"statusReason,omitempty"`

	// Tags.
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sdx cluster detail response
func (m *SdxClusterDetailResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertExpirationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorageFileSystemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackV4Response(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sdxClusterDetailResponseTypeCertExpirationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VALID","HOST_CERT_EXPIRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeCertExpirationStatePropEnum = append(sdxClusterDetailResponseTypeCertExpirationStatePropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseCertExpirationStateVALID captures enum value "VALID"
	SdxClusterDetailResponseCertExpirationStateVALID string = "VALID"

	// SdxClusterDetailResponseCertExpirationStateHOSTCERTEXPIRING captures enum value "HOST_CERT_EXPIRING"
	SdxClusterDetailResponseCertExpirationStateHOSTCERTEXPIRING string = "HOST_CERT_EXPIRING"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateCertExpirationStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeCertExpirationStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateCertExpirationState(formats strfmt.Registry) error {

	if swag.IsZero(m.CertExpirationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateCertExpirationStateEnum("certExpirationState", "body", m.CertExpirationState); err != nil {
		return err
	}

	return nil
}

var sdxClusterDetailResponseTypeCloudStorageFileSystemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WASB_INTEGRATED","GCS","WASB","ADLS","ADLS_GEN_2","S3","EFS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeCloudStorageFileSystemTypePropEnum = append(sdxClusterDetailResponseTypeCloudStorageFileSystemTypePropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseCloudStorageFileSystemTypeWASBINTEGRATED captures enum value "WASB_INTEGRATED"
	SdxClusterDetailResponseCloudStorageFileSystemTypeWASBINTEGRATED string = "WASB_INTEGRATED"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeGCS captures enum value "GCS"
	SdxClusterDetailResponseCloudStorageFileSystemTypeGCS string = "GCS"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeWASB captures enum value "WASB"
	SdxClusterDetailResponseCloudStorageFileSystemTypeWASB string = "WASB"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeADLS captures enum value "ADLS"
	SdxClusterDetailResponseCloudStorageFileSystemTypeADLS string = "ADLS"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeADLSGEN2 captures enum value "ADLS_GEN_2"
	SdxClusterDetailResponseCloudStorageFileSystemTypeADLSGEN2 string = "ADLS_GEN_2"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeS3 captures enum value "S3"
	SdxClusterDetailResponseCloudStorageFileSystemTypeS3 string = "S3"

	// SdxClusterDetailResponseCloudStorageFileSystemTypeEFS captures enum value "EFS"
	SdxClusterDetailResponseCloudStorageFileSystemTypeEFS string = "EFS"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateCloudStorageFileSystemTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeCloudStorageFileSystemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateCloudStorageFileSystemType(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorageFileSystemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudStorageFileSystemTypeEnum("cloudStorageFileSystemType", "body", m.CloudStorageFileSystemType); err != nil {
		return err
	}

	return nil
}

var sdxClusterDetailResponseTypeClusterShapePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","LIGHT_DUTY","MEDIUM_DUTY_HA","MICRO_DUTY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeClusterShapePropEnum = append(sdxClusterDetailResponseTypeClusterShapePropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseClusterShapeCUSTOM captures enum value "CUSTOM"
	SdxClusterDetailResponseClusterShapeCUSTOM string = "CUSTOM"

	// SdxClusterDetailResponseClusterShapeLIGHTDUTY captures enum value "LIGHT_DUTY"
	SdxClusterDetailResponseClusterShapeLIGHTDUTY string = "LIGHT_DUTY"

	// SdxClusterDetailResponseClusterShapeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	SdxClusterDetailResponseClusterShapeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"

	// SdxClusterDetailResponseClusterShapeMICRODUTY captures enum value "MICRO_DUTY"
	SdxClusterDetailResponseClusterShapeMICRODUTY string = "MICRO_DUTY"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateClusterShapeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeClusterShapePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateClusterShape(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterShape) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterShapeEnum("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	return nil
}

func (m *SdxClusterDetailResponse) validateFlowIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowIdentifier) { // not required
		return nil
	}

	if m.FlowIdentifier != nil {
		if err := m.FlowIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *SdxClusterDetailResponse) validateStackV4Response(formats strfmt.Registry) error {

	if swag.IsZero(m.StackV4Response) { // not required
		return nil
	}

	if m.StackV4Response != nil {
		if err := m.StackV4Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackV4Response")
			}
			return err
		}
	}

	return nil
}

var sdxClusterDetailResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","WAIT_FOR_ENVIRONMENT","ENVIRONMENT_CREATED","STACK_CREATION_IN_PROGRESS","STACK_CREATION_FINISHED","STACK_DELETED","STACK_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_STARTED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOPPED","EXTERNAL_DATABASE_CREATED","RUNNING","PROVISIONING_FAILED","REPAIR_IN_PROGRESS","REPAIR_FAILED","CHANGE_IMAGE_IN_PROGRESS","DATALAKE_UPGRADE_IN_PROGRESS","DATALAKE_UPGRADE_FAILED","RECOVERY_IN_PROGRESS","RECOVERY_FAILED","DELETE_REQUESTED","DELETED","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","START_IN_PROGRESS","START_FAILED","STOP_IN_PROGRESS","STOP_FAILED","STOPPED","CLUSTER_AMBIGUOUS","CLUSTER_UNREACHABLE","NODE_FAILURE","SYNC_FAILED","CERT_ROTATION_IN_PROGRESS","CERT_ROTATION_FAILED","CERT_ROTATION_FINISHED","CERT_RENEWAL_IN_PROGRESS","CERT_RENEWAL_FAILED","CERT_RENEWAL_FINISHED","DATALAKE_BACKUP_INPROGRESS","DATALAKE_RESTORE_INPROGRESS","DATALAKE_RESTORE_FAILED","DATALAKE_DETACHED","DATALAKE_UPGRADE_CCM_IN_PROGRESS","DATALAKE_UPGRADE_CCM_FAILED","DATAHUB_REFRESH_IN_PROGRESS","DATAHUB_REFRESH_FAILED","SALT_PASSWORD_ROTATION_IN_PROGRESS","SALT_PASSWORD_ROTATION_FAILED","SALT_PASSWORD_ROTATION_FINISHED","DATALAKE_UPGRADE_PREPARATION_IN_PROGRESS","DATALAKE_UPGRADE_PREPARATION_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterDetailResponseTypeStatusPropEnum = append(sdxClusterDetailResponseTypeStatusPropEnum, v)
	}
}

const (

	// SdxClusterDetailResponseStatusREQUESTED captures enum value "REQUESTED"
	SdxClusterDetailResponseStatusREQUESTED string = "REQUESTED"

	// SdxClusterDetailResponseStatusWAITFORENVIRONMENT captures enum value "WAIT_FOR_ENVIRONMENT"
	SdxClusterDetailResponseStatusWAITFORENVIRONMENT string = "WAIT_FOR_ENVIRONMENT"

	// SdxClusterDetailResponseStatusENVIRONMENTCREATED captures enum value "ENVIRONMENT_CREATED"
	SdxClusterDetailResponseStatusENVIRONMENTCREATED string = "ENVIRONMENT_CREATED"

	// SdxClusterDetailResponseStatusSTACKCREATIONINPROGRESS captures enum value "STACK_CREATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTACKCREATIONINPROGRESS string = "STACK_CREATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusSTACKCREATIONFINISHED captures enum value "STACK_CREATION_FINISHED"
	SdxClusterDetailResponseStatusSTACKCREATIONFINISHED string = "STACK_CREATION_FINISHED"

	// SdxClusterDetailResponseStatusSTACKDELETED captures enum value "STACK_DELETED"
	SdxClusterDetailResponseStatusSTACKDELETED string = "STACK_DELETED"

	// SdxClusterDetailResponseStatusSTACKDELETIONINPROGRESS captures enum value "STACK_DELETION_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTACKDELETIONINPROGRESS string = "STACK_DELETION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASESTARTED captures enum value "EXTERNAL_DATABASE_STARTED"
	SdxClusterDetailResponseStatusEXTERNALDATABASESTARTED string = "EXTERNAL_DATABASE_STARTED"

	// SdxClusterDetailResponseStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	SdxClusterDetailResponseStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// SdxClusterDetailResponseStatusEXTERNALDATABASESTOPPED captures enum value "EXTERNAL_DATABASE_STOPPED"
	SdxClusterDetailResponseStatusEXTERNALDATABASESTOPPED string = "EXTERNAL_DATABASE_STOPPED"

	// SdxClusterDetailResponseStatusEXTERNALDATABASECREATED captures enum value "EXTERNAL_DATABASE_CREATED"
	SdxClusterDetailResponseStatusEXTERNALDATABASECREATED string = "EXTERNAL_DATABASE_CREATED"

	// SdxClusterDetailResponseStatusRUNNING captures enum value "RUNNING"
	SdxClusterDetailResponseStatusRUNNING string = "RUNNING"

	// SdxClusterDetailResponseStatusPROVISIONINGFAILED captures enum value "PROVISIONING_FAILED"
	SdxClusterDetailResponseStatusPROVISIONINGFAILED string = "PROVISIONING_FAILED"

	// SdxClusterDetailResponseStatusREPAIRINPROGRESS captures enum value "REPAIR_IN_PROGRESS"
	SdxClusterDetailResponseStatusREPAIRINPROGRESS string = "REPAIR_IN_PROGRESS"

	// SdxClusterDetailResponseStatusREPAIRFAILED captures enum value "REPAIR_FAILED"
	SdxClusterDetailResponseStatusREPAIRFAILED string = "REPAIR_FAILED"

	// SdxClusterDetailResponseStatusCHANGEIMAGEINPROGRESS captures enum value "CHANGE_IMAGE_IN_PROGRESS"
	SdxClusterDetailResponseStatusCHANGEIMAGEINPROGRESS string = "CHANGE_IMAGE_IN_PROGRESS"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADEINPROGRESS captures enum value "DATALAKE_UPGRADE_IN_PROGRESS"
	SdxClusterDetailResponseStatusDATALAKEUPGRADEINPROGRESS string = "DATALAKE_UPGRADE_IN_PROGRESS"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADEFAILED captures enum value "DATALAKE_UPGRADE_FAILED"
	SdxClusterDetailResponseStatusDATALAKEUPGRADEFAILED string = "DATALAKE_UPGRADE_FAILED"

	// SdxClusterDetailResponseStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	SdxClusterDetailResponseStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// SdxClusterDetailResponseStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	SdxClusterDetailResponseStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// SdxClusterDetailResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	SdxClusterDetailResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// SdxClusterDetailResponseStatusDELETED captures enum value "DELETED"
	SdxClusterDetailResponseStatusDELETED string = "DELETED"

	// SdxClusterDetailResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SdxClusterDetailResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// SdxClusterDetailResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	SdxClusterDetailResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// SdxClusterDetailResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// SdxClusterDetailResponseStatusSTARTFAILED captures enum value "START_FAILED"
	SdxClusterDetailResponseStatusSTARTFAILED string = "START_FAILED"

	// SdxClusterDetailResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	SdxClusterDetailResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// SdxClusterDetailResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	SdxClusterDetailResponseStatusSTOPFAILED string = "STOP_FAILED"

	// SdxClusterDetailResponseStatusSTOPPED captures enum value "STOPPED"
	SdxClusterDetailResponseStatusSTOPPED string = "STOPPED"

	// SdxClusterDetailResponseStatusCLUSTERAMBIGUOUS captures enum value "CLUSTER_AMBIGUOUS"
	SdxClusterDetailResponseStatusCLUSTERAMBIGUOUS string = "CLUSTER_AMBIGUOUS"

	// SdxClusterDetailResponseStatusCLUSTERUNREACHABLE captures enum value "CLUSTER_UNREACHABLE"
	SdxClusterDetailResponseStatusCLUSTERUNREACHABLE string = "CLUSTER_UNREACHABLE"

	// SdxClusterDetailResponseStatusNODEFAILURE captures enum value "NODE_FAILURE"
	SdxClusterDetailResponseStatusNODEFAILURE string = "NODE_FAILURE"

	// SdxClusterDetailResponseStatusSYNCFAILED captures enum value "SYNC_FAILED"
	SdxClusterDetailResponseStatusSYNCFAILED string = "SYNC_FAILED"

	// SdxClusterDetailResponseStatusCERTROTATIONINPROGRESS captures enum value "CERT_ROTATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusCERTROTATIONINPROGRESS string = "CERT_ROTATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusCERTROTATIONFAILED captures enum value "CERT_ROTATION_FAILED"
	SdxClusterDetailResponseStatusCERTROTATIONFAILED string = "CERT_ROTATION_FAILED"

	// SdxClusterDetailResponseStatusCERTROTATIONFINISHED captures enum value "CERT_ROTATION_FINISHED"
	SdxClusterDetailResponseStatusCERTROTATIONFINISHED string = "CERT_ROTATION_FINISHED"

	// SdxClusterDetailResponseStatusCERTRENEWALINPROGRESS captures enum value "CERT_RENEWAL_IN_PROGRESS"
	SdxClusterDetailResponseStatusCERTRENEWALINPROGRESS string = "CERT_RENEWAL_IN_PROGRESS"

	// SdxClusterDetailResponseStatusCERTRENEWALFAILED captures enum value "CERT_RENEWAL_FAILED"
	SdxClusterDetailResponseStatusCERTRENEWALFAILED string = "CERT_RENEWAL_FAILED"

	// SdxClusterDetailResponseStatusCERTRENEWALFINISHED captures enum value "CERT_RENEWAL_FINISHED"
	SdxClusterDetailResponseStatusCERTRENEWALFINISHED string = "CERT_RENEWAL_FINISHED"

	// SdxClusterDetailResponseStatusDATALAKEBACKUPINPROGRESS captures enum value "DATALAKE_BACKUP_INPROGRESS"
	SdxClusterDetailResponseStatusDATALAKEBACKUPINPROGRESS string = "DATALAKE_BACKUP_INPROGRESS"

	// SdxClusterDetailResponseStatusDATALAKERESTOREINPROGRESS captures enum value "DATALAKE_RESTORE_INPROGRESS"
	SdxClusterDetailResponseStatusDATALAKERESTOREINPROGRESS string = "DATALAKE_RESTORE_INPROGRESS"

	// SdxClusterDetailResponseStatusDATALAKERESTOREFAILED captures enum value "DATALAKE_RESTORE_FAILED"
	SdxClusterDetailResponseStatusDATALAKERESTOREFAILED string = "DATALAKE_RESTORE_FAILED"

	// SdxClusterDetailResponseStatusDATALAKEDETACHED captures enum value "DATALAKE_DETACHED"
	SdxClusterDetailResponseStatusDATALAKEDETACHED string = "DATALAKE_DETACHED"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADECCMINPROGRESS captures enum value "DATALAKE_UPGRADE_CCM_IN_PROGRESS"
	SdxClusterDetailResponseStatusDATALAKEUPGRADECCMINPROGRESS string = "DATALAKE_UPGRADE_CCM_IN_PROGRESS"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADECCMFAILED captures enum value "DATALAKE_UPGRADE_CCM_FAILED"
	SdxClusterDetailResponseStatusDATALAKEUPGRADECCMFAILED string = "DATALAKE_UPGRADE_CCM_FAILED"

	// SdxClusterDetailResponseStatusDATAHUBREFRESHINPROGRESS captures enum value "DATAHUB_REFRESH_IN_PROGRESS"
	SdxClusterDetailResponseStatusDATAHUBREFRESHINPROGRESS string = "DATAHUB_REFRESH_IN_PROGRESS"

	// SdxClusterDetailResponseStatusDATAHUBREFRESHFAILED captures enum value "DATAHUB_REFRESH_FAILED"
	SdxClusterDetailResponseStatusDATAHUBREFRESHFAILED string = "DATAHUB_REFRESH_FAILED"

	// SdxClusterDetailResponseStatusSALTPASSWORDROTATIONINPROGRESS captures enum value "SALT_PASSWORD_ROTATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusSALTPASSWORDROTATIONINPROGRESS string = "SALT_PASSWORD_ROTATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusSALTPASSWORDROTATIONFAILED captures enum value "SALT_PASSWORD_ROTATION_FAILED"
	SdxClusterDetailResponseStatusSALTPASSWORDROTATIONFAILED string = "SALT_PASSWORD_ROTATION_FAILED"

	// SdxClusterDetailResponseStatusSALTPASSWORDROTATIONFINISHED captures enum value "SALT_PASSWORD_ROTATION_FINISHED"
	SdxClusterDetailResponseStatusSALTPASSWORDROTATIONFINISHED string = "SALT_PASSWORD_ROTATION_FINISHED"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADEPREPARATIONINPROGRESS captures enum value "DATALAKE_UPGRADE_PREPARATION_IN_PROGRESS"
	SdxClusterDetailResponseStatusDATALAKEUPGRADEPREPARATIONINPROGRESS string = "DATALAKE_UPGRADE_PREPARATION_IN_PROGRESS"

	// SdxClusterDetailResponseStatusDATALAKEUPGRADEPREPARATIONFAILED captures enum value "DATALAKE_UPGRADE_PREPARATION_FAILED"
	SdxClusterDetailResponseStatusDATALAKEUPGRADEPREPARATIONFAILED string = "DATALAKE_UPGRADE_PREPARATION_FAILED"
)

// prop value enum
func (m *SdxClusterDetailResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterDetailResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterDetailResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SdxClusterDetailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxClusterDetailResponse) UnmarshalBinary(b []byte) error {
	var res SdxClusterDetailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
