// Code generated by go-swagger; DO NOT EDIT.

package v4user_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutTerminatedClustersPreferencesReader is a Reader for the PutTerminatedClustersPreferences structure.
type PutTerminatedClustersPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTerminatedClustersPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutTerminatedClustersPreferencesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutTerminatedClustersPreferencesDefault creates a PutTerminatedClustersPreferencesDefault with default headers values
func NewPutTerminatedClustersPreferencesDefault(code int) *PutTerminatedClustersPreferencesDefault {
	return &PutTerminatedClustersPreferencesDefault{
		_statusCode: code,
	}
}

/*
PutTerminatedClustersPreferencesDefault handles this case with default header values.

successful operation
*/
type PutTerminatedClustersPreferencesDefault struct {
	_statusCode int
}

// Code gets the status code for the put terminated clusters preferences default response
func (o *PutTerminatedClustersPreferencesDefault) Code() int {
	return o._statusCode
}

func (o *PutTerminatedClustersPreferencesDefault) Error() string {
	return fmt.Sprintf("[PUT /v4/user_profiles/terminated_clusters_preferences][%d] putTerminatedClustersPreferences default ", o._statusCode)
}

func (o *PutTerminatedClustersPreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
