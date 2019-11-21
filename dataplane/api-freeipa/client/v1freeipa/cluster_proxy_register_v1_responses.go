// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ClusterProxyRegisterV1Reader is a Reader for the ClusterProxyRegisterV1 structure.
type ClusterProxyRegisterV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterProxyRegisterV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClusterProxyRegisterV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClusterProxyRegisterV1OK creates a ClusterProxyRegisterV1OK with default headers values
func NewClusterProxyRegisterV1OK() *ClusterProxyRegisterV1OK {
	return &ClusterProxyRegisterV1OK{}
}

/*ClusterProxyRegisterV1OK handles this case with default header values.

successful operation
*/
type ClusterProxyRegisterV1OK struct {
	Payload string
}

func (o *ClusterProxyRegisterV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/freeipa/cluster-proxy/register][%d] clusterProxyRegisterV1OK  %+v", 200, o.Payload)
}

func (o *ClusterProxyRegisterV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
