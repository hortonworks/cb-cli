// Code generated by go-swagger; DO NOT EDIT.

package v1dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddDNSARecordV1InternalReader is a Reader for the AddDNSARecordV1Internal structure.
type AddDNSARecordV1InternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDNSARecordV1InternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewAddDNSARecordV1InternalDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewAddDNSARecordV1InternalDefault creates a AddDNSARecordV1InternalDefault with default headers values
func NewAddDNSARecordV1InternalDefault(code int) *AddDNSARecordV1InternalDefault {
	return &AddDNSARecordV1InternalDefault{
		_statusCode: code,
	}
}

/*
AddDNSARecordV1InternalDefault handles this case with default header values.

successful operation
*/
type AddDNSARecordV1InternalDefault struct {
	_statusCode int
}

// Code gets the status code for the add Dns a record v1 internal default response
func (o *AddDNSARecordV1InternalDefault) Code() int {
	return o._statusCode
}

func (o *AddDNSARecordV1InternalDefault) Error() string {
	return fmt.Sprintf("[POST /v1/dns/record/a/internal][%d] addDnsARecordV1Internal default ", o._statusCode)
}

func (o *AddDNSARecordV1InternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
