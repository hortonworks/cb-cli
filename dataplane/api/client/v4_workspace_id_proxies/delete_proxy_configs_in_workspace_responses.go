// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteProxyConfigsInWorkspaceReader is a Reader for the DeleteProxyConfigsInWorkspace structure.
type DeleteProxyConfigsInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProxyConfigsInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProxyConfigsInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProxyConfigsInWorkspaceOK creates a DeleteProxyConfigsInWorkspaceOK with default headers values
func NewDeleteProxyConfigsInWorkspaceOK() *DeleteProxyConfigsInWorkspaceOK {
	return &DeleteProxyConfigsInWorkspaceOK{}
}

/*DeleteProxyConfigsInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteProxyConfigsInWorkspaceOK struct {
	Payload *model.ProxyV4Responses
}

func (o *DeleteProxyConfigsInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/proxies][%d] deleteProxyConfigsInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteProxyConfigsInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
