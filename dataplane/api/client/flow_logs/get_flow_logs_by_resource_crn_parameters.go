// Code generated by go-swagger; DO NOT EDIT.

package flow_logs

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

// NewGetFlowLogsByResourceCrnParams creates a new GetFlowLogsByResourceCrnParams object
// with the default values initialized.
func NewGetFlowLogsByResourceCrnParams() *GetFlowLogsByResourceCrnParams {
	var ()
	return &GetFlowLogsByResourceCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowLogsByResourceCrnParamsWithTimeout creates a new GetFlowLogsByResourceCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowLogsByResourceCrnParamsWithTimeout(timeout time.Duration) *GetFlowLogsByResourceCrnParams {
	var ()
	return &GetFlowLogsByResourceCrnParams{

		timeout: timeout,
	}
}

// NewGetFlowLogsByResourceCrnParamsWithContext creates a new GetFlowLogsByResourceCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowLogsByResourceCrnParamsWithContext(ctx context.Context) *GetFlowLogsByResourceCrnParams {
	var ()
	return &GetFlowLogsByResourceCrnParams{

		Context: ctx,
	}
}

// NewGetFlowLogsByResourceCrnParamsWithHTTPClient creates a new GetFlowLogsByResourceCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowLogsByResourceCrnParamsWithHTTPClient(client *http.Client) *GetFlowLogsByResourceCrnParams {
	var ()
	return &GetFlowLogsByResourceCrnParams{
		HTTPClient: client,
	}
}

/*GetFlowLogsByResourceCrnParams contains all the parameters to send to the API endpoint
for the get flow logs by resource crn operation typically these are written to a http.Request
*/
type GetFlowLogsByResourceCrnParams struct {

	/*ResourceCrn*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) WithTimeout(timeout time.Duration) *GetFlowLogsByResourceCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) WithContext(ctx context.Context) *GetFlowLogsByResourceCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) WithHTTPClient(client *http.Client) *GetFlowLogsByResourceCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceCrn adds the resourceCrn to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) WithResourceCrn(resourceCrn string) *GetFlowLogsByResourceCrnParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the get flow logs by resource crn params
func (o *GetFlowLogsByResourceCrnParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowLogsByResourceCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceCrn
	if err := r.SetPathParam("resourceCrn", o.ResourceCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
