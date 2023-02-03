// Code generated by go-swagger; DO NOT EDIT.

package v1dns

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

// NewListDNSZonesV1Params creates a new ListDNSZonesV1Params object
// with the default values initialized.
func NewListDNSZonesV1Params() *ListDNSZonesV1Params {
	var ()
	return &ListDNSZonesV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDNSZonesV1ParamsWithTimeout creates a new ListDNSZonesV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDNSZonesV1ParamsWithTimeout(timeout time.Duration) *ListDNSZonesV1Params {
	var ()
	return &ListDNSZonesV1Params{

		timeout: timeout,
	}
}

// NewListDNSZonesV1ParamsWithContext creates a new ListDNSZonesV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListDNSZonesV1ParamsWithContext(ctx context.Context) *ListDNSZonesV1Params {
	var ()
	return &ListDNSZonesV1Params{

		Context: ctx,
	}
}

// NewListDNSZonesV1ParamsWithHTTPClient creates a new ListDNSZonesV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDNSZonesV1ParamsWithHTTPClient(client *http.Client) *ListDNSZonesV1Params {
	var ()
	return &ListDNSZonesV1Params{
		HTTPClient: client,
	}
}

/*
ListDNSZonesV1Params contains all the parameters to send to the API endpoint
for the list Dns zones v1 operation typically these are written to a http.Request
*/
type ListDNSZonesV1Params struct {

	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) WithTimeout(timeout time.Duration) *ListDNSZonesV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) WithContext(ctx context.Context) *ListDNSZonesV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) WithHTTPClient(client *http.Client) *ListDNSZonesV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) WithEnvironment(environment *string) *ListDNSZonesV1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the list Dns zones v1 params
func (o *ListDNSZonesV1Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *ListDNSZonesV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
