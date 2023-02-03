// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// UpgradeDistroxClusterByCrnReader is a Reader for the UpgradeDistroxClusterByCrn structure.
type UpgradeDistroxClusterByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeDistroxClusterByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpgradeDistroxClusterByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpgradeDistroxClusterByCrnOK creates a UpgradeDistroxClusterByCrnOK with default headers values
func NewUpgradeDistroxClusterByCrnOK() *UpgradeDistroxClusterByCrnOK {
	return &UpgradeDistroxClusterByCrnOK{}
}

/*
UpgradeDistroxClusterByCrnOK handles this case with default header values.

successful operation
*/
type UpgradeDistroxClusterByCrnOK struct {
	Payload *model.DistroXUpgradeV1Response
}

func (o *UpgradeDistroxClusterByCrnOK) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/crn/{crn}/upgrade][%d] upgradeDistroxClusterByCrnOK  %+v", 200, o.Payload)
}

func (o *UpgradeDistroxClusterByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DistroXUpgradeV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
