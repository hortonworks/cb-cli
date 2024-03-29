// Code generated by go-swagger; DO NOT EDIT.

package v1diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// ListDiagnosticsCollectionsV1Reader is a Reader for the ListDiagnosticsCollectionsV1 structure.
type ListDiagnosticsCollectionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDiagnosticsCollectionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDiagnosticsCollectionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDiagnosticsCollectionsV1OK creates a ListDiagnosticsCollectionsV1OK with default headers values
func NewListDiagnosticsCollectionsV1OK() *ListDiagnosticsCollectionsV1OK {
	return &ListDiagnosticsCollectionsV1OK{}
}

/*
ListDiagnosticsCollectionsV1OK handles this case with default header values.

successful operation
*/
type ListDiagnosticsCollectionsV1OK struct {
	Payload *model.ListDiagnosticsCollectionResponse
}

func (o *ListDiagnosticsCollectionsV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/diagnostics/{environmentCrn}/collections][%d] listDiagnosticsCollectionsV1OK  %+v", 200, o.Payload)
}

func (o *ListDiagnosticsCollectionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ListDiagnosticsCollectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
