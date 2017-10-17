// Code generated by go-swagger; DO NOT EDIT.

package v1securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetSecurityGroupReader is a Reader for the GetSecurityGroup structure.
type GetSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupOK creates a GetSecurityGroupOK with default headers values
func NewGetSecurityGroupOK() *GetSecurityGroupOK {
	return &GetSecurityGroupOK{}
}

/*GetSecurityGroupOK handles this case with default header values.

successful operation
*/
type GetSecurityGroupOK struct {
	Payload *models_cloudbreak.SecurityGroupResponse
}

func (o *GetSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/securitygroups/{id}][%d] getSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
