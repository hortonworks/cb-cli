// Code generated by go-swagger; DO NOT EDIT.

package v1dns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// AddDNSZoneForSubnetIdsV1Reader is a Reader for the AddDNSZoneForSubnetIdsV1 structure.
type AddDNSZoneForSubnetIdsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDNSZoneForSubnetIdsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddDNSZoneForSubnetIdsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDNSZoneForSubnetIdsV1OK creates a AddDNSZoneForSubnetIdsV1OK with default headers values
func NewAddDNSZoneForSubnetIdsV1OK() *AddDNSZoneForSubnetIdsV1OK {
	return &AddDNSZoneForSubnetIdsV1OK{}
}

/*AddDNSZoneForSubnetIdsV1OK handles this case with default header values.

successful operation
*/
type AddDNSZoneForSubnetIdsV1OK struct {
	Payload *model.AddDNSZoneForSubnetsV1Response
}

func (o *AddDNSZoneForSubnetIdsV1OK) Error() string {
	return fmt.Sprintf("[POST /v1/dns/zone/id][%d] addDnsZoneForSubnetIdsV1OK  %+v", 200, o.Payload)
}

func (o *AddDNSZoneForSubnetIdsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AddDNSZoneForSubnetsV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
