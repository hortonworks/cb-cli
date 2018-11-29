// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPlatformNetworksForWorkspaceReader is a Reader for the GetPlatformNetworksForWorkspace structure.
type GetPlatformNetworksForWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformNetworksForWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformNetworksForWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlatformNetworksForWorkspaceOK creates a GetPlatformNetworksForWorkspaceOK with default headers values
func NewGetPlatformNetworksForWorkspaceOK() *GetPlatformNetworksForWorkspaceOK {
	return &GetPlatformNetworksForWorkspaceOK{}
}

/*GetPlatformNetworksForWorkspaceOK handles this case with default header values.

successful operation
*/
type GetPlatformNetworksForWorkspaceOK struct {
	Payload *model.PlatformNetworksResponse
}

func (o *GetPlatformNetworksForWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /v3/{workspaceId}/connectors/networks][%d] getPlatformNetworksForWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetPlatformNetworksForWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformNetworksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
