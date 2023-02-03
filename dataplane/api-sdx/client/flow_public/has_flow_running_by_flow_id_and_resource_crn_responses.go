// Code generated by go-swagger; DO NOT EDIT.

package flow_public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// HasFlowRunningByFlowIDAndResourceCrnReader is a Reader for the HasFlowRunningByFlowIDAndResourceCrn structure.
type HasFlowRunningByFlowIDAndResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasFlowRunningByFlowIDAndResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasFlowRunningByFlowIDAndResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasFlowRunningByFlowIDAndResourceCrnOK creates a HasFlowRunningByFlowIDAndResourceCrnOK with default headers values
func NewHasFlowRunningByFlowIDAndResourceCrnOK() *HasFlowRunningByFlowIDAndResourceCrnOK {
	return &HasFlowRunningByFlowIDAndResourceCrnOK{}
}

/*
HasFlowRunningByFlowIDAndResourceCrnOK handles this case with default header values.

successful operation
*/
type HasFlowRunningByFlowIDAndResourceCrnOK struct {
	Payload *model.FlowCheckResponse
}

func (o *HasFlowRunningByFlowIDAndResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /flow-public/check/flowId/{flowId}][%d] hasFlowRunningByFlowIdAndResourceCrnOK  %+v", 200, o.Payload)
}

func (o *HasFlowRunningByFlowIDAndResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
