// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_ldaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteLdapConfigInWorkspaceReader is a Reader for the DeleteLdapConfigInWorkspace structure.
type DeleteLdapConfigInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLdapConfigInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteLdapConfigInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLdapConfigInWorkspaceOK creates a DeleteLdapConfigInWorkspaceOK with default headers values
func NewDeleteLdapConfigInWorkspaceOK() *DeleteLdapConfigInWorkspaceOK {
	return &DeleteLdapConfigInWorkspaceOK{}
}

/*DeleteLdapConfigInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteLdapConfigInWorkspaceOK struct {
	Payload *model.LdapV4Response
}

func (o *DeleteLdapConfigInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/ldaps/{name}][%d] deleteLdapConfigInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteLdapConfigInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.LdapV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
