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

// GetEnvironmentV1ByCrnReader is a Reader for the GetEnvironmentV1ByCrn structure.
type GetEnvironmentV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEnvironmentV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnvironmentV1ByCrnOK creates a GetEnvironmentV1ByCrnOK with default headers values
func NewGetEnvironmentV1ByCrnOK() *GetEnvironmentV1ByCrnOK {
	return &GetEnvironmentV1ByCrnOK{}
}

/*
GetEnvironmentV1ByCrnOK handles this case with default header values.

successful operation
*/
type GetEnvironmentV1ByCrnOK struct {
	Payload *model.DetailedEnvironmentV1Response
}

func (o *GetEnvironmentV1ByCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/crn/{crn}][%d] getEnvironmentV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedEnvironmentV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
