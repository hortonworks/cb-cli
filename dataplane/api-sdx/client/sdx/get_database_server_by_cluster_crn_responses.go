// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// GetDatabaseServerByClusterCrnReader is a Reader for the GetDatabaseServerByClusterCrn structure.
type GetDatabaseServerByClusterCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabaseServerByClusterCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDatabaseServerByClusterCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDatabaseServerByClusterCrnOK creates a GetDatabaseServerByClusterCrnOK with default headers values
func NewGetDatabaseServerByClusterCrnOK() *GetDatabaseServerByClusterCrnOK {
	return &GetDatabaseServerByClusterCrnOK{}
}

/*
GetDatabaseServerByClusterCrnOK handles this case with default header values.

successful operation
*/
type GetDatabaseServerByClusterCrnOK struct {
	Payload *model.StackDatabaseServerResponse
}

func (o *GetDatabaseServerByClusterCrnOK) Error() string {
	return fmt.Sprintf("[GET /sdx/crn/{clusterCrn}/dbserver][%d] getDatabaseServerByClusterCrnOK  %+v", 200, o.Payload)
}

func (o *GetDatabaseServerByClusterCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.StackDatabaseServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
