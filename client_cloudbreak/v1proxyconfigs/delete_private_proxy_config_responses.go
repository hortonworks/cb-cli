// Code generated by go-swagger; DO NOT EDIT.

package v1proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePrivateProxyConfigReader is a Reader for the DeletePrivateProxyConfig structure.
type DeletePrivateProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePrivateProxyConfigDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePrivateProxyConfigDefault creates a DeletePrivateProxyConfigDefault with default headers values
func NewDeletePrivateProxyConfigDefault(code int) *DeletePrivateProxyConfigDefault {
	return &DeletePrivateProxyConfigDefault{
		_statusCode: code,
	}
}

/*DeletePrivateProxyConfigDefault handles this case with default header values.

successful operation
*/
type DeletePrivateProxyConfigDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private proxy config default response
func (o *DeletePrivateProxyConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateProxyConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/proxyconfigs/user/{name}][%d] deletePrivateProxyConfig default ", o._statusCode)
}

func (o *DeletePrivateProxyConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
