// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_ldapconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// ListLdapsByOrganizationReader is a Reader for the ListLdapsByOrganization structure.
type ListLdapsByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLdapsByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListLdapsByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListLdapsByOrganizationOK creates a ListLdapsByOrganizationOK with default headers values
func NewListLdapsByOrganizationOK() *ListLdapsByOrganizationOK {
	return &ListLdapsByOrganizationOK{}
}

/*ListLdapsByOrganizationOK handles this case with default header values.

successful operation
*/
type ListLdapsByOrganizationOK struct {
	Payload []*models_cloudbreak.LdapConfigResponse
}

func (o *ListLdapsByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v3/{organizationId}/ldapconfigs][%d] listLdapsByOrganizationOK  %+v", 200, o.Payload)
}

func (o *ListLdapsByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
