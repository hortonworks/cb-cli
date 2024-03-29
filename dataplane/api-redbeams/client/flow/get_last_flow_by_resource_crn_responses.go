// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// GetLastFlowByResourceCrnReader is a Reader for the GetLastFlowByResourceCrn structure.
type GetLastFlowByResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLastFlowByResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLastFlowByResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLastFlowByResourceCrnOK creates a GetLastFlowByResourceCrnOK with default headers values
func NewGetLastFlowByResourceCrnOK() *GetLastFlowByResourceCrnOK {
	return &GetLastFlowByResourceCrnOK{}
}

/*
GetLastFlowByResourceCrnOK handles this case with default header values.

successful operation
*/
type GetLastFlowByResourceCrnOK struct {
	Payload *model.FlowLogResponse
}

func (o *GetLastFlowByResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /flow/logs/resource/crn/{resourceCrn}/last][%d] getLastFlowByResourceCrnOK  %+v", 200, o.Payload)
}

func (o *GetLastFlowByResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowLogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
