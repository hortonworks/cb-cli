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

// GetLastFlowLogProgressByResourceCrnReader is a Reader for the GetLastFlowLogProgressByResourceCrn structure.
type GetLastFlowLogProgressByResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLastFlowLogProgressByResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLastFlowLogProgressByResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLastFlowLogProgressByResourceCrnOK creates a GetLastFlowLogProgressByResourceCrnOK with default headers values
func NewGetLastFlowLogProgressByResourceCrnOK() *GetLastFlowLogProgressByResourceCrnOK {
	return &GetLastFlowLogProgressByResourceCrnOK{}
}

/*
GetLastFlowLogProgressByResourceCrnOK handles this case with default header values.

successful operation
*/
type GetLastFlowLogProgressByResourceCrnOK struct {
	Payload *model.FlowProgressResponse
}

func (o *GetLastFlowLogProgressByResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /v4/progress/resource/crn/{resourceCrn}/last][%d] getLastFlowLogProgressByResourceCrnOK  %+v", 200, o.Payload)
}

func (o *GetLastFlowLogProgressByResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowProgressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
