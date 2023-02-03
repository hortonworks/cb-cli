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

// NewRotateSaltPasswordSdxByCrnParams creates a new RotateSaltPasswordSdxByCrnParams object
// with the default values initialized.
func NewRotateSaltPasswordSdxByCrnParams() *RotateSaltPasswordSdxByCrnParams {
	var ()
	return &RotateSaltPasswordSdxByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRotateSaltPasswordSdxByCrnParamsWithTimeout creates a new RotateSaltPasswordSdxByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRotateSaltPasswordSdxByCrnParamsWithTimeout(timeout time.Duration) *RotateSaltPasswordSdxByCrnParams {
	var ()
	return &RotateSaltPasswordSdxByCrnParams{

		timeout: timeout,
	}
}

// NewRotateSaltPasswordSdxByCrnParamsWithContext creates a new RotateSaltPasswordSdxByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewRotateSaltPasswordSdxByCrnParamsWithContext(ctx context.Context) *RotateSaltPasswordSdxByCrnParams {
	var ()
	return &RotateSaltPasswordSdxByCrnParams{

		Context: ctx,
	}
}

// NewRotateSaltPasswordSdxByCrnParamsWithHTTPClient creates a new RotateSaltPasswordSdxByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRotateSaltPasswordSdxByCrnParamsWithHTTPClient(client *http.Client) *RotateSaltPasswordSdxByCrnParams {
	var ()
	return &RotateSaltPasswordSdxByCrnParams{
		HTTPClient: client,
	}
}

/*
RotateSaltPasswordSdxByCrnParams contains all the parameters to send to the API endpoint
for the rotate salt password sdx by crn operation typically these are written to a http.Request
*/
type RotateSaltPasswordSdxByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) WithTimeout(timeout time.Duration) *RotateSaltPasswordSdxByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) WithContext(ctx context.Context) *RotateSaltPasswordSdxByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) WithHTTPClient(client *http.Client) *RotateSaltPasswordSdxByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) WithCrn(crn string) *RotateSaltPasswordSdxByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the rotate salt password sdx by crn params
func (o *RotateSaltPasswordSdxByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *RotateSaltPasswordSdxByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
