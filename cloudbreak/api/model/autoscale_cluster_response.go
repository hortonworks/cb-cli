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

// AutoscaleClusterResponse autoscale cluster response
// swagger:model AutoscaleClusterResponse

type AutoscaleClusterResponse struct {

	// [DEPRECATED] use RdsConfig instead! details of the external Ambari database
	AmbariDatabaseDetails *AmbariDatabaseDetails `json:"ambariDatabaseDetails,omitempty"`

	// details of the Ambari package repository
	AmbariRepoDetailsJSON *AmbariRepoDetails `json:"ambariRepoDetailsJson,omitempty"`

	// public ambari ip of the stack
	AmbariServerIP string `json:"ambariServerIp,omitempty"`

	// public ambari url
	AmbariServerURL string `json:"ambariServerUrl,omitempty"`

	// details of the Ambari stack
	AmbariStackDetails *AmbariStackDetailsResponse `json:"ambariStackDetails,omitempty"`

	// Additional information for ambari cluster
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// blueprint for the cluster
	Blueprint *BlueprintResponse `json:"blueprint,omitempty"`

	// blueprint custom properties
	BlueprintCustomProperties string `json:"blueprintCustomProperties,omitempty"`

	// blueprint id for the cluster
	BlueprintID int64 `json:"blueprintId,omitempty"`

	// blueprint inputs in the cluster
	// Unique: true
	BlueprintInputs []*BlueprintInput `json:"blueprintInputs"`

	// name of the cluster
	Cluster string `json:"cluster,omitempty"`

	// cluster exposed services for topologies
	ClusterExposedServicesForTopologies map[string][]ClusterExposedServiceResponse `json:"clusterExposedServicesForTopologies,omitempty"`

	// config recommendation strategy
	ConfigStrategy string `json:"configStrategy,omitempty"`

	// Epoch time of cluster creation finish
	CreationFinished int64 `json:"creationFinished,omitempty"`

	// custom containers
	CustomContainers *CustomContainerResponse `json:"customContainers,omitempty"`

	// custom queue for yarn orchestrator
	CustomQueue string `json:"customQueue,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// executor type of cluster
	ExecutorType string `json:"executorType,omitempty"`

	// ambari blueprint JSON, set this or the url field
	ExtendedBlueprintText string `json:"extendedBlueprintText,omitempty"`

	// filesystem for a specific stack
	FileSystemResponse *FileSystemResponse `json:"fileSystemResponse,omitempty"`

	// gateway
	Gateway *GatewayJSON `json:"gateway,omitempty"`

	// collection of hostgroups
	// Unique: true
	HostGroups []*HostGroupResponse `json:"hostGroups"`

	// duration - how long the cluster is running in hours
	HoursUp int32 `json:"hoursUp,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// kerberos response
	KerberosResponse *KerberosResponse `json:"kerberosResponse,omitempty"`

	// LDAP config for the cluster
	LdapConfig *LdapConfigResponse `json:"ldapConfig,omitempty"`

	// LDAP config id for the cluster
	LdapConfigID int64 `json:"ldapConfigId,omitempty"`

	// duration - how long the cluster is running in minutes (minus hours)
	MinutesUp int32 `json:"minutesUp,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// ambari password
	Password string `json:"password,omitempty"`

	// proxy configuration name for the cluster
	ProxyName string `json:"proxyName,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	RdsConfigIds []int64 `json:"rdsConfigIds"`

	// RDS configurations for the cluster
	// Unique: true
	RdsConfigs []*RDSConfigResponse `json:"rdsConfigs"`

	// tells wether the cluster is secured or not
	Secure *bool `json:"secure,omitempty"`

	// shared service for a specific stack
	SharedServiceResponse *SharedServiceResponse `json:"sharedServiceResponse,omitempty"`

	// status of the cluster
	Status string `json:"status,omitempty"`

	// status message of the cluster
	StatusReason string `json:"statusReason,omitempty"`

	// duration - how long the cluster is running in milliseconds
	Uptime int64 `json:"uptime,omitempty"`

	// ambari username
	UserName string `json:"userName,omitempty"`

	// workspace of the resource
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

/* polymorph AutoscaleClusterResponse ambariDatabaseDetails false */

/* polymorph AutoscaleClusterResponse ambariRepoDetailsJson false */

/* polymorph AutoscaleClusterResponse ambariServerIp false */

/* polymorph AutoscaleClusterResponse ambariServerUrl false */

/* polymorph AutoscaleClusterResponse ambariStackDetails false */

/* polymorph AutoscaleClusterResponse attributes false */

/* polymorph AutoscaleClusterResponse blueprint false */

/* polymorph AutoscaleClusterResponse blueprintCustomProperties false */

/* polymorph AutoscaleClusterResponse blueprintId false */

/* polymorph AutoscaleClusterResponse blueprintInputs false */

/* polymorph AutoscaleClusterResponse cluster false */

/* polymorph AutoscaleClusterResponse clusterExposedServicesForTopologies false */

/* polymorph AutoscaleClusterResponse configStrategy false */

/* polymorph AutoscaleClusterResponse creationFinished false */

/* polymorph AutoscaleClusterResponse customContainers false */

/* polymorph AutoscaleClusterResponse customQueue false */

/* polymorph AutoscaleClusterResponse description false */

/* polymorph AutoscaleClusterResponse executorType false */

/* polymorph AutoscaleClusterResponse extendedBlueprintText false */

/* polymorph AutoscaleClusterResponse fileSystemResponse false */

/* polymorph AutoscaleClusterResponse gateway false */

/* polymorph AutoscaleClusterResponse hostGroups false */

/* polymorph AutoscaleClusterResponse hoursUp false */

/* polymorph AutoscaleClusterResponse id false */

/* polymorph AutoscaleClusterResponse kerberosResponse false */

/* polymorph AutoscaleClusterResponse ldapConfig false */

/* polymorph AutoscaleClusterResponse ldapConfigId false */

/* polymorph AutoscaleClusterResponse minutesUp false */

/* polymorph AutoscaleClusterResponse name false */

/* polymorph AutoscaleClusterResponse password false */

/* polymorph AutoscaleClusterResponse proxyName false */

/* polymorph AutoscaleClusterResponse rdsConfigIds false */

/* polymorph AutoscaleClusterResponse rdsConfigs false */

/* polymorph AutoscaleClusterResponse secure false */

/* polymorph AutoscaleClusterResponse sharedServiceResponse false */

/* polymorph AutoscaleClusterResponse status false */

/* polymorph AutoscaleClusterResponse statusReason false */

/* polymorph AutoscaleClusterResponse uptime false */

/* polymorph AutoscaleClusterResponse userName false */

/* polymorph AutoscaleClusterResponse workspace false */

// Validate validates this autoscale cluster response
func (m *AutoscaleClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariDatabaseDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAmbariRepoDetailsJSON(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAmbariStackDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlueprint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlueprintInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterExposedServicesForTopologies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomContainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExecutorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFileSystemResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKerberosResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLdapConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSharedServiceResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscaleClusterResponse) validateAmbariDatabaseDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariDatabaseDetails) { // not required
		return nil
	}

	if m.AmbariDatabaseDetails != nil {

		if err := m.AmbariDatabaseDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariDatabaseDetails")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateAmbariRepoDetailsJSON(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariRepoDetailsJSON) { // not required
		return nil
	}

	if m.AmbariRepoDetailsJSON != nil {

		if err := m.AmbariRepoDetailsJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariRepoDetailsJson")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateAmbariStackDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariStackDetails) { // not required
		return nil
	}

	if m.AmbariStackDetails != nil {

		if err := m.AmbariStackDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariStackDetails")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateBlueprint(formats strfmt.Registry) error {

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

func (m *AutoscaleClusterResponse) validateBlueprintInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.BlueprintInputs) { // not required
		return nil
	}

	if err := validate.UniqueItems("blueprintInputs", "body", m.BlueprintInputs); err != nil {
		return err
	}

	for i := 0; i < len(m.BlueprintInputs); i++ {

		if swag.IsZero(m.BlueprintInputs[i]) { // not required
			continue
		}

		if m.BlueprintInputs[i] != nil {

			if err := m.BlueprintInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blueprintInputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoscaleClusterResponse) validateClusterExposedServicesForTopologies(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterExposedServicesForTopologies) { // not required
		return nil
	}

	if err := validate.Required("clusterExposedServicesForTopologies", "body", m.ClusterExposedServicesForTopologies); err != nil {
		return err
	}

	return nil
}

var autoscaleClusterResponseTypeConfigStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEVER_APPLY","ONLY_STACK_DEFAULTS_APPLY","ALWAYS_APPLY","ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleClusterResponseTypeConfigStrategyPropEnum = append(autoscaleClusterResponseTypeConfigStrategyPropEnum, v)
	}
}

