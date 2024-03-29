// Code generated by go-swagger; DO NOT EDIT.

package sdxutils

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

// NewCheckClientVersionOfSdxParams creates a new CheckClientVersionOfSdxParams object
// with the default values initialized.
func NewCheckClientVersionOfSdxParams() *CheckClientVersionOfSdxParams {
	var ()
	return &CheckClientVersionOfSdxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckClientVersionOfSdxParamsWithTimeout creates a new CheckClientVersionOfSdxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckClientVersionOfSdxParamsWithTimeout(timeout time.Duration) *CheckClientVersionOfSdxParams {
	var ()
	return &CheckClientVersionOfSdxParams{

		timeout: timeout,
	}
}

// NewCheckClientVersionOfSdxParamsWithContext creates a new CheckClientVersionOfSdxParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckClientVersionOfSdxParamsWithContext(ctx context.Context) *CheckClientVersionOfSdxParams {
	var ()
	return &CheckClientVersionOfSdxParams{

		Context: ctx,
	}
}

// NewCheckClientVersionOfSdxParamsWithHTTPClient creates a new CheckClientVersionOfSdxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckClientVersionOfSdxParamsWithHTTPClient(client *http.Client) *CheckClientVersionOfSdxParams {
	var ()
	return &CheckClientVersionOfSdxParams{
		HTTPClient: client,
	}
}

/*
CheckClientVersionOfSdxParams contains all the parameters to send to the API endpoint
for the check client version of sdx operation typically these are written to a http.Request
*/
type CheckClientVersionOfSdxParams struct {

	/*Version*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) WithTimeout(timeout time.Duration) *CheckClientVersionOfSdxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) WithContext(ctx context.Context) *CheckClientVersionOfSdxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) WithHTTPClient(client *http.Client) *CheckClientVersionOfSdxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) WithVersion(version *string) *CheckClientVersionOfSdxParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the check client version of sdx params
func (o *CheckClientVersionOfSdxParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CheckClientVersionOfSdxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
