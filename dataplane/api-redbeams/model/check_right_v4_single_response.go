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

// CheckRightV4SingleResponse check right v4 single response
// swagger:model CheckRightV4SingleResponse
type CheckRightV4SingleResponse struct {

	// result
	Result bool `json:"result,omitempty"`

	// right
	// Enum: [ENV_CREATE LIST_ASSIGNED_ROLES CREATE_CLUSTER_TEMPLATE CREATE_CLUSTER_DEFINITION CREATE_CREDENTIAL CREATE_RECIPE CREATE_IMAGE_CATALOG ENV_START ENV_STOP ENV_DELETE ENV_DESCRIBE ENV_EDIT CHANGE_CRED DH_CREATE UPDATE_AZURE_ENCRYPTION_RESOURCES DH_START DH_STOP DH_DELETE DH_REPAIR DH_RESIZE DH_RETRY DH_DESCRIBE DH_RECOVER DH_UPGRADE DH_REFRESH_RECIPES SDX_UPGRADE SDX_RECOVER SDX_REPAIR SDX_RETRY SDX_DESCRIBE SDX_RESIZE SDX_REFRESH_RECIPES]
	Right string `json:"right,omitempty"`
}

// Validate validates this check right v4 single response
func (m *CheckRightV4SingleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var checkRightV4SingleResponseTypeRightPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENV_CREATE","LIST_ASSIGNED_ROLES","CREATE_CLUSTER_TEMPLATE","CREATE_CLUSTER_DEFINITION","CREATE_CREDENTIAL","CREATE_RECIPE","CREATE_IMAGE_CATALOG","ENV_START","ENV_STOP","ENV_DELETE","ENV_DESCRIBE","ENV_EDIT","CHANGE_CRED","DH_CREATE","UPDATE_AZURE_ENCRYPTION_RESOURCES","DH_START","DH_STOP","DH_DELETE","DH_REPAIR","DH_RESIZE","DH_RETRY","DH_DESCRIBE","DH_RECOVER","DH_UPGRADE","DH_REFRESH_RECIPES","SDX_UPGRADE","SDX_RECOVER","SDX_REPAIR","SDX_RETRY","SDX_DESCRIBE","SDX_RESIZE","SDX_REFRESH_RECIPES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkRightV4SingleResponseTypeRightPropEnum = append(checkRightV4SingleResponseTypeRightPropEnum, v)
	}
}

