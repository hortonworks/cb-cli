// Code generated by go-swagger; DO NOT EDIT.

package v1securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicsSecurityGroupReader is a Reader for the GetPublicsSecurityGroup structure.
type GetPublicsSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsSecurityGroupOK creates a GetPublicsSecurityGroupOK with default headers values
func NewGetPublicsSecurityGroupOK() *GetPublicsSecurityGroupOK {
	return &GetPublicsSecurityGroupOK{}
}

/*GetPublicsSecurityGroupOK handles this case with default header values.

successful operation
*/
type GetPublicsSecurityGroupOK struct {
	Payload []*models_cloudbreak.SecurityGroupResponse
}

func (o *GetPublicsSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/securitygroups/account][%d] getPublicsSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetPublicsSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
