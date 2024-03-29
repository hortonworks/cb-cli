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
)

// NewRetrySdxParams creates a new RetrySdxParams object
// with the default values initialized.
func NewRetrySdxParams() *RetrySdxParams {
	var ()
	return &RetrySdxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrySdxParamsWithTimeout creates a new RetrySdxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrySdxParamsWithTimeout(timeout time.Duration) *RetrySdxParams {
	var ()
	return &RetrySdxParams{

		timeout: timeout,
	}
}

// NewRetrySdxParamsWithContext creates a new RetrySdxParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrySdxParamsWithContext(ctx context.Context) *RetrySdxParams {
	var ()
	return &RetrySdxParams{

		Context: ctx,
	}
}

// NewRetrySdxParamsWithHTTPClient creates a new RetrySdxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrySdxParamsWithHTTPClient(client *http.Client) *RetrySdxParams {
	var ()
	return &RetrySdxParams{
		HTTPClient: client,
	}
}

/*
RetrySdxParams contains all the parameters to send to the API endpoint
for the retry sdx operation typically these are written to a http.Request
*/
type RetrySdxParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retry sdx params
func (o *RetrySdxParams) WithTimeout(timeout time.Duration) *RetrySdxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retry sdx params
func (o *RetrySdxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retry sdx params
func (o *RetrySdxParams) WithContext(ctx context.Context) *RetrySdxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retry sdx params
func (o *RetrySdxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retry sdx params
func (o *RetrySdxParams) WithHTTPClient(client *http.Client) *RetrySdxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retry sdx params
func (o *RetrySdxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the retry sdx params
func (o *RetrySdxParams) WithName(name string) *RetrySdxParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the retry sdx params
func (o *RetrySdxParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RetrySdxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
