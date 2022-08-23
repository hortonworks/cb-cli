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

// AutoscaleStackV4Response autoscale stack v4 response
// swagger:model AutoscaleStackV4Response
type AutoscaleStackV4Response struct {

	// Cloudplatform of the stack
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// public ambari ip of the stack
	ClusterManagerIP string `json:"clusterManagerIp,omitempty"`

	// Cluster manager variant
	ClusterManagerVariant string `json:"clusterManagerVariant,omitempty"`

	// status of the cluster
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS EXTERNAL_DATABASE_UPGRADE_FINISHED EXTERNAL_DATABASE_UPGRADE_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED UPGRADE_CCM_IN_PROGRESS UPGRADE_CCM_FAILED UPGRADE_CCM_FINISHED]
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// CRN of the environment which the stack is assigned to
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// port of the gateway secured proxy
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// name of the stack
	// Required: true
	Name *string `json:"name"`

	// ambari password
	PasswordPath string `json:"passwordPath,omitempty"`

	// the unique crn of the resource
	StackCrn string `json:"stackCrn,omitempty"`

	// id of the stack
	StackID int64 `json:"stackId,omitempty"`

	// the typeof the resource
	// Enum: [WORKLOAD DATALAKE TEMPLATE LEGACY]
	StackType string `json:"stackType,omitempty"`

	// status of the stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS EXTERNAL_DATABASE_UPGRADE_FINISHED EXTERNAL_DATABASE_UPGRADE_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED UPGRADE_CCM_IN_PROGRESS UPGRADE_CCM_FAILED UPGRADE_CCM_FINISHED]
	Status string `json:"status,omitempty"`

	// name of the tenant
	Tenant string `json:"tenant,omitempty"`

	// Configuration that the connection going directly or with cluster proxy or with ccm and cluster proxy.
	// Enum: [DIRECT CCM CLUSTER_PROXY CCMV2 CCMV2_JUMPGATE]
	Tunnel string `json:"tunnel,omitempty"`

	// crn of the user
	UserCrn string `json:"userCrn,omitempty"`

	// id of the user
	UserID string `json:"userId,omitempty"`

	// ambari username
	UserNamePath string `json:"userNamePath,omitempty"`

	// id of the workspace
	WorkspaceID int64 `json:"workspaceId,omitempty"`
}

// Validate validates this autoscale stack v4 response
func (m *AutoscaleStackV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

var autoscaleStackV4ResponseTypeClusterStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS","EXTERNAL_DATABASE_UPGRADE_FINISHED","EXTERNAL_DATABASE_UPGRADE_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED","UPGRADE_CCM_IN_PROGRESS","UPGRADE_CCM_FAILED","UPGRADE_CCM_FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackV4ResponseTypeClusterStatusPropEnum = append(autoscaleStackV4ResponseTypeClusterStatusPropEnum, v)
	}
}

