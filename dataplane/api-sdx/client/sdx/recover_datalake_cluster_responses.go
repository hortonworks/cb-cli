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

// RecoverDatalakeClusterReader is a Reader for the RecoverDatalakeCluster structure.
type RecoverDatalakeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverDatalakeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecoverDatalakeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRecoverDatalakeClusterOK creates a RecoverDatalakeClusterOK with default headers values
func NewRecoverDatalakeClusterOK() *RecoverDatalakeClusterOK {
	return &RecoverDatalakeClusterOK{}
}

/*
RecoverDatalakeClusterOK handles this case with default header values.

successful operation
*/
type RecoverDatalakeClusterOK struct {
	Payload *model.SdxRecoveryResponse
}

func (o *RecoverDatalakeClusterOK) Error() string {
	return fmt.Sprintf("[POST /sdx/{name}/recover][%d] recoverDatalakeClusterOK  %+v", 200, o.Payload)
}

func (o *RecoverDatalakeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxRecoveryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
