// Code generated by go-swagger; DO NOT EDIT.

package v4custom_configurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetServiceTypesReader is a Reader for the GetServiceTypes structure.
type GetServiceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceTypesOK creates a GetServiceTypesOK with default headers values
func NewGetServiceTypesOK() *GetServiceTypesOK {
	return &GetServiceTypesOK{}
}

/*
GetServiceTypesOK handles this case with default header values.

successful operation
*/
type GetServiceTypesOK struct {
	Payload *model.ServiceTypeV4Response
}

func (o *GetServiceTypesOK) Error() string {
	return fmt.Sprintf("[GET /v4/custom_configurations/service_types][%d] getServiceTypesOK  %+v", 200, o.Payload)
}

func (o *GetServiceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ServiceTypeV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
