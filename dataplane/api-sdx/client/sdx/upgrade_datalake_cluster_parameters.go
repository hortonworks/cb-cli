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

// NewUpgradeDatalakeClusterParams creates a new UpgradeDatalakeClusterParams object
// with the default values initialized.
func NewUpgradeDatalakeClusterParams() *UpgradeDatalakeClusterParams {
	var ()
	return &UpgradeDatalakeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeDatalakeClusterParamsWithTimeout creates a new UpgradeDatalakeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeDatalakeClusterParamsWithTimeout(timeout time.Duration) *UpgradeDatalakeClusterParams {
	var ()
	return &UpgradeDatalakeClusterParams{

		timeout: timeout,
	}
}

// NewUpgradeDatalakeClusterParamsWithContext creates a new UpgradeDatalakeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeDatalakeClusterParamsWithContext(ctx context.Context) *UpgradeDatalakeClusterParams {
	var ()
	return &UpgradeDatalakeClusterParams{

		Context: ctx,
	}
}

// NewUpgradeDatalakeClusterParamsWithHTTPClient creates a new UpgradeDatalakeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeDatalakeClusterParamsWithHTTPClient(client *http.Client) *UpgradeDatalakeClusterParams {
	var ()
	return &UpgradeDatalakeClusterParams{
		HTTPClient: client,
	}
}

/*
UpgradeDatalakeClusterParams contains all the parameters to send to the API endpoint
for the upgrade datalake cluster operation typically these are written to a http.Request
*/
type UpgradeDatalakeClusterParams struct {

	/*Body*/
	Body *model.SdxUpgradeRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) WithTimeout(timeout time.Duration) *UpgradeDatalakeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) WithContext(ctx context.Context) *UpgradeDatalakeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) WithHTTPClient(client *http.Client) *UpgradeDatalakeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) WithBody(body *model.SdxUpgradeRequest) *UpgradeDatalakeClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) SetBody(body *model.SdxUpgradeRequest) {
	o.Body = body
}

// WithName adds the name to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) WithName(name string) *UpgradeDatalakeClusterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upgrade datalake cluster params
func (o *UpgradeDatalakeClusterParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeDatalakeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
