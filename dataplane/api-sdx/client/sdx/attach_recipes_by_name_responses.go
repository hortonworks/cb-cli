// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// AttachRecipesByNameReader is a Reader for the AttachRecipesByName structure.
type AttachRecipesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachRecipesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttachRecipesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAttachRecipesByNameOK creates a AttachRecipesByNameOK with default headers values
func NewAttachRecipesByNameOK() *AttachRecipesByNameOK {
	return &AttachRecipesByNameOK{}
}

/*
AttachRecipesByNameOK handles this case with default header values.

successful operation
*/
type AttachRecipesByNameOK struct {
	Payload *model.AttachRecipeV4Response
}

func (o *AttachRecipesByNameOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/attach_recipe][%d] attachRecipesByNameOK  %+v", 200, o.Payload)
}

func (o *AttachRecipesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AttachRecipeV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
