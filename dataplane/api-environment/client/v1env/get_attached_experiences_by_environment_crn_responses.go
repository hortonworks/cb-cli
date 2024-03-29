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

// GetAttachedExperiencesByEnvironmentCrnReader is a Reader for the GetAttachedExperiencesByEnvironmentCrn structure.
type GetAttachedExperiencesByEnvironmentCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAttachedExperiencesByEnvironmentCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAttachedExperiencesByEnvironmentCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAttachedExperiencesByEnvironmentCrnOK creates a GetAttachedExperiencesByEnvironmentCrnOK with default headers values
func NewGetAttachedExperiencesByEnvironmentCrnOK() *GetAttachedExperiencesByEnvironmentCrnOK {
	return &GetAttachedExperiencesByEnvironmentCrnOK{}
}

/*
GetAttachedExperiencesByEnvironmentCrnOK handles this case with default header values.

successful operation
*/
type GetAttachedExperiencesByEnvironmentCrnOK struct {
	Payload map[string][]string
}

func (o *GetAttachedExperiencesByEnvironmentCrnOK) Error() string {
	return fmt.Sprintf("[GET /v1/env/xp/crn/{crn}][%d] getAttachedExperiencesByEnvironmentCrnOK  %+v", 200, o.Payload)
}

func (o *GetAttachedExperiencesByEnvironmentCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
