// Code generated by go-swagger; DO NOT EDIT.

package v4customimagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetCustomImageReader is a Reader for the GetCustomImage structure.
type GetCustomImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomImageOK creates a GetCustomImageOK with default headers values
func NewGetCustomImageOK() *GetCustomImageOK {
	return &GetCustomImageOK{}
}

/*
GetCustomImageOK handles this case with default header values.

successful operation
*/
type GetCustomImageOK struct {
	Payload *model.CustomImageCatalogV4GetImageResponse
}

func (o *GetCustomImageOK) Error() string {
	return fmt.Sprintf("[GET /v4/custom_image_catalogs/{name}/image/{imageId}][%d] getCustomImageOK  %+v", 200, o.Payload)
}

func (o *GetCustomImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CustomImageCatalogV4GetImageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
