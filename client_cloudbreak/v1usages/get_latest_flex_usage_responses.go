// Code generated by go-swagger; DO NOT EDIT.

package v1usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetLatestFlexUsageReader is a Reader for the GetLatestFlexUsage structure.
type GetLatestFlexUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestFlexUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLatestFlexUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLatestFlexUsageOK creates a GetLatestFlexUsageOK with default headers values
func NewGetLatestFlexUsageOK() *GetLatestFlexUsageOK {
	return &GetLatestFlexUsageOK{}
}

/*GetLatestFlexUsageOK handles this case with default header values.

successful operation
*/
type GetLatestFlexUsageOK struct {
	Payload *models_cloudbreak.CloudbreakFlexUsage
}

func (o *GetLatestFlexUsageOK) Error() string {
	return fmt.Sprintf("[GET /v1/usages/flex/latest][%d] getLatestFlexUsageOK  %+v", 200, o.Payload)
}

func (o *GetLatestFlexUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.CloudbreakFlexUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
