// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints_util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetGeneratedTemplateReader is a Reader for the GetGeneratedTemplate structure.
type GetGeneratedTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGeneratedTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGeneratedTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGeneratedTemplateOK creates a GetGeneratedTemplateOK with default headers values
func NewGetGeneratedTemplateOK() *GetGeneratedTemplateOK {
	return &GetGeneratedTemplateOK{}
}

/*
GetGeneratedTemplateOK handles this case with default header values.

successful operation
*/
type GetGeneratedTemplateOK struct {
	Payload *model.GeneratedCmTemplateV4Response
}

func (o *GetGeneratedTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/blueprints_util/generate][%d] getGeneratedTemplateOK  %+v", 200, o.Payload)
}

func (o *GetGeneratedTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GeneratedCmTemplateV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
