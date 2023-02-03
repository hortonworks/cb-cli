// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDistroXV1ByNameReader is a Reader for the GetDistroXV1ByName structure.
type GetDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroXV1ByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroXV1ByNameOK creates a GetDistroXV1ByNameOK with default headers values
func NewGetDistroXV1ByNameOK() *GetDistroXV1ByNameOK {
	return &GetDistroXV1ByNameOK{}
}

/*
GetDistroXV1ByNameOK handles this case with default header values.

successful operation
*/
type GetDistroXV1ByNameOK struct {
	Payload *model.StackV4Response
}

func (o *GetDistroXV1ByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/name/{name}][%d] getDistroXV1ByNameOK  %+v", 200, o.Payload)
}

func (o *GetDistroXV1ByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
