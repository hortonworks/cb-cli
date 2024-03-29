// Code generated by go-swagger; DO NOT EDIT.

package v1kerberos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteKerberosConfigForEnvironmentReader is a Reader for the DeleteKerberosConfigForEnvironment structure.
type DeleteKerberosConfigForEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKerberosConfigForEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteKerberosConfigForEnvironmentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteKerberosConfigForEnvironmentDefault creates a DeleteKerberosConfigForEnvironmentDefault with default headers values
func NewDeleteKerberosConfigForEnvironmentDefault(code int) *DeleteKerberosConfigForEnvironmentDefault {
	return &DeleteKerberosConfigForEnvironmentDefault{
		_statusCode: code,
	}
}

/*
DeleteKerberosConfigForEnvironmentDefault handles this case with default header values.

successful operation
*/
type DeleteKerberosConfigForEnvironmentDefault struct {
	_statusCode int
}

// Code gets the status code for the delete kerberos config for environment default response
func (o *DeleteKerberosConfigForEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteKerberosConfigForEnvironmentDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/kerberos][%d] deleteKerberosConfigForEnvironment default ", o._statusCode)
}

func (o *DeleteKerberosConfigForEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
