// Code generated by go-swagger; DO NOT EDIT.

package v1freeipauser

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// SynchronizeUserV1Reader is a Reader for the SynchronizeUserV1 structure.
type SynchronizeUserV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SynchronizeUserV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSynchronizeUserV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSynchronizeUserV1OK creates a SynchronizeUserV1OK with default headers values
func NewSynchronizeUserV1OK() *SynchronizeUserV1OK {
	return &SynchronizeUserV1OK{}
}

/*
SynchronizeUserV1OK handles this case with default header values.

successful operation
*/
type SynchronizeUserV1OK struct {
	Payload *model.SyncOperationV1Status
}

func (o *SynchronizeUserV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/user/sync][%d] synchronizeUserV1OK  %+v", 200, o.Payload)
}

func (o *SynchronizeUserV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SyncOperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
