// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// ListSmartSenseSubscriptionsByWorkspaceReader is a Reader for the ListSmartSenseSubscriptionsByWorkspace structure.
type ListSmartSenseSubscriptionsByWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSmartSenseSubscriptionsByWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSmartSenseSubscriptionsByWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSmartSenseSubscriptionsByWorkspaceOK creates a ListSmartSenseSubscriptionsByWorkspaceOK with default headers values
func NewListSmartSenseSubscriptionsByWorkspaceOK() *ListSmartSenseSubscriptionsByWorkspaceOK {
	return &ListSmartSenseSubscriptionsByWorkspaceOK{}
}

/*ListSmartSenseSubscriptionsByWorkspaceOK handles this case with default header values.

successful operation
*/
type ListSmartSenseSubscriptionsByWorkspaceOK struct {
	Payload []*model.SmartSenseSubscriptionJSON
}

func (o *ListSmartSenseSubscriptionsByWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v3/{workspaceId}/smartsensesubscriptions][%d] listSmartSenseSubscriptionsByWorkspaceOK  %+v", 200, o.Payload)
}

func (o *ListSmartSenseSubscriptionsByWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
