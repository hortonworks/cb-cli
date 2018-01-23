// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetBlueprintRequestFromIDReader is a Reader for the GetBlueprintRequestFromID structure.
type GetBlueprintRequestFromIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintRequestFromIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlueprintRequestFromIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintRequestFromIDOK creates a GetBlueprintRequestFromIDOK with default headers values
func NewGetBlueprintRequestFromIDOK() *GetBlueprintRequestFromIDOK {
	return &GetBlueprintRequestFromIDOK{}
}

/*GetBlueprintRequestFromIDOK handles this case with default header values.

successful operation
*/
type GetBlueprintRequestFromIDOK struct {
	Payload *models_cloudbreak.BlueprintRequest
}

func (o *GetBlueprintRequestFromIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/blueprints/{id}/request][%d] getBlueprintRequestFromIdOK  %+v", 200, o.Payload)
}

func (o *GetBlueprintRequestFromIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.BlueprintRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
