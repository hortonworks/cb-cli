// Code generated by go-swagger; DO NOT EDIT.

package v1flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutPublicDefaultFlexSubscriptionByNameReader is a Reader for the PutPublicDefaultFlexSubscriptionByName structure.
type PutPublicDefaultFlexSubscriptionByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPublicDefaultFlexSubscriptionByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutPublicDefaultFlexSubscriptionByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutPublicDefaultFlexSubscriptionByNameDefault creates a PutPublicDefaultFlexSubscriptionByNameDefault with default headers values
func NewPutPublicDefaultFlexSubscriptionByNameDefault(code int) *PutPublicDefaultFlexSubscriptionByNameDefault {
	return &PutPublicDefaultFlexSubscriptionByNameDefault{
		_statusCode: code,
	}
}

/*PutPublicDefaultFlexSubscriptionByNameDefault handles this case with default header values.

successful operation
*/
type PutPublicDefaultFlexSubscriptionByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the put public default flex subscription by name default response
func (o *PutPublicDefaultFlexSubscriptionByNameDefault) Code() int {
	return o._statusCode
}

func (o *PutPublicDefaultFlexSubscriptionByNameDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/flexsubscriptions/account/setdefault/{name}][%d] putPublicDefaultFlexSubscriptionByName default ", o._statusCode)
}

func (o *PutPublicDefaultFlexSubscriptionByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
