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

// GetSdxByEnvCrnReader is a Reader for the GetSdxByEnvCrn structure.
type GetSdxByEnvCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSdxByEnvCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSdxByEnvCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSdxByEnvCrnOK creates a GetSdxByEnvCrnOK with default headers values
func NewGetSdxByEnvCrnOK() *GetSdxByEnvCrnOK {
	return &GetSdxByEnvCrnOK{}
}

/*
GetSdxByEnvCrnOK handles this case with default header values.

successful operation
*/
type GetSdxByEnvCrnOK struct {
	Payload []*model.SdxClusterResponse
}

func (o *GetSdxByEnvCrnOK) Error() string {
	return fmt.Sprintf("[GET /sdx/envcrn/{envCrn}][%d] getSdxByEnvCrnOK  %+v", 200, o.Payload)
}

func (o *GetSdxByEnvCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
