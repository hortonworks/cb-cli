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

// StopSdxByNameReader is a Reader for the StopSdxByName structure.
type StopSdxByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopSdxByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStopSdxByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopSdxByNameOK creates a StopSdxByNameOK with default headers values
func NewStopSdxByNameOK() *StopSdxByNameOK {
	return &StopSdxByNameOK{}
}

/*
StopSdxByNameOK handles this case with default header values.

successful operation
*/
type StopSdxByNameOK struct {
	Payload *model.FlowIdentifier
}

func (o *StopSdxByNameOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/stop][%d] stopSdxByNameOK  %+v", 200, o.Payload)
}

func (o *StopSdxByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
