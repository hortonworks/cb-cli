// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
)

// AttachProxyResourceToEnvironmentsReader is a Reader for the AttachProxyResourceToEnvironments structure.
type AttachProxyResourceToEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachProxyResourceToEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttachProxyResourceToEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAttachProxyResourceToEnvironmentsOK creates a AttachProxyResourceToEnvironmentsOK with default headers values
func NewAttachProxyResourceToEnvironmentsOK() *AttachProxyResourceToEnvironmentsOK {
	return &AttachProxyResourceToEnvironmentsOK{}
}

/*AttachProxyResourceToEnvironmentsOK handles this case with default header values.

successful operation
*/
type AttachProxyResourceToEnvironmentsOK struct {
	Payload *model.ProxyConfigResponse
}

func (o *AttachProxyResourceToEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /v3/{workspaceId}/proxyconfigs/{name}/attach][%d] attachProxyResourceToEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *AttachProxyResourceToEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
