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

// NewCreateCustomSdxParams creates a new CreateCustomSdxParams object
// with the default values initialized.
func NewCreateCustomSdxParams() *CreateCustomSdxParams {
	var ()
	return &CreateCustomSdxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomSdxParamsWithTimeout creates a new CreateCustomSdxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomSdxParamsWithTimeout(timeout time.Duration) *CreateCustomSdxParams {
	var ()
	return &CreateCustomSdxParams{

		timeout: timeout,
	}
}

// NewCreateCustomSdxParamsWithContext creates a new CreateCustomSdxParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomSdxParamsWithContext(ctx context.Context) *CreateCustomSdxParams {
	var ()
	return &CreateCustomSdxParams{

		Context: ctx,
	}
}

// NewCreateCustomSdxParamsWithHTTPClient creates a new CreateCustomSdxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomSdxParamsWithHTTPClient(client *http.Client) *CreateCustomSdxParams {
	var ()
	return &CreateCustomSdxParams{
		HTTPClient: client,
	}
}

/*
CreateCustomSdxParams contains all the parameters to send to the API endpoint
for the create custom sdx operation typically these are written to a http.Request
*/
type CreateCustomSdxParams struct {

	/*Body*/
	Body *model.SdxCustomClusterRequest
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create custom sdx params
func (o *CreateCustomSdxParams) WithTimeout(timeout time.Duration) *CreateCustomSdxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom sdx params
func (o *CreateCustomSdxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom sdx params
func (o *CreateCustomSdxParams) WithContext(ctx context.Context) *CreateCustomSdxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom sdx params
func (o *CreateCustomSdxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom sdx params
func (o *CreateCustomSdxParams) WithHTTPClient(client *http.Client) *CreateCustomSdxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom sdx params
func (o *CreateCustomSdxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create custom sdx params
func (o *CreateCustomSdxParams) WithBody(body *model.SdxCustomClusterRequest) *CreateCustomSdxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create custom sdx params
func (o *CreateCustomSdxParams) SetBody(body *model.SdxCustomClusterRequest) {
	o.Body = body
}

// WithName adds the name to the create custom sdx params
func (o *CreateCustomSdxParams) WithName(name string) *CreateCustomSdxParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create custom sdx params
func (o *CreateCustomSdxParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomSdxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
