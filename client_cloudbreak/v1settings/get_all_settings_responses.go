// Code generated by go-swagger; DO NOT EDIT.

package v1settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAllSettingsReader is a Reader for the GetAllSettings structure.
type GetAllSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllSettingsOK creates a GetAllSettingsOK with default headers values
func NewGetAllSettingsOK() *GetAllSettingsOK {
	return &GetAllSettingsOK{}
}

/*GetAllSettingsOK handles this case with default header values.

successful operation
*/
type GetAllSettingsOK struct {
	Payload GetAllSettingsOKBody
}

func (o *GetAllSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/settings/all][%d] getAllSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAllSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAllSettingsOKBody get all settings o k body
swagger:model GetAllSettingsOKBody
*/

type GetAllSettingsOKBody map[string]map[string]interface{}

// Validate validates this get all settings o k body
func (o GetAllSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getAllSettingsOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
