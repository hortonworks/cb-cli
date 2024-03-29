// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDistroXV1ByNameReader is a Reader for the DeleteDistroXV1ByName structure.
type DeleteDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteDistroXV1ByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteDistroXV1ByNameDefault creates a DeleteDistroXV1ByNameDefault with default headers values
func NewDeleteDistroXV1ByNameDefault(code int) *DeleteDistroXV1ByNameDefault {
	return &DeleteDistroXV1ByNameDefault{
		_statusCode: code,
	}
}

/*
DeleteDistroXV1ByNameDefault handles this case with default header values.

successful operation
*/
type DeleteDistroXV1ByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete distro x v1 by name default response
func (o *DeleteDistroXV1ByNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDistroXV1ByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/distrox/name/{name}][%d] deleteDistroXV1ByName default ", o._statusCode)
}

func (o *DeleteDistroXV1ByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
