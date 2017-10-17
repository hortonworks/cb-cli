// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPrivateRdsReader is a Reader for the PostPrivateRds structure.
type PostPrivateRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateRdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateRdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateRdsOK creates a PostPrivateRdsOK with default headers values
func NewPostPrivateRdsOK() *PostPrivateRdsOK {
	return &PostPrivateRdsOK{}
}

/*PostPrivateRdsOK handles this case with default header values.

successful operation
*/
type PostPrivateRdsOK struct {
	Payload *models_cloudbreak.RDSConfigResponse
}

func (o *PostPrivateRdsOK) Error() string {
	return fmt.Sprintf("[POST /v1/rdsconfigs/user][%d] postPrivateRdsOK  %+v", 200, o.Payload)
}

func (o *PostPrivateRdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RDSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
