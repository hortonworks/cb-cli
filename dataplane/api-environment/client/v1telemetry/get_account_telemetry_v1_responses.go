// Code generated by go-swagger; DO NOT EDIT.

package v1telemetry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetAccountTelemetryV1Reader is a Reader for the GetAccountTelemetryV1 structure.
type GetAccountTelemetryV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountTelemetryV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountTelemetryV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountTelemetryV1OK creates a GetAccountTelemetryV1OK with default headers values
func NewGetAccountTelemetryV1OK() *GetAccountTelemetryV1OK {
	return &GetAccountTelemetryV1OK{}
}

/*GetAccountTelemetryV1OK handles this case with default header values.

successful operation
*/
type GetAccountTelemetryV1OK struct {
	Payload *model.AccountTelemetryResponse
}

func (o *GetAccountTelemetryV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/telemetry][%d] getAccountTelemetryV1OK  %+v", 200, o.Payload)
}

func (o *GetAccountTelemetryV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AccountTelemetryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
