// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRegionsByCredentialParams creates a new GetRegionsByCredentialParams object
// with the default values initialized.
func NewGetRegionsByCredentialParams() *GetRegionsByCredentialParams {
	var ()
	return &GetRegionsByCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsByCredentialParamsWithTimeout creates a new GetRegionsByCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsByCredentialParamsWithTimeout(timeout time.Duration) *GetRegionsByCredentialParams {
	var ()
	return &GetRegionsByCredentialParams{

		timeout: timeout,
	}
}

// NewGetRegionsByCredentialParamsWithContext creates a new GetRegionsByCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsByCredentialParamsWithContext(ctx context.Context) *GetRegionsByCredentialParams {
	var ()
	return &GetRegionsByCredentialParams{

		Context: ctx,
	}
}

// NewGetRegionsByCredentialParamsWithHTTPClient creates a new GetRegionsByCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsByCredentialParamsWithHTTPClient(client *http.Client) *GetRegionsByCredentialParams {
	var ()
	return &GetRegionsByCredentialParams{
		HTTPClient: client,
	}
}

/*GetRegionsByCredentialParams contains all the parameters to send to the API endpoint
for the get regions by credential operation typically these are written to a http.Request
*/
type GetRegionsByCredentialParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithTimeout(timeout time.Duration) *GetRegionsByCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithContext(ctx context.Context) *GetRegionsByCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithHTTPClient(client *http.Client) *GetRegionsByCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithAvailabilityZone(availabilityZone *string) *GetRegionsByCredentialParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithCredentialCrn(credentialCrn *string) *GetRegionsByCredentialParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithCredentialName(credentialName *string) *GetRegionsByCredentialParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithPlatformVariant(platformVariant *string) *GetRegionsByCredentialParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get regions by credential params
func (o *GetRegionsByCredentialParams) WithRegion(region *string) *GetRegionsByCredentialParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get regions by credential params
func (o *GetRegionsByCredentialParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsByCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
