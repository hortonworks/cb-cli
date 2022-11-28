// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// VerticalScalingByCrnReader is a Reader for the VerticalScalingByCrn structure.
type VerticalScalingByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerticalScalingByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVerticalScalingByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVerticalScalingByCrnOK creates a VerticalScalingByCrnOK with default headers values
func NewVerticalScalingByCrnOK() *VerticalScalingByCrnOK {
	return &VerticalScalingByCrnOK{}
}

/*VerticalScalingByCrnOK handles this case with default header values.

successful operation
*/
type VerticalScalingByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *VerticalScalingByCrnOK) Error() string {
	return fmt.Sprintf("[PUT /v1/env/crn/{crn}/vertical_scaling][%d] verticalScalingByCrnOK  %+v", 200, o.Payload)
}

func (o *VerticalScalingByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
