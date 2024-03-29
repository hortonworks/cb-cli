// Code generated by go-swagger; DO NOT EDIT.

package internalsdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// RenewInternalSdxCertificateReader is a Reader for the RenewInternalSdxCertificate structure.
type RenewInternalSdxCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewInternalSdxCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRenewInternalSdxCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenewInternalSdxCertificateOK creates a RenewInternalSdxCertificateOK with default headers values
func NewRenewInternalSdxCertificateOK() *RenewInternalSdxCertificateOK {
	return &RenewInternalSdxCertificateOK{}
}

/*
RenewInternalSdxCertificateOK handles this case with default header values.

successful operation
*/
type RenewInternalSdxCertificateOK struct {
	Payload *model.FlowIdentifier
}

func (o *RenewInternalSdxCertificateOK) Error() string {
	return fmt.Sprintf("[POST /internal/sdx/crn/{crn}/renew_certificate][%d] renewInternalSdxCertificateOK  %+v", 200, o.Payload)
}

func (o *RenewInternalSdxCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
