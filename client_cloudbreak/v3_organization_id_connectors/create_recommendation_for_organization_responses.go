// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// CreateRecommendationForOrganizationReader is a Reader for the CreateRecommendationForOrganization structure.
type CreateRecommendationForOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRecommendationForOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRecommendationForOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRecommendationForOrganizationOK creates a CreateRecommendationForOrganizationOK with default headers values
func NewCreateRecommendationForOrganizationOK() *CreateRecommendationForOrganizationOK {
	return &CreateRecommendationForOrganizationOK{}
}

/*CreateRecommendationForOrganizationOK handles this case with default header values.

successful operation
*/
type CreateRecommendationForOrganizationOK struct {
	Payload *models_cloudbreak.RecommendationResponse
}

func (o *CreateRecommendationForOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /v3/{organizationId}/connectors/recommendation][%d] createRecommendationForOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateRecommendationForOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RecommendationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
