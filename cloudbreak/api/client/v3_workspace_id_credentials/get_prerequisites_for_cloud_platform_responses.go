// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/cloudbreak/api/model"
)

// GetPrerequisitesForCloudPlatformReader is a Reader for the GetPrerequisitesForCloudPlatform structure.
type GetPrerequisitesForCloudPlatformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrerequisitesForCloudPlatformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrerequisitesForCloudPlatformOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrerequisitesForCloudPlatformOK creates a GetPrerequisitesForCloudPlatformOK with default headers values
func NewGetPrerequisitesForCloudPlatformOK() *GetPrerequisitesForCloudPlatformOK {
	return &GetPrerequisitesForCloudPlatformOK{}
}

/*GetPrerequisitesForCloudPlatformOK handles this case with default header values.

successful operation
*/
type GetPrerequisitesForCloudPlatformOK struct {
	Payload *model.CredentialPrerequisites
}

func (o *GetPrerequisitesForCloudPlatformOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/credentials/prerequisites/{cloudPlatform}][%d] getPrerequisitesForCloudPlatformOK  %+v", 200, o.Payload)
}

func (o *GetPrerequisitesForCloudPlatformOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CredentialPrerequisites)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
