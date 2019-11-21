// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCreateEnvironmentForCliByNameReader is a Reader for the GetCreateEnvironmentForCliByName structure.
type GetCreateEnvironmentForCliByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreateEnvironmentForCliByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCreateEnvironmentForCliByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCreateEnvironmentForCliByNameOK creates a GetCreateEnvironmentForCliByNameOK with default headers values
func NewGetCreateEnvironmentForCliByNameOK() *GetCreateEnvironmentForCliByNameOK {
	return &GetCreateEnvironmentForCliByNameOK{}
}

/*GetCreateEnvironmentForCliByNameOK handles this case with default header values.

successful operation
*/
type GetCreateEnvironmentForCliByNameOK struct {
	Payload interface{}
}

func (o *GetCreateEnvironmentForCliByNameOK) Error() string {
	return fmt.Sprintf("[POST /v1/env/name/{name}/cli_create][%d] getCreateEnvironmentForCliByNameOK  %+v", 200, o.Payload)
}

func (o *GetCreateEnvironmentForCliByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
