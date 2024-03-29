// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetResourceGroupsByEnvReader is a Reader for the GetResourceGroupsByEnv structure.
type GetResourceGroupsByEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceGroupsByEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResourceGroupsByEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResourceGroupsByEnvOK creates a GetResourceGroupsByEnvOK with default headers values
func NewGetResourceGroupsByEnvOK() *GetResourceGroupsByEnvOK {
	return &GetResourceGroupsByEnvOK{}
}

/*
GetResourceGroupsByEnvOK handles this case with default header values.

successful operation
*/
type GetResourceGroupsByEnvOK struct {
	Payload *model.PlatformResourceGroupsResponse
}

func (o *GetResourceGroupsByEnvOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/platform_resources/resource_groups][%d] getResourceGroupsByEnvOK  %+v", 200, o.Payload)
}

func (o *GetResourceGroupsByEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformResourceGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
