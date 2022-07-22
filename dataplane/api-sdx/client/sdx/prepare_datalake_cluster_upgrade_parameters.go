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

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// NewPrepareDatalakeClusterUpgradeParams creates a new PrepareDatalakeClusterUpgradeParams object
// with the default values initialized.
func NewPrepareDatalakeClusterUpgradeParams() *PrepareDatalakeClusterUpgradeParams {
	var ()
	return &PrepareDatalakeClusterUpgradeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrepareDatalakeClusterUpgradeParamsWithTimeout creates a new PrepareDatalakeClusterUpgradeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrepareDatalakeClusterUpgradeParamsWithTimeout(timeout time.Duration) *PrepareDatalakeClusterUpgradeParams {
	var ()
	return &PrepareDatalakeClusterUpgradeParams{

		timeout: timeout,
	}
}

// NewPrepareDatalakeClusterUpgradeParamsWithContext creates a new PrepareDatalakeClusterUpgradeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrepareDatalakeClusterUpgradeParamsWithContext(ctx context.Context) *PrepareDatalakeClusterUpgradeParams {
	var ()
	return &PrepareDatalakeClusterUpgradeParams{

		Context: ctx,
	}
}

// NewPrepareDatalakeClusterUpgradeParamsWithHTTPClient creates a new PrepareDatalakeClusterUpgradeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrepareDatalakeClusterUpgradeParamsWithHTTPClient(client *http.Client) *PrepareDatalakeClusterUpgradeParams {
	var ()
	return &PrepareDatalakeClusterUpgradeParams{
		HTTPClient: client,
	}
}

/*PrepareDatalakeClusterUpgradeParams contains all the parameters to send to the API endpoint
for the prepare datalake cluster upgrade operation typically these are written to a http.Request
*/
type PrepareDatalakeClusterUpgradeParams struct {

	/*Body*/
	Body *model.SdxUpgradeRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) WithTimeout(timeout time.Duration) *PrepareDatalakeClusterUpgradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) WithContext(ctx context.Context) *PrepareDatalakeClusterUpgradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) WithHTTPClient(client *http.Client) *PrepareDatalakeClusterUpgradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) WithBody(body *model.SdxUpgradeRequest) *PrepareDatalakeClusterUpgradeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) SetBody(body *model.SdxUpgradeRequest) {
	o.Body = body
}

// WithName adds the name to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) WithName(name string) *PrepareDatalakeClusterUpgradeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the prepare datalake cluster upgrade params
func (o *PrepareDatalakeClusterUpgradeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PrepareDatalakeClusterUpgradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
