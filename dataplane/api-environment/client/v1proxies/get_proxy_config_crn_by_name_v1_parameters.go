// Code generated by go-swagger; DO NOT EDIT.

package v1proxies

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

// NewGetProxyConfigCrnByNameV1Params creates a new GetProxyConfigCrnByNameV1Params object
// with the default values initialized.
func NewGetProxyConfigCrnByNameV1Params() *GetProxyConfigCrnByNameV1Params {
	var ()
	return &GetProxyConfigCrnByNameV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxyConfigCrnByNameV1ParamsWithTimeout creates a new GetProxyConfigCrnByNameV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProxyConfigCrnByNameV1ParamsWithTimeout(timeout time.Duration) *GetProxyConfigCrnByNameV1Params {
	var ()
	return &GetProxyConfigCrnByNameV1Params{

		timeout: timeout,
	}
}

// NewGetProxyConfigCrnByNameV1ParamsWithContext creates a new GetProxyConfigCrnByNameV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetProxyConfigCrnByNameV1ParamsWithContext(ctx context.Context) *GetProxyConfigCrnByNameV1Params {
	var ()
	return &GetProxyConfigCrnByNameV1Params{

		Context: ctx,
	}
}

// NewGetProxyConfigCrnByNameV1ParamsWithHTTPClient creates a new GetProxyConfigCrnByNameV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProxyConfigCrnByNameV1ParamsWithHTTPClient(client *http.Client) *GetProxyConfigCrnByNameV1Params {
	var ()
	return &GetProxyConfigCrnByNameV1Params{
		HTTPClient: client,
	}
}

/*
GetProxyConfigCrnByNameV1Params contains all the parameters to send to the API endpoint
for the get proxy config crn by name v1 operation typically these are written to a http.Request
*/
type GetProxyConfigCrnByNameV1Params struct {

	/*AccountID*/
	AccountID string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) WithTimeout(timeout time.Duration) *GetProxyConfigCrnByNameV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) WithContext(ctx context.Context) *GetProxyConfigCrnByNameV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) WithHTTPClient(client *http.Client) *GetProxyConfigCrnByNameV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) WithAccountID(accountID string) *GetProxyConfigCrnByNameV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithName adds the name to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) WithName(name string) *GetProxyConfigCrnByNameV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get proxy config crn by name v1 params
func (o *GetProxyConfigCrnByNameV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxyConfigCrnByNameV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
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