const (
	// AutoscaleClusterResponseConfigStrategyNEVERAPPLY captures enum value "NEVER_APPLY"
	AutoscaleClusterResponseConfigStrategyNEVERAPPLY string = "NEVER_APPLY"
	// AutoscaleClusterResponseConfigStrategyONLYSTACKDEFAULTSAPPLY captures enum value "ONLY_STACK_DEFAULTS_APPLY"
	AutoscaleClusterResponseConfigStrategyONLYSTACKDEFAULTSAPPLY string = "ONLY_STACK_DEFAULTS_APPLY"
	// AutoscaleClusterResponseConfigStrategyALWAYSAPPLY captures enum value "ALWAYS_APPLY"
	AutoscaleClusterResponseConfigStrategyALWAYSAPPLY string = "ALWAYS_APPLY"
	// AutoscaleClusterResponseConfigStrategyALWAYSAPPLYDONTOVERRIDECUSTOMVALUES captures enum value "ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"
	AutoscaleClusterResponseConfigStrategyALWAYSAPPLYDONTOVERRIDECUSTOMVALUES string = "ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"
)

// prop value enum
func (m *AutoscaleClusterResponse) validateConfigStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleClusterResponseTypeConfigStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleClusterResponse) validateConfigStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigStrategyEnum("configStrategy", "body", m.ConfigStrategy); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateCustomContainers(formats strfmt.Registry) error {

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

var autoscaleClusterResponseTypeExecutorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTAINER","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleClusterResponseTypeExecutorTypePropEnum = append(autoscaleClusterResponseTypeExecutorTypePropEnum, v)
	}
}

