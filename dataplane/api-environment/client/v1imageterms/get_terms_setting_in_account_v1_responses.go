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

// GetTermsSettingInAccountV1Reader is a Reader for the GetTermsSettingInAccountV1 structure.
type GetTermsSettingInAccountV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTermsSettingInAccountV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTermsSettingInAccountV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTermsSettingInAccountV1OK creates a GetTermsSettingInAccountV1OK with default headers values
func NewGetTermsSettingInAccountV1OK() *GetTermsSettingInAccountV1OK {
	return &GetTermsSettingInAccountV1OK{}
}

/*
GetTermsSettingInAccountV1OK handles this case with default header values.

successful operation
*/
type GetTermsSettingInAccountV1OK struct {
	Payload *model.AzureMarketplaceTermsResponse
}

func (o *GetTermsSettingInAccountV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/imageterms/{accountId}][%d] getTermsSettingInAccountV1OK  %+v", 200, o.Payload)
}

func (o *GetTermsSettingInAccountV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AzureMarketplaceTermsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
