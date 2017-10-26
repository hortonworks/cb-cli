// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutstartStackV2Reader is a Reader for the PutstartStackV2 structure.
type PutstartStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutstartStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPutstartStackV2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPutstartStackV2Default creates a PutstartStackV2Default with default headers values
func NewPutstartStackV2Default(code int) *PutstartStackV2Default {
	return &PutstartStackV2Default{
		_statusCode: code,
	}
}

/*PutstartStackV2Default handles this case with default header values.

successful operation
*/
type PutstartStackV2Default struct {
	_statusCode int
}

// Code gets the status code for the putstart stack v2 default response
func (o *PutstartStackV2Default) Code() int {
	return o._statusCode
}

func (o *PutstartStackV2Default) Error() string {
	return fmt.Sprintf("[PUT /v2/stacks/start/{name}][%d] putstartStackV2 default ", o._statusCode)
}

func (o *PutstartStackV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
