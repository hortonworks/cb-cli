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

// GetClusterRecoverableByNameReader is a Reader for the GetClusterRecoverableByName structure.
type GetClusterRecoverableByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterRecoverableByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterRecoverableByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterRecoverableByNameOK creates a GetClusterRecoverableByNameOK with default headers values
func NewGetClusterRecoverableByNameOK() *GetClusterRecoverableByNameOK {
	return &GetClusterRecoverableByNameOK{}
}

/*
GetClusterRecoverableByNameOK handles this case with default header values.

successful operation
*/
type GetClusterRecoverableByNameOK struct {
	Payload *model.SdxRecoverableResponse
}

func (o *GetClusterRecoverableByNameOK) Error() string {
	return fmt.Sprintf("[GET /sdx/{name}/recoverable][%d] getClusterRecoverableByNameOK  %+v", 200, o.Payload)
}

func (o *GetClusterRecoverableByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxRecoverableResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
