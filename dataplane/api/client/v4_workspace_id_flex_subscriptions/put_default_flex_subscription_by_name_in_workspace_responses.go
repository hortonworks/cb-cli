// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_flex_subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// PutDefaultFlexSubscriptionByNameInWorkspaceReader is a Reader for the PutDefaultFlexSubscriptionByNameInWorkspace structure.
type PutDefaultFlexSubscriptionByNameInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDefaultFlexSubscriptionByNameInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDefaultFlexSubscriptionByNameInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDefaultFlexSubscriptionByNameInWorkspaceOK creates a PutDefaultFlexSubscriptionByNameInWorkspaceOK with default headers values
func NewPutDefaultFlexSubscriptionByNameInWorkspaceOK() *PutDefaultFlexSubscriptionByNameInWorkspaceOK {
	return &PutDefaultFlexSubscriptionByNameInWorkspaceOK{}
}

/*PutDefaultFlexSubscriptionByNameInWorkspaceOK handles this case with default header values.

successful operation
*/
type PutDefaultFlexSubscriptionByNameInWorkspaceOK struct {
	Payload *model.FlexSubscriptionV4Response
}

func (o *PutDefaultFlexSubscriptionByNameInWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/flex_subscriptions/{name}/default][%d] putDefaultFlexSubscriptionByNameInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutDefaultFlexSubscriptionByNameInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlexSubscriptionV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
