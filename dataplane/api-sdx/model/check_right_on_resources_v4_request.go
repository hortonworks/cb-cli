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

// CheckRightOnResourcesV4Request check right on resources v4 request
// swagger:model CheckRightOnResourcesV4Request
type CheckRightOnResourcesV4Request struct {

	// resource crns
	// Required: true
	ResourceCrns []string `json:"resourceCrns"`

	// right
	// Required: true
	// Enum: [ENV_CREATE LIST_ASSIGNED_ROLES CREATE_CLUSTER_TEMPLATE CREATE_CLUSTER_DEFINITION CREATE_CREDENTIAL CREATE_RECIPE CREATE_IMAGE_CATALOG ENV_START ENV_STOP ENV_DELETE ENV_DESCRIBE CHANGE_CRED DH_CREATE UPDATE_AZURE_ENCRYPTION_RESOURCES DH_START DH_STOP DH_DELETE DH_REPAIR DH_RESIZE DH_RETRY DH_DESCRIBE DH_RECOVER DH_UPGRADE DH_REFRESH_RECIPES SDX_UPGRADE SDX_RECOVER SDX_REPAIR SDX_RETRY SDX_DESCRIBE SDX_RESIZE SDX_REFRESH_RECIPES DISTROX_READ DISTROX_WRITE SDX_READ SDX_WRITE ENVIRONMENT_READ ENVIRONMENT_WRITE]
	Right *string `json:"right"`
}

// Validate validates this check right on resources v4 request
func (m *CheckRightOnResourcesV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceCrns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckRightOnResourcesV4Request) validateResourceCrns(formats strfmt.Registry) error {

	if err := validate.Required("resourceCrns", "body", m.ResourceCrns); err != nil {
		return err
	}

	return nil
}

var checkRightOnResourcesV4RequestTypeRightPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENV_CREATE","LIST_ASSIGNED_ROLES","CREATE_CLUSTER_TEMPLATE","CREATE_CLUSTER_DEFINITION","CREATE_CREDENTIAL","CREATE_RECIPE","CREATE_IMAGE_CATALOG","ENV_START","ENV_STOP","ENV_DELETE","ENV_DESCRIBE","CHANGE_CRED","DH_CREATE","UPDATE_AZURE_ENCRYPTION_RESOURCES","DH_START","DH_STOP","DH_DELETE","DH_REPAIR","DH_RESIZE","DH_RETRY","DH_DESCRIBE","DH_RECOVER","DH_UPGRADE","DH_REFRESH_RECIPES","SDX_UPGRADE","SDX_RECOVER","SDX_REPAIR","SDX_RETRY","SDX_DESCRIBE","SDX_RESIZE","SDX_REFRESH_RECIPES","DISTROX_READ","DISTROX_WRITE","SDX_READ","SDX_WRITE","ENVIRONMENT_READ","ENVIRONMENT_WRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkRightOnResourcesV4RequestTypeRightPropEnum = append(checkRightOnResourcesV4RequestTypeRightPropEnum, v)
	}
}

