// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PostPrivateStackV2Reader is a Reader for the PostPrivateStackV2 structure.
type PostPrivateStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateStackV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateStackV2OK creates a PostPrivateStackV2OK with default headers values
func NewPostPrivateStackV2OK() *PostPrivateStackV2OK {
	return &PostPrivateStackV2OK{}
}

/*PostPrivateStackV2OK handles this case with default header values.

successful operation
*/
type PostPrivateStackV2OK struct {
	Payload *model.StackResponse
}

func (o *PostPrivateStackV2OK) Error() string {
	return fmt.Sprintf("[POST /v2/stacks/user][%d] postPrivateStackV2OK  %+v", 200, o.Payload)
}

func (o *PostPrivateStackV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
