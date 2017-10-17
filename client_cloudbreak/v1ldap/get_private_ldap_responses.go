// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPrivateLdapReader is a Reader for the GetPrivateLdap structure.
type GetPrivateLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateLdapOK creates a GetPrivateLdapOK with default headers values
func NewGetPrivateLdapOK() *GetPrivateLdapOK {
	return &GetPrivateLdapOK{}
}

/*GetPrivateLdapOK handles this case with default header values.

successful operation
*/
type GetPrivateLdapOK struct {
	Payload *models_cloudbreak.LdapConfigResponse
}

func (o *GetPrivateLdapOK) Error() string {
	return fmt.Sprintf("[GET /v1/ldap/user/{name}][%d] getPrivateLdapOK  %+v", 200, o.Payload)
}

func (o *GetPrivateLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.LdapConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
