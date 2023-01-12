// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterV4Response cluster v4 response
// swagger:model ClusterV4Response
type ClusterV4Response struct {

	// Additional information for ambari cluster
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// blueprint for the cluster
	Blueprint *BlueprintV4Response `json:"blueprint,omitempty"`

	// Indicates the certificate status on the cluster
	// Enum: [VALID HOST_CERT_EXPIRING]
	CertExpirationState string `json:"certExpirationState,omitempty"`

	// filesystem for a specific stack
	CloudStorage *CloudStorageResponse `json:"cloudStorage,omitempty"`

	// cm
	Cm *ClouderaManagerV4Response `json:"cm,omitempty"`

	// CM password for shared usage
	CmMgmtPassword *SecretResponse `json:"cmMgmtPassword,omitempty"`

	// CM username for shared usage
	CmMgmtUser *SecretResponse `json:"cmMgmtUser,omitempty"`

	// Epoch time of cluster creation finish
	CreationFinished int64 `json:"creationFinished,omitempty"`

	// custom configurations crn for the stack
	CustomConfigurationsCrn string `json:"customConfigurationsCrn,omitempty"`

	// custom configurations name for the stack
	CustomConfigurationsName string `json:"customConfigurationsName,omitempty"`

	// custom containers
	CustomContainers *CustomContainerV4Response `json:"customContainers,omitempty"`

	// custom queue for yarn orchestrator
	CustomQueue string `json:"customQueue,omitempty"`

	// Contains valid Crn for a redbeams database server
	DatabaseServerCrn string `json:"databaseServerCrn,omitempty"`

	// Database configurations for the cluster
	Databases []*DatabaseV4Response `json:"databases"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// cluster exposed services for topologies
	ExposedServices map[string][]ClusterExposedServiceV4Response `json:"exposedServices,omitempty"`

	// blueprint, set this or the url field
	ExtendedBlueprintText string `json:"extendedBlueprintText,omitempty"`

	// gateway
	Gateway *GatewayV4Response `json:"gateway,omitempty"`

	// duration - how long the cluster is running in hours
	HoursUp int32 `json:"hoursUp,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// duration - how long the cluster is running in minutes (minus hours)
	MinutesUp int32 `json:"minutesUp,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// proxy CRN for the cluster
	ProxyConfigCrn string `json:"proxyConfigCrn,omitempty"`

	// proxy config name
	ProxyConfigName string `json:"proxyConfigName,omitempty"`

	// Enables Ranger Raz for the cluster on S3 and ADLSv2.
	RangerRazEnabled bool `json:"rangerRazEnabled,omitempty"`

	// FQDN of the gateway node for the stack
	ServerFqdn string `json:"serverFqdn,omitempty"`

	// public ambari ip of the stack
	ServerIP string `json:"serverIp,omitempty"`

	// public ambari url
	ServerURL string `json:"serverUrl,omitempty"`

	// status of the cluster
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS EXTERNAL_DATABASE_UPGRADE_FINISHED EXTERNAL_DATABASE_UPGRADE_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED UPGRADE_CCM_IN_PROGRESS UPGRADE_CCM_FAILED UPGRADE_CCM_FINISHED DETERMINE_DATALAKE_DATA_SIZES_IN_PROGRESS]
	Status string `json:"status,omitempty"`

	// status message of the cluster
	StatusReason string `json:"statusReason,omitempty"`

	// duration - how long the cluster is running in milliseconds
	Uptime int64 `json:"uptime,omitempty"`

	// workspace of the resource
	Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`
}

// Validate validates this cluster v4 response
func (m *ClusterV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertExpirationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCmMgmtPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCmMgmtUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExposedServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterV4Response) validateBlueprint(formats strfmt.Registry) error {

	if swag.IsZero(m.Blueprint) { // not required
		return nil
	}

	if m.Blueprint != nil {
		if err := m.Blueprint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blueprint")
			}
			return err
		}
	}

	return nil
}

var clusterV4ResponseTypeCertExpirationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VALID","HOST_CERT_EXPIRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterV4ResponseTypeCertExpirationStatePropEnum = append(clusterV4ResponseTypeCertExpirationStatePropEnum, v)
	}
}

const (

	// ClusterV4ResponseCertExpirationStateVALID captures enum value "VALID"
	ClusterV4ResponseCertExpirationStateVALID string = "VALID"

	// ClusterV4ResponseCertExpirationStateHOSTCERTEXPIRING captures enum value "HOST_CERT_EXPIRING"
	ClusterV4ResponseCertExpirationStateHOSTCERTEXPIRING string = "HOST_CERT_EXPIRING"
)

// prop value enum
func (m *ClusterV4Response) validateCertExpirationStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterV4ResponseTypeCertExpirationStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterV4Response) validateCertExpirationState(formats strfmt.Registry) error {

	if swag.IsZero(m.CertExpirationState) { // not required
		return nil
	}

	// value enum
	if err := m.validateCertExpirationStateEnum("certExpirationState", "body", m.CertExpirationState); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV4Response) validateCloudStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudStorage) { // not required
		return nil
	}

	if m.CloudStorage != nil {
		if err := m.CloudStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudStorage")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Response) validateCm(formats strfmt.Registry) error {

	if swag.IsZero(m.Cm) { // not required
		return nil
	}

	if m.Cm != nil {
		if err := m.Cm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cm")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Response) validateCmMgmtPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.CmMgmtPassword) { // not required
		return nil
	}

	if m.CmMgmtPassword != nil {
		if err := m.CmMgmtPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cmMgmtPassword")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Response) validateCmMgmtUser(formats strfmt.Registry) error {

	if swag.IsZero(m.CmMgmtUser) { // not required
		return nil
	}

	if m.CmMgmtUser != nil {
		if err := m.CmMgmtUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cmMgmtUser")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Response) validateCustomContainers(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomContainers) { // not required
		return nil
	}

	if m.CustomContainers != nil {
		if err := m.CustomContainers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customContainers")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterV4Response) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	for i := 0; i < len(m.Databases); i++ {
		if swag.IsZero(m.Databases[i]) { // not required
			continue
		}

		if m.Databases[i] != nil {
			if err := m.Databases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("databases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterV4Response) validateExposedServices(formats strfmt.Registry) error {

	if swag.IsZero(m.ExposedServices) { // not required
		return nil
	}

	for k := range m.ExposedServices {

		if err := validate.Required("exposedServices"+"."+k, "body", m.ExposedServices[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ExposedServices[k]); i++ {

			if err := m.ExposedServices[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exposedServices" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *ClusterV4Response) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

var clusterV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS","EXTERNAL_DATABASE_UPGRADE_FINISHED","EXTERNAL_DATABASE_UPGRADE_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED","UPGRADE_CCM_IN_PROGRESS","UPGRADE_CCM_FAILED","UPGRADE_CCM_FINISHED","DETERMINE_DATALAKE_DATA_SIZES_IN_PROGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterV4ResponseTypeStatusPropEnum = append(clusterV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	ClusterV4ResponseStatusREQUESTED string = "REQUESTED"

	// ClusterV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	ClusterV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// ClusterV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	ClusterV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// ClusterV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	ClusterV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// ClusterV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	ClusterV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// ClusterV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	ClusterV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// ClusterV4ResponseStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	ClusterV4ResponseStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// ClusterV4ResponseStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	ClusterV4ResponseStatusBACKUPFAILED string = "BACKUP_FAILED"

	// ClusterV4ResponseStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	ClusterV4ResponseStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// ClusterV4ResponseStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	ClusterV4ResponseStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// ClusterV4ResponseStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	ClusterV4ResponseStatusRESTOREFAILED string = "RESTORE_FAILED"

	// ClusterV4ResponseStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	ClusterV4ResponseStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// ClusterV4ResponseStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	ClusterV4ResponseStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// ClusterV4ResponseStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	ClusterV4ResponseStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// ClusterV4ResponseStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	ClusterV4ResponseStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// ClusterV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	ClusterV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// ClusterV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	ClusterV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// ClusterV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	ClusterV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// ClusterV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	ClusterV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// ClusterV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	ClusterV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// ClusterV4ResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	ClusterV4ResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// ClusterV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	ClusterV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// ClusterV4ResponseStatusSTOPPED captures enum value "STOPPED"
	ClusterV4ResponseStatusSTOPPED string = "STOPPED"

	// ClusterV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	ClusterV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// ClusterV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	ClusterV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// ClusterV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	ClusterV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// ClusterV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	ClusterV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// ClusterV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	ClusterV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// ClusterV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	ClusterV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// ClusterV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	ClusterV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// ClusterV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	ClusterV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// ClusterV4ResponseStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	ClusterV4ResponseStatusAMBIGUOUS string = "AMBIGUOUS"

	// ClusterV4ResponseStatusUNREACHABLE captures enum value "UNREACHABLE"
	ClusterV4ResponseStatusUNREACHABLE string = "UNREACHABLE"

	// ClusterV4ResponseStatusNODEFAILURE captures enum value "NODE_FAILURE"
	ClusterV4ResponseStatusNODEFAILURE string = "NODE_FAILURE"

	// ClusterV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	ClusterV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// ClusterV4ResponseStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	ClusterV4ResponseStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// ClusterV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	ClusterV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// ClusterV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	ClusterV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// ClusterV4ResponseStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	ClusterV4ResponseStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// ClusterV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	ClusterV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// ClusterV4ResponseStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	ClusterV4ResponseStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// ClusterV4ResponseStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	ClusterV4ResponseStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// ClusterV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	ClusterV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// ClusterV4ResponseStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	ClusterV4ResponseStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// ClusterV4ResponseStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	ClusterV4ResponseStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEINPROGRESS captures enum value "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"
	ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEINPROGRESS string = "EXTERNAL_DATABASE_UPGRADE_IN_PROGRESS"

	// ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEFINISHED captures enum value "EXTERNAL_DATABASE_UPGRADE_FINISHED"
	ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEFINISHED string = "EXTERNAL_DATABASE_UPGRADE_FINISHED"

	// ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEFAILED captures enum value "EXTERNAL_DATABASE_UPGRADE_FAILED"
	ClusterV4ResponseStatusEXTERNALDATABASEUPGRADEFAILED string = "EXTERNAL_DATABASE_UPGRADE_FAILED"

	// ClusterV4ResponseStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	ClusterV4ResponseStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// ClusterV4ResponseStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	ClusterV4ResponseStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// ClusterV4ResponseStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	ClusterV4ResponseStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"

	// ClusterV4ResponseStatusUPGRADECCMINPROGRESS captures enum value "UPGRADE_CCM_IN_PROGRESS"
	ClusterV4ResponseStatusUPGRADECCMINPROGRESS string = "UPGRADE_CCM_IN_PROGRESS"

	// ClusterV4ResponseStatusUPGRADECCMFAILED captures enum value "UPGRADE_CCM_FAILED"
	ClusterV4ResponseStatusUPGRADECCMFAILED string = "UPGRADE_CCM_FAILED"

	// ClusterV4ResponseStatusUPGRADECCMFINISHED captures enum value "UPGRADE_CCM_FINISHED"
	ClusterV4ResponseStatusUPGRADECCMFINISHED string = "UPGRADE_CCM_FINISHED"

	// ClusterV4ResponseStatusDETERMINEDATALAKEDATASIZESINPROGRESS captures enum value "DETERMINE_DATALAKE_DATA_SIZES_IN_PROGRESS"
	ClusterV4ResponseStatusDETERMINEDATALAKEDATASIZESINPROGRESS string = "DETERMINE_DATALAKE_DATA_SIZES_IN_PROGRESS"
)

// prop value enum
func (m *ClusterV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ClusterV4Response) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
