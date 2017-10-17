// Code generated by go-swagger; DO NOT EDIT.

package v1recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRecipeReader is a Reader for the DeleteRecipe structure.
type DeleteRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteRecipeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteRecipeDefault creates a DeleteRecipeDefault with default headers values
func NewDeleteRecipeDefault(code int) *DeleteRecipeDefault {
	return &DeleteRecipeDefault{
		_statusCode: code,
	}
}

/*DeleteRecipeDefault handles this case with default header values.

successful operation
*/
type DeleteRecipeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete recipe default response
func (o *DeleteRecipeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRecipeDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/recipes/{id}][%d] deleteRecipe default ", o._statusCode)
}

func (o *DeleteRecipeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
