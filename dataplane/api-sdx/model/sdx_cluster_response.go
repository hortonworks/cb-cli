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

// SdxClusterResponse sdx cluster response
// swagger:model SdxClusterResponse
type SdxClusterResponse struct {

	// cloud storage base location
	CloudStorageBaseLocation string `json:"cloudStorageBaseLocation,omitempty"`

	// cloud storage file system type
	// Enum: [WASB_INTEGRATED GCS WASB ADLS ADLS_GEN_2 S3]
	CloudStorageFileSystemType string `json:"cloudStorageFileSystemType,omitempty"`

	// cluster shape
	// Enum: [CUSTOM LIGHT_DUTY MEDIUM_DUTY_HA]
	ClusterShape string `json:"clusterShape,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// crn
	Crn string `json:"crn,omitempty"`

	// database server crn
	DatabaseServerCrn string `json:"databaseServerCrn,omitempty"`

	// environment crn
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// runtime
	Runtime string `json:"runtime,omitempty"`

	// stack crn
	StackCrn string `json:"stackCrn,omitempty"`

	// status
	// Enum: [REQUESTED WAIT_FOR_ENVIRONMENT ENVIRONMENT_CREATED STACK_CREATION_IN_PROGRESS STACK_CREATION_FINISHED STACK_DELETED STACK_DELETION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_CREATED RUNNING PROVISIONING_FAILED REPAIR_IN_PROGRESS REPAIR_FAILED CHANGE_IMAGE_IN_PROGRESS UPGRADE_IN_PROGRESS UPGRADE_FAILED DELETE_REQUESTED DELETED DELETE_FAILED START_REQUESTED START_IN_PROGRESS START_FAILED STOP_REQUESTED STOP_IN_PROGRESS STOP_FAILED STOPPED CLUSTER_AMBIGUOUS]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this sdx cluster response
func (m *SdxClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudStorageFileSystemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterShape(formats); err != nil {
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

var sdxClusterResponseTypeCloudStorageFileSystemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WASB_INTEGRATED","GCS","WASB","ADLS","ADLS_GEN_2","S3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterResponseTypeCloudStorageFileSystemTypePropEnum = append(sdxClusterResponseTypeCloudStorageFileSystemTypePropEnum, v)
	}
}

const (

	// SdxClusterResponseCloudStorageFileSystemTypeWASBINTEGRATED captures enum value "WASB_INTEGRATED"
	SdxClusterResponseCloudStorageFileSystemTypeWASBINTEGRATED string = "WASB_INTEGRATED"

	// SdxClusterResponseCloudStorageFileSystemTypeGCS captures enum value "GCS"
	SdxClusterResponseCloudStorageFileSystemTypeGCS string = "GCS"

	// SdxClusterResponseCloudStorageFileSystemTypeWASB captures enum value "WASB"
	SdxClusterResponseCloudStorageFileSystemTypeWASB string = "WASB"

	// SdxClusterResponseCloudStorageFileSystemTypeADLS captures enum value "ADLS"
	SdxClusterResponseCloudStorageFileSystemTypeADLS string = "ADLS"

	// SdxClusterResponseCloudStorageFileSystemTypeADLSGEN2 captures enum value "ADLS_GEN_2"
	SdxClusterResponseCloudStorageFileSystemTypeADLSGEN2 string = "ADLS_GEN_2"

	// SdxClusterResponseCloudStorageFileSystemTypeS3 captures enum value "S3"
	SdxClusterResponseCloudStorageFileSystemTypeS3 string = "S3"
)

// prop value enum
func (m *SdxClusterResponse) validateCloudStorageFileSystemTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterResponseTypeCloudStorageFileSystemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterResponse) validateCloudStorageFileSystemType(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorageFileSystemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudStorageFileSystemTypeEnum("cloudStorageFileSystemType", "body", m.CloudStorageFileSystemType); err != nil {
		return err
	}

	return nil
}

var sdxClusterResponseTypeClusterShapePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOM","LIGHT_DUTY","MEDIUM_DUTY_HA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterResponseTypeClusterShapePropEnum = append(sdxClusterResponseTypeClusterShapePropEnum, v)
	}
}

const (

	// SdxClusterResponseClusterShapeCUSTOM captures enum value "CUSTOM"
	SdxClusterResponseClusterShapeCUSTOM string = "CUSTOM"

	// SdxClusterResponseClusterShapeLIGHTDUTY captures enum value "LIGHT_DUTY"
	SdxClusterResponseClusterShapeLIGHTDUTY string = "LIGHT_DUTY"

	// SdxClusterResponseClusterShapeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	SdxClusterResponseClusterShapeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"
)

// prop value enum
func (m *SdxClusterResponse) validateClusterShapeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterResponseTypeClusterShapePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterResponse) validateClusterShape(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterShape) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterShapeEnum("clusterShape", "body", m.ClusterShape); err != nil {
		return err
	}

	return nil
}

var sdxClusterResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","WAIT_FOR_ENVIRONMENT","ENVIRONMENT_CREATED","STACK_CREATION_IN_PROGRESS","STACK_CREATION_FINISHED","STACK_DELETED","STACK_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_CREATED","RUNNING","PROVISIONING_FAILED","REPAIR_IN_PROGRESS","REPAIR_FAILED","CHANGE_IMAGE_IN_PROGRESS","UPGRADE_IN_PROGRESS","UPGRADE_FAILED","DELETE_REQUESTED","DELETED","DELETE_FAILED","START_REQUESTED","START_IN_PROGRESS","START_FAILED","STOP_REQUESTED","STOP_IN_PROGRESS","STOP_FAILED","STOPPED","CLUSTER_AMBIGUOUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sdxClusterResponseTypeStatusPropEnum = append(sdxClusterResponseTypeStatusPropEnum, v)
	}
}

const (

	// SdxClusterResponseStatusREQUESTED captures enum value "REQUESTED"
	SdxClusterResponseStatusREQUESTED string = "REQUESTED"

	// SdxClusterResponseStatusWAITFORENVIRONMENT captures enum value "WAIT_FOR_ENVIRONMENT"
	SdxClusterResponseStatusWAITFORENVIRONMENT string = "WAIT_FOR_ENVIRONMENT"

	// SdxClusterResponseStatusENVIRONMENTCREATED captures enum value "ENVIRONMENT_CREATED"
	SdxClusterResponseStatusENVIRONMENTCREATED string = "ENVIRONMENT_CREATED"

	// SdxClusterResponseStatusSTACKCREATIONINPROGRESS captures enum value "STACK_CREATION_IN_PROGRESS"
	SdxClusterResponseStatusSTACKCREATIONINPROGRESS string = "STACK_CREATION_IN_PROGRESS"

	// SdxClusterResponseStatusSTACKCREATIONFINISHED captures enum value "STACK_CREATION_FINISHED"
	SdxClusterResponseStatusSTACKCREATIONFINISHED string = "STACK_CREATION_FINISHED"

	// SdxClusterResponseStatusSTACKDELETED captures enum value "STACK_DELETED"
	SdxClusterResponseStatusSTACKDELETED string = "STACK_DELETED"

	// SdxClusterResponseStatusSTACKDELETIONINPROGRESS captures enum value "STACK_DELETION_IN_PROGRESS"
	SdxClusterResponseStatusSTACKDELETIONINPROGRESS string = "STACK_DELETION_IN_PROGRESS"

	// SdxClusterResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	SdxClusterResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// SdxClusterResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	SdxClusterResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// SdxClusterResponseStatusEXTERNALDATABASECREATED captures enum value "EXTERNAL_DATABASE_CREATED"
	SdxClusterResponseStatusEXTERNALDATABASECREATED string = "EXTERNAL_DATABASE_CREATED"

	// SdxClusterResponseStatusRUNNING captures enum value "RUNNING"
	SdxClusterResponseStatusRUNNING string = "RUNNING"

	// SdxClusterResponseStatusPROVISIONINGFAILED captures enum value "PROVISIONING_FAILED"
	SdxClusterResponseStatusPROVISIONINGFAILED string = "PROVISIONING_FAILED"

	// SdxClusterResponseStatusREPAIRINPROGRESS captures enum value "REPAIR_IN_PROGRESS"
	SdxClusterResponseStatusREPAIRINPROGRESS string = "REPAIR_IN_PROGRESS"

	// SdxClusterResponseStatusREPAIRFAILED captures enum value "REPAIR_FAILED"
	SdxClusterResponseStatusREPAIRFAILED string = "REPAIR_FAILED"

	// SdxClusterResponseStatusCHANGEIMAGEINPROGRESS captures enum value "CHANGE_IMAGE_IN_PROGRESS"
	SdxClusterResponseStatusCHANGEIMAGEINPROGRESS string = "CHANGE_IMAGE_IN_PROGRESS"

	// SdxClusterResponseStatusUPGRADEINPROGRESS captures enum value "UPGRADE_IN_PROGRESS"
	SdxClusterResponseStatusUPGRADEINPROGRESS string = "UPGRADE_IN_PROGRESS"

	// SdxClusterResponseStatusUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	SdxClusterResponseStatusUPGRADEFAILED string = "UPGRADE_FAILED"

	// SdxClusterResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	SdxClusterResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// SdxClusterResponseStatusDELETED captures enum value "DELETED"
	SdxClusterResponseStatusDELETED string = "DELETED"

	// SdxClusterResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	SdxClusterResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// SdxClusterResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	SdxClusterResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// SdxClusterResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	SdxClusterResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// SdxClusterResponseStatusSTARTFAILED captures enum value "START_FAILED"
	SdxClusterResponseStatusSTARTFAILED string = "START_FAILED"

	// SdxClusterResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	SdxClusterResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// SdxClusterResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	SdxClusterResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// SdxClusterResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	SdxClusterResponseStatusSTOPFAILED string = "STOP_FAILED"

	// SdxClusterResponseStatusSTOPPED captures enum value "STOPPED"
	SdxClusterResponseStatusSTOPPED string = "STOPPED"

	// SdxClusterResponseStatusCLUSTERAMBIGUOUS captures enum value "CLUSTER_AMBIGUOUS"
	SdxClusterResponseStatusCLUSTERAMBIGUOUS string = "CLUSTER_AMBIGUOUS"
)

// prop value enum
func (m *SdxClusterResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sdxClusterResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SdxClusterResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *SdxClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SdxClusterResponse) UnmarshalBinary(b []byte) error {
	var res SdxClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
