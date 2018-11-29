// Code generated by go-swagger; DO NOT EDIT.

package v3workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteWorkspaceByNameReader is a Reader for the DeleteWorkspaceByName structure.
type DeleteWorkspaceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkspaceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteWorkspaceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWorkspaceByNameOK creates a DeleteWorkspaceByNameOK with default headers values
func NewDeleteWorkspaceByNameOK() *DeleteWorkspaceByNameOK {
	return &DeleteWorkspaceByNameOK{}
}

/*DeleteWorkspaceByNameOK handles this case with default header values.

successful operation
*/
type DeleteWorkspaceByNameOK struct {
	Payload *model.WorkspaceResponse
}

func (o *DeleteWorkspaceByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /v3/workspaces/name/{name}][%d] deleteWorkspaceByNameOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkspaceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.WorkspaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