const (

	// AutoscaleStackV4ResponseClusterStatusREQUESTED captures enum value "REQUESTED"
	AutoscaleStackV4ResponseClusterStatusREQUESTED string = "REQUESTED"

	// AutoscaleStackV4ResponseClusterStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusAVAILABLE captures enum value "AVAILABLE"
	AutoscaleStackV4ResponseClusterStatusAVAILABLE string = "AVAILABLE"

	// AutoscaleStackV4ResponseClusterStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	AutoscaleStackV4ResponseClusterStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// AutoscaleStackV4ResponseClusterStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	AutoscaleStackV4ResponseClusterStatusUPDATEFAILED string = "UPDATE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	AutoscaleStackV4ResponseClusterStatusBACKUPFAILED string = "BACKUP_FAILED"

	// AutoscaleStackV4ResponseClusterStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	AutoscaleStackV4ResponseClusterStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	AutoscaleStackV4ResponseClusterStatusRESTOREFAILED string = "RESTORE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	AutoscaleStackV4ResponseClusterStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	AutoscaleStackV4ResponseClusterStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// AutoscaleStackV4ResponseClusterStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	AutoscaleStackV4ResponseClusterStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// AutoscaleStackV4ResponseClusterStatusCREATEFAILED captures enum value "CREATE_FAILED"
	AutoscaleStackV4ResponseClusterStatusCREATEFAILED string = "CREATE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	AutoscaleStackV4ResponseClusterStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// AutoscaleStackV4ResponseClusterStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusDELETEFAILED captures enum value "DELETE_FAILED"
	AutoscaleStackV4ResponseClusterStatusDELETEFAILED string = "DELETE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	AutoscaleStackV4ResponseClusterStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// AutoscaleStackV4ResponseClusterStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	AutoscaleStackV4ResponseClusterStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// AutoscaleStackV4ResponseClusterStatusSTOPPED captures enum value "STOPPED"
	AutoscaleStackV4ResponseClusterStatusSTOPPED string = "STOPPED"

	// AutoscaleStackV4ResponseClusterStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	AutoscaleStackV4ResponseClusterStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// AutoscaleStackV4ResponseClusterStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	AutoscaleStackV4ResponseClusterStatusSTARTREQUESTED string = "START_REQUESTED"

	// AutoscaleStackV4ResponseClusterStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusSTARTFAILED captures enum value "START_FAILED"
	AutoscaleStackV4ResponseClusterStatusSTARTFAILED string = "START_FAILED"

	// AutoscaleStackV4ResponseClusterStatusSTOPFAILED captures enum value "STOP_FAILED"
	AutoscaleStackV4ResponseClusterStatusSTOPFAILED string = "STOP_FAILED"

	// AutoscaleStackV4ResponseClusterStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	AutoscaleStackV4ResponseClusterStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// AutoscaleStackV4ResponseClusterStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	AutoscaleStackV4ResponseClusterStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// AutoscaleStackV4ResponseClusterStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	AutoscaleStackV4ResponseClusterStatusAMBIGUOUS string = "AMBIGUOUS"

	// AutoscaleStackV4ResponseClusterStatusUNREACHABLE captures enum value "UNREACHABLE"
	AutoscaleStackV4ResponseClusterStatusUNREACHABLE string = "UNREACHABLE"

	// AutoscaleStackV4ResponseClusterStatusNODEFAILURE captures enum value "NODE_FAILURE"
	AutoscaleStackV4ResponseClusterStatusNODEFAILURE string = "NODE_FAILURE"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEINPROGRESS captures enum value "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEINPROGRESS string = "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFINISHED captures enum value "EXTERNAL_DATABASE_UPGRADE_FINISHED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFINISHED string = "EXTERNAL_DATABASE_UPGRADE_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFAILED captures enum value "EXTERNAL_DATABASE_UPGRADE_FAILED"
	AutoscaleStackV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFAILED string = "EXTERNAL_DATABASE_UPGRADE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	AutoscaleStackV4ResponseClusterStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"

	// AutoscaleStackV4ResponseClusterStatusUPGRADECCMINPROGRESS captures enum value "UPGRADE_CCM_IN_PROGRESS"
	AutoscaleStackV4ResponseClusterStatusUPGRADECCMINPROGRESS string = "UPGRADE_CCM_IN_PROGRESS"

	// AutoscaleStackV4ResponseClusterStatusUPGRADECCMFAILED captures enum value "UPGRADE_CCM_FAILED"
	AutoscaleStackV4ResponseClusterStatusUPGRADECCMFAILED string = "UPGRADE_CCM_FAILED"

	// AutoscaleStackV4ResponseClusterStatusUPGRADECCMFINISHED captures enum value "UPGRADE_CCM_FINISHED"
	AutoscaleStackV4ResponseClusterStatusUPGRADECCMFINISHED string = "UPGRADE_CCM_FINISHED"
)

// prop value enum
func (m *AutoscaleStackV4Response) validateClusterStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackV4ResponseTypeClusterStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackV4Response) validateClusterStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterStatusEnum("clusterStatus", "body", m.ClusterStatus); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleStackV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var autoscaleStackV4ResponseTypeStackTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WORKLOAD","DATALAKE","TEMPLATE","LEGACY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackV4ResponseTypeStackTypePropEnum = append(autoscaleStackV4ResponseTypeStackTypePropEnum, v)
	}
}

const (

	// AutoscaleStackV4ResponseStackTypeWORKLOAD captures enum value "WORKLOAD"
	AutoscaleStackV4ResponseStackTypeWORKLOAD string = "WORKLOAD"

	// AutoscaleStackV4ResponseStackTypeDATALAKE captures enum value "DATALAKE"
	AutoscaleStackV4ResponseStackTypeDATALAKE string = "DATALAKE"

	// AutoscaleStackV4ResponseStackTypeTEMPLATE captures enum value "TEMPLATE"
	AutoscaleStackV4ResponseStackTypeTEMPLATE string = "TEMPLATE"

	// AutoscaleStackV4ResponseStackTypeLEGACY captures enum value "LEGACY"
	AutoscaleStackV4ResponseStackTypeLEGACY string = "LEGACY"
)

// prop value enum
func (m *AutoscaleStackV4Response) validateStackTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackV4ResponseTypeStackTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackV4Response) validateStackType(formats strfmt.Registry) error {

	if swag.IsZero(m.StackType) { // not required
		return nil
	}

	// value enum
	if err := m.validateStackTypeEnum("stackType", "body", m.StackType); err != nil {
		return err
	}

	return nil
}

