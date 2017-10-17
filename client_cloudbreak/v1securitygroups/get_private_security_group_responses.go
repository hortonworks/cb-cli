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

// GetPrivateSecurityGroupReader is a Reader for the GetPrivateSecurityGroup structure.
type GetPrivateSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSecurityGroupOK creates a GetPrivateSecurityGroupOK with default headers values
func NewGetPrivateSecurityGroupOK() *GetPrivateSecurityGroupOK {
	return &GetPrivateSecurityGroupOK{}
}

/*GetPrivateSecurityGroupOK handles this case with default header values.

successful operation
*/
type GetPrivateSecurityGroupOK struct {
	Payload *models_cloudbreak.SecurityGroupResponse
}

func (o *GetPrivateSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/securitygroups/user/{name}][%d] getPrivateSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
