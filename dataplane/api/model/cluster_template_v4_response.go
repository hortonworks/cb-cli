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

// ClusterTemplateV4Response cluster template v4 response
// swagger:model ClusterTemplateV4Response
type ClusterTemplateV4Response struct {

	// cloudplatform which this template is compatible with
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// the unique crn of the resource
	// Required: true
	Crn *string `json:"crn"`

	// datalake required which this template is compatible with. The default is OPTIONAL
	// Enum: [NONE OPTIONAL REQUIRED]
	DatalakeRequired string `json:"datalakeRequired,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// stringified template JSON
	// Required: true
	DistroXTemplate *DistroXV1Request `json:"distroXTemplate"`

	// environment crn
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// feature state
	// Enum: [PREVIEW RELEASED INTERNAL]
	FeatureState string `json:"featureState,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// last updated
	LastUpdated int64 `json:"lastUpdated,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// node count of the cluster template
	NodeCount int32 `json:"nodeCount,omitempty"`

	// stack type of cluster template
	StackType string `json:"stackType,omitempty"`

	// stack version of cluster template
	StackVersion string `json:"stackVersion,omitempty"`

	// status
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED SERVICE_MANAGED OUTDATED]
	Status string `json:"status,omitempty"`

	// type
	// Required: true
	// Enum: [SPARK HIVE EDW ETL DATASCIENCE DATAMART DATALAKE DATAENGINEERING DATAENGINEERING_HA STREAMING STREAMING_HA FLOW_MANAGEMENT FLOW_MANAGEMENT_HA OPERATIONALDATABASE DISCOVERY_DATA_AND_EXPLORATION OTHER]
	Type *string `json:"type"`
}

// Validate validates this cluster template v4 response
func (m *ClusterTemplateV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistroXTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTemplateV4Response) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

var clusterTemplateV4ResponseTypeDatalakeRequiredPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","OPTIONAL","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateV4ResponseTypeDatalakeRequiredPropEnum = append(clusterTemplateV4ResponseTypeDatalakeRequiredPropEnum, v)
	}
}

const (

	// ClusterTemplateV4ResponseDatalakeRequiredNONE captures enum value "NONE"
	ClusterTemplateV4ResponseDatalakeRequiredNONE string = "NONE"

	// ClusterTemplateV4ResponseDatalakeRequiredOPTIONAL captures enum value "OPTIONAL"
	ClusterTemplateV4ResponseDatalakeRequiredOPTIONAL string = "OPTIONAL"

	// ClusterTemplateV4ResponseDatalakeRequiredREQUIRED captures enum value "REQUIRED"
	ClusterTemplateV4ResponseDatalakeRequiredREQUIRED string = "REQUIRED"
)

// prop value enum
func (m *ClusterTemplateV4Response) validateDatalakeRequiredEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateV4ResponseTypeDatalakeRequiredPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateV4Response) validateDatalakeRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.DatalakeRequired) { // not required
		return nil
	}

	// value enum
	if err := m.validateDatalakeRequiredEnum("datalakeRequired", "body", m.DatalakeRequired); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateV4Response) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterTemplateV4Response) validateDistroXTemplate(formats strfmt.Registry) error {

	if err := validate.Required("distroXTemplate", "body", m.DistroXTemplate); err != nil {
		return err
	}

	if m.DistroXTemplate != nil {
		if err := m.DistroXTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distroXTemplate")
			}
			return err
		}
	}

	return nil
}

var clusterTemplateV4ResponseTypeFeatureStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PREVIEW","RELEASED","INTERNAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateV4ResponseTypeFeatureStatePropEnum = append(clusterTemplateV4ResponseTypeFeatureStatePropEnum, v)
	}
}

const (

	// ClusterTemplateV4ResponseFeatureStatePREVIEW captures enum value "PREVIEW"
	ClusterTemplateV4ResponseFeatureStatePREVIEW string = "PREVIEW"

	// ClusterTemplateV4ResponseFeatureStateRELEASED captures enum value "RELEASED"
	ClusterTemplateV4ResponseFeatureStateRELEASED string = "RELEASED"

	// ClusterTemplateV4ResponseFeatureStateINTERNAL captures enum value "INTERNAL"
	ClusterTemplateV4ResponseFeatureStateINTERNAL string = "INTERNAL"
)

// prop value enum
func (m *ClusterTemplateV4Response) validateFeatureStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateV4ResponseTypeFeatureStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateV4Response) validateFeatureState(formats strfmt.Registry) error {

	if swag.IsZero(m.FeatureState) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeatureStateEnum("featureState", "body", m.FeatureState); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTemplateV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[^;\/%]*$`); err != nil {
		return err
	}

	return nil
}

