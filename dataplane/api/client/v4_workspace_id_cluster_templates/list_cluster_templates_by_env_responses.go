// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_cluster_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListClusterTemplatesByEnvReader is a Reader for the ListClusterTemplatesByEnv structure.
type ListClusterTemplatesByEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterTemplatesByEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClusterTemplatesByEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClusterTemplatesByEnvOK creates a ListClusterTemplatesByEnvOK with default headers values
func NewListClusterTemplatesByEnvOK() *ListClusterTemplatesByEnvOK {
	return &ListClusterTemplatesByEnvOK{}
}

/*
ListClusterTemplatesByEnvOK handles this case with default header values.

successful operation
*/
type ListClusterTemplatesByEnvOK struct {
	Payload *model.ClusterTemplateViewV4Responses
}

func (o *ListClusterTemplatesByEnvOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/cluster_templates/env/{environmentCrn}][%d] listClusterTemplatesByEnvOK  %+v", 200, o.Payload)
}

func (o *ListClusterTemplatesByEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ClusterTemplateViewV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
