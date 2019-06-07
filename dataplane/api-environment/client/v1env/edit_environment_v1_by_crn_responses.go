// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// EditEnvironmentV1ByCrnReader is a Reader for the EditEnvironmentV1ByCrn structure.
type EditEnvironmentV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditEnvironmentV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditEnvironmentV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditEnvironmentV1ByCrnOK creates a EditEnvironmentV1ByCrnOK with default headers values
func NewEditEnvironmentV1ByCrnOK() *EditEnvironmentV1ByCrnOK {
	return &EditEnvironmentV1ByCrnOK{}
}

/*EditEnvironmentV1ByCrnOK handles this case with default header values.

successful operation
*/
type EditEnvironmentV1ByCrnOK struct {
	Payload *model.DetailedEnvironmentV1Response
}

func (o *EditEnvironmentV1ByCrnOK) Error() string {
	return fmt.Sprintf("[PUT /v1/env/crn/{crn}][%d] editEnvironmentV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *EditEnvironmentV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedEnvironmentV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
