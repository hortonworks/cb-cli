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

// UpscaleFreeIpaV1Reader is a Reader for the UpscaleFreeIpaV1 structure.
type UpscaleFreeIpaV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpscaleFreeIpaV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpscaleFreeIpaV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpscaleFreeIpaV1OK creates a UpscaleFreeIpaV1OK with default headers values
func NewUpscaleFreeIpaV1OK() *UpscaleFreeIpaV1OK {
	return &UpscaleFreeIpaV1OK{}
}

/*UpscaleFreeIpaV1OK handles this case with default header values.

successful operation
*/
type UpscaleFreeIpaV1OK struct {
	Payload *model.FreeIpaUpscaleV1Response
}

func (o *UpscaleFreeIpaV1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/freeipa/upscale][%d] upscaleFreeIpaV1OK  %+v", 200, o.Payload)
}

func (o *UpscaleFreeIpaV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FreeIpaUpscaleV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
