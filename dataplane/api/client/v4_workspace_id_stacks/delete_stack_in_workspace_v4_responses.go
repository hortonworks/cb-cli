// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteStackInWorkspaceV4Reader is a Reader for the DeleteStackInWorkspaceV4 structure.
type DeleteStackInWorkspaceV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStackInWorkspaceV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteStackInWorkspaceV4Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteStackInWorkspaceV4Default creates a DeleteStackInWorkspaceV4Default with default headers values
func NewDeleteStackInWorkspaceV4Default(code int) *DeleteStackInWorkspaceV4Default {
	return &DeleteStackInWorkspaceV4Default{
		_statusCode: code,
	}
}

/*DeleteStackInWorkspaceV4Default handles this case with default header values.

successful operation
*/
type DeleteStackInWorkspaceV4Default struct {
	_statusCode int
}

// Code gets the status code for the delete stack in workspace v4 default response
func (o *DeleteStackInWorkspaceV4Default) Code() int {
	return o._statusCode
}

func (o *DeleteStackInWorkspaceV4Default) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/stacks/{name}][%d] deleteStackInWorkspaceV4 default ", o._statusCode)
}

func (o *DeleteStackInWorkspaceV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
