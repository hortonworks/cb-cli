// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCredentialByNameV1Params creates a new GetCredentialByNameV1Params object
// with the default values initialized.
func NewGetCredentialByNameV1Params() *GetCredentialByNameV1Params {
	var ()
	return &GetCredentialByNameV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialByNameV1ParamsWithTimeout creates a new GetCredentialByNameV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCredentialByNameV1ParamsWithTimeout(timeout time.Duration) *GetCredentialByNameV1Params {
	var ()
	return &GetCredentialByNameV1Params{

		timeout: timeout,
	}
}

// NewGetCredentialByNameV1ParamsWithContext creates a new GetCredentialByNameV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCredentialByNameV1ParamsWithContext(ctx context.Context) *GetCredentialByNameV1Params {
	var ()
	return &GetCredentialByNameV1Params{

		Context: ctx,
	}
}

// NewGetCredentialByNameV1ParamsWithHTTPClient creates a new GetCredentialByNameV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCredentialByNameV1ParamsWithHTTPClient(client *http.Client) *GetCredentialByNameV1Params {
	var ()
	return &GetCredentialByNameV1Params{
		HTTPClient: client,
	}
}

/*GetCredentialByNameV1Params contains all the parameters to send to the API endpoint
for the get credential by name v1 operation typically these are written to a http.Request
*/
type GetCredentialByNameV1Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) WithTimeout(timeout time.Duration) *GetCredentialByNameV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) WithContext(ctx context.Context) *GetCredentialByNameV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) WithHTTPClient(client *http.Client) *GetCredentialByNameV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) WithName(name string) *GetCredentialByNameV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get credential by name v1 params
func (o *GetCredentialByNameV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialByNameV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
