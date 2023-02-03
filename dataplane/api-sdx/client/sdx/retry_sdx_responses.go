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

// RetrySdxReader is a Reader for the RetrySdx structure.
type RetrySdxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrySdxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrySdxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrySdxOK creates a RetrySdxOK with default headers values
func NewRetrySdxOK() *RetrySdxOK {
	return &RetrySdxOK{}
}

/*
RetrySdxOK handles this case with default header values.

successful operation
*/
type RetrySdxOK struct {
	Payload *model.FlowIdentifier
}

func (o *RetrySdxOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/retry][%d] retrySdxOK  %+v", 200, o.Payload)
}

func (o *RetrySdxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
