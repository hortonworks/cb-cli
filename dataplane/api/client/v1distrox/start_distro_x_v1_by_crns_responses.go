// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StartDistroXV1ByCrnsReader is a Reader for the StartDistroXV1ByCrns structure.
type StartDistroXV1ByCrnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDistroXV1ByCrnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStartDistroXV1ByCrnsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStartDistroXV1ByCrnsDefault creates a StartDistroXV1ByCrnsDefault with default headers values
func NewStartDistroXV1ByCrnsDefault(code int) *StartDistroXV1ByCrnsDefault {
	return &StartDistroXV1ByCrnsDefault{
		_statusCode: code,
	}
}

/*
StartDistroXV1ByCrnsDefault handles this case with default header values.

successful operation
*/
type StartDistroXV1ByCrnsDefault struct {
	_statusCode int
}

// Code gets the status code for the start distro x v1 by crns default response
func (o *StartDistroXV1ByCrnsDefault) Code() int {
	return o._statusCode
}

func (o *StartDistroXV1ByCrnsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/crn/start][%d] startDistroXV1ByCrns default ", o._statusCode)
}

func (o *StartDistroXV1ByCrnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
