// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewUpgradeDistroxClusterParams creates a new UpgradeDistroxClusterParams object
// with the default values initialized.
func NewUpgradeDistroxClusterParams() *UpgradeDistroxClusterParams {
	var ()
	return &UpgradeDistroxClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeDistroxClusterParamsWithTimeout creates a new UpgradeDistroxClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeDistroxClusterParamsWithTimeout(timeout time.Duration) *UpgradeDistroxClusterParams {
	var ()
	return &UpgradeDistroxClusterParams{

		timeout: timeout,
	}
}

// NewUpgradeDistroxClusterParamsWithContext creates a new UpgradeDistroxClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeDistroxClusterParamsWithContext(ctx context.Context) *UpgradeDistroxClusterParams {
	var ()
	return &UpgradeDistroxClusterParams{

		Context: ctx,
	}
}

// NewUpgradeDistroxClusterParamsWithHTTPClient creates a new UpgradeDistroxClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeDistroxClusterParamsWithHTTPClient(client *http.Client) *UpgradeDistroxClusterParams {
	var ()
	return &UpgradeDistroxClusterParams{
		HTTPClient: client,
	}
}

/*
UpgradeDistroxClusterParams contains all the parameters to send to the API endpoint
for the upgrade distrox cluster operation typically these are written to a http.Request
*/
type UpgradeDistroxClusterParams struct {

	/*Body*/
	Body *model.DistroXUpgradeV1Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) WithTimeout(timeout time.Duration) *UpgradeDistroxClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) WithContext(ctx context.Context) *UpgradeDistroxClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) WithHTTPClient(client *http.Client) *UpgradeDistroxClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) WithBody(body *model.DistroXUpgradeV1Request) *UpgradeDistroxClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) SetBody(body *model.DistroXUpgradeV1Request) {
	o.Body = body
}

// WithName adds the name to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) WithName(name string) *UpgradeDistroxClusterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upgrade distrox cluster params
func (o *UpgradeDistroxClusterParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeDistroxClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
