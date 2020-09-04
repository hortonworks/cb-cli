// Code generated by go-swagger; DO NOT EDIT.

package v1alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/cloudbreak/api-as/model"
)

// UpdatePrometheusAlertReader is a Reader for the UpdatePrometheusAlert structure.
type UpdatePrometheusAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrometheusAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePrometheusAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePrometheusAlertOK creates a UpdatePrometheusAlertOK with default headers values
func NewUpdatePrometheusAlertOK() *UpdatePrometheusAlertOK {
	return &UpdatePrometheusAlertOK{}
}

/*UpdatePrometheusAlertOK handles this case with default header values.

successful operation
*/
type UpdatePrometheusAlertOK struct {
	Payload *model.PrometheusAlertResponse
}

func (o *UpdatePrometheusAlertOK) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{clusterId}/alerts/prometheus/{alertId}][%d] updatePrometheusAlertOK  %+v", 200, o.Payload)
}

func (o *UpdatePrometheusAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PrometheusAlertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
