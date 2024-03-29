// Code generated by go-swagger; DO NOT EDIT.

package v1credentialsaudit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// CreateAuditCredentialV1Reader is a Reader for the CreateAuditCredentialV1 structure.
type CreateAuditCredentialV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuditCredentialV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAuditCredentialV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAuditCredentialV1OK creates a CreateAuditCredentialV1OK with default headers values
func NewCreateAuditCredentialV1OK() *CreateAuditCredentialV1OK {
	return &CreateAuditCredentialV1OK{}
}

/*
CreateAuditCredentialV1OK handles this case with default header values.

successful operation
*/
type CreateAuditCredentialV1OK struct {
	Payload *model.CredentialV1Response
}

func (o *CreateAuditCredentialV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/credentials/audit][%d] createAuditCredentialV1OK  %+v", 200, o.Payload)
}

func (o *CreateAuditCredentialV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
