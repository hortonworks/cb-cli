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

// ListRecipesByWorkspaceInternalReader is a Reader for the ListRecipesByWorkspaceInternal structure.
type ListRecipesByWorkspaceInternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRecipesByWorkspaceInternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRecipesByWorkspaceInternalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRecipesByWorkspaceInternalOK creates a ListRecipesByWorkspaceInternalOK with default headers values
func NewListRecipesByWorkspaceInternalOK() *ListRecipesByWorkspaceInternalOK {
	return &ListRecipesByWorkspaceInternalOK{}
}

/*ListRecipesByWorkspaceInternalOK handles this case with default header values.

successful operation
*/
type ListRecipesByWorkspaceInternalOK struct {
	Payload *model.RecipeViewV4Responses
}

func (o *ListRecipesByWorkspaceInternalOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/recipes/internal][%d] listRecipesByWorkspaceInternalOK  %+v", 200, o.Payload)
}

func (o *ListRecipesByWorkspaceInternalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RecipeViewV4Responses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
