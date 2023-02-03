// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// CreateE2ETestBindUserV1Reader is a Reader for the CreateE2ETestBindUserV1 structure.
type CreateE2ETestBindUserV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateE2ETestBindUserV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateE2ETestBindUserV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateE2ETestBindUserV1OK creates a CreateE2ETestBindUserV1OK with default headers values
func NewCreateE2ETestBindUserV1OK() *CreateE2ETestBindUserV1OK {
	return &CreateE2ETestBindUserV1OK{}
}

/*
CreateE2ETestBindUserV1OK handles this case with default header values.

successful operation
*/
type CreateE2ETestBindUserV1OK struct {
	Payload *model.OperationV1Status
}

func (o *CreateE2ETestBindUserV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/binduser/create/e2etest][%d] createE2ETestBindUserV1OK  %+v", 200, o.Payload)
}

func (o *CreateE2ETestBindUserV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
