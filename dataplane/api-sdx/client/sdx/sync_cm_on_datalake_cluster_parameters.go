// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewSyncCmOnDatalakeClusterParams creates a new SyncCmOnDatalakeClusterParams object
// with the default values initialized.
func NewSyncCmOnDatalakeClusterParams() *SyncCmOnDatalakeClusterParams {
	var ()
	return &SyncCmOnDatalakeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncCmOnDatalakeClusterParamsWithTimeout creates a new SyncCmOnDatalakeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncCmOnDatalakeClusterParamsWithTimeout(timeout time.Duration) *SyncCmOnDatalakeClusterParams {
	var ()
	return &SyncCmOnDatalakeClusterParams{

		timeout: timeout,
	}
}

// NewSyncCmOnDatalakeClusterParamsWithContext creates a new SyncCmOnDatalakeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncCmOnDatalakeClusterParamsWithContext(ctx context.Context) *SyncCmOnDatalakeClusterParams {
	var ()
	return &SyncCmOnDatalakeClusterParams{

		Context: ctx,
	}
}

// NewSyncCmOnDatalakeClusterParamsWithHTTPClient creates a new SyncCmOnDatalakeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncCmOnDatalakeClusterParamsWithHTTPClient(client *http.Client) *SyncCmOnDatalakeClusterParams {
	var ()
	return &SyncCmOnDatalakeClusterParams{
		HTTPClient: client,
	}
}

/*SyncCmOnDatalakeClusterParams contains all the parameters to send to the API endpoint
for the sync cm on datalake cluster operation typically these are written to a http.Request
*/
type SyncCmOnDatalakeClusterParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) WithTimeout(timeout time.Duration) *SyncCmOnDatalakeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) WithContext(ctx context.Context) *SyncCmOnDatalakeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) WithHTTPClient(client *http.Client) *SyncCmOnDatalakeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) WithName(name string) *SyncCmOnDatalakeClusterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the sync cm on datalake cluster params
func (o *SyncCmOnDatalakeClusterParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SyncCmOnDatalakeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
