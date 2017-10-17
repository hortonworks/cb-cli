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

// GetPrivateClusterReader is a Reader for the GetPrivateCluster structure.
type GetPrivateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateClusterOK creates a GetPrivateClusterOK with default headers values
func NewGetPrivateClusterOK() *GetPrivateClusterOK {
	return &GetPrivateClusterOK{}
}

/*GetPrivateClusterOK handles this case with default header values.

successful operation
*/
type GetPrivateClusterOK struct {
	Payload *models_cloudbreak.ClusterResponse
}

func (o *GetPrivateClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/user/{name}/cluster][%d] getPrivateClusterOK  %+v", 200, o.Payload)
}

func (o *GetPrivateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
