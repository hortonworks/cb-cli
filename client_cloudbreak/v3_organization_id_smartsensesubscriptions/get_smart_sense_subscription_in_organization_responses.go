// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetSmartSenseSubscriptionInOrganizationReader is a Reader for the GetSmartSenseSubscriptionInOrganization structure.
type GetSmartSenseSubscriptionInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSmartSenseSubscriptionInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSmartSenseSubscriptionInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSmartSenseSubscriptionInOrganizationOK creates a GetSmartSenseSubscriptionInOrganizationOK with default headers values
func NewGetSmartSenseSubscriptionInOrganizationOK() *GetSmartSenseSubscriptionInOrganizationOK {
	return &GetSmartSenseSubscriptionInOrganizationOK{}
}

/*GetSmartSenseSubscriptionInOrganizationOK handles this case with default header values.

successful operation
*/
type GetSmartSenseSubscriptionInOrganizationOK struct {
	Payload *models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *GetSmartSenseSubscriptionInOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v3/{organizationId}/smartsensesubscriptions/{name}][%d] getSmartSenseSubscriptionInOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetSmartSenseSubscriptionInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
