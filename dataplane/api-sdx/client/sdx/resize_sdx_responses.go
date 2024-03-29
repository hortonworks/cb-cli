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

// ResizeSdxReader is a Reader for the ResizeSdx structure.
type ResizeSdxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResizeSdxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResizeSdxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResizeSdxOK creates a ResizeSdxOK with default headers values
func NewResizeSdxOK() *ResizeSdxOK {
	return &ResizeSdxOK{}
}

/*
ResizeSdxOK handles this case with default header values.

successful operation
*/
type ResizeSdxOK struct {
	Payload *model.SdxClusterResponse
}

func (o *ResizeSdxOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/resize][%d] resizeSdxOK  %+v", 200, o.Payload)
}

func (o *ResizeSdxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
