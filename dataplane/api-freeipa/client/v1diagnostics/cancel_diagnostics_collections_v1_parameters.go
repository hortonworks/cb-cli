// Code generated by go-swagger; DO NOT EDIT.

package v1diagnostics

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

// NewCancelDiagnosticsCollectionsV1Params creates a new CancelDiagnosticsCollectionsV1Params object
// with the default values initialized.
func NewCancelDiagnosticsCollectionsV1Params() *CancelDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDiagnosticsCollectionsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelDiagnosticsCollectionsV1ParamsWithTimeout creates a new CancelDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelDiagnosticsCollectionsV1ParamsWithTimeout(timeout time.Duration) *CancelDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDiagnosticsCollectionsV1Params{

		timeout: timeout,
	}
}

// NewCancelDiagnosticsCollectionsV1ParamsWithContext creates a new CancelDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCancelDiagnosticsCollectionsV1ParamsWithContext(ctx context.Context) *CancelDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDiagnosticsCollectionsV1Params{

		Context: ctx,
	}
}

// NewCancelDiagnosticsCollectionsV1ParamsWithHTTPClient creates a new CancelDiagnosticsCollectionsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelDiagnosticsCollectionsV1ParamsWithHTTPClient(client *http.Client) *CancelDiagnosticsCollectionsV1Params {
	var ()
	return &CancelDiagnosticsCollectionsV1Params{
		HTTPClient: client,
	}
}

/*
CancelDiagnosticsCollectionsV1Params contains all the parameters to send to the API endpoint
for the cancel diagnostics collections v1 operation typically these are written to a http.Request
*/
type CancelDiagnosticsCollectionsV1Params struct {

	/*EnvironmentCrn*/
	EnvironmentCrn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) WithTimeout(timeout time.Duration) *CancelDiagnosticsCollectionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) WithContext(ctx context.Context) *CancelDiagnosticsCollectionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) WithHTTPClient(client *http.Client) *CancelDiagnosticsCollectionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentCrn adds the environmentCrn to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) WithEnvironmentCrn(environmentCrn string) *CancelDiagnosticsCollectionsV1Params {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the cancel diagnostics collections v1 params
func (o *CancelDiagnosticsCollectionsV1Params) SetEnvironmentCrn(environmentCrn string) {
	o.EnvironmentCrn = environmentCrn
}

// WriteToRequest writes these params to a swagger request
func (o *CancelDiagnosticsCollectionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentCrn
	if err := r.SetPathParam("environmentCrn", o.EnvironmentCrn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
