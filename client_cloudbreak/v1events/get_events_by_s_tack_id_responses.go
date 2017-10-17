// Code generated by go-swagger; DO NOT EDIT.

package v1events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetEventsBySTackIDReader is a Reader for the GetEventsBySTackID structure.
type GetEventsBySTackIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsBySTackIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsBySTackIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsBySTackIDOK creates a GetEventsBySTackIDOK with default headers values
func NewGetEventsBySTackIDOK() *GetEventsBySTackIDOK {
	return &GetEventsBySTackIDOK{}
}

/*GetEventsBySTackIDOK handles this case with default header values.

successful operation
*/
type GetEventsBySTackIDOK struct {
	Payload []*models_cloudbreak.CloudbreakEvent
}

func (o *GetEventsBySTackIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/{stackId}][%d] getEventsBySTackIdOK  %+v", 200, o.Payload)
}

func (o *GetEventsBySTackIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
