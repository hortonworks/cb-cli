// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/*GetRegionsOK handles this case with default header values.

successful operation
*/
type GetRegionsOK struct {
	Payload *models_cloudbreak.PlatformRegionsJSON
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /connectors/regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformRegionsJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
