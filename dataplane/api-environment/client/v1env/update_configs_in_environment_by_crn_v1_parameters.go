// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

// NewUpdateConfigsInEnvironmentByCrnV1Params creates a new UpdateConfigsInEnvironmentByCrnV1Params object
// with the default values initialized.
func NewUpdateConfigsInEnvironmentByCrnV1Params() *UpdateConfigsInEnvironmentByCrnV1Params {
	var ()
	return &UpdateConfigsInEnvironmentByCrnV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigsInEnvironmentByCrnV1ParamsWithTimeout creates a new UpdateConfigsInEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigsInEnvironmentByCrnV1ParamsWithTimeout(timeout time.Duration) *UpdateConfigsInEnvironmentByCrnV1Params {
	var ()
	return &UpdateConfigsInEnvironmentByCrnV1Params{

		timeout: timeout,
	}
}

// NewUpdateConfigsInEnvironmentByCrnV1ParamsWithContext creates a new UpdateConfigsInEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigsInEnvironmentByCrnV1ParamsWithContext(ctx context.Context) *UpdateConfigsInEnvironmentByCrnV1Params {
	var ()
	return &UpdateConfigsInEnvironmentByCrnV1Params{

		Context: ctx,
	}
}

// NewUpdateConfigsInEnvironmentByCrnV1ParamsWithHTTPClient creates a new UpdateConfigsInEnvironmentByCrnV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigsInEnvironmentByCrnV1ParamsWithHTTPClient(client *http.Client) *UpdateConfigsInEnvironmentByCrnV1Params {
	var ()
	return &UpdateConfigsInEnvironmentByCrnV1Params{
		HTTPClient: client,
	}
}

/*UpdateConfigsInEnvironmentByCrnV1Params contains all the parameters to send to the API endpoint
for the update configs in environment by crn v1 operation typically these are written to a http.Request
*/
type UpdateConfigsInEnvironmentByCrnV1Params struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) WithTimeout(timeout time.Duration) *UpdateConfigsInEnvironmentByCrnV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) WithContext(ctx context.Context) *UpdateConfigsInEnvironmentByCrnV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) WithHTTPClient(client *http.Client) *UpdateConfigsInEnvironmentByCrnV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) WithCrn(crn string) *UpdateConfigsInEnvironmentByCrnV1Params {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the update configs in environment by crn v1 params
func (o *UpdateConfigsInEnvironmentByCrnV1Params) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigsInEnvironmentByCrnV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
