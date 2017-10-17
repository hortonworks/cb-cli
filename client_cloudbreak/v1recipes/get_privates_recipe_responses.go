// Code generated by go-swagger; DO NOT EDIT.

package v1recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivatesRecipeReader is a Reader for the GetPrivatesRecipe structure.
type GetPrivatesRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesRecipeOK creates a GetPrivatesRecipeOK with default headers values
func NewGetPrivatesRecipeOK() *GetPrivatesRecipeOK {
	return &GetPrivatesRecipeOK{}
}

/*GetPrivatesRecipeOK handles this case with default header values.

successful operation
*/
type GetPrivatesRecipeOK struct {
	Payload []*models_cloudbreak.RecipeResponse
}

func (o *GetPrivatesRecipeOK) Error() string {
	return fmt.Sprintf("[GET /v1/recipes/user][%d] getPrivatesRecipeOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
