// Code generated by go-swagger; DO NOT EDIT.

package v1clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPublicClusterTemplateReader is a Reader for the PostPublicClusterTemplate structure.
type PostPublicClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicClusterTemplateOK creates a PostPublicClusterTemplateOK with default headers values
func NewPostPublicClusterTemplateOK() *PostPublicClusterTemplateOK {
	return &PostPublicClusterTemplateOK{}
}

/*PostPublicClusterTemplateOK handles this case with default header values.

successful operation
*/
type PostPublicClusterTemplateOK struct {
	Payload *models_cloudbreak.ClusterTemplateResponse
}

func (o *PostPublicClusterTemplateOK) Error() string {
	return fmt.Sprintf("[POST /v1/clustertemplates/account][%d] postPublicClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *PostPublicClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ClusterTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
