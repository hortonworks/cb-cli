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

// CheckForOsUpgradeInWorkspaceV4Reader is a Reader for the CheckForOsUpgradeInWorkspaceV4 structure.
type CheckForOsUpgradeInWorkspaceV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckForOsUpgradeInWorkspaceV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckForOsUpgradeInWorkspaceV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckForOsUpgradeInWorkspaceV4OK creates a CheckForOsUpgradeInWorkspaceV4OK with default headers values
func NewCheckForOsUpgradeInWorkspaceV4OK() *CheckForOsUpgradeInWorkspaceV4OK {
	return &CheckForOsUpgradeInWorkspaceV4OK{}
}

/*
CheckForOsUpgradeInWorkspaceV4OK handles this case with default header values.

successful operation
*/
type CheckForOsUpgradeInWorkspaceV4OK struct {
	Payload *model.UpgradeOptionV4Response
}

func (o *CheckForOsUpgradeInWorkspaceV4OK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/{name}/check_for_upgrade][%d] checkForOsUpgradeInWorkspaceV4OK  %+v", 200, o.Payload)
}

func (o *CheckForOsUpgradeInWorkspaceV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.UpgradeOptionV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
