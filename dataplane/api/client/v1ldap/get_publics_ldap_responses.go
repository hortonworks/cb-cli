// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPublicsLdapReader is a Reader for the GetPublicsLdap structure.
type GetPublicsLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsLdapOK creates a GetPublicsLdapOK with default headers values
func NewGetPublicsLdapOK() *GetPublicsLdapOK {
	return &GetPublicsLdapOK{}
}

/*GetPublicsLdapOK handles this case with default header values.

successful operation
*/
type GetPublicsLdapOK struct {
	Payload []*model.LdapConfigResponse
}

func (o *GetPublicsLdapOK) Error() string {
	return fmt.Sprintf("[GET /v1/ldap/account][%d] getPublicsLdapOK  %+v", 200, o.Payload)
}

func (o *GetPublicsLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
