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

// GetDatalakeBackupIDReader is a Reader for the GetDatalakeBackupID structure.
type GetDatalakeBackupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatalakeBackupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDatalakeBackupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDatalakeBackupIDOK creates a GetDatalakeBackupIDOK with default headers values
func NewGetDatalakeBackupIDOK() *GetDatalakeBackupIDOK {
	return &GetDatalakeBackupIDOK{}
}

/*GetDatalakeBackupIDOK handles this case with default header values.

successful operation
*/
type GetDatalakeBackupIDOK struct {
	Payload string
}

func (o *GetDatalakeBackupIDOK) Error() string {
	return fmt.Sprintf("[GET /sdx/{name}/getDatalakeBackupId][%d] getDatalakeBackupIdOK  %+v", 200, o.Payload)
}

func (o *GetDatalakeBackupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
