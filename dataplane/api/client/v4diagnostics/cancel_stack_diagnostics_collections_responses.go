// Code generated by go-swagger; DO NOT EDIT.

package v4diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CancelStackDiagnosticsCollectionsReader is a Reader for the CancelStackDiagnosticsCollections structure.
type CancelStackDiagnosticsCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelStackDiagnosticsCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewCancelStackDiagnosticsCollectionsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewCancelStackDiagnosticsCollectionsDefault creates a CancelStackDiagnosticsCollectionsDefault with default headers values
func NewCancelStackDiagnosticsCollectionsDefault(code int) *CancelStackDiagnosticsCollectionsDefault {
	return &CancelStackDiagnosticsCollectionsDefault{
		_statusCode: code,
	}
}

/*
CancelStackDiagnosticsCollectionsDefault handles this case with default header values.

successful operation
*/
type CancelStackDiagnosticsCollectionsDefault struct {
	_statusCode int
}

// Code gets the status code for the cancel stack diagnostics collections default response
func (o *CancelStackDiagnosticsCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *CancelStackDiagnosticsCollectionsDefault) Error() string {
	return fmt.Sprintf("[POST /v4/diagnostics/{crn}/collections/cancel][%d] cancelStackDiagnosticsCollections default ", o._statusCode)
}

func (o *CancelStackDiagnosticsCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
