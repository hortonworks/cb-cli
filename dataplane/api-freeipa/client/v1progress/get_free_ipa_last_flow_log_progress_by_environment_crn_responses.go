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

// GetFreeIpaLastFlowLogProgressByEnvironmentCrnReader is a Reader for the GetFreeIpaLastFlowLogProgressByEnvironmentCrn structure.
type GetFreeIpaLastFlowLogProgressByEnvironmentCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFreeIpaLastFlowLogProgressByEnvironmentCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFreeIpaLastFlowLogProgressByEnvironmentCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFreeIpaLastFlowLogProgressByEnvironmentCrnOK creates a GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK with default headers values
func NewGetFreeIpaLastFlowLogProgressByEnvironmentCrnOK() *GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK {
	return &GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK{}
}

/*GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK handles this case with default header values.

successful operation
*/
type GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK struct {
	Payload *model.FlowProgressResponse
}

func (o *GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/progress/resource/crn/{environmentCrn}/last][%d] getFreeIpaLastFlowLogProgressByEnvironmentCrnOK  %+v", 200, o.Payload)
}

func (o *GetFreeIpaLastFlowLogProgressByEnvironmentCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowProgressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
