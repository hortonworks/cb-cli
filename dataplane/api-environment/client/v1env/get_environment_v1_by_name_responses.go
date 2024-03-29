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

// GetEnvironmentV1ByNameReader is a Reader for the GetEnvironmentV1ByName structure.
type GetEnvironmentV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEnvironmentV1ByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnvironmentV1ByNameOK creates a GetEnvironmentV1ByNameOK with default headers values
func NewGetEnvironmentV1ByNameOK() *GetEnvironmentV1ByNameOK {
	return &GetEnvironmentV1ByNameOK{}
}

/*
GetEnvironmentV1ByNameOK handles this case with default header values.

successful operation
*/
type GetEnvironmentV1ByNameOK struct {
	Payload *model.DetailedEnvironmentV1Response
}

func (o *GetEnvironmentV1ByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/name/{name}][%d] getEnvironmentV1ByNameOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentV1ByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedEnvironmentV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
