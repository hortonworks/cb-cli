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

// StackDatabaseServerResponse stack database server response
// swagger:model StackDatabaseServerResponse
type StackDatabaseServerResponse struct {

	// CRN of the cluster of the database server
	ClusterCrn string `json:"clusterCrn,omitempty"`

	// Creation date / time of the database server, in epoch milliseconds
	CreationDate int64 `json:"creationDate,omitempty"`

	// CRN of the database server
	Crn string `json:"crn,omitempty"`

	// Name of the database vendor (MYSQL, POSTGRES, ...)
	DatabaseVendor string `json:"databaseVendor,omitempty"`

	// Display name of the database vendor (MySQL, PostgreSQL, ...)
	DatabaseVendorDisplayName string `json:"databaseVendorDisplayName,omitempty"`

	// Description of the database server
	Description string `json:"description,omitempty"`

	// CRN of the environment of the database server
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// Host of the database server
	Host string `json:"host,omitempty"`

	// Major version of the database server engine
	// Enum: [VERSION_10 VERSION_11 VERSION_12 VERSION_13 VERSION_14]
	MajorVersion string `json:"majorVersion,omitempty"`

	// Name of the database server
	Name string `json:"name,omitempty"`

	// Port of the database server
	Port int32 `json:"port,omitempty"`

	// Ownership status of the database server
	// Enum: [UNKNOWN SERVICE_MANAGED USER_MANAGED]
	ResourceStatus string `json:"resourceStatus,omitempty"`

	// SSL config of the database server
	SslConfig *DatabaseServerSslConfig `json:"sslConfig,omitempty"`

	// Status of the database server stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED DELETE_REQUESTED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED UPGRADE_REQUESTED UPGRADE_IN_PROGRESS UPGRADE_FAILED UNKNOWN]
	Status string `json:"status,omitempty"`

	// Additional status information about the database server stack
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this stack database server response
func (m *StackDatabaseServerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMajorVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslConfig(formats); err != nil {
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

var stackDatabaseServerResponseTypeMajorVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERSION_10","VERSION_11","VERSION_12","VERSION_13","VERSION_14"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackDatabaseServerResponseTypeMajorVersionPropEnum = append(stackDatabaseServerResponseTypeMajorVersionPropEnum, v)
	}
}

const (

	// StackDatabaseServerResponseMajorVersionVERSION10 captures enum value "VERSION_10"
	StackDatabaseServerResponseMajorVersionVERSION10 string = "VERSION_10"

	// StackDatabaseServerResponseMajorVersionVERSION11 captures enum value "VERSION_11"
	StackDatabaseServerResponseMajorVersionVERSION11 string = "VERSION_11"

	// StackDatabaseServerResponseMajorVersionVERSION12 captures enum value "VERSION_12"
	StackDatabaseServerResponseMajorVersionVERSION12 string = "VERSION_12"

	// StackDatabaseServerResponseMajorVersionVERSION13 captures enum value "VERSION_13"
	StackDatabaseServerResponseMajorVersionVERSION13 string = "VERSION_13"

	// StackDatabaseServerResponseMajorVersionVERSION14 captures enum value "VERSION_14"
	StackDatabaseServerResponseMajorVersionVERSION14 string = "VERSION_14"
)

// prop value enum
func (m *StackDatabaseServerResponse) validateMajorVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackDatabaseServerResponseTypeMajorVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackDatabaseServerResponse) validateMajorVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.MajorVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateMajorVersionEnum("majorVersion", "body", m.MajorVersion); err != nil {
		return err
	}

	return nil
}

var stackDatabaseServerResponseTypeResourceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","SERVICE_MANAGED","USER_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackDatabaseServerResponseTypeResourceStatusPropEnum = append(stackDatabaseServerResponseTypeResourceStatusPropEnum, v)
	}
}

