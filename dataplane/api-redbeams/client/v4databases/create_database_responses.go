// Code generated by go-swagger; DO NOT EDIT.

package v4databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// CreateDatabaseReader is a Reader for the CreateDatabase structure.
type CreateDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDatabaseOK creates a CreateDatabaseOK with default headers values
func NewCreateDatabaseOK() *CreateDatabaseOK {
	return &CreateDatabaseOK{}
}

/*CreateDatabaseOK handles this case with default header values.

successful operation
*/
type CreateDatabaseOK struct {
	Payload *model.DatabaseV4Response
}

func (o *CreateDatabaseOK) Error() string {
	return fmt.Sprintf("[POST /v4/databases][%d] createDatabaseOK  %+v", 200, o.Payload)
}

func (o *CreateDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
