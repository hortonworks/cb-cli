// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GenerateImageCatalogInternalReader is a Reader for the GenerateImageCatalogInternal structure.
type GenerateImageCatalogInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateImageCatalogInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenerateImageCatalogInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateImageCatalogInternalOK creates a GenerateImageCatalogInternalOK with default headers values
func NewGenerateImageCatalogInternalOK() *GenerateImageCatalogInternalOK {
	return &GenerateImageCatalogInternalOK{}
}

/*GenerateImageCatalogInternalOK handles this case with default header values.

successful operation
*/
type GenerateImageCatalogInternalOK struct {
	Payload *model.GenerateImageCatalogV4Response
}

func (o *GenerateImageCatalogInternalOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/internal/{name}/generate_image_catalog][%d] generateImageCatalogInternalOK  %+v", 200, o.Payload)
}

func (o *GenerateImageCatalogInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GenerateImageCatalogV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
