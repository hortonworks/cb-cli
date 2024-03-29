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

// NewGetCreateEnvironmentForCliByNameParams creates a new GetCreateEnvironmentForCliByNameParams object
// with the default values initialized.
func NewGetCreateEnvironmentForCliByNameParams() *GetCreateEnvironmentForCliByNameParams {
	var ()
	return &GetCreateEnvironmentForCliByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreateEnvironmentForCliByNameParamsWithTimeout creates a new GetCreateEnvironmentForCliByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreateEnvironmentForCliByNameParamsWithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliByNameParams {
	var ()
	return &GetCreateEnvironmentForCliByNameParams{

		timeout: timeout,
	}
}

// NewGetCreateEnvironmentForCliByNameParamsWithContext creates a new GetCreateEnvironmentForCliByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreateEnvironmentForCliByNameParamsWithContext(ctx context.Context) *GetCreateEnvironmentForCliByNameParams {
	var ()
	return &GetCreateEnvironmentForCliByNameParams{

		Context: ctx,
	}
}

// NewGetCreateEnvironmentForCliByNameParamsWithHTTPClient creates a new GetCreateEnvironmentForCliByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreateEnvironmentForCliByNameParamsWithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliByNameParams {
	var ()
	return &GetCreateEnvironmentForCliByNameParams{
		HTTPClient: client,
	}
}

/*
GetCreateEnvironmentForCliByNameParams contains all the parameters to send to the API endpoint
for the get create environment for cli by name operation typically these are written to a http.Request
*/
type GetCreateEnvironmentForCliByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) WithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) WithContext(ctx context.Context) *GetCreateEnvironmentForCliByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) WithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) WithName(name string) *GetCreateEnvironmentForCliByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get create environment for cli by name params
func (o *GetCreateEnvironmentForCliByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreateEnvironmentForCliByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
