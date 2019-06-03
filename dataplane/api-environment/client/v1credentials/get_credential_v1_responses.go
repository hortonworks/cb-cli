// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetCredentialV1Reader is a Reader for the GetCredentialV1 structure.
type GetCredentialV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCredentialV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCredentialV1OK creates a GetCredentialV1OK with default headers values
func NewGetCredentialV1OK() *GetCredentialV1OK {
	return &GetCredentialV1OK{}
}

/*GetCredentialV1OK handles this case with default header values.

successful operation
*/
type GetCredentialV1OK struct {
	Payload *model.CredentialV1Response
}

func (o *GetCredentialV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{name}][%d] getCredentialV1OK  %+v", 200, o.Payload)
}

func (o *GetCredentialV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
