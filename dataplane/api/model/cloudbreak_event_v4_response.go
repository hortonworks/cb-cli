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

// CloudbreakEventV4Response cloudbreak event v4 response
// swagger:model CloudbreakEventV4Response
type CloudbreakEventV4Response struct {

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// blueprint id for the cluster
	BlueprintID int64 `json:"blueprintId,omitempty"`

	// gathered from blueprintName field from the blueprint
	BlueprintName string `json:"blueprintName,omitempty"`

	// type of cloud provider
	Cloud string `json:"cloud,omitempty"`

	// id of the cluster
	ClusterID int64 `json:"clusterId,omitempty"`

	// name of the cluster
	ClusterName string `json:"clusterName,omitempty"`

	// status of the cluster
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS EXTERNAL_DATABASE_UPGRADE_FINISHED EXTERNAL_DATABASE_UPGRADE_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED UPGRADE_CCM_IN_PROGRESS UPGRADE_CCM_FAILED UPGRADE_CCM_FINISHED]
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// message of the event
	EventMessage string `json:"eventMessage,omitempty"`

	// timestamp of the event
	EventTimestamp int64 `json:"eventTimestamp,omitempty"`

	// type of the event
	EventType string `json:"eventType,omitempty"`

	// name of the instance group
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// ldap details
	LdapDetails *LdapDetails `json:"ldapDetails,omitempty"`

	// number of nodes
	NodeCount int32 `json:"nodeCount,omitempty"`

	// Type of the notification to be identifiable by the UI
	NotificationType string `json:"notificationType,omitempty"`

	// rds details
	RdsDetails *RdsDetails `json:"rdsDetails,omitempty"`

	// region of the stack
	Region string `json:"region,omitempty"`

	// the unique crn of the resource
	StackCrn string `json:"stackCrn,omitempty"`

	// name of the stack
	StackName string `json:"stackName,omitempty"`

	// status of the stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS EXTERNAL_DATABASE_UPGRADE_FINISHED EXTERNAL_DATABASE_UPGRADE_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED UPGRADE_CCM_IN_PROGRESS UPGRADE_CCM_FAILED UPGRADE_CCM_FINISHED]
	StackStatus string `json:"stackStatus,omitempty"`

	// name of the current tenant
	TenantName string `json:"tenantName,omitempty"`

	// User ID in the new authorization model
	UserID string `json:"userId,omitempty"`

	// Workspace ID of the resource
	WorkspaceID int64 `json:"workspaceId,omitempty"`
}

// Validate validates this cloudbreak event v4 response
func (m *CloudbreakEventV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdapDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdsDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloudbreakEventV4ResponseTypeClusterStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS","EXTERNAL_DATABASE_UPGRADE_FINISHED","EXTERNAL_DATABASE_UPGRADE_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED","UPGRADE_CCM_IN_PROGRESS","UPGRADE_CCM_FAILED","UPGRADE_CCM_FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudbreakEventV4ResponseTypeClusterStatusPropEnum = append(cloudbreakEventV4ResponseTypeClusterStatusPropEnum, v)
	}
}

