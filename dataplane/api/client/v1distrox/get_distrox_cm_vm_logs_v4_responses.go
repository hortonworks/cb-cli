// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDistroxCmVMLogsV4Reader is a Reader for the GetDistroxCmVMLogsV4 structure.
type GetDistroxCmVMLogsV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroxCmVMLogsV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroxCmVMLogsV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroxCmVMLogsV4OK creates a GetDistroxCmVMLogsV4OK with default headers values
func NewGetDistroxCmVMLogsV4OK() *GetDistroxCmVMLogsV4OK {
	return &GetDistroxCmVMLogsV4OK{}
}

/*
GetDistroxCmVMLogsV4OK handles this case with default header values.

successful operation
*/
type GetDistroxCmVMLogsV4OK struct {
	Payload *model.VMLogsResponse
}

func (o *GetDistroxCmVMLogsV4OK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/diagnostics/logs][%d] getDistroxCmVmLogsV4OK  %+v", 200, o.Payload)
}

func (o *GetDistroxCmVMLogsV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.VMLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
