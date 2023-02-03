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

// GetSyncOperationStatusV1Reader is a Reader for the GetSyncOperationStatusV1 structure.
type GetSyncOperationStatusV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncOperationStatusV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSyncOperationStatusV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSyncOperationStatusV1OK creates a GetSyncOperationStatusV1OK with default headers values
func NewGetSyncOperationStatusV1OK() *GetSyncOperationStatusV1OK {
	return &GetSyncOperationStatusV1OK{}
}

/*
GetSyncOperationStatusV1OK handles this case with default header values.

successful operation
*/
type GetSyncOperationStatusV1OK struct {
	Payload *model.SyncOperationV1Status
}

func (o *GetSyncOperationStatusV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/freeipa/user/status][%d] getSyncOperationStatusV1OK  %+v", 200, o.Payload)
}

func (o *GetSyncOperationStatusV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SyncOperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
