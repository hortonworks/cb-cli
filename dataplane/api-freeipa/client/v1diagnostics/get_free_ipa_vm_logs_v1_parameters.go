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

// NewGetFreeIpaVMLogsV1Params creates a new GetFreeIpaVMLogsV1Params object
// with the default values initialized.
func NewGetFreeIpaVMLogsV1Params() *GetFreeIpaVMLogsV1Params {

	return &GetFreeIpaVMLogsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFreeIpaVMLogsV1ParamsWithTimeout creates a new GetFreeIpaVMLogsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFreeIpaVMLogsV1ParamsWithTimeout(timeout time.Duration) *GetFreeIpaVMLogsV1Params {

	return &GetFreeIpaVMLogsV1Params{

		timeout: timeout,
	}
}

// NewGetFreeIpaVMLogsV1ParamsWithContext creates a new GetFreeIpaVMLogsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetFreeIpaVMLogsV1ParamsWithContext(ctx context.Context) *GetFreeIpaVMLogsV1Params {

	return &GetFreeIpaVMLogsV1Params{

		Context: ctx,
	}
}

// NewGetFreeIpaVMLogsV1ParamsWithHTTPClient creates a new GetFreeIpaVMLogsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFreeIpaVMLogsV1ParamsWithHTTPClient(client *http.Client) *GetFreeIpaVMLogsV1Params {

	return &GetFreeIpaVMLogsV1Params{
		HTTPClient: client,
	}
}

/*
GetFreeIpaVMLogsV1Params contains all the parameters to send to the API endpoint
for the get free ipa Vm logs v1 operation typically these are written to a http.Request
*/
type GetFreeIpaVMLogsV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) WithTimeout(timeout time.Duration) *GetFreeIpaVMLogsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) WithContext(ctx context.Context) *GetFreeIpaVMLogsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) WithHTTPClient(client *http.Client) *GetFreeIpaVMLogsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get free ipa Vm logs v1 params
func (o *GetFreeIpaVMLogsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFreeIpaVMLogsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
