// Code generated by go-swagger; DO NOT EDIT.

package database_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// ReleaseManagedDatabaseServerReader is a Reader for the ReleaseManagedDatabaseServer structure.
type ReleaseManagedDatabaseServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseManagedDatabaseServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReleaseManagedDatabaseServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReleaseManagedDatabaseServerOK creates a ReleaseManagedDatabaseServerOK with default headers values
func NewReleaseManagedDatabaseServerOK() *ReleaseManagedDatabaseServerOK {
	return &ReleaseManagedDatabaseServerOK{}
}

/*
ReleaseManagedDatabaseServerOK handles this case with default header values.

successful operation
*/
type ReleaseManagedDatabaseServerOK struct {
	Payload *model.DatabaseServerV4Response
}

func (o *ReleaseManagedDatabaseServerOK) Error() string {
	return fmt.Sprintf("[PUT /v4/databaseservers/{crn}/release][%d] releaseManagedDatabaseServerOK  %+v", 200, o.Payload)
}

func (o *ReleaseManagedDatabaseServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseServerV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
