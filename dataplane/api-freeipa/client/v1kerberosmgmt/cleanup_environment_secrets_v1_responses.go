// Code generated by go-swagger; DO NOT EDIT.

package v1kerberosmgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CleanupEnvironmentSecretsV1Reader is a Reader for the CleanupEnvironmentSecretsV1 structure.
type CleanupEnvironmentSecretsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CleanupEnvironmentSecretsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewCleanupEnvironmentSecretsV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewCleanupEnvironmentSecretsV1Default creates a CleanupEnvironmentSecretsV1Default with default headers values
func NewCleanupEnvironmentSecretsV1Default(code int) *CleanupEnvironmentSecretsV1Default {
	return &CleanupEnvironmentSecretsV1Default{
		_statusCode: code,
	}
}

/*
CleanupEnvironmentSecretsV1Default handles this case with default header values.

successful operation
*/
type CleanupEnvironmentSecretsV1Default struct {
	_statusCode int
}

// Code gets the status code for the cleanup environment secrets v1 default response
func (o *CleanupEnvironmentSecretsV1Default) Code() int {
	return o._statusCode
}

func (o *CleanupEnvironmentSecretsV1Default) Error() string {
	return fmt.Sprintf("[DELETE /v1/kerberosmgmt/cleanupEnvironmentSecrets][%d] cleanupEnvironmentSecretsV1 default ", o._statusCode)
}

func (o *CleanupEnvironmentSecretsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
