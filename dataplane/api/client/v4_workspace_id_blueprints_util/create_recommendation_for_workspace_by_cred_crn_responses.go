// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints_util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// CreateRecommendationForWorkspaceByCredCrnReader is a Reader for the CreateRecommendationForWorkspaceByCredCrn structure.
type CreateRecommendationForWorkspaceByCredCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRecommendationForWorkspaceByCredCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRecommendationForWorkspaceByCredCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRecommendationForWorkspaceByCredCrnOK creates a CreateRecommendationForWorkspaceByCredCrnOK with default headers values
func NewCreateRecommendationForWorkspaceByCredCrnOK() *CreateRecommendationForWorkspaceByCredCrnOK {
	return &CreateRecommendationForWorkspaceByCredCrnOK{}
}

/*
CreateRecommendationForWorkspaceByCredCrnOK handles this case with default header values.

successful operation
*/
type CreateRecommendationForWorkspaceByCredCrnOK struct {
	Payload *model.RecommendationV4Response
}

func (o *CreateRecommendationForWorkspaceByCredCrnOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/blueprints_util/recommendation_by_cred_crn][%d] createRecommendationForWorkspaceByCredCrnOK  %+v", 200, o.Payload)
}

func (o *CreateRecommendationForWorkspaceByCredCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RecommendationV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
