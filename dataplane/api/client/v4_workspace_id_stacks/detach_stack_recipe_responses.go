// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DetachStackRecipeReader is a Reader for the DetachStackRecipe structure.
type DetachStackRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachStackRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDetachStackRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetachStackRecipeOK creates a DetachStackRecipeOK with default headers values
func NewDetachStackRecipeOK() *DetachStackRecipeOK {
	return &DetachStackRecipeOK{}
}

/*DetachStackRecipeOK handles this case with default header values.

successful operation
*/
type DetachStackRecipeOK struct {
	Payload *model.DetachRecipeV4Response
}

func (o *DetachStackRecipeOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/stacks/{name}/detach_recipe][%d] detachStackRecipeOK  %+v", 200, o.Payload)
}

func (o *DetachStackRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetachRecipeV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
