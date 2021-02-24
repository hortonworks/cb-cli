// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResourceGroupsByEnvParams creates a new GetResourceGroupsByEnvParams object
// with the default values initialized.
func NewGetResourceGroupsByEnvParams() *GetResourceGroupsByEnvParams {
	var ()
	return &GetResourceGroupsByEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceGroupsByEnvParamsWithTimeout creates a new GetResourceGroupsByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceGroupsByEnvParamsWithTimeout(timeout time.Duration) *GetResourceGroupsByEnvParams {
	var ()
	return &GetResourceGroupsByEnvParams{

		timeout: timeout,
	}
}

// NewGetResourceGroupsByEnvParamsWithContext creates a new GetResourceGroupsByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceGroupsByEnvParamsWithContext(ctx context.Context) *GetResourceGroupsByEnvParams {
	var ()
	return &GetResourceGroupsByEnvParams{

		Context: ctx,
	}
}

// NewGetResourceGroupsByEnvParamsWithHTTPClient creates a new GetResourceGroupsByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceGroupsByEnvParamsWithHTTPClient(client *http.Client) *GetResourceGroupsByEnvParams {
	var ()
	return &GetResourceGroupsByEnvParams{
		HTTPClient: client,
	}
}

/*GetResourceGroupsByEnvParams contains all the parameters to send to the API endpoint
for the get resource groups by env operation typically these are written to a http.Request
*/
type GetResourceGroupsByEnvParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithTimeout(timeout time.Duration) *GetResourceGroupsByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithContext(ctx context.Context) *GetResourceGroupsByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithHTTPClient(client *http.Client) *GetResourceGroupsByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithAvailabilityZone(availabilityZone *string) *GetResourceGroupsByEnvParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithEnvironmentCrn adds the environmentCrn to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetResourceGroupsByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithPlatformVariant(platformVariant *string) *GetResourceGroupsByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) WithRegion(region *string) *GetResourceGroupsByEnvParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get resource groups by env params
func (o *GetResourceGroupsByEnvParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceGroupsByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
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
