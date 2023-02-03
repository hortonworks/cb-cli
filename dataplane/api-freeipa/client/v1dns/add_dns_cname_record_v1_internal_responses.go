// Code generated by go-swagger; DO NOT EDIT.

package v1dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddDNSCnameRecordV1InternalReader is a Reader for the AddDNSCnameRecordV1Internal structure.
type AddDNSCnameRecordV1InternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDNSCnameRecordV1InternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewAddDNSCnameRecordV1InternalDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewAddDNSCnameRecordV1InternalDefault creates a AddDNSCnameRecordV1InternalDefault with default headers values
func NewAddDNSCnameRecordV1InternalDefault(code int) *AddDNSCnameRecordV1InternalDefault {
	return &AddDNSCnameRecordV1InternalDefault{
		_statusCode: code,
	}
}

/*
AddDNSCnameRecordV1InternalDefault handles this case with default header values.

successful operation
*/
type AddDNSCnameRecordV1InternalDefault struct {
	_statusCode int
}

// Code gets the status code for the add Dns cname record v1 internal default response
func (o *AddDNSCnameRecordV1InternalDefault) Code() int {
	return o._statusCode
}

func (o *AddDNSCnameRecordV1InternalDefault) Error() string {
	return fmt.Sprintf("[POST /v1/dns/record/cname/internal][%d] addDnsCnameRecordV1Internal default ", o._statusCode)
}

func (o *AddDNSCnameRecordV1InternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
