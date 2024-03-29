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

// CheckRightInAccountReader is a Reader for the CheckRightInAccount structure.
type CheckRightInAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckRightInAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckRightInAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckRightInAccountOK creates a CheckRightInAccountOK with default headers values
func NewCheckRightInAccountOK() *CheckRightInAccountOK {
	return &CheckRightInAccountOK{}
}

/*
CheckRightInAccountOK handles this case with default header values.

successful operation
*/
type CheckRightInAccountOK struct {
	Payload *model.CheckRightV4Response
}

func (o *CheckRightInAccountOK) Error() string {
	return fmt.Sprintf("[POST /v4/utils/check_right][%d] checkRightInAccountOK  %+v", 200, o.Payload)
}

func (o *CheckRightInAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CheckRightV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
