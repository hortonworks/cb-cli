// Code generated by go-swagger; DO NOT EDIT.

package v4databaseservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// CreateDatabaseOnServerReader is a Reader for the CreateDatabaseOnServer structure.
type CreateDatabaseOnServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatabaseOnServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDatabaseOnServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDatabaseOnServerOK creates a CreateDatabaseOnServerOK with default headers values
func NewCreateDatabaseOnServerOK() *CreateDatabaseOnServerOK {
	return &CreateDatabaseOnServerOK{}
}

/*CreateDatabaseOnServerOK handles this case with default header values.

successful operation
*/
type CreateDatabaseOnServerOK struct {
	Payload *model.CreateDatabaseV4Response
}

func (o *CreateDatabaseOnServerOK) Error() string {
	return fmt.Sprintf("[POST /v4/databaseservers/createDatabase][%d] createDatabaseOnServerOK  %+v", 200, o.Payload)
}

func (o *CreateDatabaseOnServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CreateDatabaseV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
