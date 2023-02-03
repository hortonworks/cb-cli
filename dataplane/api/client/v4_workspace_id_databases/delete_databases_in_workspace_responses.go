// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteDatabasesInWorkspaceReader is a Reader for the DeleteDatabasesInWorkspace structure.
type DeleteDatabasesInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatabasesInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDatabasesInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDatabasesInWorkspaceOK creates a DeleteDatabasesInWorkspaceOK with default headers values
func NewDeleteDatabasesInWorkspaceOK() *DeleteDatabasesInWorkspaceOK {
	return &DeleteDatabasesInWorkspaceOK{}
}

/*
DeleteDatabasesInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteDatabasesInWorkspaceOK struct {
	Payload *model.DatabaseV4Responses
}

func (o *DeleteDatabasesInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/databases][%d] deleteDatabasesInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteDatabasesInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
