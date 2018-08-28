// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// DeleteCredentialInOrganizationReader is a Reader for the DeleteCredentialInOrganization structure.
type DeleteCredentialInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCredentialInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCredentialInOrganizationOK creates a DeleteCredentialInOrganizationOK with default headers values
func NewDeleteCredentialInOrganizationOK() *DeleteCredentialInOrganizationOK {
	return &DeleteCredentialInOrganizationOK{}
}

/*DeleteCredentialInOrganizationOK handles this case with default header values.

successful operation
*/
type DeleteCredentialInOrganizationOK struct {
	Payload *models_cloudbreak.CredentialResponse
}

func (o *DeleteCredentialInOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /v3/{organizationId}/credentials/{name}][%d] deleteCredentialInOrganizationOK  %+v", 200, o.Payload)
}

func (o *DeleteCredentialInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.CredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}