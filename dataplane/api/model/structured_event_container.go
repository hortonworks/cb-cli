// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StructuredEventContainer structured event container
// swagger:model StructuredEventContainer
type StructuredEventContainer struct {

	// flow
	Flow []*StructuredFlowEvent `json:"flow"`

	// notification
	Notification []*StructuredNotificationEvent `json:"notification"`

	// rest
	Rest []*StructuredRestCallEvent `json:"rest"`
}

// Validate validates this structured event container
func (m *StructuredEventContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StructuredEventContainer) validateFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	for i := 0; i < len(m.Flow); i++ {
		if swag.IsZero(m.Flow[i]) { // not required
			continue
		}

		if m.Flow[i] != nil {
			if err := m.Flow[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flow" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StructuredEventContainer) validateNotification(formats strfmt.Registry) error {

	if swag.IsZero(m.Notification) { // not required
		return nil
	}

	for i := 0; i < len(m.Notification); i++ {
		if swag.IsZero(m.Notification[i]) { // not required
			continue
		}

		if m.Notification[i] != nil {
			if err := m.Notification[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notification" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StructuredEventContainer) validateRest(formats strfmt.Registry) error {

	if swag.IsZero(m.Rest) { // not required
		return nil
	}

	for i := 0; i < len(m.Rest); i++ {
		if swag.IsZero(m.Rest[i]) { // not required
			continue
		}

		if m.Rest[i] != nil {
			if err := m.Rest[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rest" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StructuredEventContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StructuredEventContainer) UnmarshalBinary(b []byte) error {
	var res StructuredEventContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
