package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*ImageJSON image json

swagger:model ImageJson
*/
type ImageJSON struct {

	/* name of the image
	 */
	ImageName *string `json:"imageName,omitempty"`
}

// Validate validates this image json
func (m *ImageJSON) Validate(formats strfmt.Registry) error {
	return nil
}
