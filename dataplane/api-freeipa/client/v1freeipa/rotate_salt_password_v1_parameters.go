// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

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

// NewRotateSaltPasswordV1Params creates a new RotateSaltPasswordV1Params object
// with the default values initialized.
func NewRotateSaltPasswordV1Params() *RotateSaltPasswordV1Params {
	var ()
	return &RotateSaltPasswordV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRotateSaltPasswordV1ParamsWithTimeout creates a new RotateSaltPasswordV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRotateSaltPasswordV1ParamsWithTimeout(timeout time.Duration) *RotateSaltPasswordV1Params {
	var ()
	return &RotateSaltPasswordV1Params{

		timeout: timeout,
	}
}

// NewRotateSaltPasswordV1ParamsWithContext creates a new RotateSaltPasswordV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewRotateSaltPasswordV1ParamsWithContext(ctx context.Context) *RotateSaltPasswordV1Params {
	var ()
	return &RotateSaltPasswordV1Params{

		Context: ctx,
	}
}

// NewRotateSaltPasswordV1ParamsWithHTTPClient creates a new RotateSaltPasswordV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRotateSaltPasswordV1ParamsWithHTTPClient(client *http.Client) *RotateSaltPasswordV1Params {
	var ()
	return &RotateSaltPasswordV1Params{
		HTTPClient: client,
	}
}

/*
RotateSaltPasswordV1Params contains all the parameters to send to the API endpoint
for the rotate salt password v1 operation typically these are written to a http.Request
*/
type RotateSaltPasswordV1Params struct {

	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) WithTimeout(timeout time.Duration) *RotateSaltPasswordV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) WithContext(ctx context.Context) *RotateSaltPasswordV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) WithHTTPClient(client *http.Client) *RotateSaltPasswordV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) WithEnvironment(environment *string) *RotateSaltPasswordV1Params {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the rotate salt password v1 params
func (o *RotateSaltPasswordV1Params) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *RotateSaltPasswordV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
