// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// DeleteCredentialsV1Reader is a Reader for the DeleteCredentialsV1 structure.
type DeleteCredentialsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCredentialsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCredentialsV1OK creates a DeleteCredentialsV1OK with default headers values
func NewDeleteCredentialsV1OK() *DeleteCredentialsV1OK {
	return &DeleteCredentialsV1OK{}
}

/*
DeleteCredentialsV1OK handles this case with default header values.

successful operation
*/
type DeleteCredentialsV1OK struct {
	Payload *model.CredentialV1Responses
}

func (o *DeleteCredentialsV1OK) Error() string {
	return fmt.Sprintf("[DELETE /v1/credentials][%d] deleteCredentialsV1OK  %+v", 200, o.Payload)
}

func (o *DeleteCredentialsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialV1Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
