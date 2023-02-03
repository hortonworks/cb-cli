// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-redbeams/model"
)

// GetFlowLogsByFlowIdsReader is a Reader for the GetFlowLogsByFlowIds structure.
type GetFlowLogsByFlowIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLogsByFlowIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowLogsByFlowIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowLogsByFlowIdsOK creates a GetFlowLogsByFlowIdsOK with default headers values
func NewGetFlowLogsByFlowIdsOK() *GetFlowLogsByFlowIdsOK {
	return &GetFlowLogsByFlowIdsOK{}
}

/*
GetFlowLogsByFlowIdsOK handles this case with default header values.

successful operation
*/
type GetFlowLogsByFlowIdsOK struct {
	Payload *model.PageFlowLogResponse
}

func (o *GetFlowLogsByFlowIdsOK) Error() string {
	return fmt.Sprintf("[GET /flow/logs/flowIds][%d] getFlowLogsByFlowIdsOK  %+v", 200, o.Payload)
}

func (o *GetFlowLogsByFlowIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PageFlowLogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
