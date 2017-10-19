// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewUpgradeClusterParams creates a new UpgradeClusterParams object
// with the default values initialized.
func NewUpgradeClusterParams() *UpgradeClusterParams {
	var ()
	return &UpgradeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeClusterParamsWithTimeout creates a new UpgradeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeClusterParamsWithTimeout(timeout time.Duration) *UpgradeClusterParams {
	var ()
	return &UpgradeClusterParams{

		timeout: timeout,
	}
}

// NewUpgradeClusterParamsWithContext creates a new UpgradeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeClusterParamsWithContext(ctx context.Context) *UpgradeClusterParams {
	var ()
	return &UpgradeClusterParams{

		Context: ctx,
	}
}

// NewUpgradeClusterParamsWithHTTPClient creates a new UpgradeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeClusterParamsWithHTTPClient(client *http.Client) *UpgradeClusterParams {
	var ()
	return &UpgradeClusterParams{
		HTTPClient: client,
	}
}

/*UpgradeClusterParams contains all the parameters to send to the API endpoint
for the upgrade cluster operation typically these are written to a http.Request
*/
type UpgradeClusterParams struct {

	/*Body*/
	Body *models_cloudbreak.AmbariRepoDetails
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade cluster params
func (o *UpgradeClusterParams) WithTimeout(timeout time.Duration) *UpgradeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade cluster params
func (o *UpgradeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade cluster params
func (o *UpgradeClusterParams) WithContext(ctx context.Context) *UpgradeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade cluster params
func (o *UpgradeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade cluster params
func (o *UpgradeClusterParams) WithHTTPClient(client *http.Client) *UpgradeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade cluster params
func (o *UpgradeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade cluster params
func (o *UpgradeClusterParams) WithBody(body *models_cloudbreak.AmbariRepoDetails) *UpgradeClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade cluster params
func (o *UpgradeClusterParams) SetBody(body *models_cloudbreak.AmbariRepoDetails) {
	o.Body = body
}

// WithID adds the id to the upgrade cluster params
func (o *UpgradeClusterParams) WithID(id int64) *UpgradeClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the upgrade cluster params
func (o *UpgradeClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.AmbariRepoDetails)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
