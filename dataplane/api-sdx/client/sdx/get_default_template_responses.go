// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// GetDefaultTemplateReader is a Reader for the GetDefaultTemplate structure.
type GetDefaultTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDefaultTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDefaultTemplateOK creates a GetDefaultTemplateOK with default headers values
func NewGetDefaultTemplateOK() *GetDefaultTemplateOK {
	return &GetDefaultTemplateOK{}
}

/*GetDefaultTemplateOK handles this case with default header values.

successful operation
*/
type GetDefaultTemplateOK struct {
	Payload *model.SdxDefaultTemplateResponse
}

func (o *GetDefaultTemplateOK) Error() string {
	return fmt.Sprintf("[GET /sdx/default_template][%d] getDefaultTemplateOK  %+v", 200, o.Payload)
}

func (o *GetDefaultTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxDefaultTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
