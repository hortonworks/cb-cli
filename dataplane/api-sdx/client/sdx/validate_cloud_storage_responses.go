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

// ValidateCloudStorageReader is a Reader for the ValidateCloudStorage structure.
type ValidateCloudStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCloudStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateCloudStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCloudStorageOK creates a ValidateCloudStorageOK with default headers values
func NewValidateCloudStorageOK() *ValidateCloudStorageOK {
	return &ValidateCloudStorageOK{}
}

/*
ValidateCloudStorageOK handles this case with default header values.

successful operation
*/
type ValidateCloudStorageOK struct {
	Payload *model.ObjectStorageValidateResponse
}

func (o *ValidateCloudStorageOK) Error() string {
	return fmt.Sprintf("[POST /sdx/validate_cloud_storage/{name}][%d] validateCloudStorageOK  %+v", 200, o.Payload)
}

func (o *ValidateCloudStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ObjectStorageValidateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
