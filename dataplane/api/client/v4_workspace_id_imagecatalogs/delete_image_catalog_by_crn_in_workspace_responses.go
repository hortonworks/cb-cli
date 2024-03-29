// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteImageCatalogByCrnInWorkspaceReader is a Reader for the DeleteImageCatalogByCrnInWorkspace structure.
type DeleteImageCatalogByCrnInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageCatalogByCrnInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteImageCatalogByCrnInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteImageCatalogByCrnInWorkspaceOK creates a DeleteImageCatalogByCrnInWorkspaceOK with default headers values
func NewDeleteImageCatalogByCrnInWorkspaceOK() *DeleteImageCatalogByCrnInWorkspaceOK {
	return &DeleteImageCatalogByCrnInWorkspaceOK{}
}

/*
DeleteImageCatalogByCrnInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteImageCatalogByCrnInWorkspaceOK struct {
	Payload *model.ImageCatalogV4Response
}

func (o *DeleteImageCatalogByCrnInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/{workspaceId}/image_catalogs/crn/{crn}][%d] deleteImageCatalogByCrnInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteImageCatalogByCrnInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImageCatalogV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
