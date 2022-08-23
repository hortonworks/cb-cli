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

// PrepareClusterUpgradeByCrnInternalReader is a Reader for the PrepareClusterUpgradeByCrnInternal structure.
type PrepareClusterUpgradeByCrnInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrepareClusterUpgradeByCrnInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPrepareClusterUpgradeByCrnInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPrepareClusterUpgradeByCrnInternalOK creates a PrepareClusterUpgradeByCrnInternalOK with default headers values
func NewPrepareClusterUpgradeByCrnInternalOK() *PrepareClusterUpgradeByCrnInternalOK {
	return &PrepareClusterUpgradeByCrnInternalOK{}
}

/*PrepareClusterUpgradeByCrnInternalOK handles this case with default header values.

successful operation
*/
type PrepareClusterUpgradeByCrnInternalOK struct {
	Payload *model.FlowIdentifier
}

func (o *PrepareClusterUpgradeByCrnInternalOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/stacks/internal/{crn}/prepare_cluster_upgrade][%d] prepareClusterUpgradeByCrnInternalOK  %+v", 200, o.Payload)
}

func (o *PrepareClusterUpgradeByCrnInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
