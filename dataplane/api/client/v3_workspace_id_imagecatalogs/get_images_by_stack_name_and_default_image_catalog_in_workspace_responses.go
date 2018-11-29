// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetImagesByStackNameAndDefaultImageCatalogInWorkspaceReader is a Reader for the GetImagesByStackNameAndDefaultImageCatalogInWorkspace structure.
type GetImagesByStackNameAndDefaultImageCatalogInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesByStackNameAndDefaultImageCatalogInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK creates a GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK with default headers values
func NewGetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK() *GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK {
	return &GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK{}
}

/*GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK struct {
	Payload *model.ImagesResponse
}

func (o *GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/imagecatalogs/upgrade/{stackName}][%d] getImagesByStackNameAndDefaultImageCatalogInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetImagesByStackNameAndDefaultImageCatalogInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImagesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
