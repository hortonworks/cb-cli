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

// NewVerifyCredentialByEnvCrnParams creates a new VerifyCredentialByEnvCrnParams object
// with the default values initialized.
func NewVerifyCredentialByEnvCrnParams() *VerifyCredentialByEnvCrnParams {
	var ()
	return &VerifyCredentialByEnvCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyCredentialByEnvCrnParamsWithTimeout creates a new VerifyCredentialByEnvCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyCredentialByEnvCrnParamsWithTimeout(timeout time.Duration) *VerifyCredentialByEnvCrnParams {
	var ()
	return &VerifyCredentialByEnvCrnParams{

		timeout: timeout,
	}
}

// NewVerifyCredentialByEnvCrnParamsWithContext creates a new VerifyCredentialByEnvCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyCredentialByEnvCrnParamsWithContext(ctx context.Context) *VerifyCredentialByEnvCrnParams {
	var ()
	return &VerifyCredentialByEnvCrnParams{

		Context: ctx,
	}
}

// NewVerifyCredentialByEnvCrnParamsWithHTTPClient creates a new VerifyCredentialByEnvCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyCredentialByEnvCrnParamsWithHTTPClient(client *http.Client) *VerifyCredentialByEnvCrnParams {
	var ()
	return &VerifyCredentialByEnvCrnParams{
		HTTPClient: client,
	}
}

/*
VerifyCredentialByEnvCrnParams contains all the parameters to send to the API endpoint
for the verify credential by env crn operation typically these are written to a http.Request
*/
type VerifyCredentialByEnvCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) WithTimeout(timeout time.Duration) *VerifyCredentialByEnvCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) WithContext(ctx context.Context) *VerifyCredentialByEnvCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) WithHTTPClient(client *http.Client) *VerifyCredentialByEnvCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) WithCrn(crn string) *VerifyCredentialByEnvCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the verify credential by env crn params
func (o *VerifyCredentialByEnvCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyCredentialByEnvCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
