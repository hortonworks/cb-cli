// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPlatformSecurityGroupsParams creates a new GetPlatformSecurityGroupsParams object
// with the default values initialized.
func NewGetPlatformSecurityGroupsParams() *GetPlatformSecurityGroupsParams {
	var ()
	return &GetPlatformSecurityGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformSecurityGroupsParamsWithTimeout creates a new GetPlatformSecurityGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformSecurityGroupsParamsWithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsParams {
	var ()
	return &GetPlatformSecurityGroupsParams{

		timeout: timeout,
	}
}

// NewGetPlatformSecurityGroupsParamsWithContext creates a new GetPlatformSecurityGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformSecurityGroupsParamsWithContext(ctx context.Context) *GetPlatformSecurityGroupsParams {
	var ()
	return &GetPlatformSecurityGroupsParams{

		Context: ctx,
	}
}

// NewGetPlatformSecurityGroupsParamsWithHTTPClient creates a new GetPlatformSecurityGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformSecurityGroupsParamsWithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsParams {
	var ()
	return &GetPlatformSecurityGroupsParams{
		HTTPClient: client,
	}
}

/*
GetPlatformSecurityGroupsParams contains all the parameters to send to the API endpoint
for the get platform security groups operation typically these are written to a http.Request
*/
type GetPlatformSecurityGroupsParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*CredentialCrn*/
	CredentialCrn *string
	/*CredentialName*/
	CredentialName *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string
	/*SharedProjectID*/
	SharedProjectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithTimeout(timeout time.Duration) *GetPlatformSecurityGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithContext(ctx context.Context) *GetPlatformSecurityGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithHTTPClient(client *http.Client) *GetPlatformSecurityGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithAvailabilityZone(availabilityZone *string) *GetPlatformSecurityGroupsParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithCredentialCrn(credentialCrn *string) *GetPlatformSecurityGroupsParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithCredentialName(credentialName *string) *GetPlatformSecurityGroupsParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithPlatformVariant(platformVariant *string) *GetPlatformSecurityGroupsParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithRegion(region *string) *GetPlatformSecurityGroupsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetRegion(region *string) {
	o.Region = region
}

// WithSharedProjectID adds the sharedProjectID to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) WithSharedProjectID(sharedProjectID *string) *GetPlatformSecurityGroupsParams {
	o.SetSharedProjectID(sharedProjectID)
	return o
}

// SetSharedProjectID adds the sharedProjectId to the get platform security groups params
func (o *GetPlatformSecurityGroupsParams) SetSharedProjectID(sharedProjectID *string) {
	o.SharedProjectID = sharedProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformSecurityGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CredentialCrn != nil {

		// query param credentialCrn
		var qrCredentialCrn string
		if o.CredentialCrn != nil {
			qrCredentialCrn = *o.CredentialCrn
		}
		qCredentialCrn := qrCredentialCrn
		if qCredentialCrn != "" {
			if err := r.SetQueryParam("credentialCrn", qCredentialCrn); err != nil {
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

	if o.SharedProjectID != nil {

		// query param sharedProjectId
		var qrSharedProjectID string
		if o.SharedProjectID != nil {
			qrSharedProjectID = *o.SharedProjectID
		}
		qSharedProjectID := qrSharedProjectID
		if qSharedProjectID != "" {
			if err := r.SetQueryParam("sharedProjectId", qSharedProjectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
