// Code generated by go-swagger; DO NOT EDIT.

package sdx

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

// NewStartSdxByNameParams creates a new StartSdxByNameParams object
// with the default values initialized.
func NewStartSdxByNameParams() *StartSdxByNameParams {
	var ()
	return &StartSdxByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartSdxByNameParamsWithTimeout creates a new StartSdxByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartSdxByNameParamsWithTimeout(timeout time.Duration) *StartSdxByNameParams {
	var ()
	return &StartSdxByNameParams{

		timeout: timeout,
	}
}

// NewStartSdxByNameParamsWithContext creates a new StartSdxByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartSdxByNameParamsWithContext(ctx context.Context) *StartSdxByNameParams {
	var ()
	return &StartSdxByNameParams{

		Context: ctx,
	}
}

// NewStartSdxByNameParamsWithHTTPClient creates a new StartSdxByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartSdxByNameParamsWithHTTPClient(client *http.Client) *StartSdxByNameParams {
	var ()
	return &StartSdxByNameParams{
		HTTPClient: client,
	}
}

/*
StartSdxByNameParams contains all the parameters to send to the API endpoint
for the start sdx by name operation typically these are written to a http.Request
*/
type StartSdxByNameParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start sdx by name params
func (o *StartSdxByNameParams) WithTimeout(timeout time.Duration) *StartSdxByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start sdx by name params
func (o *StartSdxByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start sdx by name params
func (o *StartSdxByNameParams) WithContext(ctx context.Context) *StartSdxByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start sdx by name params
func (o *StartSdxByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start sdx by name params
func (o *StartSdxByNameParams) WithHTTPClient(client *http.Client) *StartSdxByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start sdx by name params
func (o *StartSdxByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the start sdx by name params
func (o *StartSdxByNameParams) WithName(name string) *StartSdxByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the start sdx by name params
func (o *StartSdxByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StartSdxByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
