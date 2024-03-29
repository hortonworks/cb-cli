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

// RegisterDatabaseServerReader is a Reader for the RegisterDatabaseServer structure.
type RegisterDatabaseServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterDatabaseServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegisterDatabaseServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterDatabaseServerOK creates a RegisterDatabaseServerOK with default headers values
func NewRegisterDatabaseServerOK() *RegisterDatabaseServerOK {
	return &RegisterDatabaseServerOK{}
}

/*
RegisterDatabaseServerOK handles this case with default header values.

successful operation
*/
type RegisterDatabaseServerOK struct {
	Payload *model.DatabaseServerV4Response
}

func (o *RegisterDatabaseServerOK) Error() string {
	return fmt.Sprintf("[POST /v4/databaseservers/register][%d] registerDatabaseServerOK  %+v", 200, o.Payload)
}

func (o *RegisterDatabaseServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseServerV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
