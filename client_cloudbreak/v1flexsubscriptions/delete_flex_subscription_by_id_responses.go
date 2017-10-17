// Code generated by go-swagger; DO NOT EDIT.

package v1flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFlexSubscriptionByIDReader is a Reader for the DeleteFlexSubscriptionByID structure.
type DeleteFlexSubscriptionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlexSubscriptionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteFlexSubscriptionByIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteFlexSubscriptionByIDDefault creates a DeleteFlexSubscriptionByIDDefault with default headers values
func NewDeleteFlexSubscriptionByIDDefault(code int) *DeleteFlexSubscriptionByIDDefault {
	return &DeleteFlexSubscriptionByIDDefault{
		_statusCode: code,
	}
}

/*DeleteFlexSubscriptionByIDDefault handles this case with default header values.

successful operation
*/
type DeleteFlexSubscriptionByIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete flex subscription by Id default response
func (o *DeleteFlexSubscriptionByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFlexSubscriptionByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/flexsubscriptions/{id}][%d] deleteFlexSubscriptionById default ", o._statusCode)
}

func (o *DeleteFlexSubscriptionByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
