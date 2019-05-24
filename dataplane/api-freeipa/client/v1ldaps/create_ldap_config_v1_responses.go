// Code generated by go-swagger; DO NOT EDIT.

package v1ldaps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// CreateLdapConfigV1Reader is a Reader for the CreateLdapConfigV1 structure.
type CreateLdapConfigV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLdapConfigV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateLdapConfigV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLdapConfigV1OK creates a CreateLdapConfigV1OK with default headers values
func NewCreateLdapConfigV1OK() *CreateLdapConfigV1OK {
	return &CreateLdapConfigV1OK{}
}

/*CreateLdapConfigV1OK handles this case with default header values.

successful operation
*/
type CreateLdapConfigV1OK struct {
	Payload *model.DescribeLdapConfigV1Response
}

func (o *CreateLdapConfigV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/ldaps][%d] createLdapConfigV1OK  %+v", 200, o.Payload)
}

func (o *CreateLdapConfigV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DescribeLdapConfigV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
