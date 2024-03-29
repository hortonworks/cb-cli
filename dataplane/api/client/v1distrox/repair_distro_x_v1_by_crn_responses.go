// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RepairDistroXV1ByCrnReader is a Reader for the RepairDistroXV1ByCrn structure.
type RepairDistroXV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairDistroXV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepairDistroXV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRepairDistroXV1ByCrnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepairDistroXV1ByCrnOK creates a RepairDistroXV1ByCrnOK with default headers values
func NewRepairDistroXV1ByCrnOK() *RepairDistroXV1ByCrnOK {
	return &RepairDistroXV1ByCrnOK{}
}

/*
RepairDistroXV1ByCrnOK handles this case with default header values.

successful operation
*/
type RepairDistroXV1ByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *RepairDistroXV1ByCrnOK) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/crn/{crn}/manual_repair][%d] repairDistroXV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *RepairDistroXV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepairDistroXV1ByCrnDefault creates a RepairDistroXV1ByCrnDefault with default headers values
func NewRepairDistroXV1ByCrnDefault(code int) *RepairDistroXV1ByCrnDefault {
	return &RepairDistroXV1ByCrnDefault{
		_statusCode: code,
	}
}

/*
RepairDistroXV1ByCrnDefault handles this case with default header values.

unsuccessful operation
*/
type RepairDistroXV1ByCrnDefault struct {
	_statusCode int
}

// Code gets the status code for the repair distro x v1 by crn default response
func (o *RepairDistroXV1ByCrnDefault) Code() int {
	return o._statusCode
}

func (o *RepairDistroXV1ByCrnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/crn/{crn}/manual_repair][%d] repairDistroXV1ByCrn default ", o._statusCode)
}

func (o *RepairDistroXV1ByCrnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
