// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

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

// NewGetIPPoolsCredentialIDParams creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized.
func NewGetIPPoolsCredentialIDParams() *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPPoolsCredentialIDParamsWithTimeout creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPPoolsCredentialIDParamsWithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		timeout: timeout,
	}
}

// NewGetIPPoolsCredentialIDParamsWithContext creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPPoolsCredentialIDParamsWithContext(ctx context.Context) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{

		Context: ctx,
	}
}

// NewGetIPPoolsCredentialIDParamsWithHTTPClient creates a new GetIPPoolsCredentialIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPPoolsCredentialIDParamsWithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDParams {
	var ()
	return &GetIPPoolsCredentialIDParams{
		HTTPClient: client,
	}
}

/*GetIPPoolsCredentialIDParams contains all the parameters to send to the API endpoint
for the get Ip pools credential Id operation typically these are written to a http.Request
*/
type GetIPPoolsCredentialIDParams struct {

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

// WithTimeout adds the timeout to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithTimeout(timeout time.Duration) *GetIPPoolsCredentialIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithContext(ctx context.Context) *GetIPPoolsCredentialIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithHTTPClient(client *http.Client) *GetIPPoolsCredentialIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithAvailabilityZone(availabilityZone *string) *GetIPPoolsCredentialIDParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithCredentialCrn(credentialCrn *string) *GetIPPoolsCredentialIDParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithCredentialName(credentialName *string) *GetIPPoolsCredentialIDParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithPlatformVariant(platformVariant *string) *GetIPPoolsCredentialIDParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) WithRegion(region *string) *GetIPPoolsCredentialIDParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get Ip pools credential Id params
func (o *GetIPPoolsCredentialIDParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPPoolsCredentialIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
