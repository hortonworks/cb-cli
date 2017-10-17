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

// GetLdapReader is a Reader for the GetLdap structure.
type GetLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLdapOK creates a GetLdapOK with default headers values
func NewGetLdapOK() *GetLdapOK {
	return &GetLdapOK{}
}

/*GetLdapOK handles this case with default header values.

successful operation
*/
type GetLdapOK struct {
	Payload *models_cloudbreak.LdapConfigResponse
}

func (o *GetLdapOK) Error() string {
	return fmt.Sprintf("[GET /v1/ldap/{id}][%d] getLdapOK  %+v", 200, o.Payload)
}

func (o *GetLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.LdapConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
