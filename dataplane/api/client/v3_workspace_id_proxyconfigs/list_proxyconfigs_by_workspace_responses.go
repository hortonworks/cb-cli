// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListProxyconfigsByWorkspaceReader is a Reader for the ListProxyconfigsByWorkspace structure.
type ListProxyconfigsByWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProxyconfigsByWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProxyconfigsByWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListProxyconfigsByWorkspaceOK creates a ListProxyconfigsByWorkspaceOK with default headers values
func NewListProxyconfigsByWorkspaceOK() *ListProxyconfigsByWorkspaceOK {
	return &ListProxyconfigsByWorkspaceOK{}
}

/*ListProxyconfigsByWorkspaceOK handles this case with default header values.

successful operation
*/
type ListProxyconfigsByWorkspaceOK struct {
	Payload []*model.ProxyConfigResponse
}

func (o *ListProxyconfigsByWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/proxyconfigs][%d] listProxyconfigsByWorkspaceOK  %+v", 200, o.Payload)
}

func (o *ListProxyconfigsByWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
