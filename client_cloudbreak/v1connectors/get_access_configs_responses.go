// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetAccessConfigsReader is a Reader for the GetAccessConfigs structure.
type GetAccessConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccessConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccessConfigsOK creates a GetAccessConfigsOK with default headers values
func NewGetAccessConfigsOK() *GetAccessConfigsOK {
	return &GetAccessConfigsOK{}
}

/*GetAccessConfigsOK handles this case with default header values.

successful operation
*/
type GetAccessConfigsOK struct {
	Payload *models_cloudbreak.PlatformAccessConfigsResponse
}

func (o *GetAccessConfigsOK) Error() string {
	return fmt.Sprintf("[POST /v1/connectors/accessconfigs][%d] getAccessConfigsOK  %+v", 200, o.Payload)
}

func (o *GetAccessConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformAccessConfigsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
