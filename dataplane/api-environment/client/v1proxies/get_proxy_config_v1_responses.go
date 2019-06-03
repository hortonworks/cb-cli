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

// GetProxyConfigV1Reader is a Reader for the GetProxyConfigV1 structure.
type GetProxyConfigV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxyConfigV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProxyConfigV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxyConfigV1OK creates a GetProxyConfigV1OK with default headers values
func NewGetProxyConfigV1OK() *GetProxyConfigV1OK {
	return &GetProxyConfigV1OK{}
}

/*GetProxyConfigV1OK handles this case with default header values.

successful operation
*/
type GetProxyConfigV1OK struct {
	Payload *model.ProxyResponse
}

func (o *GetProxyConfigV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/proxies/{name}][%d] getProxyConfigV1OK  %+v", 200, o.Payload)
}

func (o *GetProxyConfigV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
