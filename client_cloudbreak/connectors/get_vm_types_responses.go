// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetVMTypesReader is a Reader for the GetVMTypes structure.
type GetVMTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVMTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVMTypesOK creates a GetVMTypesOK with default headers values
func NewGetVMTypesOK() *GetVMTypesOK {
	return &GetVMTypesOK{}
}

/*GetVMTypesOK handles this case with default header values.

successful operation
*/
type GetVMTypesOK struct {
	Payload *models_cloudbreak.PlatformVirtualMachinesJSON
}

func (o *GetVMTypesOK) Error() string {
	return fmt.Sprintf("[GET /connectors/vmtypes][%d] getVmTypesOK  %+v", 200, o.Payload)
}

func (o *GetVMTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.PlatformVirtualMachinesJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
