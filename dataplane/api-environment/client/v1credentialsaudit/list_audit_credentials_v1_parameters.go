// Code generated by go-swagger; DO NOT EDIT.

package v1credentialsaudit

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

// NewListAuditCredentialsV1Params creates a new ListAuditCredentialsV1Params object
// with the default values initialized.
func NewListAuditCredentialsV1Params() *ListAuditCredentialsV1Params {

	return &ListAuditCredentialsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAuditCredentialsV1ParamsWithTimeout creates a new ListAuditCredentialsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAuditCredentialsV1ParamsWithTimeout(timeout time.Duration) *ListAuditCredentialsV1Params {

	return &ListAuditCredentialsV1Params{

		timeout: timeout,
	}
}

// NewListAuditCredentialsV1ParamsWithContext creates a new ListAuditCredentialsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListAuditCredentialsV1ParamsWithContext(ctx context.Context) *ListAuditCredentialsV1Params {

	return &ListAuditCredentialsV1Params{

		Context: ctx,
	}
}

// NewListAuditCredentialsV1ParamsWithHTTPClient creates a new ListAuditCredentialsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAuditCredentialsV1ParamsWithHTTPClient(client *http.Client) *ListAuditCredentialsV1Params {

	return &ListAuditCredentialsV1Params{
		HTTPClient: client,
	}
}

/*
ListAuditCredentialsV1Params contains all the parameters to send to the API endpoint
for the list audit credentials v1 operation typically these are written to a http.Request
*/
type ListAuditCredentialsV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) WithTimeout(timeout time.Duration) *ListAuditCredentialsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) WithContext(ctx context.Context) *ListAuditCredentialsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) WithHTTPClient(client *http.Client) *ListAuditCredentialsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list audit credentials v1 params
func (o *ListAuditCredentialsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuditCredentialsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
