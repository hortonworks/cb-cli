// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GenerateImageCatalogReader is a Reader for the GenerateImageCatalog structure.
type GenerateImageCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateImageCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenerateImageCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateImageCatalogOK creates a GenerateImageCatalogOK with default headers values
func NewGenerateImageCatalogOK() *GenerateImageCatalogOK {
	return &GenerateImageCatalogOK{}
}

/*
GenerateImageCatalogOK handles this case with default header values.

successful operation
*/
type GenerateImageCatalogOK struct {
	Payload *model.DistroXGenerateImageCatalogV1Response
}

func (o *GenerateImageCatalogOK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/name/{name}/generate_image_catalog][%d] generateImageCatalogOK  %+v", 200, o.Payload)
}

func (o *GenerateImageCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DistroXGenerateImageCatalogV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
