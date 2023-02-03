// Code generated by go-swagger; DO NOT EDIT.

package v1internaldistrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDistroXInstancesInternalV1ByCrnReader is a Reader for the GetDistroXInstancesInternalV1ByCrn structure.
type GetDistroXInstancesInternalV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroXInstancesInternalV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroXInstancesInternalV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroXInstancesInternalV1ByCrnOK creates a GetDistroXInstancesInternalV1ByCrnOK with default headers values
func NewGetDistroXInstancesInternalV1ByCrnOK() *GetDistroXInstancesInternalV1ByCrnOK {
	return &GetDistroXInstancesInternalV1ByCrnOK{}
}

/*
GetDistroXInstancesInternalV1ByCrnOK handles this case with default header values.

successful operation
*/
type GetDistroXInstancesInternalV1ByCrnOK struct {
	Payload *model.StackInstancesV4Responses
}

func (o *GetDistroXInstancesInternalV1ByCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/internal/distrox/crn/{crn}/instances][%d] getDistroXInstancesInternalV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *GetDistroXInstancesInternalV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackInstancesV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
