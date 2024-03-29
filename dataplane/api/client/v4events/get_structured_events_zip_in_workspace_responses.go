// Code generated by go-swagger; DO NOT EDIT.

package v4events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetStructuredEventsZipInWorkspaceReader is a Reader for the GetStructuredEventsZipInWorkspace structure.
type GetStructuredEventsZipInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStructuredEventsZipInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetStructuredEventsZipInWorkspaceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetStructuredEventsZipInWorkspaceDefault creates a GetStructuredEventsZipInWorkspaceDefault with default headers values
func NewGetStructuredEventsZipInWorkspaceDefault(code int) *GetStructuredEventsZipInWorkspaceDefault {
	return &GetStructuredEventsZipInWorkspaceDefault{
		_statusCode: code,
	}
}

/*
GetStructuredEventsZipInWorkspaceDefault handles this case with default header values.

successful operation
*/
type GetStructuredEventsZipInWorkspaceDefault struct {
	_statusCode int
}

// Code gets the status code for the get structured events zip in workspace default response
func (o *GetStructuredEventsZipInWorkspaceDefault) Code() int {
	return o._statusCode
}

func (o *GetStructuredEventsZipInWorkspaceDefault) Error() string {
	return fmt.Sprintf("[GET /v4/events/{name}/zip][%d] getStructuredEventsZipInWorkspace default ", o._statusCode)
}

func (o *GetStructuredEventsZipInWorkspaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
