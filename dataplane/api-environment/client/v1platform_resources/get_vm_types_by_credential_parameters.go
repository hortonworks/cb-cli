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

// NewGetVMTypesByCredentialParams creates a new GetVMTypesByCredentialParams object
// with the default values initialized.
func NewGetVMTypesByCredentialParams() *GetVMTypesByCredentialParams {
	var ()
	return &GetVMTypesByCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMTypesByCredentialParamsWithTimeout creates a new GetVMTypesByCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVMTypesByCredentialParamsWithTimeout(timeout time.Duration) *GetVMTypesByCredentialParams {
	var ()
	return &GetVMTypesByCredentialParams{

		timeout: timeout,
	}
}

// NewGetVMTypesByCredentialParamsWithContext creates a new GetVMTypesByCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVMTypesByCredentialParamsWithContext(ctx context.Context) *GetVMTypesByCredentialParams {
	var ()
	return &GetVMTypesByCredentialParams{

		Context: ctx,
	}
}

// NewGetVMTypesByCredentialParamsWithHTTPClient creates a new GetVMTypesByCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVMTypesByCredentialParamsWithHTTPClient(client *http.Client) *GetVMTypesByCredentialParams {
	var ()
	return &GetVMTypesByCredentialParams{
		HTTPClient: client,
	}
}

/*
GetVMTypesByCredentialParams contains all the parameters to send to the API endpoint
for the get Vm types by credential operation typically these are written to a http.Request
*/
type GetVMTypesByCredentialParams struct {

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
	/*ResourceType*/
	ResourceType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithTimeout(timeout time.Duration) *GetVMTypesByCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithContext(ctx context.Context) *GetVMTypesByCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithHTTPClient(client *http.Client) *GetVMTypesByCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithAvailabilityZone(availabilityZone *string) *GetVMTypesByCredentialParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithCredentialCrn(credentialCrn *string) *GetVMTypesByCredentialParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithCredentialName(credentialName *string) *GetVMTypesByCredentialParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithPlatformVariant(platformVariant *string) *GetVMTypesByCredentialParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithRegion(region *string) *GetVMTypesByCredentialParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetRegion(region *string) {
	o.Region = region
}

// WithResourceType adds the resourceType to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) WithResourceType(resourceType *string) *GetVMTypesByCredentialParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get Vm types by credential params
func (o *GetVMTypesByCredentialParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMTypesByCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
