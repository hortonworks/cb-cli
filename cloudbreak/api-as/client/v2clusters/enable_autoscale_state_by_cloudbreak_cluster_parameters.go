// Code generated by go-swagger; DO NOT EDIT.

package v2clusters

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

// NewEnableAutoscaleStateByCloudbreakClusterParams creates a new EnableAutoscaleStateByCloudbreakClusterParams object
// with the default values initialized.
func NewEnableAutoscaleStateByCloudbreakClusterParams() *EnableAutoscaleStateByCloudbreakClusterParams {
	var ()
	return &EnableAutoscaleStateByCloudbreakClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableAutoscaleStateByCloudbreakClusterParamsWithTimeout creates a new EnableAutoscaleStateByCloudbreakClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableAutoscaleStateByCloudbreakClusterParamsWithTimeout(timeout time.Duration) *EnableAutoscaleStateByCloudbreakClusterParams {
	var ()
	return &EnableAutoscaleStateByCloudbreakClusterParams{

		timeout: timeout,
	}
}

// NewEnableAutoscaleStateByCloudbreakClusterParamsWithContext creates a new EnableAutoscaleStateByCloudbreakClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableAutoscaleStateByCloudbreakClusterParamsWithContext(ctx context.Context) *EnableAutoscaleStateByCloudbreakClusterParams {
	var ()
	return &EnableAutoscaleStateByCloudbreakClusterParams{

		Context: ctx,
	}
}

// NewEnableAutoscaleStateByCloudbreakClusterParamsWithHTTPClient creates a new EnableAutoscaleStateByCloudbreakClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableAutoscaleStateByCloudbreakClusterParamsWithHTTPClient(client *http.Client) *EnableAutoscaleStateByCloudbreakClusterParams {
	var ()
	return &EnableAutoscaleStateByCloudbreakClusterParams{
		HTTPClient: client,
	}
}

/*EnableAutoscaleStateByCloudbreakClusterParams contains all the parameters to send to the API endpoint
for the enable autoscale state by cloudbreak cluster operation typically these are written to a http.Request
*/
type EnableAutoscaleStateByCloudbreakClusterParams struct {

	/*CbClusterID*/
	CbClusterID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) WithTimeout(timeout time.Duration) *EnableAutoscaleStateByCloudbreakClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) WithContext(ctx context.Context) *EnableAutoscaleStateByCloudbreakClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) WithHTTPClient(client *http.Client) *EnableAutoscaleStateByCloudbreakClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCbClusterID adds the cbClusterID to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) WithCbClusterID(cbClusterID int64) *EnableAutoscaleStateByCloudbreakClusterParams {
	o.SetCbClusterID(cbClusterID)
	return o
}

// SetCbClusterID adds the cbClusterId to the enable autoscale state by cloudbreak cluster params
func (o *EnableAutoscaleStateByCloudbreakClusterParams) SetCbClusterID(cbClusterID int64) {
	o.CbClusterID = cbClusterID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableAutoscaleStateByCloudbreakClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cbClusterId
	if err := r.SetPathParam("cbClusterId", swag.FormatInt64(o.CbClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
