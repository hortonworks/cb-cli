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

// InternalGetFreeIpaByEnvironmentV1Reader is a Reader for the InternalGetFreeIpaByEnvironmentV1 structure.
type InternalGetFreeIpaByEnvironmentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalGetFreeIpaByEnvironmentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInternalGetFreeIpaByEnvironmentV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInternalGetFreeIpaByEnvironmentV1OK creates a InternalGetFreeIpaByEnvironmentV1OK with default headers values
func NewInternalGetFreeIpaByEnvironmentV1OK() *InternalGetFreeIpaByEnvironmentV1OK {
	return &InternalGetFreeIpaByEnvironmentV1OK{}
}

/*
InternalGetFreeIpaByEnvironmentV1OK handles this case with default header values.

successful operation
*/
type InternalGetFreeIpaByEnvironmentV1OK struct {
	Payload *model.DescribeFreeIpaV1Response
}

func (o *InternalGetFreeIpaByEnvironmentV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/freeipa/internal][%d] internalGetFreeIpaByEnvironmentV1OK  %+v", 200, o.Payload)
}

func (o *InternalGetFreeIpaByEnvironmentV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DescribeFreeIpaV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
