// Code generated by go-swagger; DO NOT EDIT.

package v1operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetOperationProgressByEnvironmentCrnReader is a Reader for the GetOperationProgressByEnvironmentCrn structure.
type GetOperationProgressByEnvironmentCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperationProgressByEnvironmentCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperationProgressByEnvironmentCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOperationProgressByEnvironmentCrnOK creates a GetOperationProgressByEnvironmentCrnOK with default headers values
func NewGetOperationProgressByEnvironmentCrnOK() *GetOperationProgressByEnvironmentCrnOK {
	return &GetOperationProgressByEnvironmentCrnOK{}
}

/*
GetOperationProgressByEnvironmentCrnOK handles this case with default header values.

successful operation
*/
type GetOperationProgressByEnvironmentCrnOK struct {
	Payload *model.OperationView
}

func (o *GetOperationProgressByEnvironmentCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/operation/resource/crn/{environmentCrn}][%d] getOperationProgressByEnvironmentCrnOK  %+v", 200, o.Payload)
}

func (o *GetOperationProgressByEnvironmentCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OperationView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
