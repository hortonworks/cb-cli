// Code generated by go-swagger; DO NOT EDIT.

package v1networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicNetworkReader is a Reader for the GetPublicNetwork structure.
type GetPublicNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicNetworkOK creates a GetPublicNetworkOK with default headers values
func NewGetPublicNetworkOK() *GetPublicNetworkOK {
	return &GetPublicNetworkOK{}
}

/*GetPublicNetworkOK handles this case with default header values.

successful operation
*/
type GetPublicNetworkOK struct {
	Payload *models_cloudbreak.NetworkResponse
}

func (o *GetPublicNetworkOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/account/{name}][%d] getPublicNetworkOK  %+v", 200, o.Payload)
}

func (o *GetPublicNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}