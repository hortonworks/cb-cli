// Code generated by go-swagger; DO NOT EDIT.

package v1networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPublicNetworkReader is a Reader for the PostPublicNetwork structure.
type PostPublicNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicNetworkOK creates a PostPublicNetworkOK with default headers values
func NewPostPublicNetworkOK() *PostPublicNetworkOK {
	return &PostPublicNetworkOK{}
}

/*PostPublicNetworkOK handles this case with default header values.

successful operation
*/
type PostPublicNetworkOK struct {
	Payload *models_cloudbreak.NetworkResponse
}

func (o *PostPublicNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/networks/account][%d] postPublicNetworkOK  %+v", 200, o.Payload)
}

func (o *PostPublicNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
