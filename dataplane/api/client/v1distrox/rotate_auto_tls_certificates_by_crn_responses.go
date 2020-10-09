// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RotateAutoTLSCertificatesByCrnReader is a Reader for the RotateAutoTLSCertificatesByCrn structure.
type RotateAutoTLSCertificatesByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateAutoTLSCertificatesByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRotateAutoTLSCertificatesByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRotateAutoTLSCertificatesByCrnOK creates a RotateAutoTLSCertificatesByCrnOK with default headers values
func NewRotateAutoTLSCertificatesByCrnOK() *RotateAutoTLSCertificatesByCrnOK {
	return &RotateAutoTLSCertificatesByCrnOK{}
}

/*RotateAutoTLSCertificatesByCrnOK handles this case with default header values.

successful operation
*/
type RotateAutoTLSCertificatesByCrnOK struct {
	Payload *model.CertificatesRotationV4Response
}

func (o *RotateAutoTLSCertificatesByCrnOK) Error() string {
	return fmt.Sprintf("[PUT /v1/distrox/crn/{crn}/rotate_autotls_certificates][%d] rotateAutoTlsCertificatesByCrnOK  %+v", 200, o.Payload)
}

func (o *RotateAutoTLSCertificatesByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CertificatesRotationV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
