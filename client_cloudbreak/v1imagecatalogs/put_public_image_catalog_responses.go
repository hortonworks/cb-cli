// Code generated by go-swagger; DO NOT EDIT.

package v1imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PutPublicImageCatalogReader is a Reader for the PutPublicImageCatalog structure.
type PutPublicImageCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPublicImageCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutPublicImageCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPublicImageCatalogOK creates a PutPublicImageCatalogOK with default headers values
func NewPutPublicImageCatalogOK() *PutPublicImageCatalogOK {
	return &PutPublicImageCatalogOK{}
}

/*PutPublicImageCatalogOK handles this case with default header values.

successful operation
*/
type PutPublicImageCatalogOK struct {
	Payload *models_cloudbreak.ImageCatalogResponse
}

func (o *PutPublicImageCatalogOK) Error() string {
	return fmt.Sprintf("[PUT /v1/imagecatalogs/account][%d] putPublicImageCatalogOK  %+v", 200, o.Payload)
}

func (o *PutPublicImageCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ImageCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
