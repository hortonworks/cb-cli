// Code generated by go-swagger; DO NOT EDIT.

package v1dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDNSZoneBySubnetV1Reader is a Reader for the DeleteDNSZoneBySubnetV1 structure.
type DeleteDNSZoneBySubnetV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDNSZoneBySubnetV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteDNSZoneBySubnetV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteDNSZoneBySubnetV1Default creates a DeleteDNSZoneBySubnetV1Default with default headers values
func NewDeleteDNSZoneBySubnetV1Default(code int) *DeleteDNSZoneBySubnetV1Default {
	return &DeleteDNSZoneBySubnetV1Default{
		_statusCode: code,
	}
}

/*
DeleteDNSZoneBySubnetV1Default handles this case with default header values.

successful operation
*/
type DeleteDNSZoneBySubnetV1Default struct {
	_statusCode int
}

// Code gets the status code for the delete Dns zone by subnet v1 default response
func (o *DeleteDNSZoneBySubnetV1Default) Code() int {
	return o._statusCode
}

func (o *DeleteDNSZoneBySubnetV1Default) Error() string {
	return fmt.Sprintf("[DELETE /v1/dns/zone/cidr][%d] deleteDnsZoneBySubnetV1 default ", o._statusCode)
}

func (o *DeleteDNSZoneBySubnetV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