var clusterTemplateV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED","SERVICE_MANAGED","OUTDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateV4ResponseTypeStatusPropEnum = append(clusterTemplateV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterTemplateV4ResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterTemplateV4ResponseStatusDEFAULT string = "DEFAULT"

	// ClusterTemplateV4ResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterTemplateV4ResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterTemplateV4ResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterTemplateV4ResponseStatusUSERMANAGED string = "USER_MANAGED"

	// ClusterTemplateV4ResponseStatusSERVICEMANAGED captures enum value "SERVICE_MANAGED"
	ClusterTemplateV4ResponseStatusSERVICEMANAGED string = "SERVICE_MANAGED"

	// ClusterTemplateV4ResponseStatusOUTDATED captures enum value "OUTDATED"
	ClusterTemplateV4ResponseStatusOUTDATED string = "OUTDATED"
)

// prop value enum
func (m *ClusterTemplateV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateV4Response) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var clusterTemplateV4ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SPARK","HIVE","EDW","ETL","DATASCIENCE","DATAMART","DATALAKE","DATAENGINEERING","DATAENGINEERING_HA","STREAMING","STREAMING_HA","FLOW_MANAGEMENT","FLOW_MANAGEMENT_HA","OPERATIONALDATABASE","DISCOVERY_DATA_AND_EXPLORATION","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateV4ResponseTypeTypePropEnum = append(clusterTemplateV4ResponseTypeTypePropEnum, v)
	}
}

const (

	// ClusterTemplateV4ResponseTypeSPARK captures enum value "SPARK"
	ClusterTemplateV4ResponseTypeSPARK string = "SPARK"

	// ClusterTemplateV4ResponseTypeHIVE captures enum value "HIVE"
	ClusterTemplateV4ResponseTypeHIVE string = "HIVE"

	// ClusterTemplateV4ResponseTypeEDW captures enum value "EDW"
	ClusterTemplateV4ResponseTypeEDW string = "EDW"

	// ClusterTemplateV4ResponseTypeETL captures enum value "ETL"
	ClusterTemplateV4ResponseTypeETL string = "ETL"

	// ClusterTemplateV4ResponseTypeDATASCIENCE captures enum value "DATASCIENCE"
	ClusterTemplateV4ResponseTypeDATASCIENCE string = "DATASCIENCE"

	// ClusterTemplateV4ResponseTypeDATAMART captures enum value "DATAMART"
	ClusterTemplateV4ResponseTypeDATAMART string = "DATAMART"

	// ClusterTemplateV4ResponseTypeDATALAKE captures enum value "DATALAKE"
	ClusterTemplateV4ResponseTypeDATALAKE string = "DATALAKE"

	// ClusterTemplateV4ResponseTypeDATAENGINEERING captures enum value "DATAENGINEERING"
	ClusterTemplateV4ResponseTypeDATAENGINEERING string = "DATAENGINEERING"

	// ClusterTemplateV4ResponseTypeDATAENGINEERINGHA captures enum value "DATAENGINEERING_HA"
	ClusterTemplateV4ResponseTypeDATAENGINEERINGHA string = "DATAENGINEERING_HA"

	// ClusterTemplateV4ResponseTypeSTREAMING captures enum value "STREAMING"
	ClusterTemplateV4ResponseTypeSTREAMING string = "STREAMING"

	// ClusterTemplateV4ResponseTypeSTREAMINGHA captures enum value "STREAMING_HA"
	ClusterTemplateV4ResponseTypeSTREAMINGHA string = "STREAMING_HA"

	// ClusterTemplateV4ResponseTypeFLOWMANAGEMENT captures enum value "FLOW_MANAGEMENT"
	ClusterTemplateV4ResponseTypeFLOWMANAGEMENT string = "FLOW_MANAGEMENT"

	// ClusterTemplateV4ResponseTypeFLOWMANAGEMENTHA captures enum value "FLOW_MANAGEMENT_HA"
	ClusterTemplateV4ResponseTypeFLOWMANAGEMENTHA string = "FLOW_MANAGEMENT_HA"

	// ClusterTemplateV4ResponseTypeOPERATIONALDATABASE captures enum value "OPERATIONALDATABASE"
	ClusterTemplateV4ResponseTypeOPERATIONALDATABASE string = "OPERATIONALDATABASE"

	// ClusterTemplateV4ResponseTypeDISCOVERYDATAANDEXPLORATION captures enum value "DISCOVERY_DATA_AND_EXPLORATION"
	ClusterTemplateV4ResponseTypeDISCOVERYDATAANDEXPLORATION string = "DISCOVERY_DATA_AND_EXPLORATION"

	// ClusterTemplateV4ResponseTypeOTHER captures enum value "OTHER"
	ClusterTemplateV4ResponseTypeOTHER string = "OTHER"
)

// prop value enum
func (m *ClusterTemplateV4Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateV4ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateV4Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTemplateV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplateV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterTemplateV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
