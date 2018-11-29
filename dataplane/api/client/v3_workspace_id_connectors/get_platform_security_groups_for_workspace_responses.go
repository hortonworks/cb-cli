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

// GetPlatformSecurityGroupsForWorkspaceReader is a Reader for the GetPlatformSecurityGroupsForWorkspace structure.
type GetPlatformSecurityGroupsForWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformSecurityGroupsForWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformSecurityGroupsForWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlatformSecurityGroupsForWorkspaceOK creates a GetPlatformSecurityGroupsForWorkspaceOK with default headers values
func NewGetPlatformSecurityGroupsForWorkspaceOK() *GetPlatformSecurityGroupsForWorkspaceOK {
	return &GetPlatformSecurityGroupsForWorkspaceOK{}
}

/*GetPlatformSecurityGroupsForWorkspaceOK handles this case with default header values.

successful operation
*/
type GetPlatformSecurityGroupsForWorkspaceOK struct {
	Payload *model.PlatformSecurityGroupsResponse
}

func (o *GetPlatformSecurityGroupsForWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /v3/{workspaceId}/connectors/securitygroups][%d] getPlatformSecurityGroupsForWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetPlatformSecurityGroupsForWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformSecurityGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
