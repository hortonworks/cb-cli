// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_autoscale"
)

// GetAlertsReader is a Reader for the GetAlerts structure.
type GetAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlertsOK creates a GetAlertsOK with default headers values
func NewGetAlertsOK() *GetAlertsOK {
	return &GetAlertsOK{}
}

/*GetAlertsOK handles this case with default header values.

successful operation
*/
type GetAlertsOK struct {
	Payload []*models_autoscale.MetricAlert
}

func (o *GetAlertsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{clusterId}/alerts/metric][%d] getAlertsOK  %+v", 200, o.Payload)
}

func (o *GetAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
