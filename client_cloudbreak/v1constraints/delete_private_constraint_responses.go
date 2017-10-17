// Code generated by go-swagger; DO NOT EDIT.

package v1constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateConstraintReader is a Reader for the DeletePrivateConstraint structure.
type DeletePrivateConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateConstraintDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateConstraintDefault creates a DeletePrivateConstraintDefault with default headers values
func NewDeletePrivateConstraintDefault(code int) *DeletePrivateConstraintDefault {
	return &DeletePrivateConstraintDefault{
		_statusCode: code,
	}
}

/*DeletePrivateConstraintDefault handles this case with default header values.

successful operation
*/
type DeletePrivateConstraintDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private constraint default response
func (o *DeletePrivateConstraintDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateConstraintDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/constraints/user/{name}][%d] deletePrivateConstraint default ", o._statusCode)
}

func (o *DeletePrivateConstraintDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
