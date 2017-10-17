// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetFullClusterReader is a Reader for the GetFullCluster structure.
type GetFullClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFullClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFullClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFullClusterOK creates a GetFullClusterOK with default headers values
func NewGetFullClusterOK() *GetFullClusterOK {
	return &GetFullClusterOK{}
}

/*GetFullClusterOK handles this case with default header values.

successful operation
*/
type GetFullClusterOK struct {
	Payload *models_cloudbreak.AutoscaleClusterResponse
}

func (o *GetFullClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/{id}/cluster/full][%d] getFullClusterOK  %+v", 200, o.Payload)
}

func (o *GetFullClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.AutoscaleClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
