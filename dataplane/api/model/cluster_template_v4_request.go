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

// ClusterTemplateV4Request cluster template v4 request
// swagger:model ClusterTemplateV4Request
type ClusterTemplateV4Request struct {

	// cloudplatform which this template is compatible with
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// stringified template JSON
	// Required: true
	DistroXTemplate *DistroXV1Request `json:"distroXTemplate"`

	// name of the resource
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// type
	// Required: true
	// Enum: [SPARK HIVE EDW ETL DATASCIENCE DATAMART DATAMART_HA DATALAKE DATAENGINEERING DATAENGINEERING_HA OTHER]
	Type *string `json:"type"`
}

// Validate validates this cluster template v4 request
func (m *ClusterTemplateV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistroXTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ClusterTemplateV4Request) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterTemplateV4Request) validateDistroXTemplate(formats strfmt.Registry) error {

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

func (m *ClusterTemplateV4Request) validateName(formats strfmt.Registry) error {

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

var clusterTemplateV4RequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SPARK","HIVE","EDW","ETL","DATASCIENCE","DATAMART","DATAMART_HA","DATALAKE","DATAENGINEERING","DATAENGINEERING_HA","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTemplateV4RequestTypeTypePropEnum = append(clusterTemplateV4RequestTypeTypePropEnum, v)
	}
}

const (

	// ClusterTemplateV4RequestTypeSPARK captures enum value "SPARK"
	ClusterTemplateV4RequestTypeSPARK string = "SPARK"

	// ClusterTemplateV4RequestTypeHIVE captures enum value "HIVE"
	ClusterTemplateV4RequestTypeHIVE string = "HIVE"

	// ClusterTemplateV4RequestTypeEDW captures enum value "EDW"
	ClusterTemplateV4RequestTypeEDW string = "EDW"

	// ClusterTemplateV4RequestTypeETL captures enum value "ETL"
	ClusterTemplateV4RequestTypeETL string = "ETL"

	// ClusterTemplateV4RequestTypeDATASCIENCE captures enum value "DATASCIENCE"
	ClusterTemplateV4RequestTypeDATASCIENCE string = "DATASCIENCE"

	// ClusterTemplateV4RequestTypeDATAMART captures enum value "DATAMART"
	ClusterTemplateV4RequestTypeDATAMART string = "DATAMART"

	// ClusterTemplateV4RequestTypeDATAMARTHA captures enum value "DATAMART_HA"
	ClusterTemplateV4RequestTypeDATAMARTHA string = "DATAMART_HA"

	// ClusterTemplateV4RequestTypeDATALAKE captures enum value "DATALAKE"
	ClusterTemplateV4RequestTypeDATALAKE string = "DATALAKE"

	// ClusterTemplateV4RequestTypeDATAENGINEERING captures enum value "DATAENGINEERING"
	ClusterTemplateV4RequestTypeDATAENGINEERING string = "DATAENGINEERING"

	// ClusterTemplateV4RequestTypeDATAENGINEERINGHA captures enum value "DATAENGINEERING_HA"
	ClusterTemplateV4RequestTypeDATAENGINEERINGHA string = "DATAENGINEERING_HA"

	// ClusterTemplateV4RequestTypeOTHER captures enum value "OTHER"
	ClusterTemplateV4RequestTypeOTHER string = "OTHER"
)

// prop value enum
func (m *ClusterTemplateV4Request) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterTemplateV4RequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterTemplateV4Request) validateType(formats strfmt.Registry) error {

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
func (m *ClusterTemplateV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTemplateV4Request) UnmarshalBinary(b []byte) error {
	var res ClusterTemplateV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
