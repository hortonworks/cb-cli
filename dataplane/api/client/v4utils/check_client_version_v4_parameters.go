// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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

// NewCheckClientVersionV4Params creates a new CheckClientVersionV4Params object
// with the default values initialized.
func NewCheckClientVersionV4Params() *CheckClientVersionV4Params {
	var ()
	return &CheckClientVersionV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckClientVersionV4ParamsWithTimeout creates a new CheckClientVersionV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckClientVersionV4ParamsWithTimeout(timeout time.Duration) *CheckClientVersionV4Params {
	var ()
	return &CheckClientVersionV4Params{

		timeout: timeout,
	}
}

// NewCheckClientVersionV4ParamsWithContext creates a new CheckClientVersionV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewCheckClientVersionV4ParamsWithContext(ctx context.Context) *CheckClientVersionV4Params {
	var ()
	return &CheckClientVersionV4Params{

		Context: ctx,
	}
}

// NewCheckClientVersionV4ParamsWithHTTPClient creates a new CheckClientVersionV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckClientVersionV4ParamsWithHTTPClient(client *http.Client) *CheckClientVersionV4Params {
	var ()
	return &CheckClientVersionV4Params{
		HTTPClient: client,
	}
}

/*CheckClientVersionV4Params contains all the parameters to send to the API endpoint
for the check client version v4 operation typically these are written to a http.Request
*/
type CheckClientVersionV4Params struct {

	/*Version*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check client version v4 params
func (o *CheckClientVersionV4Params) WithTimeout(timeout time.Duration) *CheckClientVersionV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check client version v4 params
func (o *CheckClientVersionV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check client version v4 params
func (o *CheckClientVersionV4Params) WithContext(ctx context.Context) *CheckClientVersionV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check client version v4 params
func (o *CheckClientVersionV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check client version v4 params
func (o *CheckClientVersionV4Params) WithHTTPClient(client *http.Client) *CheckClientVersionV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check client version v4 params
func (o *CheckClientVersionV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the check client version v4 params
func (o *CheckClientVersionV4Params) WithVersion(version *string) *CheckClientVersionV4Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the check client version v4 params
func (o *CheckClientVersionV4Params) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CheckClientVersionV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