const (

	// CheckRightV4SingleResponseRightENVCREATE captures enum value "ENV_CREATE"
	CheckRightV4SingleResponseRightENVCREATE string = "ENV_CREATE"

	// CheckRightV4SingleResponseRightLISTASSIGNEDROLES captures enum value "LIST_ASSIGNED_ROLES"
	CheckRightV4SingleResponseRightLISTASSIGNEDROLES string = "LIST_ASSIGNED_ROLES"

	// CheckRightV4SingleResponseRightCREATECLUSTERTEMPLATE captures enum value "CREATE_CLUSTER_TEMPLATE"
	CheckRightV4SingleResponseRightCREATECLUSTERTEMPLATE string = "CREATE_CLUSTER_TEMPLATE"

	// CheckRightV4SingleResponseRightCREATECLUSTERDEFINITION captures enum value "CREATE_CLUSTER_DEFINITION"
	CheckRightV4SingleResponseRightCREATECLUSTERDEFINITION string = "CREATE_CLUSTER_DEFINITION"

	// CheckRightV4SingleResponseRightCREATECREDENTIAL captures enum value "CREATE_CREDENTIAL"
	CheckRightV4SingleResponseRightCREATECREDENTIAL string = "CREATE_CREDENTIAL"

	// CheckRightV4SingleResponseRightCREATERECIPE captures enum value "CREATE_RECIPE"
	CheckRightV4SingleResponseRightCREATERECIPE string = "CREATE_RECIPE"

	// CheckRightV4SingleResponseRightCREATEIMAGECATALOG captures enum value "CREATE_IMAGE_CATALOG"
	CheckRightV4SingleResponseRightCREATEIMAGECATALOG string = "CREATE_IMAGE_CATALOG"

	// CheckRightV4SingleResponseRightENVSTART captures enum value "ENV_START"
	CheckRightV4SingleResponseRightENVSTART string = "ENV_START"

	// CheckRightV4SingleResponseRightENVSTOP captures enum value "ENV_STOP"
	CheckRightV4SingleResponseRightENVSTOP string = "ENV_STOP"

	// CheckRightV4SingleResponseRightENVDELETE captures enum value "ENV_DELETE"
	CheckRightV4SingleResponseRightENVDELETE string = "ENV_DELETE"

	// CheckRightV4SingleResponseRightENVDESCRIBE captures enum value "ENV_DESCRIBE"
	CheckRightV4SingleResponseRightENVDESCRIBE string = "ENV_DESCRIBE"

	// CheckRightV4SingleResponseRightENVEDIT captures enum value "ENV_EDIT"
	CheckRightV4SingleResponseRightENVEDIT string = "ENV_EDIT"

	// CheckRightV4SingleResponseRightCHANGECRED captures enum value "CHANGE_CRED"
	CheckRightV4SingleResponseRightCHANGECRED string = "CHANGE_CRED"

	// CheckRightV4SingleResponseRightDHCREATE captures enum value "DH_CREATE"
	CheckRightV4SingleResponseRightDHCREATE string = "DH_CREATE"

	// CheckRightV4SingleResponseRightUPDATEAZUREENCRYPTIONRESOURCES captures enum value "UPDATE_AZURE_ENCRYPTION_RESOURCES"
	CheckRightV4SingleResponseRightUPDATEAZUREENCRYPTIONRESOURCES string = "UPDATE_AZURE_ENCRYPTION_RESOURCES"

	// CheckRightV4SingleResponseRightDHSTART captures enum value "DH_START"
	CheckRightV4SingleResponseRightDHSTART string = "DH_START"

	// CheckRightV4SingleResponseRightDHSTOP captures enum value "DH_STOP"
	CheckRightV4SingleResponseRightDHSTOP string = "DH_STOP"

	// CheckRightV4SingleResponseRightDHDELETE captures enum value "DH_DELETE"
	CheckRightV4SingleResponseRightDHDELETE string = "DH_DELETE"

	// CheckRightV4SingleResponseRightDHREPAIR captures enum value "DH_REPAIR"
	CheckRightV4SingleResponseRightDHREPAIR string = "DH_REPAIR"

	// CheckRightV4SingleResponseRightDHRESIZE captures enum value "DH_RESIZE"
	CheckRightV4SingleResponseRightDHRESIZE string = "DH_RESIZE"

	// CheckRightV4SingleResponseRightDHRETRY captures enum value "DH_RETRY"
	CheckRightV4SingleResponseRightDHRETRY string = "DH_RETRY"

	// CheckRightV4SingleResponseRightDHDESCRIBE captures enum value "DH_DESCRIBE"
	CheckRightV4SingleResponseRightDHDESCRIBE string = "DH_DESCRIBE"

	// CheckRightV4SingleResponseRightDHRECOVER captures enum value "DH_RECOVER"
	CheckRightV4SingleResponseRightDHRECOVER string = "DH_RECOVER"

	// CheckRightV4SingleResponseRightDHUPGRADE captures enum value "DH_UPGRADE"
	CheckRightV4SingleResponseRightDHUPGRADE string = "DH_UPGRADE"

	// CheckRightV4SingleResponseRightDHREFRESHRECIPES captures enum value "DH_REFRESH_RECIPES"
	CheckRightV4SingleResponseRightDHREFRESHRECIPES string = "DH_REFRESH_RECIPES"

	// CheckRightV4SingleResponseRightSDXUPGRADE captures enum value "SDX_UPGRADE"
	CheckRightV4SingleResponseRightSDXUPGRADE string = "SDX_UPGRADE"

	// CheckRightV4SingleResponseRightSDXRECOVER captures enum value "SDX_RECOVER"
	CheckRightV4SingleResponseRightSDXRECOVER string = "SDX_RECOVER"

	// CheckRightV4SingleResponseRightSDXREPAIR captures enum value "SDX_REPAIR"
	CheckRightV4SingleResponseRightSDXREPAIR string = "SDX_REPAIR"

	// CheckRightV4SingleResponseRightSDXRETRY captures enum value "SDX_RETRY"
	CheckRightV4SingleResponseRightSDXRETRY string = "SDX_RETRY"

	// CheckRightV4SingleResponseRightSDXDESCRIBE captures enum value "SDX_DESCRIBE"
	CheckRightV4SingleResponseRightSDXDESCRIBE string = "SDX_DESCRIBE"

	// CheckRightV4SingleResponseRightSDXRESIZE captures enum value "SDX_RESIZE"
	CheckRightV4SingleResponseRightSDXRESIZE string = "SDX_RESIZE"

	// CheckRightV4SingleResponseRightSDXREFRESHRECIPES captures enum value "SDX_REFRESH_RECIPES"
	CheckRightV4SingleResponseRightSDXREFRESHRECIPES string = "SDX_REFRESH_RECIPES"
)

// prop value enum
func (m *CheckRightV4SingleResponse) validateRightEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkRightV4SingleResponseTypeRightPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckRightV4SingleResponse) validateRight(formats strfmt.Registry) error {

	if swag.IsZero(m.Right) { // not required
		return nil
	}

	// value enum
	if err := m.validateRightEnum("right", "body", m.Right); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckRightV4SingleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRightV4SingleResponse) UnmarshalBinary(b []byte) error {
	var res CheckRightV4SingleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
