// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetDistroXRequestFromNameV1Reader is a Reader for the GetDistroXRequestFromNameV1 structure.
type GetDistroXRequestFromNameV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistroXRequestFromNameV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDistroXRequestFromNameV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDistroXRequestFromNameV1OK creates a GetDistroXRequestFromNameV1OK with default headers values
func NewGetDistroXRequestFromNameV1OK() *GetDistroXRequestFromNameV1OK {
	return &GetDistroXRequestFromNameV1OK{}
}

/*GetDistroXRequestFromNameV1OK handles this case with default header values.

successful operation
*/
type GetDistroXRequestFromNameV1OK struct {
	Payload *model.DistroXV1Request
}

func (o *GetDistroXRequestFromNameV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/distrox/{name}/request][%d] getDistroXRequestFromNameV1OK  %+v", 200, o.Payload)
}

func (o *GetDistroXRequestFromNameV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DistroXV1Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
