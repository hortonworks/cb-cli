// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DetachChildEnvironmentV1Reader is a Reader for the DetachChildEnvironmentV1 structure.
type DetachChildEnvironmentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachChildEnvironmentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDetachChildEnvironmentV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDetachChildEnvironmentV1Default creates a DetachChildEnvironmentV1Default with default headers values
func NewDetachChildEnvironmentV1Default(code int) *DetachChildEnvironmentV1Default {
	return &DetachChildEnvironmentV1Default{
		_statusCode: code,
	}
}

/*
DetachChildEnvironmentV1Default handles this case with default header values.

successful operation
*/
type DetachChildEnvironmentV1Default struct {
	_statusCode int
}

// Code gets the status code for the detach child environment v1 default response
func (o *DetachChildEnvironmentV1Default) Code() int {
	return o._statusCode
}

func (o *DetachChildEnvironmentV1Default) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/detach_child_environment][%d] detachChildEnvironmentV1 default ", o._statusCode)
}

func (o *DetachChildEnvironmentV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
