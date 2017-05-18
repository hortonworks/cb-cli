package smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteSmartsensesubscriptionsUserSubscriptionIDReader is a Reader for the DeleteSmartsensesubscriptionsUserSubscriptionID structure.
type DeleteSmartsensesubscriptionsUserSubscriptionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteSmartsensesubscriptionsUserSubscriptionIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteSmartsensesubscriptionsUserSubscriptionIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteSmartsensesubscriptionsUserSubscriptionIDDefault creates a DeleteSmartsensesubscriptionsUserSubscriptionIDDefault with default headers values
func NewDeleteSmartsensesubscriptionsUserSubscriptionIDDefault(code int) *DeleteSmartsensesubscriptionsUserSubscriptionIDDefault {
	return &DeleteSmartsensesubscriptionsUserSubscriptionIDDefault{
		_statusCode: code,
	}
}

/*DeleteSmartsensesubscriptionsUserSubscriptionIDDefault handles this case with default header values.

successful operation
*/
type DeleteSmartsensesubscriptionsUserSubscriptionIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete smartsensesubscriptions user subscription ID default response
func (o *DeleteSmartsensesubscriptionsUserSubscriptionIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSmartsensesubscriptionsUserSubscriptionIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /smartsensesubscriptions/user/{subscriptionId}][%d] DeleteSmartsensesubscriptionsUserSubscriptionID default ", o._statusCode)
}

func (o *DeleteSmartsensesubscriptionsUserSubscriptionIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