const (

	// CheckRightOnResourcesV4RequestRightENVCREATE captures enum value "ENV_CREATE"
	CheckRightOnResourcesV4RequestRightENVCREATE string = "ENV_CREATE"

	// CheckRightOnResourcesV4RequestRightLISTASSIGNEDROLES captures enum value "LIST_ASSIGNED_ROLES"
	CheckRightOnResourcesV4RequestRightLISTASSIGNEDROLES string = "LIST_ASSIGNED_ROLES"

	// CheckRightOnResourcesV4RequestRightCREATECLUSTERTEMPLATE captures enum value "CREATE_CLUSTER_TEMPLATE"
	CheckRightOnResourcesV4RequestRightCREATECLUSTERTEMPLATE string = "CREATE_CLUSTER_TEMPLATE"

	// CheckRightOnResourcesV4RequestRightCREATECLUSTERDEFINITION captures enum value "CREATE_CLUSTER_DEFINITION"
	CheckRightOnResourcesV4RequestRightCREATECLUSTERDEFINITION string = "CREATE_CLUSTER_DEFINITION"

	// CheckRightOnResourcesV4RequestRightCREATECREDENTIAL captures enum value "CREATE_CREDENTIAL"
	CheckRightOnResourcesV4RequestRightCREATECREDENTIAL string = "CREATE_CREDENTIAL"

	// CheckRightOnResourcesV4RequestRightCREATERECIPE captures enum value "CREATE_RECIPE"
	CheckRightOnResourcesV4RequestRightCREATERECIPE string = "CREATE_RECIPE"

	// CheckRightOnResourcesV4RequestRightCREATEIMAGECATALOG captures enum value "CREATE_IMAGE_CATALOG"
	CheckRightOnResourcesV4RequestRightCREATEIMAGECATALOG string = "CREATE_IMAGE_CATALOG"

	// CheckRightOnResourcesV4RequestRightENVSTART captures enum value "ENV_START"
	CheckRightOnResourcesV4RequestRightENVSTART string = "ENV_START"

	// CheckRightOnResourcesV4RequestRightENVSTOP captures enum value "ENV_STOP"
	CheckRightOnResourcesV4RequestRightENVSTOP string = "ENV_STOP"

	// CheckRightOnResourcesV4RequestRightENVDELETE captures enum value "ENV_DELETE"
	CheckRightOnResourcesV4RequestRightENVDELETE string = "ENV_DELETE"

	// CheckRightOnResourcesV4RequestRightENVDESCRIBE captures enum value "ENV_DESCRIBE"
	CheckRightOnResourcesV4RequestRightENVDESCRIBE string = "ENV_DESCRIBE"

	// CheckRightOnResourcesV4RequestRightCHANGECRED captures enum value "CHANGE_CRED"
	CheckRightOnResourcesV4RequestRightCHANGECRED string = "CHANGE_CRED"

	// CheckRightOnResourcesV4RequestRightDHCREATE captures enum value "DH_CREATE"
	CheckRightOnResourcesV4RequestRightDHCREATE string = "DH_CREATE"

	// CheckRightOnResourcesV4RequestRightUPDATEAZUREENCRYPTIONRESOURCES captures enum value "UPDATE_AZURE_ENCRYPTION_RESOURCES"
	CheckRightOnResourcesV4RequestRightUPDATEAZUREENCRYPTIONRESOURCES string = "UPDATE_AZURE_ENCRYPTION_RESOURCES"

	// CheckRightOnResourcesV4RequestRightDHSTART captures enum value "DH_START"
	CheckRightOnResourcesV4RequestRightDHSTART string = "DH_START"

	// CheckRightOnResourcesV4RequestRightDHSTOP captures enum value "DH_STOP"
	CheckRightOnResourcesV4RequestRightDHSTOP string = "DH_STOP"

	// CheckRightOnResourcesV4RequestRightDHDELETE captures enum value "DH_DELETE"
	CheckRightOnResourcesV4RequestRightDHDELETE string = "DH_DELETE"

	// CheckRightOnResourcesV4RequestRightDHREPAIR captures enum value "DH_REPAIR"
	CheckRightOnResourcesV4RequestRightDHREPAIR string = "DH_REPAIR"

	// CheckRightOnResourcesV4RequestRightDHRESIZE captures enum value "DH_RESIZE"
	CheckRightOnResourcesV4RequestRightDHRESIZE string = "DH_RESIZE"

	// CheckRightOnResourcesV4RequestRightDHRETRY captures enum value "DH_RETRY"
	CheckRightOnResourcesV4RequestRightDHRETRY string = "DH_RETRY"

	// CheckRightOnResourcesV4RequestRightDHDESCRIBE captures enum value "DH_DESCRIBE"
	CheckRightOnResourcesV4RequestRightDHDESCRIBE string = "DH_DESCRIBE"

	// CheckRightOnResourcesV4RequestRightDHRECOVER captures enum value "DH_RECOVER"
	CheckRightOnResourcesV4RequestRightDHRECOVER string = "DH_RECOVER"

	// CheckRightOnResourcesV4RequestRightDHUPGRADE captures enum value "DH_UPGRADE"
	CheckRightOnResourcesV4RequestRightDHUPGRADE string = "DH_UPGRADE"

	// CheckRightOnResourcesV4RequestRightDHREFRESHRECIPES captures enum value "DH_REFRESH_RECIPES"
	CheckRightOnResourcesV4RequestRightDHREFRESHRECIPES string = "DH_REFRESH_RECIPES"

	// CheckRightOnResourcesV4RequestRightSDXUPGRADE captures enum value "SDX_UPGRADE"
	CheckRightOnResourcesV4RequestRightSDXUPGRADE string = "SDX_UPGRADE"

	// CheckRightOnResourcesV4RequestRightSDXRECOVER captures enum value "SDX_RECOVER"
	CheckRightOnResourcesV4RequestRightSDXRECOVER string = "SDX_RECOVER"

	// CheckRightOnResourcesV4RequestRightSDXREPAIR captures enum value "SDX_REPAIR"
	CheckRightOnResourcesV4RequestRightSDXREPAIR string = "SDX_REPAIR"

	// CheckRightOnResourcesV4RequestRightSDXRETRY captures enum value "SDX_RETRY"
	CheckRightOnResourcesV4RequestRightSDXRETRY string = "SDX_RETRY"

	// CheckRightOnResourcesV4RequestRightSDXDESCRIBE captures enum value "SDX_DESCRIBE"
	CheckRightOnResourcesV4RequestRightSDXDESCRIBE string = "SDX_DESCRIBE"

	// CheckRightOnResourcesV4RequestRightSDXRESIZE captures enum value "SDX_RESIZE"
	CheckRightOnResourcesV4RequestRightSDXRESIZE string = "SDX_RESIZE"

	// CheckRightOnResourcesV4RequestRightSDXREFRESHRECIPES captures enum value "SDX_REFRESH_RECIPES"
	CheckRightOnResourcesV4RequestRightSDXREFRESHRECIPES string = "SDX_REFRESH_RECIPES"

	// CheckRightOnResourcesV4RequestRightDISTROXREAD captures enum value "DISTROX_READ"
	CheckRightOnResourcesV4RequestRightDISTROXREAD string = "DISTROX_READ"

	// CheckRightOnResourcesV4RequestRightDISTROXWRITE captures enum value "DISTROX_WRITE"
	CheckRightOnResourcesV4RequestRightDISTROXWRITE string = "DISTROX_WRITE"

	// CheckRightOnResourcesV4RequestRightSDXREAD captures enum value "SDX_READ"
	CheckRightOnResourcesV4RequestRightSDXREAD string = "SDX_READ"

	// CheckRightOnResourcesV4RequestRightSDXWRITE captures enum value "SDX_WRITE"
	CheckRightOnResourcesV4RequestRightSDXWRITE string = "SDX_WRITE"

	// CheckRightOnResourcesV4RequestRightENVIRONMENTREAD captures enum value "ENVIRONMENT_READ"
	CheckRightOnResourcesV4RequestRightENVIRONMENTREAD string = "ENVIRONMENT_READ"

	// CheckRightOnResourcesV4RequestRightENVIRONMENTWRITE captures enum value "ENVIRONMENT_WRITE"
	CheckRightOnResourcesV4RequestRightENVIRONMENTWRITE string = "ENVIRONMENT_WRITE"
)

// prop value enum
func (m *CheckRightOnResourcesV4Request) validateRightEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkRightOnResourcesV4RequestTypeRightPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckRightOnResourcesV4Request) validateRight(formats strfmt.Registry) error {

	if err := validate.Required("right", "body", m.Right); err != nil {
		return err
	}

	// value enum
	if err := m.validateRightEnum("right", "body", *m.Right); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckRightOnResourcesV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRightOnResourcesV4Request) UnmarshalBinary(b []byte) error {
	var res CheckRightOnResourcesV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
