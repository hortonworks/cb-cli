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

// DeleteMultipleDatabaseServersReader is a Reader for the DeleteMultipleDatabaseServers structure.
type DeleteMultipleDatabaseServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMultipleDatabaseServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMultipleDatabaseServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMultipleDatabaseServersOK creates a DeleteMultipleDatabaseServersOK with default headers values
func NewDeleteMultipleDatabaseServersOK() *DeleteMultipleDatabaseServersOK {
	return &DeleteMultipleDatabaseServersOK{}
}

/*DeleteMultipleDatabaseServersOK handles this case with default header values.

successful operation
*/
type DeleteMultipleDatabaseServersOK struct {
	Payload *model.DatabaseServerV4Responses
}

func (o *DeleteMultipleDatabaseServersOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/databaseservers][%d] deleteMultipleDatabaseServersOK  %+v", 200, o.Payload)
}

func (o *DeleteMultipleDatabaseServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseServerV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
