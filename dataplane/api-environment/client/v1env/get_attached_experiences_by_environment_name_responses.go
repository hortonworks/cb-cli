// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAttachedExperiencesByEnvironmentNameReader is a Reader for the GetAttachedExperiencesByEnvironmentName structure.
type GetAttachedExperiencesByEnvironmentNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAttachedExperiencesByEnvironmentNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAttachedExperiencesByEnvironmentNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAttachedExperiencesByEnvironmentNameOK creates a GetAttachedExperiencesByEnvironmentNameOK with default headers values
func NewGetAttachedExperiencesByEnvironmentNameOK() *GetAttachedExperiencesByEnvironmentNameOK {
	return &GetAttachedExperiencesByEnvironmentNameOK{}
}

/*GetAttachedExperiencesByEnvironmentNameOK handles this case with default header values.

successful operation
*/
type GetAttachedExperiencesByEnvironmentNameOK struct {
	Payload map[string][]string
}

func (o *GetAttachedExperiencesByEnvironmentNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/xp/name/{name}][%d] getAttachedExperiencesByEnvironmentNameOK  %+v", 200, o.Payload)
}

func (o *GetAttachedExperiencesByEnvironmentNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
