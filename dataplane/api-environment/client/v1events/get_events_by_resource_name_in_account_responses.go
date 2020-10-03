// Code generated by go-swagger; DO NOT EDIT.

package v1events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetEventsByResourceNameInAccountReader is a Reader for the GetEventsByResourceNameInAccount structure.
type GetEventsByResourceNameInAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsByResourceNameInAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsByResourceNameInAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsByResourceNameInAccountOK creates a GetEventsByResourceNameInAccountOK with default headers values
func NewGetEventsByResourceNameInAccountOK() *GetEventsByResourceNameInAccountOK {
	return &GetEventsByResourceNameInAccountOK{}
}

/*GetEventsByResourceNameInAccountOK handles this case with default header values.

successful operation
*/
type GetEventsByResourceNameInAccountOK struct {
	Payload *model.PageCDPEventV1Response
}

func (o *GetEventsByResourceNameInAccountOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/{name}][%d] getEventsByResourceNameInAccountOK  %+v", 200, o.Payload)
}

func (o *GetEventsByResourceNameInAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PageCDPEventV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
