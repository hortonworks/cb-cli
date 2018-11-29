// Code generated by go-swagger; DO NOT EDIT.

package v1proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPublicsProxyConfigReader is a Reader for the GetPublicsProxyConfig structure.
type GetPublicsProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsProxyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsProxyConfigOK creates a GetPublicsProxyConfigOK with default headers values
func NewGetPublicsProxyConfigOK() *GetPublicsProxyConfigOK {
	return &GetPublicsProxyConfigOK{}
}

/*GetPublicsProxyConfigOK handles this case with default header values.

successful operation
*/
type GetPublicsProxyConfigOK struct {
	Payload []*model.ProxyConfigResponse
}

func (o *GetPublicsProxyConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/proxyconfigs/account][%d] getPublicsProxyConfigOK  %+v", 200, o.Payload)
}

func (o *GetPublicsProxyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
