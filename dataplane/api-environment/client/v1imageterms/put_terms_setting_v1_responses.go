// Code generated by go-swagger; DO NOT EDIT.

package v1imageterms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// PutTermsSettingV1Reader is a Reader for the PutTermsSettingV1 structure.
type PutTermsSettingV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTermsSettingV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutTermsSettingV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutTermsSettingV1OK creates a PutTermsSettingV1OK with default headers values
func NewPutTermsSettingV1OK() *PutTermsSettingV1OK {
	return &PutTermsSettingV1OK{}
}

/*
PutTermsSettingV1OK handles this case with default header values.

successful operation
*/
type PutTermsSettingV1OK struct {
	Payload *model.AzureMarketplaceTermsResponse
}

func (o *PutTermsSettingV1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/imageterms][%d] putTermsSettingV1OK  %+v", 200, o.Payload)
}

func (o *PutTermsSettingV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AzureMarketplaceTermsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
