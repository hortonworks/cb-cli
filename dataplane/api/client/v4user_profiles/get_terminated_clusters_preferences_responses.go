// Code generated by go-swagger; DO NOT EDIT.

package v4user_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetTerminatedClustersPreferencesReader is a Reader for the GetTerminatedClustersPreferences structure.
type GetTerminatedClustersPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerminatedClustersPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTerminatedClustersPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTerminatedClustersPreferencesOK creates a GetTerminatedClustersPreferencesOK with default headers values
func NewGetTerminatedClustersPreferencesOK() *GetTerminatedClustersPreferencesOK {
	return &GetTerminatedClustersPreferencesOK{}
}

/*GetTerminatedClustersPreferencesOK handles this case with default header values.

successful operation
*/
type GetTerminatedClustersPreferencesOK struct {
	Payload *model.ShowTerminatedClusterPreferencesV4Response
}

func (o *GetTerminatedClustersPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /v4/user_profiles/terminated_clusters_preferences][%d] getTerminatedClustersPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetTerminatedClustersPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ShowTerminatedClusterPreferencesV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
