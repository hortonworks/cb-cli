// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints_util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateRecommendationForWorkspaceParams creates a new CreateRecommendationForWorkspaceParams object
// with the default values initialized.
func NewCreateRecommendationForWorkspaceParams() *CreateRecommendationForWorkspaceParams {
	var ()
	return &CreateRecommendationForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRecommendationForWorkspaceParamsWithTimeout creates a new CreateRecommendationForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRecommendationForWorkspaceParamsWithTimeout(timeout time.Duration) *CreateRecommendationForWorkspaceParams {
	var ()
	return &CreateRecommendationForWorkspaceParams{

		timeout: timeout,
	}
}

// NewCreateRecommendationForWorkspaceParamsWithContext creates a new CreateRecommendationForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRecommendationForWorkspaceParamsWithContext(ctx context.Context) *CreateRecommendationForWorkspaceParams {
	var ()
	return &CreateRecommendationForWorkspaceParams{

		Context: ctx,
	}
}

// NewCreateRecommendationForWorkspaceParamsWithHTTPClient creates a new CreateRecommendationForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRecommendationForWorkspaceParamsWithHTTPClient(client *http.Client) *CreateRecommendationForWorkspaceParams {
	var ()
	return &CreateRecommendationForWorkspaceParams{
		HTTPClient: client,
	}
}

/*CreateRecommendationForWorkspaceParams contains all the parameters to send to the API endpoint
for the create recommendation for workspace operation typically these are written to a http.Request
*/
type CreateRecommendationForWorkspaceParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*BlueprintName*/
	BlueprintName *string
	/*CredentialName*/
	CredentialName *string
	/*DefinitionName*/
	DefinitionName *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string
	/*ResourceType*/
	ResourceType *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithTimeout(timeout time.Duration) *CreateRecommendationForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithContext(ctx context.Context) *CreateRecommendationForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithHTTPClient(client *http.Client) *CreateRecommendationForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithAvailabilityZone(availabilityZone *string) *CreateRecommendationForWorkspaceParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithBlueprintName adds the blueprintName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithBlueprintName(blueprintName *string) *CreateRecommendationForWorkspaceParams {
	o.SetBlueprintName(blueprintName)
	return o
}

// SetBlueprintName adds the blueprintName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetBlueprintName(blueprintName *string) {
	o.BlueprintName = blueprintName
}

// WithCredentialName adds the credentialName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithCredentialName(credentialName *string) *CreateRecommendationForWorkspaceParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithDefinitionName adds the definitionName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithDefinitionName(definitionName *string) *CreateRecommendationForWorkspaceParams {
	o.SetDefinitionName(definitionName)
	return o
}

// SetDefinitionName adds the definitionName to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetDefinitionName(definitionName *string) {
	o.DefinitionName = definitionName
}

// WithPlatformVariant adds the platformVariant to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithPlatformVariant(platformVariant *string) *CreateRecommendationForWorkspaceParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithRegion(region *string) *CreateRecommendationForWorkspaceParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetRegion(region *string) {
	o.Region = region
}

// WithResourceType adds the resourceType to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithResourceType(resourceType *string) *CreateRecommendationForWorkspaceParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithWorkspaceID adds the workspaceID to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) WithWorkspaceID(workspaceID int64) *CreateRecommendationForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the create recommendation for workspace params
func (o *CreateRecommendationForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRecommendationForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityZone != nil {

		// query param availabilityZone
		var qrAvailabilityZone string
		if o.AvailabilityZone != nil {
			qrAvailabilityZone = *o.AvailabilityZone
		}
		qAvailabilityZone := qrAvailabilityZone
		if qAvailabilityZone != "" {
			if err := r.SetQueryParam("availabilityZone", qAvailabilityZone); err != nil {
				return err
			}
		}

	}

	if o.BlueprintName != nil {

		// query param blueprintName
		var qrBlueprintName string
		if o.BlueprintName != nil {
			qrBlueprintName = *o.BlueprintName
		}
		qBlueprintName := qrBlueprintName
		if qBlueprintName != "" {
			if err := r.SetQueryParam("blueprintName", qBlueprintName); err != nil {
				return err
			}
		}

	}

	if o.CredentialName != nil {

		// query param credentialName
		var qrCredentialName string
		if o.CredentialName != nil {
			qrCredentialName = *o.CredentialName
		}
		qCredentialName := qrCredentialName
		if qCredentialName != "" {
			if err := r.SetQueryParam("credentialName", qCredentialName); err != nil {
				return err
			}
		}

	}

	if o.DefinitionName != nil {

		// query param definitionName
		var qrDefinitionName string
		if o.DefinitionName != nil {
			qrDefinitionName = *o.DefinitionName
		}
		qDefinitionName := qrDefinitionName
		if qDefinitionName != "" {
			if err := r.SetQueryParam("definitionName", qDefinitionName); err != nil {
				return err
			}
		}

	}

	if o.PlatformVariant != nil {

		// query param platformVariant
		var qrPlatformVariant string
		if o.PlatformVariant != nil {
			qrPlatformVariant = *o.PlatformVariant
		}
		qPlatformVariant := qrPlatformVariant
		if qPlatformVariant != "" {
			if err := r.SetQueryParam("platformVariant", qPlatformVariant); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.ResourceType != nil {

		// query param resourceType
		var qrResourceType string
		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {
			if err := r.SetQueryParam("resourceType", qResourceType); err != nil {
				return err
			}
		}

	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
