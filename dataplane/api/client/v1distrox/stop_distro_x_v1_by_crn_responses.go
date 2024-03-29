// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// StopDistroXV1ByCrnReader is a Reader for the StopDistroXV1ByCrn structure.
type StopDistroXV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDistroXV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStopDistroXV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopDistroXV1ByCrnOK creates a StopDistroXV1ByCrnOK with default headers values
func NewStopDistroXV1ByCrnOK() *StopDistroXV1ByCrnOK {
	return &StopDistroXV1ByCrnOK{}
}

/*
StopDistroXV1ByCrnOK handles this case with default header values.

successful operation
*/
type StopDistroXV1ByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *StopDistroXV1ByCrnOK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/crn/{crn}/stop][%d] stopDistroXV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *StopDistroXV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
