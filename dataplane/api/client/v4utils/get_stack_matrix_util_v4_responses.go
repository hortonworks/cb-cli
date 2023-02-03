// Code generated by go-swagger; DO NOT EDIT.

package v4utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetStackMatrixUtilV4Reader is a Reader for the GetStackMatrixUtilV4 structure.
type GetStackMatrixUtilV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackMatrixUtilV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStackMatrixUtilV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStackMatrixUtilV4OK creates a GetStackMatrixUtilV4OK with default headers values
func NewGetStackMatrixUtilV4OK() *GetStackMatrixUtilV4OK {
	return &GetStackMatrixUtilV4OK{}
}

/*
GetStackMatrixUtilV4OK handles this case with default header values.

successful operation
*/
type GetStackMatrixUtilV4OK struct {
	Payload *model.StackMatrixV4Response
}

func (o *GetStackMatrixUtilV4OK) Error() string {
	return fmt.Sprintf("[GET /v4/utils/stack_matrix][%d] getStackMatrixUtilV4OK  %+v", 200, o.Payload)
}

func (o *GetStackMatrixUtilV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackMatrixV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
