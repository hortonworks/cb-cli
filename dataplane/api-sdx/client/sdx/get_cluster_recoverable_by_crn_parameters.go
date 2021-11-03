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

// NewGetClusterRecoverableByCrnParams creates a new GetClusterRecoverableByCrnParams object
// with the default values initialized.
func NewGetClusterRecoverableByCrnParams() *GetClusterRecoverableByCrnParams {
	var ()
	return &GetClusterRecoverableByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterRecoverableByCrnParamsWithTimeout creates a new GetClusterRecoverableByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterRecoverableByCrnParamsWithTimeout(timeout time.Duration) *GetClusterRecoverableByCrnParams {
	var ()
	return &GetClusterRecoverableByCrnParams{

		timeout: timeout,
	}
}

// NewGetClusterRecoverableByCrnParamsWithContext creates a new GetClusterRecoverableByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterRecoverableByCrnParamsWithContext(ctx context.Context) *GetClusterRecoverableByCrnParams {
	var ()
	return &GetClusterRecoverableByCrnParams{

		Context: ctx,
	}
}

// NewGetClusterRecoverableByCrnParamsWithHTTPClient creates a new GetClusterRecoverableByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterRecoverableByCrnParamsWithHTTPClient(client *http.Client) *GetClusterRecoverableByCrnParams {
	var ()
	return &GetClusterRecoverableByCrnParams{
		HTTPClient: client,
	}
}

/*GetClusterRecoverableByCrnParams contains all the parameters to send to the API endpoint
for the get cluster recoverable by crn operation typically these are written to a http.Request
*/
type GetClusterRecoverableByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) WithTimeout(timeout time.Duration) *GetClusterRecoverableByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) WithContext(ctx context.Context) *GetClusterRecoverableByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) WithHTTPClient(client *http.Client) *GetClusterRecoverableByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) WithCrn(crn string) *GetClusterRecoverableByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get cluster recoverable by crn params
func (o *GetClusterRecoverableByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterRecoverableByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
