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

// NewPutStackForAutoscaleStartParams creates a new PutStackForAutoscaleStartParams object
// with the default values initialized.
func NewPutStackForAutoscaleStartParams() *PutStackForAutoscaleStartParams {
	var ()
	return &PutStackForAutoscaleStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutStackForAutoscaleStartParamsWithTimeout creates a new PutStackForAutoscaleStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutStackForAutoscaleStartParamsWithTimeout(timeout time.Duration) *PutStackForAutoscaleStartParams {
	var ()
	return &PutStackForAutoscaleStartParams{

		timeout: timeout,
	}
}

// NewPutStackForAutoscaleStartParamsWithContext creates a new PutStackForAutoscaleStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutStackForAutoscaleStartParamsWithContext(ctx context.Context) *PutStackForAutoscaleStartParams {
	var ()
	return &PutStackForAutoscaleStartParams{

		Context: ctx,
	}
}

// NewPutStackForAutoscaleStartParamsWithHTTPClient creates a new PutStackForAutoscaleStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutStackForAutoscaleStartParamsWithHTTPClient(client *http.Client) *PutStackForAutoscaleStartParams {
	var ()
	return &PutStackForAutoscaleStartParams{
		HTTPClient: client,
	}
}

/*PutStackForAutoscaleStartParams contains all the parameters to send to the API endpoint
for the put stack for autoscale start operation typically these are written to a http.Request
*/
type PutStackForAutoscaleStartParams struct {

	/*Body*/
	Body *model.UpdateStackV4Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) WithTimeout(timeout time.Duration) *PutStackForAutoscaleStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) WithContext(ctx context.Context) *PutStackForAutoscaleStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) WithHTTPClient(client *http.Client) *PutStackForAutoscaleStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) WithBody(body *model.UpdateStackV4Request) *PutStackForAutoscaleStartParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) SetBody(body *model.UpdateStackV4Request) {
	o.Body = body
}

// WithCrn adds the crn to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) WithCrn(crn string) *PutStackForAutoscaleStartParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the put stack for autoscale start params
func (o *PutStackForAutoscaleStartParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *PutStackForAutoscaleStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
