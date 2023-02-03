// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// VerticalScalingByNameReader is a Reader for the VerticalScalingByName structure.
type VerticalScalingByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerticalScalingByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVerticalScalingByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVerticalScalingByNameOK creates a VerticalScalingByNameOK with default headers values
func NewVerticalScalingByNameOK() *VerticalScalingByNameOK {
	return &VerticalScalingByNameOK{}
}

/*
VerticalScalingByNameOK handles this case with default header values.

successful operation
*/
type VerticalScalingByNameOK struct {
	Payload *model.FlowIdentifier
}

func (o *VerticalScalingByNameOK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/name/{name}/vertical_scaling][%d] verticalScalingByNameOK  %+v", 200, o.Payload)
}

func (o *VerticalScalingByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
