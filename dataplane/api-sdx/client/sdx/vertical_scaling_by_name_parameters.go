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

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// NewVerticalScalingByNameParams creates a new VerticalScalingByNameParams object
// with the default values initialized.
func NewVerticalScalingByNameParams() *VerticalScalingByNameParams {
	var ()
	return &VerticalScalingByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerticalScalingByNameParamsWithTimeout creates a new VerticalScalingByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerticalScalingByNameParamsWithTimeout(timeout time.Duration) *VerticalScalingByNameParams {
	var ()
	return &VerticalScalingByNameParams{

		timeout: timeout,
	}
}

// NewVerticalScalingByNameParamsWithContext creates a new VerticalScalingByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerticalScalingByNameParamsWithContext(ctx context.Context) *VerticalScalingByNameParams {
	var ()
	return &VerticalScalingByNameParams{

		Context: ctx,
	}
}

// NewVerticalScalingByNameParamsWithHTTPClient creates a new VerticalScalingByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerticalScalingByNameParamsWithHTTPClient(client *http.Client) *VerticalScalingByNameParams {
	var ()
	return &VerticalScalingByNameParams{
		HTTPClient: client,
	}
}

/*
VerticalScalingByNameParams contains all the parameters to send to the API endpoint
for the vertical scaling by name operation typically these are written to a http.Request
*/
type VerticalScalingByNameParams struct {

	/*Body*/
	Body *model.StackVerticalScaleV4Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vertical scaling by name params
func (o *VerticalScalingByNameParams) WithTimeout(timeout time.Duration) *VerticalScalingByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vertical scaling by name params
func (o *VerticalScalingByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vertical scaling by name params
func (o *VerticalScalingByNameParams) WithContext(ctx context.Context) *VerticalScalingByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vertical scaling by name params
func (o *VerticalScalingByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vertical scaling by name params
func (o *VerticalScalingByNameParams) WithHTTPClient(client *http.Client) *VerticalScalingByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vertical scaling by name params
func (o *VerticalScalingByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the vertical scaling by name params
func (o *VerticalScalingByNameParams) WithBody(body *model.StackVerticalScaleV4Request) *VerticalScalingByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the vertical scaling by name params
func (o *VerticalScalingByNameParams) SetBody(body *model.StackVerticalScaleV4Request) {
	o.Body = body
}

// WithName adds the name to the vertical scaling by name params
func (o *VerticalScalingByNameParams) WithName(name string) *VerticalScalingByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the vertical scaling by name params
func (o *VerticalScalingByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VerticalScalingByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
