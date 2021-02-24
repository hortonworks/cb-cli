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

// NewGetAccessConfigsByEnvParams creates a new GetAccessConfigsByEnvParams object
// with the default values initialized.
func NewGetAccessConfigsByEnvParams() *GetAccessConfigsByEnvParams {
	var (
		accessConfigTypeDefault = string("INSTANCE_PROFILE")
	)
	return &GetAccessConfigsByEnvParams{
		AccessConfigType: &accessConfigTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessConfigsByEnvParamsWithTimeout creates a new GetAccessConfigsByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessConfigsByEnvParamsWithTimeout(timeout time.Duration) *GetAccessConfigsByEnvParams {
	var (
		accessConfigTypeDefault = string("INSTANCE_PROFILE")
	)
	return &GetAccessConfigsByEnvParams{
		AccessConfigType: &accessConfigTypeDefault,

		timeout: timeout,
	}
}

// NewGetAccessConfigsByEnvParamsWithContext creates a new GetAccessConfigsByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessConfigsByEnvParamsWithContext(ctx context.Context) *GetAccessConfigsByEnvParams {
	var (
		accessConfigTypeDefault = string("INSTANCE_PROFILE")
	)
	return &GetAccessConfigsByEnvParams{
		AccessConfigType: &accessConfigTypeDefault,

		Context: ctx,
	}
}

// NewGetAccessConfigsByEnvParamsWithHTTPClient creates a new GetAccessConfigsByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessConfigsByEnvParamsWithHTTPClient(client *http.Client) *GetAccessConfigsByEnvParams {
	var (
		accessConfigTypeDefault = string("INSTANCE_PROFILE")
	)
	return &GetAccessConfigsByEnvParams{
		AccessConfigType: &accessConfigTypeDefault,
		HTTPClient:       client,
	}
}

/*GetAccessConfigsByEnvParams contains all the parameters to send to the API endpoint
for the get access configs by env operation typically these are written to a http.Request
*/
type GetAccessConfigsByEnvParams struct {

	/*AccessConfigType*/
	AccessConfigType *string
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

// WithTimeout adds the timeout to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithTimeout(timeout time.Duration) *GetAccessConfigsByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithContext(ctx context.Context) *GetAccessConfigsByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithHTTPClient(client *http.Client) *GetAccessConfigsByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessConfigType adds the accessConfigType to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithAccessConfigType(accessConfigType *string) *GetAccessConfigsByEnvParams {
	o.SetAccessConfigType(accessConfigType)
	return o
}

// SetAccessConfigType adds the accessConfigType to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetAccessConfigType(accessConfigType *string) {
	o.AccessConfigType = accessConfigType
}

// WithAvailabilityZone adds the availabilityZone to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithAvailabilityZone(availabilityZone *string) *GetAccessConfigsByEnvParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithEnvironmentCrn adds the environmentCrn to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetAccessConfigsByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithPlatformVariant(platformVariant *string) *GetAccessConfigsByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) WithRegion(region *string) *GetAccessConfigsByEnvParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get access configs by env params
func (o *GetAccessConfigsByEnvParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessConfigsByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessConfigType != nil {

		// query param accessConfigType
		var qrAccessConfigType string
		if o.AccessConfigType != nil {
			qrAccessConfigType = *o.AccessConfigType
		}
		qAccessConfigType := qrAccessConfigType
		if qAccessConfigType != "" {
			if err := r.SetQueryParam("accessConfigType", qAccessConfigType); err != nil {
				return err
			}
		}

	}

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
