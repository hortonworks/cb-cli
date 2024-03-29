// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutStackForAutoscaleStartByCrnParams creates a new PutStackForAutoscaleStartByCrnParams object
// with the default values initialized.
func NewPutStackForAutoscaleStartByCrnParams() *PutStackForAutoscaleStartByCrnParams {
	var ()
	return &PutStackForAutoscaleStartByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutStackForAutoscaleStartByCrnParamsWithTimeout creates a new PutStackForAutoscaleStartByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutStackForAutoscaleStartByCrnParamsWithTimeout(timeout time.Duration) *PutStackForAutoscaleStartByCrnParams {
	var ()
	return &PutStackForAutoscaleStartByCrnParams{

		timeout: timeout,
	}
}

// NewPutStackForAutoscaleStartByCrnParamsWithContext creates a new PutStackForAutoscaleStartByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutStackForAutoscaleStartByCrnParamsWithContext(ctx context.Context) *PutStackForAutoscaleStartByCrnParams {
	var ()
	return &PutStackForAutoscaleStartByCrnParams{

		Context: ctx,
	}
}

// NewPutStackForAutoscaleStartByCrnParamsWithHTTPClient creates a new PutStackForAutoscaleStartByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutStackForAutoscaleStartByCrnParamsWithHTTPClient(client *http.Client) *PutStackForAutoscaleStartByCrnParams {
	var ()
	return &PutStackForAutoscaleStartByCrnParams{
		HTTPClient: client,
	}
}

/*
PutStackForAutoscaleStartByCrnParams contains all the parameters to send to the API endpoint
for the put stack for autoscale start by crn operation typically these are written to a http.Request
*/
type PutStackForAutoscaleStartByCrnParams struct {

	/*Body*/
	Body *model.UpdateStackV4Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) WithTimeout(timeout time.Duration) *PutStackForAutoscaleStartByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) WithContext(ctx context.Context) *PutStackForAutoscaleStartByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) WithHTTPClient(client *http.Client) *PutStackForAutoscaleStartByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) WithBody(body *model.UpdateStackV4Request) *PutStackForAutoscaleStartByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) SetBody(body *model.UpdateStackV4Request) {
	o.Body = body
}

// WithCrn adds the crn to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) WithCrn(crn string) *PutStackForAutoscaleStartByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the put stack for autoscale start by crn params
func (o *PutStackForAutoscaleStartByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *PutStackForAutoscaleStartByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
