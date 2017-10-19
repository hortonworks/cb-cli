// Code generated by go-swagger; DO NOT EDIT.

package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPublicRdsReader is a Reader for the PostPublicRds structure.
type PostPublicRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicRdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicRdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicRdsOK creates a PostPublicRdsOK with default headers values
func NewPostPublicRdsOK() *PostPublicRdsOK {
	return &PostPublicRdsOK{}
}

/*PostPublicRdsOK handles this case with default header values.

successful operation
*/
type PostPublicRdsOK struct {
	Payload *models_cloudbreak.RDSConfigResponse
}

func (o *PostPublicRdsOK) Error() string {
	return fmt.Sprintf("[POST /rdsconfigs/account][%d] postPublicRdsOK  %+v", 200, o.Payload)
}

func (o *PostPublicRdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RDSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
