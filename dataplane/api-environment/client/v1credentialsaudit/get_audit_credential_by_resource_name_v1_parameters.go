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

// NewGetAuditCredentialByResourceNameV1Params creates a new GetAuditCredentialByResourceNameV1Params object
// with the default values initialized.
func NewGetAuditCredentialByResourceNameV1Params() *GetAuditCredentialByResourceNameV1Params {
	var ()
	return &GetAuditCredentialByResourceNameV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditCredentialByResourceNameV1ParamsWithTimeout creates a new GetAuditCredentialByResourceNameV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuditCredentialByResourceNameV1ParamsWithTimeout(timeout time.Duration) *GetAuditCredentialByResourceNameV1Params {
	var ()
	return &GetAuditCredentialByResourceNameV1Params{

		timeout: timeout,
	}
}

// NewGetAuditCredentialByResourceNameV1ParamsWithContext creates a new GetAuditCredentialByResourceNameV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuditCredentialByResourceNameV1ParamsWithContext(ctx context.Context) *GetAuditCredentialByResourceNameV1Params {
	var ()
	return &GetAuditCredentialByResourceNameV1Params{

		Context: ctx,
	}
}

// NewGetAuditCredentialByResourceNameV1ParamsWithHTTPClient creates a new GetAuditCredentialByResourceNameV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuditCredentialByResourceNameV1ParamsWithHTTPClient(client *http.Client) *GetAuditCredentialByResourceNameV1Params {
	var ()
	return &GetAuditCredentialByResourceNameV1Params{
		HTTPClient: client,
	}
}

/*
GetAuditCredentialByResourceNameV1Params contains all the parameters to send to the API endpoint
for the get audit credential by resource name v1 operation typically these are written to a http.Request
*/
type GetAuditCredentialByResourceNameV1Params struct {

	/*AccountID*/
	AccountID *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) WithTimeout(timeout time.Duration) *GetAuditCredentialByResourceNameV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) WithContext(ctx context.Context) *GetAuditCredentialByResourceNameV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) WithHTTPClient(client *http.Client) *GetAuditCredentialByResourceNameV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) WithAccountID(accountID *string) *GetAuditCredentialByResourceNameV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithName adds the name to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) WithName(name string) *GetAuditCredentialByResourceNameV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get audit credential by resource name v1 params
func (o *GetAuditCredentialByResourceNameV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditCredentialByResourceNameV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
