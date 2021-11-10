// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

// NewUpgradeCcmByEnvironmentNameV1Params creates a new UpgradeCcmByEnvironmentNameV1Params object
// with the default values initialized.
func NewUpgradeCcmByEnvironmentNameV1Params() *UpgradeCcmByEnvironmentNameV1Params {
	var ()
	return &UpgradeCcmByEnvironmentNameV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCcmByEnvironmentNameV1ParamsWithTimeout creates a new UpgradeCcmByEnvironmentNameV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeCcmByEnvironmentNameV1ParamsWithTimeout(timeout time.Duration) *UpgradeCcmByEnvironmentNameV1Params {
	var ()
	return &UpgradeCcmByEnvironmentNameV1Params{

		timeout: timeout,
	}
}

// NewUpgradeCcmByEnvironmentNameV1ParamsWithContext creates a new UpgradeCcmByEnvironmentNameV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeCcmByEnvironmentNameV1ParamsWithContext(ctx context.Context) *UpgradeCcmByEnvironmentNameV1Params {
	var ()
	return &UpgradeCcmByEnvironmentNameV1Params{

		Context: ctx,
	}
}

// NewUpgradeCcmByEnvironmentNameV1ParamsWithHTTPClient creates a new UpgradeCcmByEnvironmentNameV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeCcmByEnvironmentNameV1ParamsWithHTTPClient(client *http.Client) *UpgradeCcmByEnvironmentNameV1Params {
	var ()
	return &UpgradeCcmByEnvironmentNameV1Params{
		HTTPClient: client,
	}
}

/*UpgradeCcmByEnvironmentNameV1Params contains all the parameters to send to the API endpoint
for the upgrade ccm by environment name v1 operation typically these are written to a http.Request
*/
type UpgradeCcmByEnvironmentNameV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) WithTimeout(timeout time.Duration) *UpgradeCcmByEnvironmentNameV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) WithContext(ctx context.Context) *UpgradeCcmByEnvironmentNameV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) WithHTTPClient(client *http.Client) *UpgradeCcmByEnvironmentNameV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) WithName(name string) *UpgradeCcmByEnvironmentNameV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the upgrade ccm by environment name v1 params
func (o *UpgradeCcmByEnvironmentNameV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCcmByEnvironmentNameV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
