// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

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

// NewGetPrivateDNSZonesByEnvParams creates a new GetPrivateDNSZonesByEnvParams object
// with the default values initialized.
func NewGetPrivateDNSZonesByEnvParams() *GetPrivateDNSZonesByEnvParams {
	var ()
	return &GetPrivateDNSZonesByEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateDNSZonesByEnvParamsWithTimeout creates a new GetPrivateDNSZonesByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateDNSZonesByEnvParamsWithTimeout(timeout time.Duration) *GetPrivateDNSZonesByEnvParams {
	var ()
	return &GetPrivateDNSZonesByEnvParams{

		timeout: timeout,
	}
}

// NewGetPrivateDNSZonesByEnvParamsWithContext creates a new GetPrivateDNSZonesByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateDNSZonesByEnvParamsWithContext(ctx context.Context) *GetPrivateDNSZonesByEnvParams {
	var ()
	return &GetPrivateDNSZonesByEnvParams{

		Context: ctx,
	}
}

// NewGetPrivateDNSZonesByEnvParamsWithHTTPClient creates a new GetPrivateDNSZonesByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateDNSZonesByEnvParamsWithHTTPClient(client *http.Client) *GetPrivateDNSZonesByEnvParams {
	var ()
	return &GetPrivateDNSZonesByEnvParams{
		HTTPClient: client,
	}
}

/*GetPrivateDNSZonesByEnvParams contains all the parameters to send to the API endpoint
for the get private Dns zones by env operation typically these are written to a http.Request
*/
type GetPrivateDNSZonesByEnvParams struct {

	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*PlatformVariant*/
	PlatformVariant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) WithTimeout(timeout time.Duration) *GetPrivateDNSZonesByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) WithContext(ctx context.Context) *GetPrivateDNSZonesByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) WithHTTPClient(client *http.Client) *GetPrivateDNSZonesByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetPrivateDNSZonesByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) WithPlatformVariant(platformVariant *string) *GetPrivateDNSZonesByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get private Dns zones by env params
func (o *GetPrivateDNSZonesByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateDNSZonesByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
