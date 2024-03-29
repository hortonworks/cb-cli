// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CodeGrantFlowOnExistingCredentialV1Reader is a Reader for the CodeGrantFlowOnExistingCredentialV1 structure.
type CodeGrantFlowOnExistingCredentialV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CodeGrantFlowOnExistingCredentialV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewCodeGrantFlowOnExistingCredentialV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewCodeGrantFlowOnExistingCredentialV1Default creates a CodeGrantFlowOnExistingCredentialV1Default with default headers values
func NewCodeGrantFlowOnExistingCredentialV1Default(code int) *CodeGrantFlowOnExistingCredentialV1Default {
	return &CodeGrantFlowOnExistingCredentialV1Default{
		_statusCode: code,
	}
}

/*
CodeGrantFlowOnExistingCredentialV1Default handles this case with default header values.

successful operation
*/
type CodeGrantFlowOnExistingCredentialV1Default struct {
	_statusCode int
}

// Code gets the status code for the code grant flow on existing credential v1 default response
func (o *CodeGrantFlowOnExistingCredentialV1Default) Code() int {
	return o._statusCode
}

func (o *CodeGrantFlowOnExistingCredentialV1Default) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/code_grant_flow/init/{name}][%d] codeGrantFlowOnExistingCredentialV1 default ", o._statusCode)
}

func (o *CodeGrantFlowOnExistingCredentialV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
