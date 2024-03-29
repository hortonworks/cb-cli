// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetStackForAutoscaleReader is a Reader for the GetStackForAutoscale structure.
type GetStackForAutoscaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackForAutoscaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStackForAutoscaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStackForAutoscaleOK creates a GetStackForAutoscaleOK with default headers values
func NewGetStackForAutoscaleOK() *GetStackForAutoscaleOK {
	return &GetStackForAutoscaleOK{}
}

/*
GetStackForAutoscaleOK handles this case with default header values.

successful operation
*/
type GetStackForAutoscaleOK struct {
	Payload *model.StackV4Response
}

func (o *GetStackForAutoscaleOK) Error() string {
	return fmt.Sprintf("[GET /autoscale/stack/crn/{crn}][%d] getStackForAutoscaleOK  %+v", 200, o.Payload)
}

func (o *GetStackForAutoscaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
