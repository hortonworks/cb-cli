// Code generated by go-swagger; DO NOT EDIT.

package v1diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// CollectFreeIpaDiagnosticsV1Reader is a Reader for the CollectFreeIpaDiagnosticsV1 structure.
type CollectFreeIpaDiagnosticsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectFreeIpaDiagnosticsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollectFreeIpaDiagnosticsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollectFreeIpaDiagnosticsV1OK creates a CollectFreeIpaDiagnosticsV1OK with default headers values
func NewCollectFreeIpaDiagnosticsV1OK() *CollectFreeIpaDiagnosticsV1OK {
	return &CollectFreeIpaDiagnosticsV1OK{}
}

/*CollectFreeIpaDiagnosticsV1OK handles this case with default header values.

successful operation
*/
type CollectFreeIpaDiagnosticsV1OK struct {
	Payload *model.FlowIdentifier
}

func (o *CollectFreeIpaDiagnosticsV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/diagnostics][%d] collectFreeIpaDiagnosticsV1OK  %+v", 200, o.Payload)
}

func (o *CollectFreeIpaDiagnosticsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
