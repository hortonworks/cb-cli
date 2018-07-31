// Code generated by go-swagger; DO NOT EDIT.

package v1organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// UpdateOrganizationUsersReader is a Reader for the UpdateOrganizationUsers structure.
type UpdateOrganizationUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOrganizationUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrganizationUsersOK creates a UpdateOrganizationUsersOK with default headers values
func NewUpdateOrganizationUsersOK() *UpdateOrganizationUsersOK {
	return &UpdateOrganizationUsersOK{}
}

/*UpdateOrganizationUsersOK handles this case with default header values.

successful operation
*/
type UpdateOrganizationUsersOK struct {
	Payload []*models_cloudbreak.UserResponseJSON
}

func (o *UpdateOrganizationUsersOK) Error() string {
	return fmt.Sprintf("[PUT /v1/organizations/name/{name}/updateUsers][%d] updateOrganizationUsersOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
