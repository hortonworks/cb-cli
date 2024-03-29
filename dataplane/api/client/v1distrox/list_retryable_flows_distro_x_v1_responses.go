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

// ListRetryableFlowsDistroXV1Reader is a Reader for the ListRetryableFlowsDistroXV1 structure.
type ListRetryableFlowsDistroXV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRetryableFlowsDistroXV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRetryableFlowsDistroXV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRetryableFlowsDistroXV1OK creates a ListRetryableFlowsDistroXV1OK with default headers values
func NewListRetryableFlowsDistroXV1OK() *ListRetryableFlowsDistroXV1OK {
	return &ListRetryableFlowsDistroXV1OK{}
}

/*
ListRetryableFlowsDistroXV1OK handles this case with default header values.

successful operation
*/
type ListRetryableFlowsDistroXV1OK struct {
	Payload []*model.RetryableFlowResponse
}

func (o *ListRetryableFlowsDistroXV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/name/{name}/retry][%d] listRetryableFlowsDistroXV1OK  %+v", 200, o.Payload)
}

func (o *ListRetryableFlowsDistroXV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