const (

	// CloudbreakEventV4ResponseClusterStatusREQUESTED captures enum value "REQUESTED"
	CloudbreakEventV4ResponseClusterStatusREQUESTED string = "REQUESTED"

	// CloudbreakEventV4ResponseClusterStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusAVAILABLE captures enum value "AVAILABLE"
	CloudbreakEventV4ResponseClusterStatusAVAILABLE string = "AVAILABLE"

	// CloudbreakEventV4ResponseClusterStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	CloudbreakEventV4ResponseClusterStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// CloudbreakEventV4ResponseClusterStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	CloudbreakEventV4ResponseClusterStatusUPDATEFAILED string = "UPDATE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	CloudbreakEventV4ResponseClusterStatusBACKUPFAILED string = "BACKUP_FAILED"

	// CloudbreakEventV4ResponseClusterStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	CloudbreakEventV4ResponseClusterStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	CloudbreakEventV4ResponseClusterStatusRESTOREFAILED string = "RESTORE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	CloudbreakEventV4ResponseClusterStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	CloudbreakEventV4ResponseClusterStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// CloudbreakEventV4ResponseClusterStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	CloudbreakEventV4ResponseClusterStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// CloudbreakEventV4ResponseClusterStatusCREATEFAILED captures enum value "CREATE_FAILED"
	CloudbreakEventV4ResponseClusterStatusCREATEFAILED string = "CREATE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	CloudbreakEventV4ResponseClusterStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// CloudbreakEventV4ResponseClusterStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusDELETEFAILED captures enum value "DELETE_FAILED"
	CloudbreakEventV4ResponseClusterStatusDELETEFAILED string = "DELETE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	CloudbreakEventV4ResponseClusterStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// CloudbreakEventV4ResponseClusterStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	CloudbreakEventV4ResponseClusterStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// CloudbreakEventV4ResponseClusterStatusSTOPPED captures enum value "STOPPED"
	CloudbreakEventV4ResponseClusterStatusSTOPPED string = "STOPPED"

	// CloudbreakEventV4ResponseClusterStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	CloudbreakEventV4ResponseClusterStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// CloudbreakEventV4ResponseClusterStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	CloudbreakEventV4ResponseClusterStatusSTARTREQUESTED string = "START_REQUESTED"

	// CloudbreakEventV4ResponseClusterStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusSTARTFAILED captures enum value "START_FAILED"
	CloudbreakEventV4ResponseClusterStatusSTARTFAILED string = "START_FAILED"

	// CloudbreakEventV4ResponseClusterStatusSTOPFAILED captures enum value "STOP_FAILED"
	CloudbreakEventV4ResponseClusterStatusSTOPFAILED string = "STOP_FAILED"

	// CloudbreakEventV4ResponseClusterStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	CloudbreakEventV4ResponseClusterStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// CloudbreakEventV4ResponseClusterStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	CloudbreakEventV4ResponseClusterStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// CloudbreakEventV4ResponseClusterStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	CloudbreakEventV4ResponseClusterStatusAMBIGUOUS string = "AMBIGUOUS"

	// CloudbreakEventV4ResponseClusterStatusUNREACHABLE captures enum value "UNREACHABLE"
	CloudbreakEventV4ResponseClusterStatusUNREACHABLE string = "UNREACHABLE"

	// CloudbreakEventV4ResponseClusterStatusNODEFAILURE captures enum value "NODE_FAILURE"
	CloudbreakEventV4ResponseClusterStatusNODEFAILURE string = "NODE_FAILURE"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEINPROGRESS captures enum value "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEINPROGRESS string = "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFINISHED captures enum value "EXTERNAL_DATABASE_UPGRADE_FINISHED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFINISHED string = "EXTERNAL_DATABASE_UPGRADE_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFAILED captures enum value "EXTERNAL_DATABASE_UPGRADE_FAILED"
	CloudbreakEventV4ResponseClusterStatusEXTERNALDATABASEUPGRADEFAILED string = "EXTERNAL_DATABASE_UPGRADE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	CloudbreakEventV4ResponseClusterStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"

	// CloudbreakEventV4ResponseClusterStatusUPGRADECCMINPROGRESS captures enum value "UPGRADE_CCM_IN_PROGRESS"
	CloudbreakEventV4ResponseClusterStatusUPGRADECCMINPROGRESS string = "UPGRADE_CCM_IN_PROGRESS"

	// CloudbreakEventV4ResponseClusterStatusUPGRADECCMFAILED captures enum value "UPGRADE_CCM_FAILED"
	CloudbreakEventV4ResponseClusterStatusUPGRADECCMFAILED string = "UPGRADE_CCM_FAILED"

	// CloudbreakEventV4ResponseClusterStatusUPGRADECCMFINISHED captures enum value "UPGRADE_CCM_FINISHED"
	CloudbreakEventV4ResponseClusterStatusUPGRADECCMFINISHED string = "UPGRADE_CCM_FINISHED"
)

// prop value enum
func (m *CloudbreakEventV4Response) validateClusterStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudbreakEventV4ResponseTypeClusterStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudbreakEventV4Response) validateClusterStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterStatusEnum("clusterStatus", "body", m.ClusterStatus); err != nil {
		return err
	}

	return nil
}

func (m *CloudbreakEventV4Response) validateLdapDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapDetails) { // not required
		return nil
	}

	if m.LdapDetails != nil {
		if err := m.LdapDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapDetails")
			}
			return err
		}
	}

	return nil
}

func (m *CloudbreakEventV4Response) validateRdsDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsDetails) { // not required
		return nil
	}

	if m.RdsDetails != nil {
		if err := m.RdsDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rdsDetails")
			}
			return err
		}
	}

	return nil
}

var cloudbreakEventV4ResponseTypeStackStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS","EXTERNAL_DATABASE_UPGRADE_FINISHED","EXTERNAL_DATABASE_UPGRADE_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED","UPGRADE_CCM_IN_PROGRESS","UPGRADE_CCM_FAILED","UPGRADE_CCM_FINISHED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudbreakEventV4ResponseTypeStackStatusPropEnum = append(cloudbreakEventV4ResponseTypeStackStatusPropEnum, v)
	}
}

