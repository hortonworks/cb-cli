// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicStackReader is a Reader for the GetPublicStack structure.
type GetPublicStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicStackOK creates a GetPublicStackOK with default headers values
func NewGetPublicStackOK() *GetPublicStackOK {
	return &GetPublicStackOK{}
}

/*GetPublicStackOK handles this case with default header values.

successful operation
*/
type GetPublicStackOK struct {
	Payload *models_cloudbreak.StackResponse
}

func (o *GetPublicStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/account/{name}][%d] getPublicStackOK  %+v", 200, o.Payload)
}

func (o *GetPublicStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}