// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SyncDistroXV1ByNameReader is a Reader for the SyncDistroXV1ByName structure.
type SyncDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewSyncDistroXV1ByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewSyncDistroXV1ByNameDefault creates a SyncDistroXV1ByNameDefault with default headers values
func NewSyncDistroXV1ByNameDefault(code int) *SyncDistroXV1ByNameDefault {
	return &SyncDistroXV1ByNameDefault{
		_statusCode: code,
	}
}

/*
SyncDistroXV1ByNameDefault handles this case with default header values.

successful operation
*/
type SyncDistroXV1ByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the sync distro x v1 by name default response
func (o *SyncDistroXV1ByNameDefault) Code() int {
	return o._statusCode
}

func (o *SyncDistroXV1ByNameDefault) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/name/{name}/sync][%d] syncDistroXV1ByName default ", o._statusCode)
}

func (o *SyncDistroXV1ByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
