// Code generated by go-swagger; DO NOT EDIT.

package v1util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// TestAmbariDatabaseUtilReader is a Reader for the TestAmbariDatabaseUtil structure.
type TestAmbariDatabaseUtilReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestAmbariDatabaseUtilReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestAmbariDatabaseUtilOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestAmbariDatabaseUtilOK creates a TestAmbariDatabaseUtilOK with default headers values
func NewTestAmbariDatabaseUtilOK() *TestAmbariDatabaseUtilOK {
	return &TestAmbariDatabaseUtilOK{}
}

/*TestAmbariDatabaseUtilOK handles this case with default header values.

successful operation
*/
type TestAmbariDatabaseUtilOK struct {
	Payload *models_cloudbreak.AmbariDatabaseTestResult
}

func (o *TestAmbariDatabaseUtilOK) Error() string {
	return fmt.Sprintf("[POST /v1/util/ambari-database][%d] testAmbariDatabaseUtilOK  %+v", 200, o.Payload)
}

func (o *TestAmbariDatabaseUtilOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.AmbariDatabaseTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
