// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

// NewStatusDistroXV1ByCrnParams creates a new StatusDistroXV1ByCrnParams object
// with the default values initialized.
func NewStatusDistroXV1ByCrnParams() *StatusDistroXV1ByCrnParams {
	var ()
	return &StatusDistroXV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusDistroXV1ByCrnParamsWithTimeout creates a new StatusDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusDistroXV1ByCrnParamsWithTimeout(timeout time.Duration) *StatusDistroXV1ByCrnParams {
	var ()
	return &StatusDistroXV1ByCrnParams{

		timeout: timeout,
	}
}

// NewStatusDistroXV1ByCrnParamsWithContext creates a new StatusDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusDistroXV1ByCrnParamsWithContext(ctx context.Context) *StatusDistroXV1ByCrnParams {
	var ()
	return &StatusDistroXV1ByCrnParams{

		Context: ctx,
	}
}

// NewStatusDistroXV1ByCrnParamsWithHTTPClient creates a new StatusDistroXV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatusDistroXV1ByCrnParamsWithHTTPClient(client *http.Client) *StatusDistroXV1ByCrnParams {
	var ()
	return &StatusDistroXV1ByCrnParams{
		HTTPClient: client,
	}
}

/*StatusDistroXV1ByCrnParams contains all the parameters to send to the API endpoint
for the status distro x v1 by crn operation typically these are written to a http.Request
*/
type StatusDistroXV1ByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) WithTimeout(timeout time.Duration) *StatusDistroXV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) WithContext(ctx context.Context) *StatusDistroXV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) WithHTTPClient(client *http.Client) *StatusDistroXV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) WithCrn(crn string) *StatusDistroXV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the status distro x v1 by crn params
func (o *StatusDistroXV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *StatusDistroXV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
