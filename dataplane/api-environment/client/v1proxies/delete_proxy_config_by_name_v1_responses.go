// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// DeleteProxyConfigByNameV1Reader is a Reader for the DeleteProxyConfigByNameV1 structure.
type DeleteProxyConfigByNameV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProxyConfigByNameV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProxyConfigByNameV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProxyConfigByNameV1OK creates a DeleteProxyConfigByNameV1OK with default headers values
func NewDeleteProxyConfigByNameV1OK() *DeleteProxyConfigByNameV1OK {
	return &DeleteProxyConfigByNameV1OK{}
}

/*
DeleteProxyConfigByNameV1OK handles this case with default header values.

successful operation
*/
type DeleteProxyConfigByNameV1OK struct {
	Payload *model.ProxyResponse
}

func (o *DeleteProxyConfigByNameV1OK) Error() string {
	return fmt.Sprintf("[DELETE /v1/proxies/name/{name}][%d] deleteProxyConfigByNameV1OK  %+v", 200, o.Payload)
}

func (o *DeleteProxyConfigByNameV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
