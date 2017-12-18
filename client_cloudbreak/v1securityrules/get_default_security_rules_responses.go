// Code generated by go-swagger; DO NOT EDIT.

package v1securityrules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetDefaultSecurityRulesReader is a Reader for the GetDefaultSecurityRules structure.
type GetDefaultSecurityRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultSecurityRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDefaultSecurityRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDefaultSecurityRulesOK creates a GetDefaultSecurityRulesOK with default headers values
func NewGetDefaultSecurityRulesOK() *GetDefaultSecurityRulesOK {
	return &GetDefaultSecurityRulesOK{}
}

/*GetDefaultSecurityRulesOK handles this case with default header values.

successful operation
*/
type GetDefaultSecurityRulesOK struct {
	Payload *models_cloudbreak.SecurityRulesResponse
}

func (o *GetDefaultSecurityRulesOK) Error() string {
	return fmt.Sprintf("[GET /v1/securityrules/defaultsecurityrules][%d] getDefaultSecurityRulesOK  %+v", 200, o.Payload)
}

func (o *GetDefaultSecurityRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
