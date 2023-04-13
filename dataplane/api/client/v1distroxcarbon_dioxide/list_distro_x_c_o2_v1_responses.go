// Code generated by go-swagger; DO NOT EDIT.

package v1distroxcarbon_dioxide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListDistroXCO2V1Reader is a Reader for the ListDistroXCO2V1 structure.
type ListDistroXCO2V1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDistroXCO2V1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDistroXCO2V1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDistroXCO2V1OK creates a ListDistroXCO2V1OK with default headers values
func NewListDistroXCO2V1OK() *ListDistroXCO2V1OK {
	return &ListDistroXCO2V1OK{}
}

/*
ListDistroXCO2V1OK handles this case with default header values.

successful operation
*/
type ListDistroXCO2V1OK struct {
	Payload *model.RealTimeCO2Response
}

func (o *ListDistroXCO2V1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/carbon_dioxide][%d] listDistroXCO2V1OK  %+v", 200, o.Payload)
}

func (o *ListDistroXCO2V1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RealTimeCO2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
