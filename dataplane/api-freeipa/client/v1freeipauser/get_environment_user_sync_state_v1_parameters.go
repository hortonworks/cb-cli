// Code generated by go-swagger; DO NOT EDIT.

package v1freeipauser

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

// NewGetEnvironmentUserSyncStateV1Params creates a new GetEnvironmentUserSyncStateV1Params object
// with the default values initialized.
func NewGetEnvironmentUserSyncStateV1Params() *GetEnvironmentUserSyncStateV1Params {
	var ()
	return &GetEnvironmentUserSyncStateV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentUserSyncStateV1ParamsWithTimeout creates a new GetEnvironmentUserSyncStateV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentUserSyncStateV1ParamsWithTimeout(timeout time.Duration) *GetEnvironmentUserSyncStateV1Params {
	var ()
	return &GetEnvironmentUserSyncStateV1Params{

		timeout: timeout,
	}
}

// NewGetEnvironmentUserSyncStateV1ParamsWithContext creates a new GetEnvironmentUserSyncStateV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentUserSyncStateV1ParamsWithContext(ctx context.Context) *GetEnvironmentUserSyncStateV1Params {
	var ()
	return &GetEnvironmentUserSyncStateV1Params{

		Context: ctx,
	}
}

// NewGetEnvironmentUserSyncStateV1ParamsWithHTTPClient creates a new GetEnvironmentUserSyncStateV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentUserSyncStateV1ParamsWithHTTPClient(client *http.Client) *GetEnvironmentUserSyncStateV1Params {
	var ()
	return &GetEnvironmentUserSyncStateV1Params{
		HTTPClient: client,
	}
}

/*GetEnvironmentUserSyncStateV1Params contains all the parameters to send to the API endpoint
for the get environment user sync state v1 operation typically these are written to a http.Request
*/
type GetEnvironmentUserSyncStateV1Params struct {

	/*EnvironmentCrn*/
	EnvironmentCrn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) WithTimeout(timeout time.Duration) *GetEnvironmentUserSyncStateV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) WithContext(ctx context.Context) *GetEnvironmentUserSyncStateV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) WithHTTPClient(client *http.Client) *GetEnvironmentUserSyncStateV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) WithEnvironmentCrn(environmentCrn *string) *GetEnvironmentUserSyncStateV1Params {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get environment user sync state v1 params
func (o *GetEnvironmentUserSyncStateV1Params) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentUserSyncStateV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
