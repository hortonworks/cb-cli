// Code generated by go-swagger; DO NOT EDIT.

package v1utils

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

// NewCheckClientVersionOfEnvironmentV1Params creates a new CheckClientVersionOfEnvironmentV1Params object
// with the default values initialized.
func NewCheckClientVersionOfEnvironmentV1Params() *CheckClientVersionOfEnvironmentV1Params {
	var ()
	return &CheckClientVersionOfEnvironmentV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckClientVersionOfEnvironmentV1ParamsWithTimeout creates a new CheckClientVersionOfEnvironmentV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckClientVersionOfEnvironmentV1ParamsWithTimeout(timeout time.Duration) *CheckClientVersionOfEnvironmentV1Params {
	var ()
	return &CheckClientVersionOfEnvironmentV1Params{

		timeout: timeout,
	}
}

// NewCheckClientVersionOfEnvironmentV1ParamsWithContext creates a new CheckClientVersionOfEnvironmentV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCheckClientVersionOfEnvironmentV1ParamsWithContext(ctx context.Context) *CheckClientVersionOfEnvironmentV1Params {
	var ()
	return &CheckClientVersionOfEnvironmentV1Params{

		Context: ctx,
	}
}

// NewCheckClientVersionOfEnvironmentV1ParamsWithHTTPClient creates a new CheckClientVersionOfEnvironmentV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckClientVersionOfEnvironmentV1ParamsWithHTTPClient(client *http.Client) *CheckClientVersionOfEnvironmentV1Params {
	var ()
	return &CheckClientVersionOfEnvironmentV1Params{
		HTTPClient: client,
	}
}

/*
CheckClientVersionOfEnvironmentV1Params contains all the parameters to send to the API endpoint
for the check client version of environment v1 operation typically these are written to a http.Request
*/
type CheckClientVersionOfEnvironmentV1Params struct {

	/*Version*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) WithTimeout(timeout time.Duration) *CheckClientVersionOfEnvironmentV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) WithContext(ctx context.Context) *CheckClientVersionOfEnvironmentV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) WithHTTPClient(client *http.Client) *CheckClientVersionOfEnvironmentV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) WithVersion(version *string) *CheckClientVersionOfEnvironmentV1Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the check client version of environment v1 params
func (o *CheckClientVersionOfEnvironmentV1Params) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CheckClientVersionOfEnvironmentV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
