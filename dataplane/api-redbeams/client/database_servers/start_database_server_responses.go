// Code generated by go-swagger; DO NOT EDIT.

package database_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StartDatabaseServerReader is a Reader for the StartDatabaseServer structure.
type StartDatabaseServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDatabaseServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStartDatabaseServerDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStartDatabaseServerDefault creates a StartDatabaseServerDefault with default headers values
func NewStartDatabaseServerDefault(code int) *StartDatabaseServerDefault {
	return &StartDatabaseServerDefault{
		_statusCode: code,
	}
}

/*
StartDatabaseServerDefault handles this case with default header values.

successful operation
*/
type StartDatabaseServerDefault struct {
	_statusCode int
}

// Code gets the status code for the start database server default response
func (o *StartDatabaseServerDefault) Code() int {
	return o._statusCode
}

func (o *StartDatabaseServerDefault) Error() string {
	return fmt.Sprintf("[PUT /v4/databaseservers/{crn}/start][%d] startDatabaseServer default ", o._statusCode)
}

func (o *StartDatabaseServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
