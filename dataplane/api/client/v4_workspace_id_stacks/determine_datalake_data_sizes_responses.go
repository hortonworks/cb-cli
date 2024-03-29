// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DetermineDatalakeDataSizesReader is a Reader for the DetermineDatalakeDataSizes structure.
type DetermineDatalakeDataSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetermineDatalakeDataSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDetermineDatalakeDataSizesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDetermineDatalakeDataSizesDefault creates a DetermineDatalakeDataSizesDefault with default headers values
func NewDetermineDatalakeDataSizesDefault(code int) *DetermineDatalakeDataSizesDefault {
	return &DetermineDatalakeDataSizesDefault{
		_statusCode: code,
	}
}

/*
DetermineDatalakeDataSizesDefault handles this case with default header values.

successful operation
*/
type DetermineDatalakeDataSizesDefault struct {
	_statusCode int
}

// Code gets the status code for the determine datalake data sizes default response
func (o *DetermineDatalakeDataSizesDefault) Code() int {
	return o._statusCode
}

func (o *DetermineDatalakeDataSizesDefault) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/stacks/{name}/determine_datalake_data_sizes][%d] determineDatalakeDataSizes default ", o._statusCode)
}

func (o *DetermineDatalakeDataSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
