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

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// NewChangeImageCatalogParams creates a new ChangeImageCatalogParams object
// with the default values initialized.
func NewChangeImageCatalogParams() *ChangeImageCatalogParams {
	var ()
	return &ChangeImageCatalogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeImageCatalogParamsWithTimeout creates a new ChangeImageCatalogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeImageCatalogParamsWithTimeout(timeout time.Duration) *ChangeImageCatalogParams {
	var ()
	return &ChangeImageCatalogParams{

		timeout: timeout,
	}
}

// NewChangeImageCatalogParamsWithContext creates a new ChangeImageCatalogParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeImageCatalogParamsWithContext(ctx context.Context) *ChangeImageCatalogParams {
	var ()
	return &ChangeImageCatalogParams{

		Context: ctx,
	}
}

// NewChangeImageCatalogParamsWithHTTPClient creates a new ChangeImageCatalogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeImageCatalogParamsWithHTTPClient(client *http.Client) *ChangeImageCatalogParams {
	var ()
	return &ChangeImageCatalogParams{
		HTTPClient: client,
	}
}

/*
ChangeImageCatalogParams contains all the parameters to send to the API endpoint
for the change image catalog operation typically these are written to a http.Request
*/
type ChangeImageCatalogParams struct {

	/*Body*/
	Body *model.ChangeImageCatalogRequest
	/*Environment*/
	Environment *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change image catalog params
func (o *ChangeImageCatalogParams) WithTimeout(timeout time.Duration) *ChangeImageCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change image catalog params
func (o *ChangeImageCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change image catalog params
func (o *ChangeImageCatalogParams) WithContext(ctx context.Context) *ChangeImageCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change image catalog params
func (o *ChangeImageCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change image catalog params
func (o *ChangeImageCatalogParams) WithHTTPClient(client *http.Client) *ChangeImageCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change image catalog params
func (o *ChangeImageCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change image catalog params
func (o *ChangeImageCatalogParams) WithBody(body *model.ChangeImageCatalogRequest) *ChangeImageCatalogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change image catalog params
func (o *ChangeImageCatalogParams) SetBody(body *model.ChangeImageCatalogRequest) {
	o.Body = body
}

// WithEnvironment adds the environment to the change image catalog params
func (o *ChangeImageCatalogParams) WithEnvironment(environment *string) *ChangeImageCatalogParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the change image catalog params
func (o *ChangeImageCatalogParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeImageCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
