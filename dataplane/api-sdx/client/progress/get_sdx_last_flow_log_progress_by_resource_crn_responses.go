// Code generated by go-swagger; DO NOT EDIT.

package progress

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// GetSdxLastFlowLogProgressByResourceCrnReader is a Reader for the GetSdxLastFlowLogProgressByResourceCrn structure.
type GetSdxLastFlowLogProgressByResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSdxLastFlowLogProgressByResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSdxLastFlowLogProgressByResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSdxLastFlowLogProgressByResourceCrnOK creates a GetSdxLastFlowLogProgressByResourceCrnOK with default headers values
func NewGetSdxLastFlowLogProgressByResourceCrnOK() *GetSdxLastFlowLogProgressByResourceCrnOK {
	return &GetSdxLastFlowLogProgressByResourceCrnOK{}
}

/*
GetSdxLastFlowLogProgressByResourceCrnOK handles this case with default header values.

successful operation
*/
type GetSdxLastFlowLogProgressByResourceCrnOK struct {
	Payload *model.SdxProgressResponse
}

func (o *GetSdxLastFlowLogProgressByResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /progress/resource/crn/{resourceCrn}/last][%d] getSdxLastFlowLogProgressByResourceCrnOK  %+v", 200, o.Payload)
}

func (o *GetSdxLastFlowLogProgressByResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxProgressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
