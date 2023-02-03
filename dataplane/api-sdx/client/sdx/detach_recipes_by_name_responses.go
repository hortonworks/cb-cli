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

// DetachRecipesByNameReader is a Reader for the DetachRecipesByName structure.
type DetachRecipesByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachRecipesByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDetachRecipesByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetachRecipesByNameOK creates a DetachRecipesByNameOK with default headers values
func NewDetachRecipesByNameOK() *DetachRecipesByNameOK {
	return &DetachRecipesByNameOK{}
}

/*
DetachRecipesByNameOK handles this case with default header values.

successful operation
*/
type DetachRecipesByNameOK struct {
	Payload *model.DetachRecipeV4Response
}

func (o *DetachRecipesByNameOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/detach_recipe][%d] detachRecipesByNameOK  %+v", 200, o.Payload)
}

func (o *DetachRecipesByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DetachRecipeV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
