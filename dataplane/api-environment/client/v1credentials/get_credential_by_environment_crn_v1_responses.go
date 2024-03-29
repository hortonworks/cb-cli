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

// GetCredentialByEnvironmentCrnV1Reader is a Reader for the GetCredentialByEnvironmentCrnV1 structure.
type GetCredentialByEnvironmentCrnV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialByEnvironmentCrnV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCredentialByEnvironmentCrnV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCredentialByEnvironmentCrnV1OK creates a GetCredentialByEnvironmentCrnV1OK with default headers values
func NewGetCredentialByEnvironmentCrnV1OK() *GetCredentialByEnvironmentCrnV1OK {
	return &GetCredentialByEnvironmentCrnV1OK{}
}

/*
GetCredentialByEnvironmentCrnV1OK handles this case with default header values.

successful operation
*/
type GetCredentialByEnvironmentCrnV1OK struct {
	Payload *model.CredentialV1Response
}

func (o *GetCredentialByEnvironmentCrnV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/environment/crn/{environmentCrn}][%d] getCredentialByEnvironmentCrnV1OK  %+v", 200, o.Payload)
}

func (o *GetCredentialByEnvironmentCrnV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
