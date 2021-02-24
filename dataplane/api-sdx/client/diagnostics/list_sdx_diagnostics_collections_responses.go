// Code generated by go-swagger; DO NOT EDIT.

package diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// ListSdxDiagnosticsCollectionsReader is a Reader for the ListSdxDiagnosticsCollections structure.
type ListSdxDiagnosticsCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSdxDiagnosticsCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSdxDiagnosticsCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSdxDiagnosticsCollectionsOK creates a ListSdxDiagnosticsCollectionsOK with default headers values
func NewListSdxDiagnosticsCollectionsOK() *ListSdxDiagnosticsCollectionsOK {
	return &ListSdxDiagnosticsCollectionsOK{}
}

/*ListSdxDiagnosticsCollectionsOK handles this case with default header values.

successful operation
*/
type ListSdxDiagnosticsCollectionsOK struct {
	Payload *model.ListDiagnosticsCollectionResponse
}

func (o *ListSdxDiagnosticsCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /diagnostics/{crn}/collections][%d] listSdxDiagnosticsCollectionsOK  %+v", 200, o.Payload)
}

func (o *ListSdxDiagnosticsCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ListDiagnosticsCollectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
