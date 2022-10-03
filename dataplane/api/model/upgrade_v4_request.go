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

// UpgradeV4Request upgrade v4 request
// swagger:model UpgradeV4Request
type UpgradeV4Request struct {

	// Checks the eligibility of an image to upgrade
	DryRun bool `json:"dryRun,omitempty"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// UUID of the image to upgrade
	ImageID string `json:"imageId,omitempty"`

	// Upgrades to image with the same version of stack and clustermanager, if available
	LockComponents bool `json:"lockComponents,omitempty"`

	// replace vms
	ReplaceVms bool `json:"replaceVms,omitempty"`

	// Cloudera Runtime version
	Runtime string `json:"runtime,omitempty"`

	// Returns the list of images that are eligible for the upgrade
	// Enum: [SHOW LATEST_ONLY DO_NOT_SHOW]
	ShowAvailableImages string `json:"showAvailableImages,omitempty"`

	// With this option, the Data Lake upgrade can be performed with running Data Hub clusters. The usage of this option can cause problems on the running Data Hub clusters during the Data Lake upgrade.
	SkipDataHubValidation bool `json:"skipDataHubValidation,omitempty"`
}

// Validate validates this upgrade v4 request
func (m *UpgradeV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShowAvailableImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var upgradeV4RequestTypeShowAvailableImagesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SHOW","LATEST_ONLY","DO_NOT_SHOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		upgradeV4RequestTypeShowAvailableImagesPropEnum = append(upgradeV4RequestTypeShowAvailableImagesPropEnum, v)
	}
}

const (

	// UpgradeV4RequestShowAvailableImagesSHOW captures enum value "SHOW"
	UpgradeV4RequestShowAvailableImagesSHOW string = "SHOW"

	// UpgradeV4RequestShowAvailableImagesLATESTONLY captures enum value "LATEST_ONLY"
	UpgradeV4RequestShowAvailableImagesLATESTONLY string = "LATEST_ONLY"

	// UpgradeV4RequestShowAvailableImagesDONOTSHOW captures enum value "DO_NOT_SHOW"
	UpgradeV4RequestShowAvailableImagesDONOTSHOW string = "DO_NOT_SHOW"
)

// prop value enum
func (m *UpgradeV4Request) validateShowAvailableImagesEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, upgradeV4RequestTypeShowAvailableImagesPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpgradeV4Request) validateShowAvailableImages(formats strfmt.Registry) error {

	if swag.IsZero(m.ShowAvailableImages) { // not required
		return nil
	}

	// value enum
	if err := m.validateShowAvailableImagesEnum("showAvailableImages", "body", m.ShowAvailableImages); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeV4Request) UnmarshalBinary(b []byte) error {
	var res UpgradeV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
