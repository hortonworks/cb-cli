// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StartDistroXV1ByNamesReader is a Reader for the StartDistroXV1ByNames structure.
type StartDistroXV1ByNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDistroXV1ByNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStartDistroXV1ByNamesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStartDistroXV1ByNamesDefault creates a StartDistroXV1ByNamesDefault with default headers values
func NewStartDistroXV1ByNamesDefault(code int) *StartDistroXV1ByNamesDefault {
	return &StartDistroXV1ByNamesDefault{
		_statusCode: code,
	}
}

/*
StartDistroXV1ByNamesDefault handles this case with default header values.

successful operation
*/
type StartDistroXV1ByNamesDefault struct {
	_statusCode int
}

// Code gets the status code for the start distro x v1 by names default response
func (o *StartDistroXV1ByNamesDefault) Code() int {
	return o._statusCode
}

func (o *StartDistroXV1ByNamesDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/name/start][%d] startDistroXV1ByNames default ", o._statusCode)
}

func (o *StartDistroXV1ByNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
