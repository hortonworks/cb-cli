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

// StackV4Response stack v4 response
// swagger:model StackV4Response
type StackV4Response struct {

	// stack related authentication
	Authentication *StackAuthenticationV4Response `json:"authentication,omitempty"`

	// aws specific parameters for stack
	Aws AwsStackV4Parameters `json:"aws,omitempty"`

	// azure specific parameters for stack
	Azure *AzureStackV4Parameters `json:"azure,omitempty"`

	// Cloudplatform of the stack
	// Enum: [AWS GCP AZURE YARN MOCK OPENSTACK]
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// details of the Cloudbreak that provisioned the stack
	CloudbreakDetails *CloudbreakDetailsV4Response `json:"cloudbreakDetails,omitempty"`

	// related events for a cloudbreak stack
	CloudbreakEvents []*CloudbreakEventV4Response `json:"cloudbreakEvents"`

	// cluster object on stack
	Cluster *ClusterV4Response `json:"cluster,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// credential crn
	CredentialCrn string `json:"credentialCrn,omitempty"`

	// credential name
	CredentialName string `json:"credentialName,omitempty"`

	// crn
	Crn string `json:"crn,omitempty"`

	// custom domains
	CustomDomains *CustomDomainSettingsV4Response `json:"customDomains,omitempty"`

	// environment crn
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// External database parameters for the stack.
	ExternalDatabase *DatabaseResponse `json:"externalDatabase,omitempty"`

	// Flow identifier for the current stack creation. Only returned during the stack create request/response.
	FlowIdentifier *FlowIdentifier `json:"flowIdentifier,omitempty"`

	// port of the gateway secured proxy
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// gcp specific parameters for stack
	Gcp GcpStackV4Parameters `json:"gcp,omitempty"`

	// hardware information where pairing hostmetadata with instancemetadata
	// Unique: true
	HardwareInfoGroups []*HardwareInfoGroupV4Response `json:"hardwareInfoGroups"`

	// id
	ID int64 `json:"id,omitempty"`

	// image of the stack
	Image *StackImageV4Response `json:"image,omitempty"`

	// instance groups
	InstanceGroups []*InstanceGroupV4Response `json:"instanceGroups"`

	// Any load balancers routing traffic both to and inside of the cluster.
	LoadBalancers []*LoadBalancerResponse `json:"loadBalancers"`

	// mock
	Mock MockStackV4Parameters `json:"mock,omitempty"`

	// name of the stack
	// Required: true
	Name *string `json:"name"`

	// stack related network
	Network *NetworkV4Response `json:"network,omitempty"`

	// node count of the stack
	NodeCount int32 `json:"nodeCount,omitempty"`

	// placement configuration parameters for a cluster (e.g. 'region', 'availabilityZone')
	// Required: true
	Placement *PlacementSettingsV4Response `json:"placement"`

	// Shared service request
	SharedService *SharedServiceV4Response `json:"sharedService,omitempty"`

	// status of the stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED BACKUP_IN_PROGRESS BACKUP_FAILED BACKUP_FINISHED RESTORE_IN_PROGRESS RESTORE_FAILED RESTORE_FINISHED RECOVERY_IN_PROGRESS RECOVERY_REQUESTED RECOVERY_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETED_ON_PROVIDER_SIDE DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED AMBIGUOUS UNREACHABLE NODE_FAILURE EXTERNAL_DATABASE_CREATION_IN_PROGRESS EXTERNAL_DATABASE_CREATION_FAILED EXTERNAL_DATABASE_DELETION_IN_PROGRESS EXTERNAL_DATABASE_DELETION_FINISHED EXTERNAL_DATABASE_DELETION_FAILED EXTERNAL_DATABASE_START_IN_PROGRESS EXTERNAL_DATABASE_START_FINISHED EXTERNAL_DATABASE_START_FAILED EXTERNAL_DATABASE_STOP_IN_PROGRESS EXTERNAL_DATABASE_STOP_FINISHED EXTERNAL_DATABASE_STOP_FAILED LOAD_BALANCER_UPDATE_IN_PROGRESS LOAD_BALANCER_UPDATE_FINISHED LOAD_BALANCER_UPDATE_FAILED]
	Status string `json:"status,omitempty"`

	// status message of the stack
	StatusReason string `json:"statusReason,omitempty"`

	// stack related tags
	Tags *TagsV4Response `json:"tags,omitempty"`

	// stack related telemetry settings
	Telemetry *TelemetryResponse `json:"telemetry,omitempty"`

	// termination completion time of stack in long
	Terminated int64 `json:"terminated,omitempty"`

	// time to live
	TimeToLive int64 `json:"timeToLive,omitempty"`

	// Configuration that the connection going directly or with cluster proxy or with ccm and cluster proxy.
	// Enum: [DIRECT CCM CLUSTER_PROXY CCMV2 CCMV2_JUMPGATE]
	Tunnel string `json:"tunnel,omitempty"`

	// Variant of the stack
	Variant string `json:"variant,omitempty"`

	// workspace of the resource
	Workspace *WorkspaceResourceV4Response `json:"workspace,omitempty"`

	// yarn specific parameters for stack
	Yarn *YarnStackV4Parameters `json:"yarn,omitempty"`
}

// Validate validates this stack v4 response
func (m *StackV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudbreakDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudbreakEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareInfoGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackV4Response) validateAuthentication(formats strfmt.Registry) error {

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

func (m *StackV4Response) validateAzure(formats strfmt.Registry) error {

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

var stackV4ResponseTypeCloudPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP","AZURE","YARN","MOCK","OPENSTACK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackV4ResponseTypeCloudPlatformPropEnum = append(stackV4ResponseTypeCloudPlatformPropEnum, v)
	}
}

const (

	// StackV4ResponseCloudPlatformAWS captures enum value "AWS"
	StackV4ResponseCloudPlatformAWS string = "AWS"

	// StackV4ResponseCloudPlatformGCP captures enum value "GCP"
	StackV4ResponseCloudPlatformGCP string = "GCP"

	// StackV4ResponseCloudPlatformAZURE captures enum value "AZURE"
	StackV4ResponseCloudPlatformAZURE string = "AZURE"

	// StackV4ResponseCloudPlatformYARN captures enum value "YARN"
	StackV4ResponseCloudPlatformYARN string = "YARN"

	// StackV4ResponseCloudPlatformMOCK captures enum value "MOCK"
	StackV4ResponseCloudPlatformMOCK string = "MOCK"

	// StackV4ResponseCloudPlatformOPENSTACK captures enum value "OPENSTACK"
	StackV4ResponseCloudPlatformOPENSTACK string = "OPENSTACK"
)

// prop value enum
func (m *StackV4Response) validateCloudPlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackV4ResponseTypeCloudPlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackV4Response) validateCloudPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudPlatformEnum("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Response) validateCloudbreakDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudbreakDetails) { // not required
		return nil
	}

	if m.CloudbreakDetails != nil {
		if err := m.CloudbreakDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudbreakDetails")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateCloudbreakEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudbreakEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudbreakEvents); i++ {
		if swag.IsZero(m.CloudbreakEvents[i]) { // not required
			continue
		}

		if m.CloudbreakEvents[i] != nil {
			if err := m.CloudbreakEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudbreakEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV4Response) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateCustomDomains(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomDomains) { // not required
		return nil
	}

	if m.CustomDomains != nil {
		if err := m.CustomDomains.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDomains")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateExternalDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDatabase) { // not required
		return nil
	}

	if m.ExternalDatabase != nil {
		if err := m.ExternalDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDatabase")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateFlowIdentifier(formats strfmt.Registry) error {

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

func (m *StackV4Response) validateHardwareInfoGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HardwareInfoGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hardwareInfoGroups", "body", m.HardwareInfoGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HardwareInfoGroups); i++ {
		if swag.IsZero(m.HardwareInfoGroups[i]) { // not required
			continue
		}

		if m.HardwareInfoGroups[i] != nil {
			if err := m.HardwareInfoGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hardwareInfoGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV4Response) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateInstanceGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV4Response) validateLoadBalancers(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalancers) { // not required
		return nil
	}

	for i := 0; i < len(m.LoadBalancers); i++ {
		if swag.IsZero(m.LoadBalancers[i]) { // not required
			continue
		}

		if m.LoadBalancers[i] != nil {
			if err := m.LoadBalancers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("loadBalancers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Response) validateNetwork(formats strfmt.Registry) error {

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

func (m *StackV4Response) validatePlacement(formats strfmt.Registry) error {

	if err := validate.Required("placement", "body", m.Placement); err != nil {
		return err
	}

	if m.Placement != nil {
		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

func (m *StackV4Response) validateSharedService(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedService) { // not required
		return nil
	}

	if m.SharedService != nil {
		if err := m.SharedService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedService")
			}
			return err
		}
	}

	return nil
}

var stackV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackV4ResponseTypeStatusPropEnum = append(stackV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// StackV4ResponseStatusREQUESTED captures enum value "REQUESTED"
	StackV4ResponseStatusREQUESTED string = "REQUESTED"

	// StackV4ResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackV4ResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// StackV4ResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackV4ResponseStatusAVAILABLE string = "AVAILABLE"

	// StackV4ResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackV4ResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// StackV4ResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackV4ResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// StackV4ResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackV4ResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// StackV4ResponseStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	StackV4ResponseStatusBACKUPINPROGRESS string = "BACKUP_IN_PROGRESS"

	// StackV4ResponseStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	StackV4ResponseStatusBACKUPFAILED string = "BACKUP_FAILED"

	// StackV4ResponseStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	StackV4ResponseStatusBACKUPFINISHED string = "BACKUP_FINISHED"

	// StackV4ResponseStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	StackV4ResponseStatusRESTOREINPROGRESS string = "RESTORE_IN_PROGRESS"

	// StackV4ResponseStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	StackV4ResponseStatusRESTOREFAILED string = "RESTORE_FAILED"

	// StackV4ResponseStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	StackV4ResponseStatusRESTOREFINISHED string = "RESTORE_FINISHED"

	// StackV4ResponseStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	StackV4ResponseStatusRECOVERYINPROGRESS string = "RECOVERY_IN_PROGRESS"

	// StackV4ResponseStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	StackV4ResponseStatusRECOVERYREQUESTED string = "RECOVERY_REQUESTED"

	// StackV4ResponseStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	StackV4ResponseStatusRECOVERYFAILED string = "RECOVERY_FAILED"

	// StackV4ResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackV4ResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// StackV4ResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackV4ResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// StackV4ResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackV4ResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// StackV4ResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackV4ResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// StackV4ResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackV4ResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// StackV4ResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	StackV4ResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// StackV4ResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackV4ResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// StackV4ResponseStatusSTOPPED captures enum value "STOPPED"
	StackV4ResponseStatusSTOPPED string = "STOPPED"

	// StackV4ResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackV4ResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// StackV4ResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackV4ResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// StackV4ResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackV4ResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// StackV4ResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackV4ResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// StackV4ResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackV4ResponseStatusSTARTFAILED string = "START_FAILED"

	// StackV4ResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackV4ResponseStatusSTOPFAILED string = "STOP_FAILED"

	// StackV4ResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackV4ResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// StackV4ResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackV4ResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// StackV4ResponseStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	StackV4ResponseStatusAMBIGUOUS string = "AMBIGUOUS"

	// StackV4ResponseStatusUNREACHABLE captures enum value "UNREACHABLE"
	StackV4ResponseStatusUNREACHABLE string = "UNREACHABLE"

	// StackV4ResponseStatusNODEFAILURE captures enum value "NODE_FAILURE"
	StackV4ResponseStatusNODEFAILURE string = "NODE_FAILURE"

	// StackV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	StackV4ResponseStatusEXTERNALDATABASECREATIONINPROGRESS string = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// StackV4ResponseStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	StackV4ResponseStatusEXTERNALDATABASECREATIONFAILED string = "EXTERNAL_DATABASE_CREATION_FAILED"

	// StackV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	StackV4ResponseStatusEXTERNALDATABASEDELETIONINPROGRESS string = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// StackV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	StackV4ResponseStatusEXTERNALDATABASEDELETIONFINISHED string = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// StackV4ResponseStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	StackV4ResponseStatusEXTERNALDATABASEDELETIONFAILED string = "EXTERNAL_DATABASE_DELETION_FAILED"

	// StackV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	StackV4ResponseStatusEXTERNALDATABASESTARTINPROGRESS string = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// StackV4ResponseStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	StackV4ResponseStatusEXTERNALDATABASESTARTFINISHED string = "EXTERNAL_DATABASE_START_FINISHED"

	// StackV4ResponseStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	StackV4ResponseStatusEXTERNALDATABASESTARTFAILED string = "EXTERNAL_DATABASE_START_FAILED"

	// StackV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	StackV4ResponseStatusEXTERNALDATABASESTOPINPROGRESS string = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// StackV4ResponseStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	StackV4ResponseStatusEXTERNALDATABASESTOPFINISHED string = "EXTERNAL_DATABASE_STOP_FINISHED"

	// StackV4ResponseStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	StackV4ResponseStatusEXTERNALDATABASESTOPFAILED string = "EXTERNAL_DATABASE_STOP_FAILED"

	// StackV4ResponseStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	StackV4ResponseStatusLOADBALANCERUPDATEINPROGRESS string = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// StackV4ResponseStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	StackV4ResponseStatusLOADBALANCERUPDATEFINISHED string = "LOAD_BALANCER_UPDATE_FINISHED"

	// StackV4ResponseStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	StackV4ResponseStatusLOADBALANCERUPDATEFAILED string = "LOAD_BALANCER_UPDATE_FAILED"
)

// prop value enum
func (m *StackV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Response) validateTags(formats strfmt.Registry) error {

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

func (m *StackV4Response) validateTelemetry(formats strfmt.Registry) error {

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

var stackV4ResponseTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM","CLUSTER_PROXY","CCMV2","CCMV2_JUMPGATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackV4ResponseTypeTunnelPropEnum = append(stackV4ResponseTypeTunnelPropEnum, v)
	}
}

const (

	// StackV4ResponseTunnelDIRECT captures enum value "DIRECT"
	StackV4ResponseTunnelDIRECT string = "DIRECT"

	// StackV4ResponseTunnelCCM captures enum value "CCM"
	StackV4ResponseTunnelCCM string = "CCM"

	// StackV4ResponseTunnelCLUSTERPROXY captures enum value "CLUSTER_PROXY"
	StackV4ResponseTunnelCLUSTERPROXY string = "CLUSTER_PROXY"

	// StackV4ResponseTunnelCCMV2 captures enum value "CCMV2"
	StackV4ResponseTunnelCCMV2 string = "CCMV2"

	// StackV4ResponseTunnelCCMV2JUMPGATE captures enum value "CCMV2_JUMPGATE"
	StackV4ResponseTunnelCCMV2JUMPGATE string = "CCMV2_JUMPGATE"
)

// prop value enum
func (m *StackV4Response) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackV4ResponseTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackV4Response) validateTunnel(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnel) { // not required
		return nil
	}

	// value enum
	if err := m.validateTunnelEnum("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

func (m *StackV4Response) validateWorkspace(formats strfmt.Registry) error {

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

func (m *StackV4Response) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackV4Response) UnmarshalBinary(b []byte) error {
	var res StackV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
