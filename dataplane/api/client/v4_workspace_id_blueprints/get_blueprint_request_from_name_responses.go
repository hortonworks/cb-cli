// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetBlueprintRequestFromNameReader is a Reader for the GetBlueprintRequestFromName structure.
type GetBlueprintRequestFromNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintRequestFromNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlueprintRequestFromNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintRequestFromNameOK creates a GetBlueprintRequestFromNameOK with default headers values
func NewGetBlueprintRequestFromNameOK() *GetBlueprintRequestFromNameOK {
	return &GetBlueprintRequestFromNameOK{}
}

/*
GetBlueprintRequestFromNameOK handles this case with default header values.

successful operation
*/
type GetBlueprintRequestFromNameOK struct {
	Payload *model.BlueprintV4Request
}

func (o *GetBlueprintRequestFromNameOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/blueprints/{name}/request][%d] getBlueprintRequestFromNameOK  %+v", 200, o.Payload)
}

func (o *GetBlueprintRequestFromNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.BlueprintV4Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
