// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetRecipeRequestFromNameInWorkspaceReader is a Reader for the GetRecipeRequestFromNameInWorkspace structure.
type GetRecipeRequestFromNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecipeRequestFromNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecipeRequestFromNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecipeRequestFromNameInWorkspaceOK creates a GetRecipeRequestFromNameInWorkspaceOK with default headers values
func NewGetRecipeRequestFromNameInWorkspaceOK() *GetRecipeRequestFromNameInWorkspaceOK {
	return &GetRecipeRequestFromNameInWorkspaceOK{}
}

/*
GetRecipeRequestFromNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetRecipeRequestFromNameInWorkspaceOK struct {
	Payload *model.RecipeV4Request
}

func (o *GetRecipeRequestFromNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/recipes/{name}/request][%d] getRecipeRequestFromNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetRecipeRequestFromNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RecipeV4Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
