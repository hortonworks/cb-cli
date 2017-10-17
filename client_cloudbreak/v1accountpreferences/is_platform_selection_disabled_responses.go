// Code generated by go-swagger; DO NOT EDIT.

package v1accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// IsPlatformSelectionDisabledReader is a Reader for the IsPlatformSelectionDisabled structure.
type IsPlatformSelectionDisabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsPlatformSelectionDisabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIsPlatformSelectionDisabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsPlatformSelectionDisabledOK creates a IsPlatformSelectionDisabledOK with default headers values
func NewIsPlatformSelectionDisabledOK() *IsPlatformSelectionDisabledOK {
	return &IsPlatformSelectionDisabledOK{}
}

/*IsPlatformSelectionDisabledOK handles this case with default header values.

successful operation
*/
type IsPlatformSelectionDisabledOK struct {
	Payload IsPlatformSelectionDisabledOKBody
}

func (o *IsPlatformSelectionDisabledOK) Error() string {
	return fmt.Sprintf("[GET /v1/accountpreferences/isplatformselectiondisabled][%d] isPlatformSelectionDisabledOK  %+v", 200, o.Payload)
}

func (o *IsPlatformSelectionDisabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IsPlatformSelectionDisabledOKBody is platform selection disabled o k body
swagger:model IsPlatformSelectionDisabledOKBody
*/

type IsPlatformSelectionDisabledOKBody map[string]bool

// Validate validates this is platform selection disabled o k body
func (o IsPlatformSelectionDisabledOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