const (

	// CloudbreakEventV4ResponseStackStatusREQUESTED captures enum value "REQUESTED"
	CloudbreakEventV4ResponseStackStatusREQUESTED string = "REQUESTED"

	// CloudbreakEventV4ResponseStackStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusAVAILABLE captures enum value "AVAILABLE"
	CloudbreakEventV4ResponseStackStatusAVAILABLE string = "AVAILABLE"

	// CloudbreakEventV4ResponseStackStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	CloudbreakEventV4ResponseStackStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// CloudbreakEventV4ResponseStackStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	CloudbreakEventV4ResponseStackStatusUPDATEFAILED string = "UPDATE_FAILED"

	// CloudbreakEventV4ResponseStackStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	CloudbreakEventV4ResponseStackStatusBACKUPFAILED string = "BACKUP_FAILED"

	// CloudbreakEventV4ResponseStackStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	CloudbreakEventV4ResponseStackStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// CloudbreakEventV4ResponseStackStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	CloudbreakEventV4ResponseStackStatusRESTOREFAILED string = "RESTORE_FAILED"

	// CloudbreakEventV4ResponseStackStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	CloudbreakEventV4ResponseStackStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// CloudbreakEventV4ResponseStackStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	CloudbreakEventV4ResponseStackStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// CloudbreakEventV4ResponseStackStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	CloudbreakEventV4ResponseStackStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// CloudbreakEventV4ResponseStackStatusCREATEFAILED captures enum value "CREATE_FAILED"
	CloudbreakEventV4ResponseStackStatusCREATEFAILED string = "CREATE_FAILED"

	// CloudbreakEventV4ResponseStackStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	CloudbreakEventV4ResponseStackStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// CloudbreakEventV4ResponseStackStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusDELETEFAILED captures enum value "DELETE_FAILED"
	CloudbreakEventV4ResponseStackStatusDELETEFAILED string = "DELETE_FAILED"

	// CloudbreakEventV4ResponseStackStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	CloudbreakEventV4ResponseStackStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// CloudbreakEventV4ResponseStackStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	CloudbreakEventV4ResponseStackStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// CloudbreakEventV4ResponseStackStatusSTOPPED captures enum value "STOPPED"
	CloudbreakEventV4ResponseStackStatusSTOPPED string = "STOPPED"

	// CloudbreakEventV4ResponseStackStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	CloudbreakEventV4ResponseStackStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// CloudbreakEventV4ResponseStackStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	CloudbreakEventV4ResponseStackStatusSTARTREQUESTED string = "START_REQUESTED"

	// CloudbreakEventV4ResponseStackStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusSTARTFAILED captures enum value "START_FAILED"
	CloudbreakEventV4ResponseStackStatusSTARTFAILED string = "START_FAILED"

	// CloudbreakEventV4ResponseStackStatusSTOPFAILED captures enum value "STOP_FAILED"
	CloudbreakEventV4ResponseStackStatusSTOPFAILED string = "STOP_FAILED"

	// CloudbreakEventV4ResponseStackStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	CloudbreakEventV4ResponseStackStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// CloudbreakEventV4ResponseStackStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	CloudbreakEventV4ResponseStackStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// CloudbreakEventV4ResponseStackStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	CloudbreakEventV4ResponseStackStatusAMBIGUOUS string = "AMBIGUOUS"

	// CloudbreakEventV4ResponseStackStatusUNREACHABLE captures enum value "UNREACHABLE"
	CloudbreakEventV4ResponseStackStatusUNREACHABLE string = "UNREACHABLE"

	// CloudbreakEventV4ResponseStackStatusNODEFAILURE captures enum value "NODE_FAILURE"
	CloudbreakEventV4ResponseStackStatusNODEFAILURE string = "NODE_FAILURE"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEINPROGRESS captures enum value "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEINPROGRESS string = "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEFINISHED captures enum value "EXTERNAL_DATABASE_UPGRADE_FINISHED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEFINISHED string = "EXTERNAL_DATABASE_UPGRADE_FINISHED"

	// CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEFAILED captures enum value "EXTERNAL_DATABASE_UPGRADE_FAILED"
	CloudbreakEventV4ResponseStackStatusEXTERNALDATABASEUPGRADEFAILED string = "EXTERNAL_DATABASE_UPGRADE_FAILED"

	// CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	CloudbreakEventV4ResponseStackStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"

	// CloudbreakEventV4ResponseStackStatusUPGRADECCMINPROGRESS captures enum value "UPGRADE_CCM_IN_PROGRESS"
	CloudbreakEventV4ResponseStackStatusUPGRADECCMINPROGRESS string = "UPGRADE_CCM_IN_PROGRESS"

	// CloudbreakEventV4ResponseStackStatusUPGRADECCMFAILED captures enum value "UPGRADE_CCM_FAILED"
	CloudbreakEventV4ResponseStackStatusUPGRADECCMFAILED string = "UPGRADE_CCM_FAILED"

	// CloudbreakEventV4ResponseStackStatusUPGRADECCMFINISHED captures enum value "UPGRADE_CCM_FINISHED"
	CloudbreakEventV4ResponseStackStatusUPGRADECCMFINISHED string = "UPGRADE_CCM_FINISHED"
)

// prop value enum
func (m *CloudbreakEventV4Response) validateStackStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudbreakEventV4ResponseTypeStackStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudbreakEventV4Response) validateStackStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.StackStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateStackStatusEnum("stackStatus", "body", m.StackStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudbreakEventV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudbreakEventV4Response) UnmarshalBinary(b []byte) error {
	var res CloudbreakEventV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
