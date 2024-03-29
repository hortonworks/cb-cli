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

// NewDeleteProxyConfigByCrnV1Params creates a new DeleteProxyConfigByCrnV1Params object
// with the default values initialized.
func NewDeleteProxyConfigByCrnV1Params() *DeleteProxyConfigByCrnV1Params {
	var ()
	return &DeleteProxyConfigByCrnV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProxyConfigByCrnV1ParamsWithTimeout creates a new DeleteProxyConfigByCrnV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProxyConfigByCrnV1ParamsWithTimeout(timeout time.Duration) *DeleteProxyConfigByCrnV1Params {
	var ()
	return &DeleteProxyConfigByCrnV1Params{

		timeout: timeout,
	}
}

// NewDeleteProxyConfigByCrnV1ParamsWithContext creates a new DeleteProxyConfigByCrnV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProxyConfigByCrnV1ParamsWithContext(ctx context.Context) *DeleteProxyConfigByCrnV1Params {
	var ()
	return &DeleteProxyConfigByCrnV1Params{

		Context: ctx,
	}
}

// NewDeleteProxyConfigByCrnV1ParamsWithHTTPClient creates a new DeleteProxyConfigByCrnV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProxyConfigByCrnV1ParamsWithHTTPClient(client *http.Client) *DeleteProxyConfigByCrnV1Params {
	var ()
	return &DeleteProxyConfigByCrnV1Params{
		HTTPClient: client,
	}
}

/*
DeleteProxyConfigByCrnV1Params contains all the parameters to send to the API endpoint
for the delete proxy config by crn v1 operation typically these are written to a http.Request
*/
type DeleteProxyConfigByCrnV1Params struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) WithTimeout(timeout time.Duration) *DeleteProxyConfigByCrnV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) WithContext(ctx context.Context) *DeleteProxyConfigByCrnV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) WithHTTPClient(client *http.Client) *DeleteProxyConfigByCrnV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) WithCrn(crn string) *DeleteProxyConfigByCrnV1Params {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the delete proxy config by crn v1 params
func (o *DeleteProxyConfigByCrnV1Params) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProxyConfigByCrnV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
