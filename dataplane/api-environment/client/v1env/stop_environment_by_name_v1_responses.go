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

// StopEnvironmentByNameV1Reader is a Reader for the StopEnvironmentByNameV1 structure.
type StopEnvironmentByNameV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopEnvironmentByNameV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStopEnvironmentByNameV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopEnvironmentByNameV1OK creates a StopEnvironmentByNameV1OK with default headers values
func NewStopEnvironmentByNameV1OK() *StopEnvironmentByNameV1OK {
	return &StopEnvironmentByNameV1OK{}
}

/*StopEnvironmentByNameV1OK handles this case with default header values.

successful operation
*/
type StopEnvironmentByNameV1OK struct {
	Payload *model.FlowIdentifier
}

func (o *StopEnvironmentByNameV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/env/name/{name}/stop][%d] stopEnvironmentByNameV1OK  %+v", 200, o.Payload)
}

func (o *StopEnvironmentByNameV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
