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

// NewGetNoSQLTablesParams creates a new GetNoSQLTablesParams object
// with the default values initialized.
func NewGetNoSQLTablesParams() *GetNoSQLTablesParams {
	var ()
	return &GetNoSQLTablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNoSQLTablesParamsWithTimeout creates a new GetNoSQLTablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNoSQLTablesParamsWithTimeout(timeout time.Duration) *GetNoSQLTablesParams {
	var ()
	return &GetNoSQLTablesParams{

		timeout: timeout,
	}
}

// NewGetNoSQLTablesParamsWithContext creates a new GetNoSQLTablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNoSQLTablesParamsWithContext(ctx context.Context) *GetNoSQLTablesParams {
	var ()
	return &GetNoSQLTablesParams{

		Context: ctx,
	}
}

// NewGetNoSQLTablesParamsWithHTTPClient creates a new GetNoSQLTablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNoSQLTablesParamsWithHTTPClient(client *http.Client) *GetNoSQLTablesParams {
	var ()
	return &GetNoSQLTablesParams{
		HTTPClient: client,
	}
}

/*GetNoSQLTablesParams contains all the parameters to send to the API endpoint
for the get no Sql tables operation typically these are written to a http.Request
*/
type GetNoSQLTablesParams struct {

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

// WithTimeout adds the timeout to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithTimeout(timeout time.Duration) *GetNoSQLTablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithContext(ctx context.Context) *GetNoSQLTablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithHTTPClient(client *http.Client) *GetNoSQLTablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithAvailabilityZone(availabilityZone *string) *GetNoSQLTablesParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithCredentialCrn(credentialCrn *string) *GetNoSQLTablesParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithCredentialName(credentialName *string) *GetNoSQLTablesParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithPlatformVariant(platformVariant *string) *GetNoSQLTablesParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get no Sql tables params
func (o *GetNoSQLTablesParams) WithRegion(region *string) *GetNoSQLTablesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get no Sql tables params
func (o *GetNoSQLTablesParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetNoSQLTablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
