// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteClusterTemplateByNameInWorkspaceReader is a Reader for the DeleteClusterTemplateByNameInWorkspace structure.
type DeleteClusterTemplateByNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterTemplateByNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterTemplateByNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClusterTemplateByNameInWorkspaceOK creates a DeleteClusterTemplateByNameInWorkspaceOK with default headers values
func NewDeleteClusterTemplateByNameInWorkspaceOK() *DeleteClusterTemplateByNameInWorkspaceOK {
	return &DeleteClusterTemplateByNameInWorkspaceOK{}
}

/*DeleteClusterTemplateByNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteClusterTemplateByNameInWorkspaceOK struct {
	Payload *model.ClusterTemplateV4Response
}

func (o *DeleteClusterTemplateByNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/cluster_templates/name/{name}][%d] deleteClusterTemplateByNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteClusterTemplateByNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ClusterTemplateV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
