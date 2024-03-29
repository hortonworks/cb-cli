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

// CheckRightOnResourcesV4Response check right on resources v4 response
// swagger:model CheckRightOnResourcesV4Response
type CheckRightOnResourcesV4Response struct {

	// responses
	Responses []*CheckResourceRightV4Response `json:"responses"`

	// right
	// Enum: [ENV_CREATE LIST_ASSIGNED_ROLES CREATE_CLUSTER_TEMPLATE CREATE_CLUSTER_DEFINITION CREATE_CREDENTIAL CREATE_RECIPE CREATE_IMAGE_CATALOG CREATE_PROXY ENV_START ENV_STOP ENV_DELETE ENV_DESCRIBE ENV_EDIT CHANGE_CRED DH_CREATE UPDATE_AZURE_ENCRYPTION_RESOURCES ENV_VERTICAL_SCALING DH_START DH_STOP DH_DELETE DH_REPAIR DH_RESIZE DH_RETRY DH_DESCRIBE DH_RECOVER DH_UPGRADE DH_REFRESH_RECIPES DH_VERTICAL_SCALING SDX_UPGRADE SDX_RECOVER SDX_REPAIR SDX_RETRY SDX_DESCRIBE SDX_RESIZE SDX_VERTICAL_SCALING SDX_REFRESH_RECIPES]
	Right string `json:"right,omitempty"`
}

