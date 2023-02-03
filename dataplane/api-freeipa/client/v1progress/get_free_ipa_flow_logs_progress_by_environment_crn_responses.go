// Code generated by go-swagger; DO NOT EDIT.

package v1progress

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetFreeIpaFlowLogsProgressByEnvironmentCrnReader is a Reader for the GetFreeIpaFlowLogsProgressByEnvironmentCrn structure.
type GetFreeIpaFlowLogsProgressByEnvironmentCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFreeIpaFlowLogsProgressByEnvironmentCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFreeIpaFlowLogsProgressByEnvironmentCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFreeIpaFlowLogsProgressByEnvironmentCrnOK creates a GetFreeIpaFlowLogsProgressByEnvironmentCrnOK with default headers values
func NewGetFreeIpaFlowLogsProgressByEnvironmentCrnOK() *GetFreeIpaFlowLogsProgressByEnvironmentCrnOK {
	return &GetFreeIpaFlowLogsProgressByEnvironmentCrnOK{}
}

/*
GetFreeIpaFlowLogsProgressByEnvironmentCrnOK handles this case with default header values.

successful operation
*/
type GetFreeIpaFlowLogsProgressByEnvironmentCrnOK struct {
	Payload []*model.FlowProgressResponse
}

func (o *GetFreeIpaFlowLogsProgressByEnvironmentCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/progress/resource/crn/{environmentCrn}][%d] getFreeIpaFlowLogsProgressByEnvironmentCrnOK  %+v", 200, o.Payload)
}

func (o *GetFreeIpaFlowLogsProgressByEnvironmentCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
