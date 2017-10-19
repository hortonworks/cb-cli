// Code generated by go-swagger; DO NOT EDIT.

package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPublicSecurityGroupReader is a Reader for the PostPublicSecurityGroup structure.
type PostPublicSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublicSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPublicSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPublicSecurityGroupOK creates a PostPublicSecurityGroupOK with default headers values
func NewPostPublicSecurityGroupOK() *PostPublicSecurityGroupOK {
	return &PostPublicSecurityGroupOK{}
}

/*PostPublicSecurityGroupOK handles this case with default header values.

successful operation
*/
type PostPublicSecurityGroupOK struct {
	Payload *models_cloudbreak.SecurityGroupResponse
}

func (o *PostPublicSecurityGroupOK) Error() string {
	return fmt.Sprintf("[POST /securitygroups/account][%d] postPublicSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *PostPublicSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
