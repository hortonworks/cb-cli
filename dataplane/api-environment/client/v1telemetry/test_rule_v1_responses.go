// Code generated by go-swagger; DO NOT EDIT.

package v1telemetry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// TestRuleV1Reader is a Reader for the TestRuleV1 structure.
type TestRuleV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestRuleV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestRuleV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestRuleV1OK creates a TestRuleV1OK with default headers values
func NewTestRuleV1OK() *TestRuleV1OK {
	return &TestRuleV1OK{}
}

/*TestRuleV1OK handles this case with default header values.

successful operation
*/
type TestRuleV1OK struct {
	Payload *model.TestAnonymizationRuleResponse
}

func (o *TestRuleV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/telemetry/rules/test][%d] testRuleV1OK  %+v", 200, o.Payload)
}

func (o *TestRuleV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.TestAnonymizationRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
