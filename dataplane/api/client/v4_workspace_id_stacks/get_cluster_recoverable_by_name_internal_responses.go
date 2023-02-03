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

// GetClusterRecoverableByNameInternalReader is a Reader for the GetClusterRecoverableByNameInternal structure.
type GetClusterRecoverableByNameInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterRecoverableByNameInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterRecoverableByNameInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterRecoverableByNameInternalOK creates a GetClusterRecoverableByNameInternalOK with default headers values
func NewGetClusterRecoverableByNameInternalOK() *GetClusterRecoverableByNameInternalOK {
	return &GetClusterRecoverableByNameInternalOK{}
}

/*
GetClusterRecoverableByNameInternalOK handles this case with default header values.

successful operation
*/
type GetClusterRecoverableByNameInternalOK struct {
	Payload *model.RecoveryValidationV4Response
}

func (o *GetClusterRecoverableByNameInternalOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/internal/{name}/cluster_recoverable][%d] getClusterRecoverableByNameInternalOK  %+v", 200, o.Payload)
}

func (o *GetClusterRecoverableByNameInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RecoveryValidationV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
