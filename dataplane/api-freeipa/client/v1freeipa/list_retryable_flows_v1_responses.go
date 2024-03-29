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

// ListRetryableFlowsV1Reader is a Reader for the ListRetryableFlowsV1 structure.
type ListRetryableFlowsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRetryableFlowsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRetryableFlowsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRetryableFlowsV1OK creates a ListRetryableFlowsV1OK with default headers values
func NewListRetryableFlowsV1OK() *ListRetryableFlowsV1OK {
	return &ListRetryableFlowsV1OK{}
}

/*
ListRetryableFlowsV1OK handles this case with default header values.

successful operation
*/
type ListRetryableFlowsV1OK struct {
	Payload []*model.RetryableFlowResponse
}

func (o *ListRetryableFlowsV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/freeipa/retry][%d] listRetryableFlowsV1OK  %+v", 200, o.Payload)
}

func (o *ListRetryableFlowsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
