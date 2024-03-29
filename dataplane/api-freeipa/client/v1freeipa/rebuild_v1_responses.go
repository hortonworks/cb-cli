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

// RebuildV1Reader is a Reader for the RebuildV1 structure.
type RebuildV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebuildV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRebuildV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRebuildV1OK creates a RebuildV1OK with default headers values
func NewRebuildV1OK() *RebuildV1OK {
	return &RebuildV1OK{}
}

/*
RebuildV1OK handles this case with default header values.

successful operation
*/
type RebuildV1OK struct {
	Payload *model.DescribeFreeIpaV1Response
}

func (o *RebuildV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/rebuild][%d] rebuildV1OK  %+v", 200, o.Payload)
}

func (o *RebuildV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DescribeFreeIpaV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
