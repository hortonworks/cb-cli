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

// UpgradeClusterByNameInternalReader is a Reader for the UpgradeClusterByNameInternal structure.
type UpgradeClusterByNameInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeClusterByNameInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpgradeClusterByNameInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpgradeClusterByNameInternalOK creates a UpgradeClusterByNameInternalOK with default headers values
func NewUpgradeClusterByNameInternalOK() *UpgradeClusterByNameInternalOK {
	return &UpgradeClusterByNameInternalOK{}
}

/*
UpgradeClusterByNameInternalOK handles this case with default header values.

successful operation
*/
type UpgradeClusterByNameInternalOK struct {
	Payload *model.FlowIdentifier
}

func (o *UpgradeClusterByNameInternalOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/stacks/internal/{name}/cluster_upgrade][%d] upgradeClusterByNameInternalOK  %+v", 200, o.Payload)
}

func (o *UpgradeClusterByNameInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
