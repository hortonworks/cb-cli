// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// DownscaleFreeIpaV1Reader is a Reader for the DownscaleFreeIpaV1 structure.
type DownscaleFreeIpaV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownscaleFreeIpaV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDownscaleFreeIpaV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDownscaleFreeIpaV1OK creates a DownscaleFreeIpaV1OK with default headers values
func NewDownscaleFreeIpaV1OK() *DownscaleFreeIpaV1OK {
	return &DownscaleFreeIpaV1OK{}
}

/*DownscaleFreeIpaV1OK handles this case with default header values.

successful operation
*/
type DownscaleFreeIpaV1OK struct {
	Payload *model.FreeIpaDownscaleV1Response
}

func (o *DownscaleFreeIpaV1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/freeipa/downscale][%d] downscaleFreeIpaV1OK  %+v", 200, o.Payload)
}

func (o *DownscaleFreeIpaV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FreeIpaDownscaleV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
