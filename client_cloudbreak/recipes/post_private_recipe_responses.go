// Code generated by go-swagger; DO NOT EDIT.

package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPrivateRecipeReader is a Reader for the PostPrivateRecipe structure.
type PostPrivateRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateRecipeOK creates a PostPrivateRecipeOK with default headers values
func NewPostPrivateRecipeOK() *PostPrivateRecipeOK {
	return &PostPrivateRecipeOK{}
}

/*PostPrivateRecipeOK handles this case with default header values.

successful operation
*/
type PostPrivateRecipeOK struct {
	Payload *models_cloudbreak.RecipeResponse
}

func (o *PostPrivateRecipeOK) Error() string {
	return fmt.Sprintf("[POST /recipes/user][%d] postPrivateRecipeOK  %+v", 200, o.Payload)
}

func (o *PostPrivateRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RecipeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
