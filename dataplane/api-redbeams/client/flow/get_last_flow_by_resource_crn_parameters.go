// Code generated by go-swagger; DO NOT EDIT.

package flow

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

// NewGetLastFlowByResourceCrnParams creates a new GetLastFlowByResourceCrnParams object
// with the default values initialized.
func NewGetLastFlowByResourceCrnParams() *GetLastFlowByResourceCrnParams {
	var ()
	return &GetLastFlowByResourceCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLastFlowByResourceCrnParamsWithTimeout creates a new GetLastFlowByResourceCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLastFlowByResourceCrnParamsWithTimeout(timeout time.Duration) *GetLastFlowByResourceCrnParams {
	var ()
	return &GetLastFlowByResourceCrnParams{

		timeout: timeout,
	}
}

// NewGetLastFlowByResourceCrnParamsWithContext creates a new GetLastFlowByResourceCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLastFlowByResourceCrnParamsWithContext(ctx context.Context) *GetLastFlowByResourceCrnParams {
	var ()
	return &GetLastFlowByResourceCrnParams{

		Context: ctx,
	}
}

// NewGetLastFlowByResourceCrnParamsWithHTTPClient creates a new GetLastFlowByResourceCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLastFlowByResourceCrnParamsWithHTTPClient(client *http.Client) *GetLastFlowByResourceCrnParams {
	var ()
	return &GetLastFlowByResourceCrnParams{
		HTTPClient: client,
	}
}

/*
GetLastFlowByResourceCrnParams contains all the parameters to send to the API endpoint
for the get last flow by resource crn operation typically these are written to a http.Request
*/
type GetLastFlowByResourceCrnParams struct {

	/*ResourceCrn*/
	ResourceCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) WithTimeout(timeout time.Duration) *GetLastFlowByResourceCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) WithContext(ctx context.Context) *GetLastFlowByResourceCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) WithHTTPClient(client *http.Client) *GetLastFlowByResourceCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceCrn adds the resourceCrn to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) WithResourceCrn(resourceCrn string) *GetLastFlowByResourceCrnParams {
	o.SetResourceCrn(resourceCrn)
	return o
}

// SetResourceCrn adds the resourceCrn to the get last flow by resource crn params
func (o *GetLastFlowByResourceCrnParams) SetResourceCrn(resourceCrn string) {
	o.ResourceCrn = resourceCrn
}

// WriteToRequest writes these params to a swagger request
func (o *GetLastFlowByResourceCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
