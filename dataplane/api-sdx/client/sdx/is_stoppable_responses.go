// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// IsStoppableReader is a Reader for the IsStoppable structure.
type IsStoppableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsStoppableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIsStoppableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsStoppableOK creates a IsStoppableOK with default headers values
func NewIsStoppableOK() *IsStoppableOK {
	return &IsStoppableOK{}
}

/*IsStoppableOK handles this case with default header values.

successful operation
*/
type IsStoppableOK struct {
	Payload *model.SdxStopValidationResponse
}

func (o *IsStoppableOK) Error() string {
	return fmt.Sprintf("[GET /sdx/crn/{crn}/internal/stoppable][%d] isStoppableOK  %+v", 200, o.Payload)
}

func (o *IsStoppableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxStopValidationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
