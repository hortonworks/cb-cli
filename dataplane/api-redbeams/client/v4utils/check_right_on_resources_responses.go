// Code generated by go-swagger; DO NOT EDIT.

package v4utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// CheckRightOnResourcesReader is a Reader for the CheckRightOnResources structure.
type CheckRightOnResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckRightOnResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckRightOnResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckRightOnResourcesOK creates a CheckRightOnResourcesOK with default headers values
func NewCheckRightOnResourcesOK() *CheckRightOnResourcesOK {
	return &CheckRightOnResourcesOK{}
}

/*
CheckRightOnResourcesOK handles this case with default header values.

successful operation
*/
type CheckRightOnResourcesOK struct {
	Payload *model.CheckRightOnResourcesV4Response
}

func (o *CheckRightOnResourcesOK) Error() string {
	return fmt.Sprintf("[POST /v4/utils/check_right_on_resources][%d] checkRightOnResourcesOK  %+v", 200, o.Payload)
}

func (o *CheckRightOnResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CheckRightOnResourcesV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
