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

// GetByCrnReader is a Reader for the GetByCrn structure.
type GetByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetByCrnOK creates a GetByCrnOK with default headers values
func NewGetByCrnOK() *GetByCrnOK {
	return &GetByCrnOK{}
}

/*
GetByCrnOK handles this case with default header values.

successful operation
*/
type GetByCrnOK struct {
	Payload *model.CustomConfigurationsV4Response
}

func (o *GetByCrnOK) Error() string {
	return fmt.Sprintf("[GET /v4/custom_configurations/crn/{crn}][%d] getByCrnOK  %+v", 200, o.Payload)
}

func (o *GetByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CustomConfigurationsV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
