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

// InternalModifyProxyConfigByEnvironmentV1Reader is a Reader for the InternalModifyProxyConfigByEnvironmentV1 structure.
type InternalModifyProxyConfigByEnvironmentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalModifyProxyConfigByEnvironmentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInternalModifyProxyConfigByEnvironmentV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInternalModifyProxyConfigByEnvironmentV1OK creates a InternalModifyProxyConfigByEnvironmentV1OK with default headers values
func NewInternalModifyProxyConfigByEnvironmentV1OK() *InternalModifyProxyConfigByEnvironmentV1OK {
	return &InternalModifyProxyConfigByEnvironmentV1OK{}
}

/*
InternalModifyProxyConfigByEnvironmentV1OK handles this case with default header values.

successful operation
*/
type InternalModifyProxyConfigByEnvironmentV1OK struct {
	Payload *model.OperationV1Status
}

func (o *InternalModifyProxyConfigByEnvironmentV1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/freeipa/internal/modify_proxy][%d] internalModifyProxyConfigByEnvironmentV1OK  %+v", 200, o.Payload)
}

func (o *InternalModifyProxyConfigByEnvironmentV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
