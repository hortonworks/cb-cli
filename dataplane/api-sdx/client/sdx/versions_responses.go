// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VersionsReader is a Reader for the Versions structure.
type VersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVersionsOK creates a VersionsOK with default headers values
func NewVersionsOK() *VersionsOK {
	return &VersionsOK{}
}

/*VersionsOK handles this case with default header values.

successful operation
*/
type VersionsOK struct {
	Payload []string
}

func (o *VersionsOK) Error() string {
	return fmt.Sprintf("[GET /sdx/versions][%d] versionsOK  %+v", 200, o.Payload)
}

func (o *VersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