// Validate validates this check right on resources v4 response
func (m *CheckRightOnResourcesV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponses(formats); err != nil {
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

func (m *CheckRightOnResourcesV4Response) validateResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	for i := 0; i < len(m.Responses); i++ {
		if swag.IsZero(m.Responses[i]) { // not required
			continue
		}

		if m.Responses[i] != nil {
			if err := m.Responses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var checkRightOnResourcesV4ResponseTypeRightPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENV_CREATE","LIST_ASSIGNED_ROLES","CREATE_CLUSTER_TEMPLATE","CREATE_CLUSTER_DEFINITION","CREATE_CREDENTIAL","CREATE_RECIPE","CREATE_IMAGE_CATALOG","CREATE_PROXY","ENV_START","ENV_STOP","ENV_DELETE","ENV_DESCRIBE","ENV_EDIT","CHANGE_CRED","DH_CREATE","UPDATE_AZURE_ENCRYPTION_RESOURCES","ENV_VERTICAL_SCALING","DH_START","DH_STOP","DH_DELETE","DH_REPAIR","DH_RESIZE","DH_RETRY","DH_DESCRIBE","DH_RECOVER","DH_UPGRADE","DH_REFRESH_RECIPES","DH_VERTICAL_SCALING","SDX_UPGRADE","SDX_RECOVER","SDX_REPAIR","SDX_RETRY","SDX_DESCRIBE","SDX_RESIZE","SDX_VERTICAL_SCALING","SDX_REFRESH_RECIPES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkRightOnResourcesV4ResponseTypeRightPropEnum = append(checkRightOnResourcesV4ResponseTypeRightPropEnum, v)
	}
}

const (

	// CheckRightOnResourcesV4ResponseRightENVCREATE captures enum value "ENV_CREATE"
	CheckRightOnResourcesV4ResponseRightENVCREATE string = "ENV_CREATE"

	// CheckRightOnResourcesV4ResponseRightLISTASSIGNEDROLES captures enum value "LIST_ASSIGNED_ROLES"
	CheckRightOnResourcesV4ResponseRightLISTASSIGNEDROLES string = "LIST_ASSIGNED_ROLES"

	// CheckRightOnResourcesV4ResponseRightCREATECLUSTERTEMPLATE captures enum value "CREATE_CLUSTER_TEMPLATE"
	CheckRightOnResourcesV4ResponseRightCREATECLUSTERTEMPLATE string = "CREATE_CLUSTER_TEMPLATE"

	// CheckRightOnResourcesV4ResponseRightCREATECLUSTERDEFINITION captures enum value "CREATE_CLUSTER_DEFINITION"
	CheckRightOnResourcesV4ResponseRightCREATECLUSTERDEFINITION string = "CREATE_CLUSTER_DEFINITION"

	// CheckRightOnResourcesV4ResponseRightCREATECREDENTIAL captures enum value "CREATE_CREDENTIAL"
	CheckRightOnResourcesV4ResponseRightCREATECREDENTIAL string = "CREATE_CREDENTIAL"

	// CheckRightOnResourcesV4ResponseRightCREATERECIPE captures enum value "CREATE_RECIPE"
	CheckRightOnResourcesV4ResponseRightCREATERECIPE string = "CREATE_RECIPE"

	// CheckRightOnResourcesV4ResponseRightCREATEIMAGECATALOG captures enum value "CREATE_IMAGE_CATALOG"
	CheckRightOnResourcesV4ResponseRightCREATEIMAGECATALOG string = "CREATE_IMAGE_CATALOG"

	// CheckRightOnResourcesV4ResponseRightCREATEPROXY captures enum value "CREATE_PROXY"
	CheckRightOnResourcesV4ResponseRightCREATEPROXY string = "CREATE_PROXY"

	// CheckRightOnResourcesV4ResponseRightENVSTART captures enum value "ENV_START"
	CheckRightOnResourcesV4ResponseRightENVSTART string = "ENV_START"

	// CheckRightOnResourcesV4ResponseRightENVSTOP captures enum value "ENV_STOP"
	CheckRightOnResourcesV4ResponseRightENVSTOP string = "ENV_STOP"

	// CheckRightOnResourcesV4ResponseRightENVDELETE captures enum value "ENV_DELETE"
	CheckRightOnResourcesV4ResponseRightENVDELETE string = "ENV_DELETE"

	// CheckRightOnResourcesV4ResponseRightENVDESCRIBE captures enum value "ENV_DESCRIBE"
	CheckRightOnResourcesV4ResponseRightENVDESCRIBE string = "ENV_DESCRIBE"

	// CheckRightOnResourcesV4ResponseRightENVEDIT captures enum value "ENV_EDIT"
	CheckRightOnResourcesV4ResponseRightENVEDIT string = "ENV_EDIT"

	// CheckRightOnResourcesV4ResponseRightCHANGECRED captures enum value "CHANGE_CRED"
	CheckRightOnResourcesV4ResponseRightCHANGECRED string = "CHANGE_CRED"

	// CheckRightOnResourcesV4ResponseRightDHCREATE captures enum value "DH_CREATE"
	CheckRightOnResourcesV4ResponseRightDHCREATE string = "DH_CREATE"

	// CheckRightOnResourcesV4ResponseRightUPDATEAZUREENCRYPTIONRESOURCES captures enum value "UPDATE_AZURE_ENCRYPTION_RESOURCES"
	CheckRightOnResourcesV4ResponseRightUPDATEAZUREENCRYPTIONRESOURCES string = "UPDATE_AZURE_ENCRYPTION_RESOURCES"

	// CheckRightOnResourcesV4ResponseRightENVVERTICALSCALING captures enum value "ENV_VERTICAL_SCALING"
	CheckRightOnResourcesV4ResponseRightENVVERTICALSCALING string = "ENV_VERTICAL_SCALING"

	// CheckRightOnResourcesV4ResponseRightDHSTART captures enum value "DH_START"
	CheckRightOnResourcesV4ResponseRightDHSTART string = "DH_START"

	// CheckRightOnResourcesV4ResponseRightDHSTOP captures enum value "DH_STOP"
	CheckRightOnResourcesV4ResponseRightDHSTOP string = "DH_STOP"

	// CheckRightOnResourcesV4ResponseRightDHDELETE captures enum value "DH_DELETE"
	CheckRightOnResourcesV4ResponseRightDHDELETE string = "DH_DELETE"

	// CheckRightOnResourcesV4ResponseRightDHREPAIR captures enum value "DH_REPAIR"
	CheckRightOnResourcesV4ResponseRightDHREPAIR string = "DH_REPAIR"

	// CheckRightOnResourcesV4ResponseRightDHRESIZE captures enum value "DH_RESIZE"
	CheckRightOnResourcesV4ResponseRightDHRESIZE string = "DH_RESIZE"

	// CheckRightOnResourcesV4ResponseRightDHRETRY captures enum value "DH_RETRY"
	CheckRightOnResourcesV4ResponseRightDHRETRY string = "DH_RETRY"

	// CheckRightOnResourcesV4ResponseRightDHDESCRIBE captures enum value "DH_DESCRIBE"
	CheckRightOnResourcesV4ResponseRightDHDESCRIBE string = "DH_DESCRIBE"

	// CheckRightOnResourcesV4ResponseRightDHRECOVER captures enum value "DH_RECOVER"
	CheckRightOnResourcesV4ResponseRightDHRECOVER string = "DH_RECOVER"

	// CheckRightOnResourcesV4ResponseRightDHUPGRADE captures enum value "DH_UPGRADE"
	CheckRightOnResourcesV4ResponseRightDHUPGRADE string = "DH_UPGRADE"

	// CheckRightOnResourcesV4ResponseRightDHREFRESHRECIPES captures enum value "DH_REFRESH_RECIPES"
	CheckRightOnResourcesV4ResponseRightDHREFRESHRECIPES string = "DH_REFRESH_RECIPES"

	// CheckRightOnResourcesV4ResponseRightDHVERTICALSCALING captures enum value "DH_VERTICAL_SCALING"
	CheckRightOnResourcesV4ResponseRightDHVERTICALSCALING string = "DH_VERTICAL_SCALING"

	// CheckRightOnResourcesV4ResponseRightSDXUPGRADE captures enum value "SDX_UPGRADE"
	CheckRightOnResourcesV4ResponseRightSDXUPGRADE string = "SDX_UPGRADE"

	// CheckRightOnResourcesV4ResponseRightSDXRECOVER captures enum value "SDX_RECOVER"
	CheckRightOnResourcesV4ResponseRightSDXRECOVER string = "SDX_RECOVER"

	// CheckRightOnResourcesV4ResponseRightSDXREPAIR captures enum value "SDX_REPAIR"
	CheckRightOnResourcesV4ResponseRightSDXREPAIR string = "SDX_REPAIR"

	// CheckRightOnResourcesV4ResponseRightSDXRETRY captures enum value "SDX_RETRY"
	CheckRightOnResourcesV4ResponseRightSDXRETRY string = "SDX_RETRY"

	// CheckRightOnResourcesV4ResponseRightSDXDESCRIBE captures enum value "SDX_DESCRIBE"
	CheckRightOnResourcesV4ResponseRightSDXDESCRIBE string = "SDX_DESCRIBE"

	// CheckRightOnResourcesV4ResponseRightSDXRESIZE captures enum value "SDX_RESIZE"
	CheckRightOnResourcesV4ResponseRightSDXRESIZE string = "SDX_RESIZE"

	// CheckRightOnResourcesV4ResponseRightSDXVERTICALSCALING captures enum value "SDX_VERTICAL_SCALING"
	CheckRightOnResourcesV4ResponseRightSDXVERTICALSCALING string = "SDX_VERTICAL_SCALING"

	// CheckRightOnResourcesV4ResponseRightSDXREFRESHRECIPES captures enum value "SDX_REFRESH_RECIPES"
	CheckRightOnResourcesV4ResponseRightSDXREFRESHRECIPES string = "SDX_REFRESH_RECIPES"
)

// prop value enum
func (m *CheckRightOnResourcesV4Response) validateRightEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkRightOnResourcesV4ResponseTypeRightPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckRightOnResourcesV4Response) validateRight(formats strfmt.Registry) error {

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
func (m *CheckRightOnResourcesV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRightOnResourcesV4Response) UnmarshalBinary(b []byte) error {
	var res CheckRightOnResourcesV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
