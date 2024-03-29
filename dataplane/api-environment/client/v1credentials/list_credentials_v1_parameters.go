// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

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

// NewListCredentialsV1Params creates a new ListCredentialsV1Params object
// with the default values initialized.
func NewListCredentialsV1Params() *ListCredentialsV1Params {

	return &ListCredentialsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCredentialsV1ParamsWithTimeout creates a new ListCredentialsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCredentialsV1ParamsWithTimeout(timeout time.Duration) *ListCredentialsV1Params {

	return &ListCredentialsV1Params{

		timeout: timeout,
	}
}

// NewListCredentialsV1ParamsWithContext creates a new ListCredentialsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListCredentialsV1ParamsWithContext(ctx context.Context) *ListCredentialsV1Params {

	return &ListCredentialsV1Params{

		Context: ctx,
	}
}

// NewListCredentialsV1ParamsWithHTTPClient creates a new ListCredentialsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCredentialsV1ParamsWithHTTPClient(client *http.Client) *ListCredentialsV1Params {

	return &ListCredentialsV1Params{
		HTTPClient: client,
	}
}

/*
ListCredentialsV1Params contains all the parameters to send to the API endpoint
for the list credentials v1 operation typically these are written to a http.Request
*/
type ListCredentialsV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list credentials v1 params
func (o *ListCredentialsV1Params) WithTimeout(timeout time.Duration) *ListCredentialsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list credentials v1 params
func (o *ListCredentialsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list credentials v1 params
func (o *ListCredentialsV1Params) WithContext(ctx context.Context) *ListCredentialsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list credentials v1 params
func (o *ListCredentialsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list credentials v1 params
func (o *ListCredentialsV1Params) WithHTTPClient(client *http.Client) *ListCredentialsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list credentials v1 params
func (o *ListCredentialsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCredentialsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
