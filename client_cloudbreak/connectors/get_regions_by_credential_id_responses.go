// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetRegionsByCredentialIDReader is a Reader for the GetRegionsByCredentialID structure.
type GetRegionsByCredentialIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsByCredentialIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsByCredentialIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegionsByCredentialIDOK creates a GetRegionsByCredentialIDOK with default headers values
func NewGetRegionsByCredentialIDOK() *GetRegionsByCredentialIDOK {
	return &GetRegionsByCredentialIDOK{}
}

/*GetRegionsByCredentialIDOK handles this case with default header values.

successful operation
*/
type GetRegionsByCredentialIDOK struct {
	Payload *models_cloudbreak.RegionResponse
}

func (o *GetRegionsByCredentialIDOK) Error() string {
	return fmt.Sprintf("[POST /connectors/v2/regions][%d] getRegionsByCredentialIdOK  %+v", 200, o.Payload)
}

func (o *GetRegionsByCredentialIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RegionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
