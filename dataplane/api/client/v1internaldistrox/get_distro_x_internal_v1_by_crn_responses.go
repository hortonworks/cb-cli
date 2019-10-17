// Code generated by go-swagger; DO NOT EDIT.

package v1internaldistrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDistroXInternalV1ByCrnReader is a Reader for the GetDistroXInternalV1ByCrn structure.
type GetDistroXInternalV1ByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroXInternalV1ByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroXInternalV1ByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroXInternalV1ByCrnOK creates a GetDistroXInternalV1ByCrnOK with default headers values
func NewGetDistroXInternalV1ByCrnOK() *GetDistroXInternalV1ByCrnOK {
	return &GetDistroXInternalV1ByCrnOK{}
}

/*GetDistroXInternalV1ByCrnOK handles this case with default header values.

successful operation
*/
type GetDistroXInternalV1ByCrnOK struct {
	Payload *model.StackViewV4Response
}

func (o *GetDistroXInternalV1ByCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/internal/distrox/crn/{crn}][%d] getDistroXInternalV1ByCrnOK  %+v", 200, o.Payload)
}

func (o *GetDistroXInternalV1ByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackViewV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
