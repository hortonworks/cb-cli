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

// UpgradeDistroxClusterReader is a Reader for the UpgradeDistroxCluster structure.
type UpgradeDistroxClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeDistroxClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpgradeDistroxClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpgradeDistroxClusterOK creates a UpgradeDistroxClusterOK with default headers values
func NewUpgradeDistroxClusterOK() *UpgradeDistroxClusterOK {
	return &UpgradeDistroxClusterOK{}
}

/*
UpgradeDistroxClusterOK handles this case with default header values.

successful operation
*/
type UpgradeDistroxClusterOK struct {
	Payload *model.DistroXUpgradeV1Response
}

func (o *UpgradeDistroxClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/{name}/upgrade][%d] upgradeDistroxClusterOK  %+v", 200, o.Payload)
}

func (o *UpgradeDistroxClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.DistroXUpgradeV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
