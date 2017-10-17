// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetAllStackV2Reader is a Reader for the GetAllStackV2 structure.
type GetAllStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllStackV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllStackV2OK creates a GetAllStackV2OK with default headers values
func NewGetAllStackV2OK() *GetAllStackV2OK {
	return &GetAllStackV2OK{}
}

/*GetAllStackV2OK handles this case with default header values.

successful operation
*/
type GetAllStackV2OK struct {
	Payload []*models_cloudbreak.AutoscaleStackResponse
}

func (o *GetAllStackV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/stacks/all][%d] getAllStackV2OK  %+v", 200, o.Payload)
}

func (o *GetAllStackV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
