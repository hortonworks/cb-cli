// Code generated by go-swagger; DO NOT EDIT.

package v1envcarbon_dioxide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// ListEnvironmentCO2V1Reader is a Reader for the ListEnvironmentCO2V1 structure.
type ListEnvironmentCO2V1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEnvironmentCO2V1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListEnvironmentCO2V1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListEnvironmentCO2V1OK creates a ListEnvironmentCO2V1OK with default headers values
func NewListEnvironmentCO2V1OK() *ListEnvironmentCO2V1OK {
	return &ListEnvironmentCO2V1OK{}
}

/*
ListEnvironmentCO2V1OK handles this case with default header values.

successful operation
*/
type ListEnvironmentCO2V1OK struct {
	Payload *model.EnvironmentRealTimeCO2Response
}

func (o *ListEnvironmentCO2V1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/env/carbon_dioxide][%d] listEnvironmentCO2V1OK  %+v", 200, o.Payload)
}

func (o *ListEnvironmentCO2V1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.EnvironmentRealTimeCO2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
