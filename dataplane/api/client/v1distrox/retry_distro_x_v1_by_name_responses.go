// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RetryDistroXV1ByNameReader is a Reader for the RetryDistroXV1ByName structure.
type RetryDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRetryDistroXV1ByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRetryDistroXV1ByNameDefault creates a RetryDistroXV1ByNameDefault with default headers values
func NewRetryDistroXV1ByNameDefault(code int) *RetryDistroXV1ByNameDefault {
	return &RetryDistroXV1ByNameDefault{
		_statusCode: code,
	}
}

/*RetryDistroXV1ByNameDefault handles this case with default header values.

successful operation
*/
type RetryDistroXV1ByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the retry distro x v1 by name default response
func (o *RetryDistroXV1ByNameDefault) Code() int {
	return o._statusCode
}

func (o *RetryDistroXV1ByNameDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/name/{name}/retry][%d] retryDistroXV1ByName default ", o._statusCode)
}

func (o *RetryDistroXV1ByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
