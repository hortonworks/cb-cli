// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPublicsStackReader is a Reader for the GetPublicsStack structure.
type GetPublicsStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsStackOK creates a GetPublicsStackOK with default headers values
func NewGetPublicsStackOK() *GetPublicsStackOK {
	return &GetPublicsStackOK{}
}

/*GetPublicsStackOK handles this case with default header values.

successful operation
*/
type GetPublicsStackOK struct {
	Payload []*models_cloudbreak.StackResponse
}

func (o *GetPublicsStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/account][%d] getPublicsStackOK  %+v", 200, o.Payload)
}

func (o *GetPublicsStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
