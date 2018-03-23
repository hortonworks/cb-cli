// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostLdapConnectionTestReader is a Reader for the PostLdapConnectionTest structure.
type PostLdapConnectionTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLdapConnectionTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLdapConnectionTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLdapConnectionTestOK creates a PostLdapConnectionTestOK with default headers values
func NewPostLdapConnectionTestOK() *PostLdapConnectionTestOK {
	return &PostLdapConnectionTestOK{}
}

/*PostLdapConnectionTestOK handles this case with default header values.

successful operation
*/
type PostLdapConnectionTestOK struct {
	Payload *models_cloudbreak.LdapTestResult
}

func (o *PostLdapConnectionTestOK) Error() string {
	return fmt.Sprintf("[POST /v1/ldap/testconnect][%d] postLdapConnectionTestOK  %+v", 200, o.Payload)
}

func (o *PostLdapConnectionTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.LdapTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
