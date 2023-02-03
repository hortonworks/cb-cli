// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// CreateBlueprintInWorkspaceInternalReader is a Reader for the CreateBlueprintInWorkspaceInternal structure.
type CreateBlueprintInWorkspaceInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlueprintInWorkspaceInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateBlueprintInWorkspaceInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateBlueprintInWorkspaceInternalOK creates a CreateBlueprintInWorkspaceInternalOK with default headers values
func NewCreateBlueprintInWorkspaceInternalOK() *CreateBlueprintInWorkspaceInternalOK {
	return &CreateBlueprintInWorkspaceInternalOK{}
}

/*
CreateBlueprintInWorkspaceInternalOK handles this case with default header values.

successful operation
*/
type CreateBlueprintInWorkspaceInternalOK struct {
	Payload *model.BlueprintV4Response
}

func (o *CreateBlueprintInWorkspaceInternalOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/blueprints/internal][%d] createBlueprintInWorkspaceInternalOK  %+v", 200, o.Payload)
}

func (o *CreateBlueprintInWorkspaceInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.BlueprintV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
