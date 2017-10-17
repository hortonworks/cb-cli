// Code generated by go-swagger; DO NOT EDIT.

package v1users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HasResourcesUserReader is a Reader for the HasResourcesUser structure.
type HasResourcesUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HasResourcesUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHasResourcesUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHasResourcesUserOK creates a HasResourcesUserOK with default headers values
func NewHasResourcesUserOK() *HasResourcesUserOK {
	return &HasResourcesUserOK{}
}

/*HasResourcesUserOK handles this case with default header values.

successful operation
*/
type HasResourcesUserOK struct {
	Payload bool
}

func (o *HasResourcesUserOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}/resources][%d] hasResourcesUserOK  %+v", 200, o.Payload)
}

func (o *HasResourcesUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
