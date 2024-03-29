// Code generated by go-swagger; DO NOT EDIT.

package v4carbon_dioxide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListClusterCO2V4Reader is a Reader for the ListClusterCO2V4 structure.
type ListClusterCO2V4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterCO2V4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClusterCO2V4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClusterCO2V4OK creates a ListClusterCO2V4OK with default headers values
func NewListClusterCO2V4OK() *ListClusterCO2V4OK {
	return &ListClusterCO2V4OK{}
}

/*
ListClusterCO2V4OK handles this case with default header values.

successful operation
*/
type ListClusterCO2V4OK struct {
	Payload *model.RealTimeCO2Response
}

func (o *ListClusterCO2V4OK) Error() string {
	return fmt.Sprintf("[PUT /v4/carbon_dioxide][%d] listClusterCO2V4OK  %+v", 200, o.Payload)
}

func (o *ListClusterCO2V4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RealTimeCO2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