var autoscaleStackV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS","EXTERNAL_DATABASE_UPGRADE_FINISHED","EXTERNAL_DATABASE_UPGRADE_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED","UPGRADE_CCM_IN_PROGRESS","UPGRADE_CCM_FAILED","UPGRADE_CCM_FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackV4ResponseTypeStatusPropEnum = append(autoscaleStackV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// AutoscaleStackV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	AutoscaleStackV4ResponseStatusREQUESTED string = "REQUESTED"

	// AutoscaleStackV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	AutoscaleStackV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// AutoscaleStackV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	AutoscaleStackV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// AutoscaleStackV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	AutoscaleStackV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// AutoscaleStackV4ResponseStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	AutoscaleStackV4ResponseStatusBACKUPFAILED string = "BACKUP_FAILED"

	// AutoscaleStackV4ResponseStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	AutoscaleStackV4ResponseStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// AutoscaleStackV4ResponseStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	AutoscaleStackV4ResponseStatusRESTOREFAILED string = "RESTORE_FAILED"

	// AutoscaleStackV4ResponseStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	AutoscaleStackV4ResponseStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// AutoscaleStackV4ResponseStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	AutoscaleStackV4ResponseStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// AutoscaleStackV4ResponseStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	AutoscaleStackV4ResponseStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// AutoscaleStackV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	AutoscaleStackV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// AutoscaleStackV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	AutoscaleStackV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// AutoscaleStackV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	AutoscaleStackV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// AutoscaleStackV4ResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	AutoscaleStackV4ResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// AutoscaleStackV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	AutoscaleStackV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// AutoscaleStackV4ResponseStatusSTOPPED captures enum value "STOPPED"
	AutoscaleStackV4ResponseStatusSTOPPED string = "STOPPED"

	// AutoscaleStackV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	AutoscaleStackV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// AutoscaleStackV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	AutoscaleStackV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// AutoscaleStackV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	AutoscaleStackV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// AutoscaleStackV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	AutoscaleStackV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// AutoscaleStackV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	AutoscaleStackV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// AutoscaleStackV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	AutoscaleStackV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// AutoscaleStackV4ResponseStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	AutoscaleStackV4ResponseStatusAMBIGUOUS string = "AMBIGUOUS"

	// AutoscaleStackV4ResponseStatusUNREACHABLE captures enum value "UNREACHABLE"
	AutoscaleStackV4ResponseStatusUNREACHABLE string = "UNREACHABLE"

	// AutoscaleStackV4ResponseStatusNODEFAILURE captures enum value "NODE_FAILURE"
	AutoscaleStackV4ResponseStatusNODEFAILURE string = "NODE_FAILURE"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEINPROGRESS captures enum value "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEINPROGRESS string = "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEFINISHED captures enum value "EXTERNAL_DATABASE_UPGRADE_FINISHED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEFINISHED string = "EXTERNAL_DATABASE_UPGRADE_FINISHED"

	// AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEFAILED captures enum value "EXTERNAL_DATABASE_UPGRADE_FAILED"
	AutoscaleStackV4ResponseStatusEXTERNALDATABASEUPGRADEFAILED string = "EXTERNAL_DATABASE_UPGRADE_FAILED"

	// AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	AutoscaleStackV4ResponseStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"

	// AutoscaleStackV4ResponseStatusUPGRADECCMINPROGRESS captures enum value "UPGRADE_CCM_IN_PROGRESS"
	AutoscaleStackV4ResponseStatusUPGRADECCMINPROGRESS string = "UPGRADE_CCM_IN_PROGRESS"

	// AutoscaleStackV4ResponseStatusUPGRADECCMFAILED captures enum value "UPGRADE_CCM_FAILED"
	AutoscaleStackV4ResponseStatusUPGRADECCMFAILED string = "UPGRADE_CCM_FAILED"

	// AutoscaleStackV4ResponseStatusUPGRADECCMFINISHED captures enum value "UPGRADE_CCM_FINISHED"
	AutoscaleStackV4ResponseStatusUPGRADECCMFINISHED string = "UPGRADE_CCM_FINISHED"
)

// prop value enum
func (m *AutoscaleStackV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var autoscaleStackV4ResponseTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM","CLUSTER_PROXY","CCMV2","CCMV2_JUMPGATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackV4ResponseTypeTunnelPropEnum = append(autoscaleStackV4ResponseTypeTunnelPropEnum, v)
	}
}

const (

	// AutoscaleStackV4ResponseTunnelDIRECT captures enum value "DIRECT"
	AutoscaleStackV4ResponseTunnelDIRECT string = "DIRECT"

	// AutoscaleStackV4ResponseTunnelCCM captures enum value "CCM"
	AutoscaleStackV4ResponseTunnelCCM string = "CCM"

	// AutoscaleStackV4ResponseTunnelCLUSTERPROXY captures enum value "CLUSTER_PROXY"
	AutoscaleStackV4ResponseTunnelCLUSTERPROXY string = "CLUSTER_PROXY"

	// AutoscaleStackV4ResponseTunnelCCMV2 captures enum value "CCMV2"
	AutoscaleStackV4ResponseTunnelCCMV2 string = "CCMV2"

	// AutoscaleStackV4ResponseTunnelCCMV2JUMPGATE captures enum value "CCMV2_JUMPGATE"
	AutoscaleStackV4ResponseTunnelCCMV2JUMPGATE string = "CCMV2_JUMPGATE"
)

// prop value enum
func (m *AutoscaleStackV4Response) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackV4ResponseTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackV4Response) validateTunnel(formats strfmt.Registry) error {

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
func (m *AutoscaleStackV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscaleStackV4Response) UnmarshalBinary(b []byte) error {
	var res AutoscaleStackV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
