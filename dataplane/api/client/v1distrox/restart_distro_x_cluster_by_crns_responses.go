// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RestartDistroXClusterByCrnsReader is a Reader for the RestartDistroXClusterByCrns structure.
type RestartDistroXClusterByCrnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDistroXClusterByCrnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRestartDistroXClusterByCrnsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRestartDistroXClusterByCrnsDefault creates a RestartDistroXClusterByCrnsDefault with default headers values
func NewRestartDistroXClusterByCrnsDefault(code int) *RestartDistroXClusterByCrnsDefault {
	return &RestartDistroXClusterByCrnsDefault{
		_statusCode: code,
	}
}

/*
RestartDistroXClusterByCrnsDefault handles this case with default header values.

successful operation
*/
type RestartDistroXClusterByCrnsDefault struct {
	_statusCode int
}

// Code gets the status code for the restart distro x cluster by crns default response
func (o *RestartDistroXClusterByCrnsDefault) Code() int {
	return o._statusCode
}

func (o *RestartDistroXClusterByCrnsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/crn/restartCluster][%d] restartDistroXClusterByCrns default ", o._statusCode)
}

func (o *RestartDistroXClusterByCrnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
