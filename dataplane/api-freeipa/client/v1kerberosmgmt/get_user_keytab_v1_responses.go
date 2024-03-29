// Code generated by go-swagger; DO NOT EDIT.

package v1kerberosmgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetUserKeytabV1Reader is a Reader for the GetUserKeytabV1 structure.
type GetUserKeytabV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserKeytabV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserKeytabV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserKeytabV1OK creates a GetUserKeytabV1OK with default headers values
func NewGetUserKeytabV1OK() *GetUserKeytabV1OK {
	return &GetUserKeytabV1OK{}
}

/*
GetUserKeytabV1OK handles this case with default header values.

successful operation
*/
type GetUserKeytabV1OK struct {
	Payload string
}

func (o *GetUserKeytabV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/kerberosmgmt/userkeytab][%d] getUserKeytabV1OK  %+v", 200, o.Payload)
}

func (o *GetUserKeytabV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
