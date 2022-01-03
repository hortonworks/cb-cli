// Code generated by go-swagger; DO NOT EDIT.

package v1internaldistrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RenewInternalDistroXCertificateReader is a Reader for the RenewInternalDistroXCertificate structure.
type RenewInternalDistroXCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewInternalDistroXCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRenewInternalDistroXCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenewInternalDistroXCertificateOK creates a RenewInternalDistroXCertificateOK with default headers values
func NewRenewInternalDistroXCertificateOK() *RenewInternalDistroXCertificateOK {
	return &RenewInternalDistroXCertificateOK{}
}

/*RenewInternalDistroXCertificateOK handles this case with default header values.

successful operation
*/
type RenewInternalDistroXCertificateOK struct {
	Payload *model.FlowIdentifier
}

func (o *RenewInternalDistroXCertificateOK) Error() string {
	return fmt.Sprintf("[POST /v1/internal/distrox/crn/{crn}/renew_certificate][%d] renewInternalDistroXCertificateOK  %+v", 200, o.Payload)
}

func (o *RenewInternalDistroXCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
