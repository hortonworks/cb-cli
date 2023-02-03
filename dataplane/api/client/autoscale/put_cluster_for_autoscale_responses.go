// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PutClusterForAutoscaleReader is a Reader for the PutClusterForAutoscale structure.
type PutClusterForAutoscaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterForAutoscaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutClusterForAutoscaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutClusterForAutoscaleOK creates a PutClusterForAutoscaleOK with default headers values
func NewPutClusterForAutoscaleOK() *PutClusterForAutoscaleOK {
	return &PutClusterForAutoscaleOK{}
}

/*
PutClusterForAutoscaleOK handles this case with default header values.

successful operation
*/
type PutClusterForAutoscaleOK struct {
	Payload *model.FlowIdentifier
}

func (o *PutClusterForAutoscaleOK) Error() string {
	return fmt.Sprintf("[PUT /autoscale/stack/crn/{crn}/{userId}/cluster][%d] putClusterForAutoscaleOK  %+v", 200, o.Payload)
}

func (o *PutClusterForAutoscaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
