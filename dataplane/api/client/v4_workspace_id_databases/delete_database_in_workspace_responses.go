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

// DeleteDatabaseInWorkspaceReader is a Reader for the DeleteDatabaseInWorkspace structure.
type DeleteDatabaseInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatabaseInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDatabaseInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDatabaseInWorkspaceOK creates a DeleteDatabaseInWorkspaceOK with default headers values
func NewDeleteDatabaseInWorkspaceOK() *DeleteDatabaseInWorkspaceOK {
	return &DeleteDatabaseInWorkspaceOK{}
}

/*
DeleteDatabaseInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteDatabaseInWorkspaceOK struct {
	Payload *model.DatabaseV4Response
}

func (o *DeleteDatabaseInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/databases/{name}][%d] deleteDatabaseInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteDatabaseInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DatabaseV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
