// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UsedImageV1Response used image v1 response
// swagger:model UsedImageV1Response
type UsedImageV1Response struct {

	// image Id
	ImageID string `json:"imageId,omitempty"`
}

// Validate validates this used image v1 response
func (m *UsedImageV1Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsedImageV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsedImageV1Response) UnmarshalBinary(b []byte) error {
	var res UsedImageV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
