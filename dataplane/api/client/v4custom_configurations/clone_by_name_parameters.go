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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCloneByNameParams creates a new CloneByNameParams object
// with the default values initialized.
func NewCloneByNameParams() *CloneByNameParams {
	var ()
	return &CloneByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloneByNameParamsWithTimeout creates a new CloneByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloneByNameParamsWithTimeout(timeout time.Duration) *CloneByNameParams {
	var ()
	return &CloneByNameParams{

		timeout: timeout,
	}
}

// NewCloneByNameParamsWithContext creates a new CloneByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloneByNameParamsWithContext(ctx context.Context) *CloneByNameParams {
	var ()
	return &CloneByNameParams{

		Context: ctx,
	}
}

// NewCloneByNameParamsWithHTTPClient creates a new CloneByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloneByNameParamsWithHTTPClient(client *http.Client) *CloneByNameParams {
	var ()
	return &CloneByNameParams{
		HTTPClient: client,
	}
}

/*
CloneByNameParams contains all the parameters to send to the API endpoint
for the clone by name operation typically these are written to a http.Request
*/
type CloneByNameParams struct {

	/*Body*/
	Body *model.CloneCustomConfigurationsV4Request
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clone by name params
func (o *CloneByNameParams) WithTimeout(timeout time.Duration) *CloneByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone by name params
func (o *CloneByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone by name params
func (o *CloneByNameParams) WithContext(ctx context.Context) *CloneByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone by name params
func (o *CloneByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone by name params
func (o *CloneByNameParams) WithHTTPClient(client *http.Client) *CloneByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone by name params
func (o *CloneByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the clone by name params
func (o *CloneByNameParams) WithBody(body *model.CloneCustomConfigurationsV4Request) *CloneByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the clone by name params
func (o *CloneByNameParams) SetBody(body *model.CloneCustomConfigurationsV4Request) {
	o.Body = body
}

// WithName adds the name to the clone by name params
func (o *CloneByNameParams) WithName(name string) *CloneByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the clone by name params
func (o *CloneByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CloneByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
