// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SubmitDatalakeDataSizesInternalReader is a Reader for the SubmitDatalakeDataSizesInternal structure.
type SubmitDatalakeDataSizesInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitDatalakeDataSizesInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewSubmitDatalakeDataSizesInternalDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewSubmitDatalakeDataSizesInternalDefault creates a SubmitDatalakeDataSizesInternalDefault with default headers values
func NewSubmitDatalakeDataSizesInternalDefault(code int) *SubmitDatalakeDataSizesInternalDefault {
	return &SubmitDatalakeDataSizesInternalDefault{
		_statusCode: code,
	}
}

/*
SubmitDatalakeDataSizesInternalDefault handles this case with default header values.

successful operation
*/
type SubmitDatalakeDataSizesInternalDefault struct {
	_statusCode int
}

// Code gets the status code for the submit datalake data sizes internal default response
func (o *SubmitDatalakeDataSizesInternalDefault) Code() int {
	return o._statusCode
}

func (o *SubmitDatalakeDataSizesInternalDefault) Error() string {
	return fmt.Sprintf("[PUT /sdx/crn/{crn}/submitDatalakeDataSizes/internal][%d] submitDatalakeDataSizesInternal default ", o._statusCode)
}

func (o *SubmitDatalakeDataSizesInternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
