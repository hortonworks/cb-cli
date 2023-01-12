// Code generated by go-swagger; DO NOT EDIT.

package internalsdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// ModifyInternalSdxProxyConfigReader is a Reader for the ModifyInternalSdxProxyConfig structure.
type ModifyInternalSdxProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyInternalSdxProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyInternalSdxProxyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyInternalSdxProxyConfigOK creates a ModifyInternalSdxProxyConfigOK with default headers values
func NewModifyInternalSdxProxyConfigOK() *ModifyInternalSdxProxyConfigOK {
	return &ModifyInternalSdxProxyConfigOK{}
}

/*ModifyInternalSdxProxyConfigOK handles this case with default header values.

successful operation
*/
type ModifyInternalSdxProxyConfigOK struct {
	Payload *model.FlowIdentifier
}

func (o *ModifyInternalSdxProxyConfigOK) Error() string {
	return fmt.Sprintf("[PUT /internal/sdx/crn/{crn}/modify_proxy][%d] modifyInternalSdxProxyConfigOK  %+v", 200, o.Payload)
}

func (o *ModifyInternalSdxProxyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
