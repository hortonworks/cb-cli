// Code generated by go-swagger; DO NOT EDIT.

package flow_public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// HasFlowRunningByChainIDAndResourceCrnReader is a Reader for the HasFlowRunningByChainIDAndResourceCrn structure.
type HasFlowRunningByChainIDAndResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasFlowRunningByChainIDAndResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasFlowRunningByChainIDAndResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasFlowRunningByChainIDAndResourceCrnOK creates a HasFlowRunningByChainIDAndResourceCrnOK with default headers values
func NewHasFlowRunningByChainIDAndResourceCrnOK() *HasFlowRunningByChainIDAndResourceCrnOK {
	return &HasFlowRunningByChainIDAndResourceCrnOK{}
}

/*
HasFlowRunningByChainIDAndResourceCrnOK handles this case with default header values.

successful operation
*/
type HasFlowRunningByChainIDAndResourceCrnOK struct {
	Payload *model.FlowCheckResponse
}

func (o *HasFlowRunningByChainIDAndResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /flow-public/check/chainId/{chainId}][%d] hasFlowRunningByChainIdAndResourceCrnOK  %+v", 200, o.Payload)
}

func (o *HasFlowRunningByChainIDAndResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