const (
	// AutoscaleClusterResponseExecutorTypeCONTAINER captures enum value "CONTAINER"
	AutoscaleClusterResponseExecutorTypeCONTAINER string = "CONTAINER"
	// AutoscaleClusterResponseExecutorTypeDEFAULT captures enum value "DEFAULT"
	AutoscaleClusterResponseExecutorTypeDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *AutoscaleClusterResponse) validateExecutorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleClusterResponseTypeExecutorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleClusterResponse) validateExecutorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutorTypeEnum("executorType", "body", m.ExecutorType); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateFileSystemResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.FileSystemResponse) { // not required
		return nil
	}

	if m.FileSystemResponse != nil {

		if err := m.FileSystemResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileSystemResponse")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateGateway(formats strfmt.Registry) error {

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

func (m *AutoscaleClusterResponse) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {

		if swag.IsZero(m.HostGroups[i]) { // not required
			continue
		}

		if m.HostGroups[i] != nil {

			if err := m.HostGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoscaleClusterResponse) validateKerberosResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.KerberosResponse) { // not required
		return nil
	}

	if m.KerberosResponse != nil {

		if err := m.KerberosResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberosResponse")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateLdapConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapConfig) { // not required
		return nil
	}

	if m.LdapConfig != nil {

		if err := m.LdapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateRdsConfigIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigIds", "body", m.RdsConfigIds); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateRdsConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigs) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigs", "body", m.RdsConfigs); err != nil {
		return err
	}

	for i := 0; i < len(m.RdsConfigs); i++ {

		if swag.IsZero(m.RdsConfigs[i]) { // not required
			continue
		}

		if m.RdsConfigs[i] != nil {

			if err := m.RdsConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rdsConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AutoscaleClusterResponse) validateSharedServiceResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedServiceResponse) { // not required
		return nil
	}

	if m.SharedServiceResponse != nil {

		if err := m.SharedServiceResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedServiceResponse")
			}
			return err
		}
	}

	return nil
}

var autoscaleClusterResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleClusterResponseTypeStatusPropEnum = append(autoscaleClusterResponseTypeStatusPropEnum, v)
	}
}

const (
	// AutoscaleClusterResponseStatusREQUESTED captures enum value "REQUESTED"
	AutoscaleClusterResponseStatusREQUESTED string = "REQUESTED"
	// AutoscaleClusterResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	AutoscaleClusterResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// AutoscaleClusterResponseStatusAVAILABLE captures enum value "AVAILABLE"
	AutoscaleClusterResponseStatusAVAILABLE string = "AVAILABLE"
	// AutoscaleClusterResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	AutoscaleClusterResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// AutoscaleClusterResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	AutoscaleClusterResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// AutoscaleClusterResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	AutoscaleClusterResponseStatusUPDATEFAILED string = "UPDATE_FAILED"
	// AutoscaleClusterResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	AutoscaleClusterResponseStatusCREATEFAILED string = "CREATE_FAILED"
	// AutoscaleClusterResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	AutoscaleClusterResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// AutoscaleClusterResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	AutoscaleClusterResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// AutoscaleClusterResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	AutoscaleClusterResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// AutoscaleClusterResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	AutoscaleClusterResponseStatusDELETEFAILED string = "DELETE_FAILED"
	// AutoscaleClusterResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	AutoscaleClusterResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// AutoscaleClusterResponseStatusSTOPPED captures enum value "STOPPED"
	AutoscaleClusterResponseStatusSTOPPED string = "STOPPED"
	// AutoscaleClusterResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	AutoscaleClusterResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// AutoscaleClusterResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	AutoscaleClusterResponseStatusSTARTREQUESTED string = "START_REQUESTED"
	// AutoscaleClusterResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	AutoscaleClusterResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// AutoscaleClusterResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	AutoscaleClusterResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// AutoscaleClusterResponseStatusSTARTFAILED captures enum value "START_FAILED"
	AutoscaleClusterResponseStatusSTARTFAILED string = "START_FAILED"
	// AutoscaleClusterResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	AutoscaleClusterResponseStatusSTOPFAILED string = "STOP_FAILED"
	// AutoscaleClusterResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	AutoscaleClusterResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *AutoscaleClusterResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleClusterResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleClusterResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleClusterResponse) validateWorkspace(formats strfmt.Registry) error {

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
func (m *AutoscaleClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscaleClusterResponse) UnmarshalBinary(b []byte) error {
	var res AutoscaleClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