const (

	// StackDatabaseServerResponseResourceStatusUNKNOWN captures enum value "UNKNOWN"
	StackDatabaseServerResponseResourceStatusUNKNOWN string = "UNKNOWN"

	// StackDatabaseServerResponseResourceStatusSERVICEMANAGED captures enum value "SERVICE_MANAGED"
	StackDatabaseServerResponseResourceStatusSERVICEMANAGED string = "SERVICE_MANAGED"

	// StackDatabaseServerResponseResourceStatusUSERMANAGED captures enum value "USER_MANAGED"
	StackDatabaseServerResponseResourceStatusUSERMANAGED string = "USER_MANAGED"
)

// prop value enum
func (m *StackDatabaseServerResponse) validateResourceStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackDatabaseServerResponseTypeResourceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackDatabaseServerResponse) validateResourceStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceStatusEnum("resourceStatus", "body", m.ResourceStatus); err != nil {
		return err
	}

	return nil
}

func (m *StackDatabaseServerResponse) validateSslConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SslConfig) { // not required
		return nil
	}

	if m.SslConfig != nil {
		if err := m.SslConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sslConfig")
			}
			return err
		}
	}

	return nil
}

var stackDatabaseServerResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","DELETE_REQUESTED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","UPGRADE_REQUESTED","UPGRADE_IN_PROGRESS","UPGRADE_FAILED","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackDatabaseServerResponseTypeStatusPropEnum = append(stackDatabaseServerResponseTypeStatusPropEnum, v)
	}
}

const (

	// StackDatabaseServerResponseStatusREQUESTED captures enum value "REQUESTED"
	StackDatabaseServerResponseStatusREQUESTED string = "REQUESTED"

	// StackDatabaseServerResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackDatabaseServerResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// StackDatabaseServerResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackDatabaseServerResponseStatusAVAILABLE string = "AVAILABLE"

	// StackDatabaseServerResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackDatabaseServerResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// StackDatabaseServerResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackDatabaseServerResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// StackDatabaseServerResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackDatabaseServerResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// StackDatabaseServerResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackDatabaseServerResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// StackDatabaseServerResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackDatabaseServerResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// StackDatabaseServerResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	StackDatabaseServerResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// StackDatabaseServerResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackDatabaseServerResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// StackDatabaseServerResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackDatabaseServerResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// StackDatabaseServerResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackDatabaseServerResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// StackDatabaseServerResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackDatabaseServerResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// StackDatabaseServerResponseStatusSTOPPED captures enum value "STOPPED"
	StackDatabaseServerResponseStatusSTOPPED string = "STOPPED"

	// StackDatabaseServerResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackDatabaseServerResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// StackDatabaseServerResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackDatabaseServerResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// StackDatabaseServerResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackDatabaseServerResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// StackDatabaseServerResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackDatabaseServerResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// StackDatabaseServerResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackDatabaseServerResponseStatusSTARTFAILED string = "START_FAILED"

	// StackDatabaseServerResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackDatabaseServerResponseStatusSTOPFAILED string = "STOP_FAILED"

	// StackDatabaseServerResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackDatabaseServerResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// StackDatabaseServerResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackDatabaseServerResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// StackDatabaseServerResponseStatusUPGRADEREQUESTED captures enum value "UPGRADE_REQUESTED"
	StackDatabaseServerResponseStatusUPGRADEREQUESTED string = "UPGRADE_REQUESTED"

	// StackDatabaseServerResponseStatusUPGRADEINPROGRESS captures enum value "UPGRADE_IN_PROGRESS"
	StackDatabaseServerResponseStatusUPGRADEINPROGRESS string = "UPGRADE_IN_PROGRESS"

	// StackDatabaseServerResponseStatusUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	StackDatabaseServerResponseStatusUPGRADEFAILED string = "UPGRADE_FAILED"

	// StackDatabaseServerResponseStatusUNKNOWN captures enum value "UNKNOWN"
	StackDatabaseServerResponseStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *StackDatabaseServerResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackDatabaseServerResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackDatabaseServerResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *StackDatabaseServerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackDatabaseServerResponse) UnmarshalBinary(b []byte) error {
	var res StackDatabaseServerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
