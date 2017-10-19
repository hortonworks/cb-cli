// Code generated by go-swagger; DO NOT EDIT.

package rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetRdsReader is a Reader for the GetRds structure.
type GetRdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRdsOK creates a GetRdsOK with default headers values
func NewGetRdsOK() *GetRdsOK {
	return &GetRdsOK{}
}

/*GetRdsOK handles this case with default header values.

successful operation
*/
type GetRdsOK struct {
	Payload *models_cloudbreak.RDSConfigResponse
}

func (o *GetRdsOK) Error() string {
	return fmt.Sprintf("[GET /rdsconfigs/{id}][%d] getRdsOK  %+v", 200, o.Payload)
}

func (o *GetRdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RDSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
