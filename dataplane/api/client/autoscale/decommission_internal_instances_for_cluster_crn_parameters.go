// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDecommissionInternalInstancesForClusterCrnParams creates a new DecommissionInternalInstancesForClusterCrnParams object
// with the default values initialized.
func NewDecommissionInternalInstancesForClusterCrnParams() *DecommissionInternalInstancesForClusterCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DecommissionInternalInstancesForClusterCrnParams{
		Forced: &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDecommissionInternalInstancesForClusterCrnParamsWithTimeout creates a new DecommissionInternalInstancesForClusterCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDecommissionInternalInstancesForClusterCrnParamsWithTimeout(timeout time.Duration) *DecommissionInternalInstancesForClusterCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DecommissionInternalInstancesForClusterCrnParams{
		Forced: &forcedDefault,

		timeout: timeout,
	}
}

// NewDecommissionInternalInstancesForClusterCrnParamsWithContext creates a new DecommissionInternalInstancesForClusterCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDecommissionInternalInstancesForClusterCrnParamsWithContext(ctx context.Context) *DecommissionInternalInstancesForClusterCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DecommissionInternalInstancesForClusterCrnParams{
		Forced: &forcedDefault,

		Context: ctx,
	}
}

// NewDecommissionInternalInstancesForClusterCrnParamsWithHTTPClient creates a new DecommissionInternalInstancesForClusterCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDecommissionInternalInstancesForClusterCrnParamsWithHTTPClient(client *http.Client) *DecommissionInternalInstancesForClusterCrnParams {
	var (
		forcedDefault = bool(false)
	)
	return &DecommissionInternalInstancesForClusterCrnParams{
		Forced:     &forcedDefault,
		HTTPClient: client,
	}
}

/*
DecommissionInternalInstancesForClusterCrnParams contains all the parameters to send to the API endpoint
for the decommission internal instances for cluster crn operation typically these are written to a http.Request
*/
type DecommissionInternalInstancesForClusterCrnParams struct {

	/*Body*/
	Body []string
	/*Crn*/
	Crn string
	/*Forced*/
	Forced *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithTimeout(timeout time.Duration) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithContext(ctx context.Context) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithHTTPClient(client *http.Client) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithBody(body []string) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetBody(body []string) {
	o.Body = body
}

// WithCrn adds the crn to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithCrn(crn string) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WithForced adds the forced to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) WithForced(forced *bool) *DecommissionInternalInstancesForClusterCrnParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the decommission internal instances for cluster crn params
func (o *DecommissionInternalInstancesForClusterCrnParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WriteToRequest writes these params to a swagger request
func (o *DecommissionInternalInstancesForClusterCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
