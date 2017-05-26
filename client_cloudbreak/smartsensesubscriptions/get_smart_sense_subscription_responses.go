package smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetSmartSenseSubscriptionReader is a Reader for the GetSmartSenseSubscription structure.
type GetSmartSenseSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSmartSenseSubscriptionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSmartSenseSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSmartSenseSubscriptionOK creates a GetSmartSenseSubscriptionOK with default headers values
func NewGetSmartSenseSubscriptionOK() *GetSmartSenseSubscriptionOK {
	return &GetSmartSenseSubscriptionOK{}
}

/*GetSmartSenseSubscriptionOK handles this case with default header values.

successful operation
*/
type GetSmartSenseSubscriptionOK struct {
	Payload *models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *GetSmartSenseSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /smartsensesubscriptions][%d] getSmartSenseSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSmartSenseSubscriptionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
