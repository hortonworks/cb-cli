// Code generated by go-swagger; DO NOT EDIT.

package v4custom_configurations

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

// NewGetByNameParams creates a new GetByNameParams object
// with the default values initialized.
func NewGetByNameParams() *GetByNameParams {
	var ()
	return &GetByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetByNameParamsWithTimeout creates a new GetByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetByNameParamsWithTimeout(timeout time.Duration) *GetByNameParams {
	var ()
	return &GetByNameParams{

		timeout: timeout,
	}
}

// NewGetByNameParamsWithContext creates a new GetByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetByNameParamsWithContext(ctx context.Context) *GetByNameParams {
	var ()
	return &GetByNameParams{

		Context: ctx,
	}
}

// NewGetByNameParamsWithHTTPClient creates a new GetByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetByNameParamsWithHTTPClient(client *http.Client) *GetByNameParams {
	var ()
	return &GetByNameParams{
		HTTPClient: client,
	}
}

/*
GetByNameParams contains all the parameters to send to the API endpoint
for the get by name operation typically these are written to a http.Request
*/
type GetByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get by name params
func (o *GetByNameParams) WithTimeout(timeout time.Duration) *GetByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get by name params
func (o *GetByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get by name params
func (o *GetByNameParams) WithContext(ctx context.Context) *GetByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get by name params
func (o *GetByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get by name params
func (o *GetByNameParams) WithHTTPClient(client *http.Client) *GetByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get by name params
func (o *GetByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get by name params
func (o *GetByNameParams) WithName(name string) *GetByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get by name params
func (o *GetByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
