// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// UpdateGatewayTopologiesReader is a Reader for the UpdateGatewayTopologies structure.
type UpdateGatewayTopologiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGatewayTopologiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateGatewayTopologiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateGatewayTopologiesOK creates a UpdateGatewayTopologiesOK with default headers values
func NewUpdateGatewayTopologiesOK() *UpdateGatewayTopologiesOK {
	return &UpdateGatewayTopologiesOK{}
}

/*UpdateGatewayTopologiesOK handles this case with default header values.

successful operation
*/
type UpdateGatewayTopologiesOK struct {
	Payload *models_cloudbreak.GatewayJSON
}

func (o *UpdateGatewayTopologiesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/stacks/{id}/cluster/gateway][%d] updateGatewayTopologiesOK  %+v", 200, o.Payload)
}

func (o *UpdateGatewayTopologiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.GatewayJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
