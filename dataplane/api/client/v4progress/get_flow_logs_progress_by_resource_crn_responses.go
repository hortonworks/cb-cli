// Code generated by go-swagger; DO NOT EDIT.

package v4progress

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetFlowLogsProgressByResourceCrnReader is a Reader for the GetFlowLogsProgressByResourceCrn structure.
type GetFlowLogsProgressByResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLogsProgressByResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowLogsProgressByResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowLogsProgressByResourceCrnOK creates a GetFlowLogsProgressByResourceCrnOK with default headers values
func NewGetFlowLogsProgressByResourceCrnOK() *GetFlowLogsProgressByResourceCrnOK {
	return &GetFlowLogsProgressByResourceCrnOK{}
}

/*
GetFlowLogsProgressByResourceCrnOK handles this case with default header values.

successful operation
*/
type GetFlowLogsProgressByResourceCrnOK struct {
	Payload []*model.FlowProgressResponse
}

func (o *GetFlowLogsProgressByResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /v4/progress/resource/crn/{resourceCrn}][%d] getFlowLogsProgressByResourceCrnOK  %+v", 200, o.Payload)
}

func (o *GetFlowLogsProgressByResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
