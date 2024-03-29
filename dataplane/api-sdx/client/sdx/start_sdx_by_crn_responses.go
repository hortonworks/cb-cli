// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// StartSdxByCrnReader is a Reader for the StartSdxByCrn structure.
type StartSdxByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartSdxByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartSdxByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartSdxByCrnOK creates a StartSdxByCrnOK with default headers values
func NewStartSdxByCrnOK() *StartSdxByCrnOK {
	return &StartSdxByCrnOK{}
}

/*
StartSdxByCrnOK handles this case with default header values.

successful operation
*/
type StartSdxByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *StartSdxByCrnOK) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/start][%d] startSdxByCrnOK  %+v", 200, o.Payload)
}

func (o *StartSdxByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
