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

// NewPrepareDatalakeClusterUpgradeByCrnParams creates a new PrepareDatalakeClusterUpgradeByCrnParams object
// with the default values initialized.
func NewPrepareDatalakeClusterUpgradeByCrnParams() *PrepareDatalakeClusterUpgradeByCrnParams {
	var ()
	return &PrepareDatalakeClusterUpgradeByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrepareDatalakeClusterUpgradeByCrnParamsWithTimeout creates a new PrepareDatalakeClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrepareDatalakeClusterUpgradeByCrnParamsWithTimeout(timeout time.Duration) *PrepareDatalakeClusterUpgradeByCrnParams {
	var ()
	return &PrepareDatalakeClusterUpgradeByCrnParams{

		timeout: timeout,
	}
}

// NewPrepareDatalakeClusterUpgradeByCrnParamsWithContext creates a new PrepareDatalakeClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrepareDatalakeClusterUpgradeByCrnParamsWithContext(ctx context.Context) *PrepareDatalakeClusterUpgradeByCrnParams {
	var ()
	return &PrepareDatalakeClusterUpgradeByCrnParams{

		Context: ctx,
	}
}

// NewPrepareDatalakeClusterUpgradeByCrnParamsWithHTTPClient creates a new PrepareDatalakeClusterUpgradeByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrepareDatalakeClusterUpgradeByCrnParamsWithHTTPClient(client *http.Client) *PrepareDatalakeClusterUpgradeByCrnParams {
	var ()
	return &PrepareDatalakeClusterUpgradeByCrnParams{
		HTTPClient: client,
	}
}

/*
PrepareDatalakeClusterUpgradeByCrnParams contains all the parameters to send to the API endpoint
for the prepare datalake cluster upgrade by crn operation typically these are written to a http.Request
*/
type PrepareDatalakeClusterUpgradeByCrnParams struct {

	/*Body*/
	Body *model.SdxUpgradeRequest
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WithTimeout(timeout time.Duration) *PrepareDatalakeClusterUpgradeByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WithContext(ctx context.Context) *PrepareDatalakeClusterUpgradeByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WithHTTPClient(client *http.Client) *PrepareDatalakeClusterUpgradeByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WithBody(body *model.SdxUpgradeRequest) *PrepareDatalakeClusterUpgradeByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) SetBody(body *model.SdxUpgradeRequest) {
	o.Body = body
}

// WithCrn adds the crn to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WithCrn(crn string) *PrepareDatalakeClusterUpgradeByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the prepare datalake cluster upgrade by crn params
func (o *PrepareDatalakeClusterUpgradeByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *PrepareDatalakeClusterUpgradeByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
