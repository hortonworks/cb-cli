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

// PostPrivateSecurityGroupReader is a Reader for the PostPrivateSecurityGroup structure.
type PostPrivateSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateSecurityGroupOK creates a PostPrivateSecurityGroupOK with default headers values
func NewPostPrivateSecurityGroupOK() *PostPrivateSecurityGroupOK {
	return &PostPrivateSecurityGroupOK{}
}

/*PostPrivateSecurityGroupOK handles this case with default header values.

successful operation
*/
type PostPrivateSecurityGroupOK struct {
	Payload *models_cloudbreak.SecurityGroupResponse
}

func (o *PostPrivateSecurityGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/securitygroups/user][%d] postPrivateSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *PostPrivateSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
