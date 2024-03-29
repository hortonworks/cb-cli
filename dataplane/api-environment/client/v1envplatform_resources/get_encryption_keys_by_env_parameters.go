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

// NewGetEncryptionKeysByEnvParams creates a new GetEncryptionKeysByEnvParams object
// with the default values initialized.
func NewGetEncryptionKeysByEnvParams() *GetEncryptionKeysByEnvParams {
	var ()
	return &GetEncryptionKeysByEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEncryptionKeysByEnvParamsWithTimeout creates a new GetEncryptionKeysByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEncryptionKeysByEnvParamsWithTimeout(timeout time.Duration) *GetEncryptionKeysByEnvParams {
	var ()
	return &GetEncryptionKeysByEnvParams{

		timeout: timeout,
	}
}

// NewGetEncryptionKeysByEnvParamsWithContext creates a new GetEncryptionKeysByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEncryptionKeysByEnvParamsWithContext(ctx context.Context) *GetEncryptionKeysByEnvParams {
	var ()
	return &GetEncryptionKeysByEnvParams{

		Context: ctx,
	}
}

// NewGetEncryptionKeysByEnvParamsWithHTTPClient creates a new GetEncryptionKeysByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEncryptionKeysByEnvParamsWithHTTPClient(client *http.Client) *GetEncryptionKeysByEnvParams {
	var ()
	return &GetEncryptionKeysByEnvParams{
		HTTPClient: client,
	}
}

/*
GetEncryptionKeysByEnvParams contains all the parameters to send to the API endpoint
for the get encryption keys by env operation typically these are written to a http.Request
*/
type GetEncryptionKeysByEnvParams struct {

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

// WithTimeout adds the timeout to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithTimeout(timeout time.Duration) *GetEncryptionKeysByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithContext(ctx context.Context) *GetEncryptionKeysByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithHTTPClient(client *http.Client) *GetEncryptionKeysByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithAvailabilityZone(availabilityZone *string) *GetEncryptionKeysByEnvParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithEnvironmentCrn adds the environmentCrn to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetEncryptionKeysByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithPlatformVariant(platformVariant *string) *GetEncryptionKeysByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) WithRegion(region *string) *GetEncryptionKeysByEnvParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get encryption keys by env params
func (o *GetEncryptionKeysByEnvParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetEncryptionKeysByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
