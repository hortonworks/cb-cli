// Code generated by go-swagger; DO NOT EDIT.

package v1operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetOperationStatusV1Reader is a Reader for the GetOperationStatusV1 structure.
type GetOperationStatusV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperationStatusV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperationStatusV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOperationStatusV1OK creates a GetOperationStatusV1OK with default headers values
func NewGetOperationStatusV1OK() *GetOperationStatusV1OK {
	return &GetOperationStatusV1OK{}
}

/*
GetOperationStatusV1OK handles this case with default header values.

successful operation
*/
type GetOperationStatusV1OK struct {
	Payload *model.OperationV1Status
}

func (o *GetOperationStatusV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/operation][%d] getOperationStatusV1OK  %+v", 200, o.Payload)
}

func (o *GetOperationStatusV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.OperationV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
