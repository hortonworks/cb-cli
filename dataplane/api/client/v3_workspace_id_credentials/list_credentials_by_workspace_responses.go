// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListCredentialsByWorkspaceReader is a Reader for the ListCredentialsByWorkspace structure.
type ListCredentialsByWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCredentialsByWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCredentialsByWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCredentialsByWorkspaceOK creates a ListCredentialsByWorkspaceOK with default headers values
func NewListCredentialsByWorkspaceOK() *ListCredentialsByWorkspaceOK {
	return &ListCredentialsByWorkspaceOK{}
}

/*ListCredentialsByWorkspaceOK handles this case with default header values.

successful operation
*/
type ListCredentialsByWorkspaceOK struct {
	Payload []*model.CredentialResponse
}

func (o *ListCredentialsByWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/credentials][%d] listCredentialsByWorkspaceOK  %+v", 200, o.Payload)
}

func (o *ListCredentialsByWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
