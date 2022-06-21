// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PutStackForAutoscaleStartByNameReader is a Reader for the PutStackForAutoscaleStartByName structure.
type PutStackForAutoscaleStartByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutStackForAutoscaleStartByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutStackForAutoscaleStartByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutStackForAutoscaleStartByNameOK creates a PutStackForAutoscaleStartByNameOK with default headers values
func NewPutStackForAutoscaleStartByNameOK() *PutStackForAutoscaleStartByNameOK {
	return &PutStackForAutoscaleStartByNameOK{}
}

/*PutStackForAutoscaleStartByNameOK handles this case with default header values.

successful operation
*/
type PutStackForAutoscaleStartByNameOK struct {
	Payload *model.FlowIdentifier
}

func (o *PutStackForAutoscaleStartByNameOK) Error() string {
	return fmt.Sprintf("[PUT /autoscale/stack/startNodes/name/{name}][%d] putStackForAutoscaleStartByNameOK  %+v", 200, o.Payload)
}

func (o *PutStackForAutoscaleStartByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
