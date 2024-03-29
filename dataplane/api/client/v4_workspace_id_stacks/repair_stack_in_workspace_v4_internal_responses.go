// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RepairStackInWorkspaceV4InternalReader is a Reader for the RepairStackInWorkspaceV4Internal structure.
type RepairStackInWorkspaceV4InternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairStackInWorkspaceV4InternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepairStackInWorkspaceV4InternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepairStackInWorkspaceV4InternalOK creates a RepairStackInWorkspaceV4InternalOK with default headers values
func NewRepairStackInWorkspaceV4InternalOK() *RepairStackInWorkspaceV4InternalOK {
	return &RepairStackInWorkspaceV4InternalOK{}
}

/*
RepairStackInWorkspaceV4InternalOK handles this case with default header values.

successful operation
*/
type RepairStackInWorkspaceV4InternalOK struct {
	Payload *model.FlowIdentifier
}

func (o *RepairStackInWorkspaceV4InternalOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/stacks/internal/{name}/manual_repair][%d] repairStackInWorkspaceV4InternalOK  %+v", 200, o.Payload)
}

func (o *RepairStackInWorkspaceV4InternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
