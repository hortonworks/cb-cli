// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetGatewaysCredentialIDReader is a Reader for the GetGatewaysCredentialID structure.
type GetGatewaysCredentialIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGatewaysCredentialIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGatewaysCredentialIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGatewaysCredentialIDOK creates a GetGatewaysCredentialIDOK with default headers values
func NewGetGatewaysCredentialIDOK() *GetGatewaysCredentialIDOK {
	return &GetGatewaysCredentialIDOK{}
}

/*GetGatewaysCredentialIDOK handles this case with default header values.

successful operation
*/
type GetGatewaysCredentialIDOK struct {
	Payload *models_cloudbreak.PlatformGatewaysResponse
}

func (o *GetGatewaysCredentialIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/connectors/gateways][%d] getGatewaysCredentialIdOK  %+v", 200, o.Payload)
}

func (o *GetGatewaysCredentialIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformGatewaysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
