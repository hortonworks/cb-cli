// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// ChangeTelemetryFeaturesInEnvironmentV1ByNameReader is a Reader for the ChangeTelemetryFeaturesInEnvironmentV1ByName structure.
type ChangeTelemetryFeaturesInEnvironmentV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeTelemetryFeaturesInEnvironmentV1ByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeTelemetryFeaturesInEnvironmentV1ByNameOK creates a ChangeTelemetryFeaturesInEnvironmentV1ByNameOK with default headers values
func NewChangeTelemetryFeaturesInEnvironmentV1ByNameOK() *ChangeTelemetryFeaturesInEnvironmentV1ByNameOK {
	return &ChangeTelemetryFeaturesInEnvironmentV1ByNameOK{}
}

/*
ChangeTelemetryFeaturesInEnvironmentV1ByNameOK handles this case with default header values.

successful operation
*/
type ChangeTelemetryFeaturesInEnvironmentV1ByNameOK struct {
	Payload *model.DetailedEnvironmentV1Response
}

func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameOK) Error() string {
	return fmt.Sprintf("[PUT /v1/env/name/{name}/change_telemetry_features][%d] changeTelemetryFeaturesInEnvironmentV1ByNameOK  %+v", 200, o.Payload)
}

func (o *ChangeTelemetryFeaturesInEnvironmentV1ByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetailedEnvironmentV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
