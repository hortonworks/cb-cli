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

// CheckRightV4Request check right v4 request
// swagger:model CheckRightV4Request
type CheckRightV4Request struct {

	// rights
	Rights []string `json:"rights"`
}

// Validate validates this check right v4 request
func (m *CheckRightV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var checkRightV4RequestRightsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENV_CREATE","LIST_ASSIGNED_ROLES","CREATE_CLUSTER_TEMPLATE","CREATE_CLUSTER_DEFINITION","CREATE_CREDENTIAL","CREATE_RECIPE","CREATE_IMAGE_CATALOG","ENV_START","ENV_STOP","ENV_DELETE","ENV_DESCRIBE","CHANGE_CRED","DH_CREATE","UPDATE_AZURE_ENCRYPTION_RESOURCES","UPDATE_AWS_DISK_ENCRYPTION_PARAMETERS","DH_START","DH_STOP","DH_DELETE","DH_REPAIR","DH_RESIZE","DH_RETRY","DH_DESCRIBE","DH_RECOVER","DH_UPGRADE","DH_REFRESH_RECIPES","SDX_UPGRADE","SDX_RECOVER","SDX_REPAIR","SDX_RETRY","SDX_DESCRIBE","SDX_RESIZE","SDX_REFRESH_RECIPES","DISTROX_READ","DISTROX_WRITE","SDX_READ","SDX_WRITE","ENVIRONMENT_READ","ENVIRONMENT_WRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkRightV4RequestRightsItemsEnum = append(checkRightV4RequestRightsItemsEnum, v)
	}
}

func (m *CheckRightV4Request) validateRightsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkRightV4RequestRightsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckRightV4Request) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {

		// value enum
		if err := m.validateRightsItemsEnum("rights"+"."+strconv.Itoa(i), "body", m.Rights[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckRightV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRightV4Request) UnmarshalBinary(b []byte) error {
	var res CheckRightV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
